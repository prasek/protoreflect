package desc

import (
	"testing"

	"github.com/jhump/protoreflect/internal/testutil"
)

func TestServiceDescriptors(t *testing.T) {
	fd, err := LoadFileDescriptor("desc_test_proto3.proto")
	testutil.Ok(t, err)
	sd := fd.FindSymbol("testprotos.TestService").(*ServiceDescriptor)
	// check the descriptor graph for this service and its descendants
	checkDescriptor(t, "service", 0, sd, fd, fd, descCase{
		name: "testprotos.TestService",
		references: map[string]childCases{
			"methods": {(*ServiceDescriptor).GetMethods, []descCase{
				{
					name: "testprotos.TestService.DoSomething",
					references: map[string]childCases{
						"request":  {(*MethodDescriptor).GetInputType, refs("testprotos.TestRequest")},
						"response": {(*MethodDescriptor).GetOutputType, refs("jhump.protoreflect.desc.Bar")},
					},
				},
				{
					name: "testprotos.TestService.DoSomethingElse",
					references: map[string]childCases{
						"request":  {(*MethodDescriptor).GetInputType, refs("testprotos.TestMessage")},
						"response": {(*MethodDescriptor).GetOutputType, refs("testprotos.TestResponse")},
					},
				},
				{
					name: "testprotos.TestService.DoSomethingAgain",
					references: map[string]childCases{
						"request":  {(*MethodDescriptor).GetInputType, refs("jhump.protoreflect.desc.Bar")},
						"response": {(*MethodDescriptor).GetOutputType, refs("testprotos.AnotherTestMessage")},
					},
				},
				{
					name: "testprotos.TestService.DoSomethingForever",
					references: map[string]childCases{
						"request":  {(*MethodDescriptor).GetInputType, refs("testprotos.TestRequest")},
						"response": {(*MethodDescriptor).GetOutputType, refs("testprotos.TestResponse")},
					},
				},
			}},
		},
	})
	// now verify that FindMethodByName works correctly
	for _, md := range sd.GetMethods() {
		found := sd.FindMethodByName(md.GetName())
		testutil.Eq(t, md, found)
	}
	testutil.Eq(t, (*MethodDescriptor)(nil), sd.FindMethodByName("junk name"))
}
