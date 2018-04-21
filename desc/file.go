package desc

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// FileDescriptor describes a proto source file.
type FileDescriptor struct {
	proto      *dpb.FileDescriptorProto
	symbols    map[string]Descriptor
	deps       []*FileDescriptor
	publicDeps []*FileDescriptor
	weakDeps   []*FileDescriptor
	messages   []*MessageDescriptor
	enums      []*EnumDescriptor
	extensions []*FieldDescriptor
	services   []*ServiceDescriptor
	fieldIndex map[string]map[int32]*FieldDescriptor
	isProto3   bool
}

// CreateFileDescriptor instantiates a new file descriptor for the given descriptor proto.
// The file's direct dependencies must be provided. If the given dependencies do not include
// all of the file's dependencies or if the contents of the descriptors are internally
// inconsistent (e.g. contain unresolvable symbols) then an error is returned.
func CreateFileDescriptor(fd *dpb.FileDescriptorProto, deps ...*FileDescriptor) (*FileDescriptor, error) {
	ret := &FileDescriptor{proto: fd, symbols: map[string]Descriptor{}, fieldIndex: map[string]map[int32]*FieldDescriptor{}}
	pkg := fd.GetPackage()

	// populate references to file descriptor dependencies
	files := map[string]*FileDescriptor{}
	for _, f := range deps {
		files[f.proto.GetName()] = f
	}
	ret.deps = make([]*FileDescriptor, len(fd.GetDependency()))
	for i, d := range fd.GetDependency() {
		ret.deps[i] = files[d]
		if ret.deps[i] == nil {
			return nil, fmt.Errorf("Given dependencies did not include %q", d)
		}
	}
	ret.publicDeps = make([]*FileDescriptor, len(fd.GetPublicDependency()))
	for i, pd := range fd.GetPublicDependency() {
		ret.publicDeps[i] = ret.deps[pd]
	}
	ret.weakDeps = make([]*FileDescriptor, len(fd.GetWeakDependency()))
	for i, wd := range fd.GetWeakDependency() {
		ret.weakDeps[i] = ret.deps[wd]
	}
	ret.isProto3 = fd.GetSyntax() == "proto3"

	// populate all tables of child descriptors
	for _, m := range fd.GetMessageType() {
		md, n := createMessageDescriptor(ret, ret, pkg, m, ret.symbols)
		ret.symbols[n] = md
		ret.messages = append(ret.messages, md)
	}
	for _, e := range fd.GetEnumType() {
		ed, n := createEnumDescriptor(ret, ret, pkg, e, ret.symbols)
		ret.symbols[n] = ed
		ret.enums = append(ret.enums, ed)
	}
	for _, ex := range fd.GetExtension() {
		exd, n := createFieldDescriptor(ret, ret, pkg, ex)
		ret.symbols[n] = exd
		ret.extensions = append(ret.extensions, exd)
	}
	for _, s := range fd.GetService() {
		sd, n := createServiceDescriptor(ret, pkg, s, ret.symbols)
		ret.symbols[n] = sd
		ret.services = append(ret.services, sd)
	}
	sourceCodeInfo := createsourceInfoMap(fd)

	// now we can resolve all type references and source code info
	scopes := []scope{fileScope(ret)}
	path := make([]int32, 1, 8)
	path[0] = File_messagesTag
	for i, md := range ret.messages {
		if err := md.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}
	path[0] = File_enumsTag
	for i, ed := range ret.enums {
		ed.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	path[0] = File_extensionsTag
	for i, exd := range ret.extensions {
		if err := exd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}
	path[0] = File_servicesTag
	for i, sd := range ret.services {
		if err := sd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// CreateFileDescriptors constructs a set of descriptors, one for each of the
// given descriptor protos. The given set of descriptor protos must include all
// transitive dependencies for every file.
func CreateFileDescriptors(fds []*dpb.FileDescriptorProto) (map[string]*FileDescriptor, error) {
	if len(fds) == 0 {
		return nil, nil
	}
	files := map[string]*dpb.FileDescriptorProto{}
	resolved := map[string]*FileDescriptor{}
	var name string
	for _, fd := range fds {
		name = fd.GetName()
		files[name] = fd
	}
	for _, fd := range fds {
		_, err := createFromSet(fd.GetName(), nil, files, resolved)
		if err != nil {
			return nil, err
		}
	}
	return resolved, nil
}

// CreateFileDescriptorFromSet creates a descriptor from the given file descriptor set. The
// set's *last* file will be the returned descriptor. The set's remaining files must comprise
// the full set of transitive dependencies of that last file. This is the same format and
// order used by protoc when emitting a FileDescriptorSet file with an invocation like so:
//    protoc --descriptor_set_out=./test.protoset --include_imports -I. test.proto
func CreateFileDescriptorFromSet(fds *dpb.FileDescriptorSet) (*FileDescriptor, error) {
	files := fds.GetFile()
	if len(files) == 0 {
		return nil, errors.New("file descriptor set is empty")
	}
	resolved, err := CreateFileDescriptors(files)
	if err != nil {
		return nil, err
	}
	lastFilename := files[len(files)-1].GetName()
	return resolved[lastFilename], nil
}

// createFromSet creates a descriptor for the given filename. It recursively
// creates descriptors for the given file's dependencies.
func createFromSet(filename string, seen []string, files map[string]*dpb.FileDescriptorProto, resolved map[string]*FileDescriptor) (*FileDescriptor, error) {
	for _, s := range seen {
		if filename == s {
			return nil, fmt.Errorf("cycle in imports: %s", strings.Join(append(seen, filename), " -> "))
		}
	}
	seen = append(seen, filename)

	if d, ok := resolved[filename]; ok {
		return d, nil
	}
	fdp := files[filename]
	if fdp == nil {
		return nil, fmt.Errorf("file descriptor set missing a dependency: %s", filename)
	}
	deps := make([]*FileDescriptor, len(fdp.GetDependency()))
	for i, depName := range fdp.GetDependency() {
		if dep, err := createFromSet(depName, seen, files, resolved); err != nil {
			return nil, err
		} else {
			deps[i] = dep
		}
	}
	d, err := CreateFileDescriptor(fdp, deps...)
	if err != nil {
		return nil, err
	}
	resolved[filename] = d
	return d, nil
}

func (fd *FileDescriptor) registerField(field *FieldDescriptor) {
	fields := fd.fieldIndex[field.owner.GetFullyQualifiedName()]
	if fields == nil {
		fields = map[int32]*FieldDescriptor{}
		fd.fieldIndex[field.owner.GetFullyQualifiedName()] = fields
	}
	fields[field.GetNumber()] = field
}

func (fd *FileDescriptor) GetName() string {
	return fd.proto.GetName()
}

func (fd *FileDescriptor) GetFullyQualifiedName() string {
	return fd.proto.GetName()
}

func (fd *FileDescriptor) GetPackage() string {
	return fd.proto.GetPackage()
}

func (fd *FileDescriptor) GetParent() Descriptor {
	return nil
}

func (fd *FileDescriptor) GetFile() *FileDescriptor {
	return fd
}

func (fd *FileDescriptor) GetOptions() proto.Message {
	return fd.proto.GetOptions()
}

func (fd *FileDescriptor) GetFileOptions() *dpb.FileOptions {
	return fd.proto.GetOptions()
}

func (fd *FileDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return nil
}

func (fd *FileDescriptor) AsProto() proto.Message {
	return fd.proto
}

func (fd *FileDescriptor) AsFileDescriptorProto() *dpb.FileDescriptorProto {
	return fd.proto
}

func (fd *FileDescriptor) String() string {
	return fd.proto.String()
}

// IsProto3 returns true if the file declares a syntax of "proto3".
func (fd *FileDescriptor) IsProto3() bool {
	return fd.isProto3
}

// GetDependencies returns all of this file's dependencies. These correspond to
// import statements in the file.
func (fd *FileDescriptor) GetDependencies() []*FileDescriptor {
	return fd.deps
}

// GetPublicDependencies returns all of this file's public dependencies. These
// correspond to public import statements in the file.
func (fd *FileDescriptor) GetPublicDependencies() []*FileDescriptor {
	return fd.publicDeps
}

// GetWeakDependencies returns all of this file's weak dependencies. These
// correspond to weak import statements in the file.
func (fd *FileDescriptor) GetWeakDependencies() []*FileDescriptor {
	return fd.weakDeps
}

// GetMessageTypes returns all top-level messages declared in this file.
func (fd *FileDescriptor) GetMessageTypes() []*MessageDescriptor {
	return fd.messages
}

// GetEnumTypes returns all top-level enums declared in this file.
func (fd *FileDescriptor) GetEnumTypes() []*EnumDescriptor {
	return fd.enums
}

// GetExtensions returns all top-level extensions declared in this file.
func (fd *FileDescriptor) GetExtensions() []*FieldDescriptor {
	return fd.extensions
}

// GetServices returns all services declared in this file.
func (fd *FileDescriptor) GetServices() []*ServiceDescriptor {
	return fd.services
}

// FindSymbol returns the descriptor contained within this file for the
// element with the given fully-qualified symbol name. If no such element
// exists then this method returns nil.
func (fd *FileDescriptor) FindSymbol(symbol string) Descriptor {
	return fd.symbols[symbol]
}

// FindMessage finds the message with the given fully-qualified name. If no
// such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindMessage(msgName string) *MessageDescriptor {
	if md, ok := fd.symbols[msgName].(*MessageDescriptor); ok {
		return md
	} else {
		return nil
	}
}

// FindEnum finds the enum with the given fully-qualified name. If no such
// element exists in this file then nil is returned.
func (fd *FileDescriptor) FindEnum(enumName string) *EnumDescriptor {
	if ed, ok := fd.symbols[enumName].(*EnumDescriptor); ok {
		return ed
	} else {
		return nil
	}
}

// FindService finds the service with the given fully-qualified name. If no
// such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindService(serviceName string) *ServiceDescriptor {
	if sd, ok := fd.symbols[serviceName].(*ServiceDescriptor); ok {
		return sd
	} else {
		return nil
	}
}

// FindExtension finds the extension field for the given extended type name and
// tag number. If no such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindExtension(extendeeName string, tagNumber int32) *FieldDescriptor {
	if exd, ok := fd.fieldIndex[extendeeName][tagNumber]; ok && exd.IsExtension() {
		return exd
	} else {
		return nil
	}
}

