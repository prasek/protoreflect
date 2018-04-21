package desc

import (
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

var _ Descriptor = (*MethodDescriptor)(nil)

// MethodDescriptor describes an RPC method declared in a proto file.
type MethodDescriptor struct {
	proto      *dpb.MethodDescriptorProto
	parent     *ServiceDescriptor
	file       *FileDescriptor
	inType     *MessageDescriptor
	outType    *MessageDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

func createMethodDescriptor(fd *FileDescriptor, parent *ServiceDescriptor, enclosing string, md *dpb.MethodDescriptorProto) (*MethodDescriptor, string) {
	// request and response types get resolved later
	methodName := merge(enclosing, md.GetName())
	return &MethodDescriptor{proto: md, parent: parent, file: fd, fqn: methodName}, methodName
}

func (md *MethodDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	md.sourceInfo = sourceCodeInfo.Get(path)
	if desc, err := resolve(md.file, md.proto.GetInputType(), scopes); err != nil {
		return err
	} else {
		md.inType = desc.(*MessageDescriptor)
	}
	if desc, err := resolve(md.file, md.proto.GetOutputType(), scopes); err != nil {
		return err
	} else {
		md.outType = desc.(*MessageDescriptor)
	}
	return nil
}

func (md *MethodDescriptor) GetName() string {
	return md.proto.GetName()
}

func (md *MethodDescriptor) GetFullyQualifiedName() string {
	return md.fqn
}

func (md *MethodDescriptor) GetParent() Descriptor {
	return md.parent
}

// GetService returns the RPC service in which this method is declared.
func (md *MethodDescriptor) GetService() *ServiceDescriptor {
	return md.parent
}

func (md *MethodDescriptor) GetFile() *FileDescriptor {
	return md.file
}

func (md *MethodDescriptor) GetMethodOptions() *dpb.MethodOptions {
	return md.proto.GetOptions()
}

func (md *MethodDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return md.sourceInfo
}

func (md *MethodDescriptor) AsMethodDescriptorProto() *dpb.MethodDescriptorProto {
	return md.proto
}

func (md *MethodDescriptor) String() string {
	return md.proto.String()
}

// IsServerStreaming returns true if this is a server-streaming method.
func (md *MethodDescriptor) IsServerStreaming() bool {
	return md.proto.GetServerStreaming()
}

// IsClientStreaming returns true if this is a client-streaming method.
func (md *MethodDescriptor) IsClientStreaming() bool {
	return md.proto.GetClientStreaming()
}

// GetInputType returns the input type, or request type, of the RPC method.
func (md *MethodDescriptor) GetInputType() *MessageDescriptor {
	return md.inType
}

// GetOutputType returns the output type, or response type, of the RPC method.
func (md *MethodDescriptor) GetOutputType() *MessageDescriptor {
	return md.outType
}
