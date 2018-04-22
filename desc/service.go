package desc

import (
	"fmt"
)

var _ Descriptor = (*ServiceDescriptor)(nil)

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