// FindExtensionByName finds the extension field with the given fully-qualified
// name. If no such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindExtensionByName(extName string) *FieldDescriptor {
	if exd, ok := fd.symbols[extName].(*FieldDescriptor); ok && exd.IsExtension() {
		return exd
	} else {
		return nil
	}
}

func fileScope(fd *FileDescriptor) scope {
	// we search symbols in this file, but also symbols in other files that have
	// the same package as this file or a "parent" package (in protobuf,
	// packages are a hierarchy like C++ namespaces)
	prefixes := createPrefixList(fd.proto.GetPackage())
	return func(name string) Descriptor {
		for _, prefix := range prefixes {
			n := merge(prefix, name)
			d := findSymbol(fd, n, false)
			if d != nil {
				return d
			}
		}
		return nil
	}
}

func resolve(fd *FileDescriptor, name string, scopes []scope) (Descriptor, error) {
	if strings.HasPrefix(name, ".") {
		// already fully-qualified
		d := findSymbol(fd, name[1:], false)
		if d != nil {
			return d, nil
		}
	} else {
		// unqualified, so we look in the enclosing (last) scope first and move
		// towards outermost (first) scope, trying to resolve the symbol
		for i := len(scopes) - 1; i >= 0; i-- {
			d := scopes[i](name)
			if d != nil {
				return d, nil
			}
		}
	}
	return nil, fmt.Errorf("File %q included an unresolvable reference to %q", fd.proto.GetName(), name)
}

func findSymbol(fd *FileDescriptor, name string, public bool) Descriptor {
	d := fd.symbols[name]
	if d != nil {
		return d
	}

	// When public = false, we are searching only directly imported symbols. But we
	// also need to search transitive public imports due to semantics of public imports.
	var deps []*FileDescriptor
	if public {
		deps = fd.publicDeps
	} else {
		deps = fd.deps
	}
	for _, dep := range deps {
		d = findSymbol(dep, name, true)
		if d != nil {
			return d
		}
	}

	return nil
}
