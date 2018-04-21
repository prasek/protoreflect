package desc

import (
	"testing"

	"google.golang.org/grpc"

	"github.com/jhump/protoreflect/internal/testprotos"
	"github.com/stretchr/testify/require"
)

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.TestService",
	HandlerType: (*testprotos.TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName: "DoSomethingElse",
		},
		{
			StreamName: "DoSomethingAgain",
		},
		{
			StreamName: "DoSomethingForever",
		},
	},
	Metadata: "desc_test_proto3.proto",
}

type testService struct {
	testprotos.TestServiceServer
}

func TestMethodExtension(t *testing.T) {
	/*
	   func (gio *GrpcInterceptorOptions) IsTx(info *grpc.UnaryServerInfo) bool {
	   	method := gio.methods[info.FullMethod]
	   	tx := proto.GetBoolExtension(method.GetMethodOptions(), stub.E_Tx, false)
	   	return tx
	   }
	*/

}

func TestLoadServiceDescriptors(t *testing.T) {
	s := grpc.NewServer()
	testprotos.RegisterTestServiceServer(s, testService{})
	sds, err := LoadServiceDescriptors(s)
	require.Nil(t, err, err)
	require.Equal(t, 1, len(sds), "service descriptor len")
	sd := sds["testprotos.TestService"]

	cases := []struct{ method, request, response string }{
		{"DoSomething", "testprotos.TestRequest", "jhump.protoreflect.desc.Bar"},
		{"DoSomethingElse", "testprotos.TestMessage", "testprotos.TestResponse"},
		{"DoSomethingAgain", "jhump.protoreflect.desc.Bar", "testprotos.AnotherTestMessage"},
		{"DoSomethingForever", "testprotos.TestRequest", "testprotos.TestResponse"},
	}

	require.Equal(t, len(cases), len(sd.GetMethods()))

	for i, c := range cases {
		md := sd.GetMethods()[i]
		require.Equal(t, c.method, md.GetName())
		require.Equal(t, c.request, md.GetInputType().GetFullyQualifiedName())
		require.Equal(t, c.response, md.GetOutputType().GetFullyQualifiedName())
	}
}

func TestLoadServiceDescriptor(t *testing.T) {
	s := grpc.NewServer()
	testprotos.RegisterTestServiceServer(s, testService{})
	sd, err := LoadServiceDescriptor(&_TestService_serviceDesc)
	require.Nil(t, err)
	require.NotNil(t, sd, "Service descriptor missing")

	cases := []struct{ method, request, response string }{
		{"DoSomething", "testprotos.TestRequest", "jhump.protoreflect.desc.Bar"},
		{"DoSomethingElse", "testprotos.TestMessage", "testprotos.TestResponse"},
		{"DoSomethingAgain", "jhump.protoreflect.desc.Bar", "testprotos.AnotherTestMessage"},
		{"DoSomethingForever", "testprotos.TestRequest", "testprotos.TestResponse"},
	}

	require.Equal(t, len(cases), len(sd.GetMethods()))

	for i, c := range cases {
		md := sd.GetMethods()[i]
		require.Equal(t, c.method, md.GetName())
		require.Equal(t, c.request, md.GetInputType().GetFullyQualifiedName())
		require.Equal(t, c.response, md.GetOutputType().GetFullyQualifiedName())
	}
}
