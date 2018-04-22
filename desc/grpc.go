package desc

import (
	"fmt"

	"google.golang.org/grpc"
)

func LoadInterceptorInfo(svr *grpc.Server) (*InterceptorInfo, error) {
	info := newInterceptorInfo()
	sds, err := LoadServiceDescriptors(svr)
	if err != nil {
		return nil, err
	}
	for _, sd := range sds {
		for _, md := range sd.GetMethods() {
			fqn := fmt.Sprintf("/%s/%s", sd.GetFullyQualifiedName(), md.GetName())
			info.register(fqn, sd, md)
		}
	}
	return info, nil
}

func newInterceptorInfo() *InterceptorInfo {
	return &InterceptorInfo{
		methods: make(map[string]*InterceptorInfoEntry),
	}
}

type InterceptorInfo struct {
	methods map[string]*InterceptorInfoEntry
}

type InterceptorInfoEntry struct {
	Service *ServiceDescriptor
	Method  *MethodDescriptor
}

func (info *InterceptorInfo) register(fqn string, sd *ServiceDescriptor, md *MethodDescriptor) {
	info.methods[fqn] = &InterceptorInfoEntry{Service: sd, Method: md}
}

// LoadServiceDescriptors loads the service descriptors for all services exposed by the
// given GRPC server.
func LoadServiceDescriptors(s *grpc.Server) (map[string]*ServiceDescriptor, error) {
	descs := map[string]*ServiceDescriptor{}
	for name, info := range s.GetServiceInfo() {
		file, ok := info.Metadata.(string)
		if !ok {
			return nil, fmt.Errorf("Service %q has unexpected metadata. Expecting a string, got %v", name, info.Metadata)
		}
		fd, err := LoadFileDescriptor(file)
		if err != nil {
			return nil, err
		}
		d := fd.FindSymbol(name)
		if d == nil {
			return nil, fmt.Errorf("File descriptor for %q has no element named %q", file, name)
		}
		sd, ok := d.(*ServiceDescriptor)
		if !ok {
			return nil, fmt.Errorf("File descriptor for %q has incorrect element named %q. Expecting a service descriptor, got %v", file, name, d)
		}
		descs[name] = sd
	}
	return descs, nil
}

// LoadServiceDescriptor loads a rich descriptor for a given service description
// generated by protoc-gen-gogo. Generated code contains an unexported symbol with
// a name like "_<Service>_serviceDesc" which is the service's description. It
// is used internally to register a service implementation with a GRPC server.
// But it can also be used by this package to retrieve the rich descriptor for
// the service.
func LoadServiceDescriptor(svc *grpc.ServiceDesc) (*ServiceDescriptor, error) {
	file, ok := svc.Metadata.(string)
	if !ok {
		return nil, fmt.Errorf("Service %q has unexpected metadata. Expecting a string, got %v", svc.ServiceName, svc.Metadata)
	}
	fd, err := LoadFileDescriptor(file)
	if err != nil {
		return nil, err
	}
	d := fd.FindSymbol(svc.ServiceName)
	if d == nil {
		return nil, fmt.Errorf("File descriptor for %q has no element named %q", file, svc.ServiceName)
	}
	sd, ok := d.(*ServiceDescriptor)
	if !ok {
		return nil, fmt.Errorf("File descriptor for %q has incorrect element named %q. Expecting a service descriptor, got %v", file, svc.ServiceName, d)
	}
	return sd, nil
}
