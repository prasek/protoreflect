package desc

import (
	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

var _ Descriptor = (*MethodDescriptor)(nil)

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

//fqn comes from grpc server info.FullMethod
func (md *MethodDescriptor) GetBoolExtension(field int32, def bool) bool {
	bpb, err := proto.Marshal(md.proto)
	if err != nil {
		return def
	}
	return getBoolExtension(bpb, field, def)
}

/*
func (md *MethodDescriptor) GetBoolExtension(field int32, def bool) bool {
	m := proto.RegisteredExtensions(md.proto.Options)
	if m == nil {
		return def
	}
	ext, ok := m[field]
	if !ok {
		return def
	}
	return proto.GetBoolExtension(md.proto.GetOptions(), ext, def)
}
*/

func getBoolExtension(protoMessage []byte, extField int32, def bool) bool {
	mdp := &dpb.MethodDescriptorProto{}
	err := proto.Unmarshal(protoMessage, mdp)
	if err != nil {
		return def
	}

	m := proto.RegisteredExtensions(mdp.Options)
	if m == nil {
		return def
	}
	ext, ok := m[extField]
	if !ok {
		return def
	}
	return proto.GetBoolExtension(mdp.GetOptions(), ext, def)
}
