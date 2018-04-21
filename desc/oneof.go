package desc

import (
	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

var _ Descriptor = (*OneOfDescriptor)(nil)

// OneOfDescriptor describes a one-of field set declared in a protocol buffer message.
type OneOfDescriptor struct {
	proto      *dpb.OneofDescriptorProto
	parent     *MessageDescriptor
	file       *FileDescriptor
	choices    []*FieldDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

func createOneOfDescriptor(fd *FileDescriptor, parent *MessageDescriptor, index int, enclosing string, od *dpb.OneofDescriptorProto) (*OneOfDescriptor, string) {
	oneOfName := merge(enclosing, od.GetName())
	ret := &OneOfDescriptor{proto: od, parent: parent, file: fd, fqn: oneOfName}
	for _, f := range parent.fields {
		oi := f.proto.OneofIndex
		if oi != nil && *oi == int32(index) {
			f.oneOf = ret
			ret.choices = append(ret.choices, f)
		}
	}
	return ret, oneOfName
}

func (od *OneOfDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	od.sourceInfo = sourceCodeInfo.Get(path)
}

func (od *OneOfDescriptor) GetName() string {
	return od.proto.GetName()
}

func (od *OneOfDescriptor) GetFullyQualifiedName() string {
	return od.fqn
}

func (od *OneOfDescriptor) GetParent() Descriptor {
	return od.parent
}

// GetOwner returns the message to which this one-of field set belongs.
func (od *OneOfDescriptor) GetOwner() *MessageDescriptor {
	return od.parent
}

func (od *OneOfDescriptor) GetFile() *FileDescriptor {
	return od.file
}

func (od *OneOfDescriptor) GetOptions() proto.Message {
	return od.proto.GetOptions()
}

func (od *OneOfDescriptor) GetOneOfOptions() *dpb.OneofOptions {
	return od.proto.GetOptions()
}

func (od *OneOfDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return od.sourceInfo
}

func (od *OneOfDescriptor) AsProto() proto.Message {
	return od.proto
}

func (od *OneOfDescriptor) AsOneofDescriptorProto() *dpb.OneofDescriptorProto {
	return od.proto
}

func (od *OneOfDescriptor) String() string {
	return od.proto.String()
}

// GetChoices returns the fields that are part of the one-of field set. At most one of
// these fields may be set for a given message.
func (od *OneOfDescriptor) GetChoices() []*FieldDescriptor {
	return od.choices
}
