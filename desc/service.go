package desc

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// ServiceDescriptor describes an RPC service declared in a proto file.
type ServiceDescriptor struct {
	proto      *dpb.ServiceDescriptorProto
	file       *FileDescriptor
	methods    []*MethodDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

func createServiceDescriptor(fd *FileDescriptor, enclosing string, sd *dpb.ServiceDescriptorProto, symbols map[string]Descriptor) (*ServiceDescriptor, string) {
	serviceName := merge(enclosing, sd.GetName())
	ret := &ServiceDescriptor{proto: sd, file: fd, fqn: serviceName}
	for _, m := range sd.GetMethod() {
		md, n := createMethodDescriptor(fd, ret, serviceName, m)
		symbols[n] = md
		ret.methods = append(ret.methods, md)
	}
	return ret, serviceName
}

func (sd *ServiceDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	sd.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Service_methodsTag)
	for i, md := range sd.methods {
		if err := md.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	return nil
}

func (sd *ServiceDescriptor) GetName() string {
	return sd.proto.GetName()
}

func (sd *ServiceDescriptor) GetFullyQualifiedName() string {
	return sd.fqn
}

func (sd *ServiceDescriptor) GetParent() Descriptor {
	return sd.file
}

func (sd *ServiceDescriptor) GetFile() *FileDescriptor {
	return sd.file
}

func (sd *ServiceDescriptor) GetOptions() proto.Message {
	return sd.proto.GetOptions()
}

func (sd *ServiceDescriptor) GetServiceOptions() *dpb.ServiceOptions {
	return sd.proto.GetOptions()
}

func (sd *ServiceDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return sd.sourceInfo
}

func (sd *ServiceDescriptor) AsProto() proto.Message {
	return sd.proto
}

func (sd *ServiceDescriptor) AsServiceDescriptorProto() *dpb.ServiceDescriptorProto {
	return sd.proto
}

func (sd *ServiceDescriptor) String() string {
	return sd.proto.String()
}

// GetMethods returns all of the RPC methods for this service.
func (sd *ServiceDescriptor) GetMethods() []*MethodDescriptor {
	return sd.methods
}

// FindMethodByName finds the method with the given name. If no such method exists
// then nil is returned.
func (sd *ServiceDescriptor) FindMethodByName(name string) *MethodDescriptor {
	fqn := fmt.Sprintf("%s.%s", sd.fqn, name)
	if md, ok := sd.file.symbols[fqn].(*MethodDescriptor); ok {
		return md
	} else {
		return nil
	}
}
