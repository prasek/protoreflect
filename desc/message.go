package desc

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// MessageDescriptor describes a protocol buffer message.
type MessageDescriptor struct {
	proto      *dpb.DescriptorProto
	parent     Descriptor
	file       *FileDescriptor
	fields     []*FieldDescriptor
	nested     []*MessageDescriptor
	enums      []*EnumDescriptor
	extensions []*FieldDescriptor
	oneOfs     []*OneOfDescriptor
	extRanges  extRanges
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
	isProto3   bool
	isMapEntry bool
}

func createMessageDescriptor(fd *FileDescriptor, parent Descriptor, enclosing string, md *dpb.DescriptorProto, symbols map[string]Descriptor) (*MessageDescriptor, string) {
	msgName := merge(enclosing, md.GetName())
	ret := &MessageDescriptor{proto: md, parent: parent, file: fd, fqn: msgName}
	for _, f := range md.GetField() {
		fld, n := createFieldDescriptor(fd, ret, msgName, f)
		symbols[n] = fld
		ret.fields = append(ret.fields, fld)
	}
	for _, nm := range md.NestedType {
		nmd, n := createMessageDescriptor(fd, ret, msgName, nm, symbols)
		symbols[n] = nmd
		ret.nested = append(ret.nested, nmd)
	}
	for _, e := range md.EnumType {
		ed, n := createEnumDescriptor(fd, ret, msgName, e, symbols)
		symbols[n] = ed
		ret.enums = append(ret.enums, ed)
	}
	for _, ex := range md.GetExtension() {
		exd, n := createFieldDescriptor(fd, ret, msgName, ex)
		symbols[n] = exd
		ret.extensions = append(ret.extensions, exd)
	}
	for i, o := range md.GetOneofDecl() {
		od, n := createOneOfDescriptor(fd, ret, i, msgName, o)
		symbols[n] = od
		ret.oneOfs = append(ret.oneOfs, od)
	}
	for _, r := range md.GetExtensionRange() {
		// proto.ExtensionRange is inclusive (and that's how extension ranges are defined in code).
		// but protoc converts range to exclusive end in descriptor, so we must convert back
		end := r.GetEnd() - 1
		ret.extRanges = append(ret.extRanges, proto.ExtensionRange{
			Start: r.GetStart(),
			End:   end})
	}
	sort.Sort(ret.extRanges)
	ret.isProto3 = fd.isProto3
	ret.isMapEntry = md.GetOptions().GetMapEntry() &&
		len(ret.fields) == 2 &&
		ret.fields[0].GetNumber() == 1 &&
		ret.fields[1].GetNumber() == 2

	return ret, msgName
}

func (md *MessageDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	md.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Message_nestedMessagesTag)
	scopes = append(scopes, messageScope(md))
	for i, nmd := range md.nested {
		if err := nmd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_enumsTag
	for i, ed := range md.enums {
		ed.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	path[len(path)-1] = Message_fieldsTag
	for i, fld := range md.fields {
		if err := fld.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_extensionsTag
	for i, exd := range md.extensions {
		if err := exd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_oneOfsTag
	for i, od := range md.oneOfs {
		od.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	return nil
}

func (md *MessageDescriptor) GetName() string {
	return md.proto.GetName()
}

func (md *MessageDescriptor) GetFullyQualifiedName() string {
	return md.fqn
}

func (md *MessageDescriptor) GetParent() Descriptor {
	return md.parent
}

func (md *MessageDescriptor) GetFile() *FileDescriptor {
	return md.file
}

func (md *MessageDescriptor) GetOptions() proto.Message {
	return md.proto.GetOptions()
}

func (md *MessageDescriptor) GetMessageOptions() *dpb.MessageOptions {
	return md.proto.GetOptions()
}

func (md *MessageDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return md.sourceInfo
}

func (md *MessageDescriptor) AsProto() proto.Message {
	return md.proto
}

func (md *MessageDescriptor) AsDescriptorProto() *dpb.DescriptorProto {
	return md.proto
}

func (md *MessageDescriptor) String() string {
	return md.proto.String()
}

// IsMapEntry returns true if this is a synthetic message type that represents an entry
// in a map field.
func (md *MessageDescriptor) IsMapEntry() bool {
	return md.isMapEntry
}

// GetFields returns all of the fields for this message.
func (md *MessageDescriptor) GetFields() []*FieldDescriptor {
	return md.fields
}

// GetNestedMessageTypes returns all of the message types declared inside this message.
func (md *MessageDescriptor) GetNestedMessageTypes() []*MessageDescriptor {
	return md.nested
}

// GetNestedEnumTypes returns all of the enums declared inside this message.
func (md *MessageDescriptor) GetNestedEnumTypes() []*EnumDescriptor {
	return md.enums
}

// GetNestedExtensions returns all of the extensions declared inside this message.
func (md *MessageDescriptor) GetNestedExtensions() []*FieldDescriptor {
	return md.extensions
}

// GetOneOfs returns all of the one-of field sets declared inside this message.
func (md *MessageDescriptor) GetOneOfs() []*OneOfDescriptor {
	return md.oneOfs
}

// IsProto3 returns true if the file in which this message is defined declares a syntax of "proto3".
func (md *MessageDescriptor) IsProto3() bool {
	return md.isProto3
}

// GetExtensionRanges returns the ranges of extension field numbers for this message.
func (md *MessageDescriptor) GetExtensionRanges() []proto.ExtensionRange {
	return md.extRanges
}

// IsExtendable returns true if this message has any extension ranges.
func (md *MessageDescriptor) IsExtendable() bool {
	return len(md.extRanges) > 0
}

// IsExtension returns true if the given tag number is within any of this message's
// extension ranges.
func (md *MessageDescriptor) IsExtension(tagNumber int32) bool {
	return md.extRanges.IsExtension(tagNumber)
}

type extRanges []proto.ExtensionRange

func (er extRanges) String() string {
	var buf bytes.Buffer
	first := true
	for _, r := range er {
		if first {
			first = false
		} else {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d..%d", r.Start, r.End)
	}
	return buf.String()
}

func (er extRanges) IsExtension(tagNumber int32) bool {
	i := sort.Search(len(er), func(i int) bool { return er[i].End >= tagNumber })
	return i < len(er) && tagNumber >= er[i].Start
}

func (er extRanges) Len() int {
	return len(er)
}

func (er extRanges) Less(i, j int) bool {
	return er[i].Start < er[j].Start
}

func (er extRanges) Swap(i, j int) {
	er[i], er[j] = er[j], er[i]
}

// FindFieldByName finds the field with the given name. If no such field exists
// then nil is returned. Only regular fields are returned, not extensions.
func (md *MessageDescriptor) FindFieldByName(fieldName string) *FieldDescriptor {
	fqn := fmt.Sprintf("%s.%s", md.fqn, fieldName)
	if fd, ok := md.file.symbols[fqn].(*FieldDescriptor); ok && !fd.IsExtension() {
		return fd
	} else {
		return nil
	}
}

// FindFieldByNumber finds the field with the given tag number. If no such field
// exists then nil is returned. Only regular fields are returned, not extensions.
func (md *MessageDescriptor) FindFieldByNumber(tagNumber int32) *FieldDescriptor {
	if fd, ok := md.file.fieldIndex[md.fqn][tagNumber]; ok && !fd.IsExtension() {
		return fd
	} else {
		return nil
	}
}

func messageScope(md *MessageDescriptor) scope {
	return func(name string) Descriptor {
		n := merge(md.fqn, name)
		if d, ok := md.file.symbols[n]; ok {
			return d
		}
		return nil
	}
}
