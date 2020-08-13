package protoparse

import (
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	_ "github.com/gogo/protobuf/types"

	"github.com/jhump/protoreflect/internal"
)

// All files that are included with protoc are also included with this package
// so that clients do not need to explicitly supply a copy of these protos (just
// like callers of protoc do not need to supply them).
var standardImports map[string]*dpb.FileDescriptorProto

func init() {
	standardFilenames := []string{
		"google/protobuf/any.proto",
		"google/protobuf/api.proto",
		"google/protobuf/compiler/plugin.proto",
		"google/protobuf/descriptor.proto",
		"google/protobuf/duration.proto",
		"google/protobuf/empty.proto",
		"google/protobuf/field_mask.proto",
		"google/protobuf/source_context.proto",
		"google/protobuf/struct.proto",
		"google/protobuf/timestamp.proto",
		"google/protobuf/type.proto",
		"google/protobuf/wrappers.proto",
	}

	standardImports = map[string]*dpb.FileDescriptorProto{}
	for _, fn := range standardFilenames {
		fd, err := internal.LoadFileDescriptor(fn)
		if err != nil {
			panic(err.Error())
		}
		standardImports[fn] = fd
	}
}
