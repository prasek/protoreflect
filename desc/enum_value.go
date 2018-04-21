package desc

import dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// EnumValueDescriptor describes an allowed value of an enum declared in a proto file.
type EnumValueDescriptor struct {
	proto      *dpb.EnumValueDescriptorProto
	parent     *EnumDescriptor
	file       *FileDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

func createEnumValueDescriptor(fd *FileDescriptor, parent *EnumDescriptor, enclosing string, evd *dpb.EnumValueDescriptorProto) (*EnumValueDescriptor, string) {
	valName := merge(enclosing, evd.GetName())
	return &EnumValueDescriptor{proto: evd, parent: parent, file: fd, fqn: valName}, valName
}

func (vd *EnumValueDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	vd.sourceInfo = sourceCodeInfo.Get(path)
}

func (vd *EnumValueDescriptor) GetName() string {
	return vd.proto.GetName()
}

// GetNumber returns the numeric value associated with this enum value.
func (vd *EnumValueDescriptor) GetNumber() int32 {
	return vd.proto.GetNumber()
}

func (vd *EnumValueDescriptor) GetFullyQualifiedName() string {
	return vd.fqn
}

func (vd *EnumValueDescriptor) GetParent() Descriptor {
	return vd.parent
}

// GetEnum returns the enum in which this enum value is defined.
func (vd *EnumValueDescriptor) GetEnum() *EnumDescriptor {
	return vd.parent
}

func (vd *EnumValueDescriptor) GetFile() *FileDescriptor {
	return vd.file
}

func (vd *EnumValueDescriptor) String() string {
	return vd.proto.String()
}
