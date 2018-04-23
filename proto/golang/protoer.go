package golang

import (
	"fmt"
	"reflect"
	"strings"

	golang "github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/golang/protobuf/protoc-gen-go/plugin"
	_ "github.com/golang/protobuf/ptypes/empty"
	"github.com/jhump/protoreflect/proto"
	_ "google.golang.org/genproto/protobuf/api"
	_ "google.golang.org/genproto/protobuf/field_mask"
	_ "google.golang.org/genproto/protobuf/ptype"
	_ "google.golang.org/genproto/protobuf/source_context"
)

var defaultGolangAliases = map[string]string{

	// Files for the github.com/golang/protobuf/ptypes package at one point were
	// registered using the path where the proto files are mirrored in GOPATH,
	// inside the golang/protobuf repo.
	// (Fixed as of https://github.com/golang/protobuf/pull/412)
	"google/protobuf/any.proto":       "github.com/golang/protobuf/ptypes/any/any.proto",
	"google/protobuf/duration.proto":  "github.com/golang/protobuf/ptypes/duration/duration.proto",
	"google/protobuf/empty.proto":     "github.com/golang/protobuf/ptypes/empty/empty.proto",
	"google/protobuf/struct.proto":    "github.com/golang/protobuf/ptypes/struct/struct.proto",
	"google/protobuf/timestamp.proto": "github.com/golang/protobuf/ptypes/timestamp/timestamp.proto",
	"google/protobuf/wrappers.proto":  "github.com/golang/protobuf/ptypes/wrappers/wrappers.proto",
	// Files for the google.golang.org/genproto/protobuf package at one point
	// were registered with an anomalous "src/" prefix.
	// (Fixed as of https://github.com/google/go-genproto/pull/31)
	"google/protobuf/api.proto":            "src/google/protobuf/api.proto",
	"google/protobuf/field_mask.proto":     "src/google/protobuf/field_mask.proto",
	"google/protobuf/source_context.proto": "src/google/protobuf/source_context.proto",
	"google/protobuf/type.proto":           "src/google/protobuf/type.proto",

	// Other standard files (descriptor.proto and compiler/plugin.proto) are
	// registered correctly, so we don't need rules for them here.
}

func NewProtoer(aliases map[string]string) proto.Protoer {
	g := &protoer{
		aliases: make(map[string]string),
	}

	if aliases == nil {
		aliases = defaultGolangAliases
	}

	// bidirectional mapping
	for k, v := range aliases {
		g.aliases[k] = v
		g.aliases[v] = k
	}
	return g
}

type protoer struct {
	aliases map[string]string
}

func (p *protoer) Aliases() map[string]string {
	return p.aliases
}

func (p *protoer) MessageName(m proto.Message) string {
	return golang.MessageName(m)
}

func (p *protoer) MessageType(name string) reflect.Type {
	return golang.MessageType(name)
}

func (p *protoer) FileDescriptor(file string) []byte {
	fdb := golang.FileDescriptor(file)
	if fdb == nil {
		var ok bool
		alias, ok := p.aliases[file]
		if ok {
			if fdb = golang.FileDescriptor(alias); fdb == nil {
				return nil
			}
		} else {
			return nil
		}
	}
	return fdb
}

func (p *protoer) Unmarshal(b []byte, m proto.Message) error {
	return golang.Unmarshal(b, m)
}

func (p *protoer) Marshal(m proto.Message) ([]byte, error) {
	return golang.Marshal(m)
}

type extendableProto interface {
	golang.Message
	ExtensionRangeArray() []golang.ExtensionRange
}

func (g *protoer) GetExtension(pb proto.Message, extField int32) (interface{}, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return nil, fmt.Errorf("no pb message")
	}

	pb, err := g.EnsureNativeMessage(pb)
	if err != nil {
		return nil, err
	}

	m := golang.RegisteredExtensions(pb)
	if m == nil {
		return nil, fmt.Errorf("no registered proto extensions")
	}
	ext, ok := m[extField]
	if !ok {
		return nil, fmt.Errorf("Field %d not found.", extField)
	}

	return golang.GetExtension(pb, ext)
}

func (g *protoer) EnsureNativeMessage(pb proto.Message) (proto.Message, error) {
	if _, ok := pb.(extendableProto); ok {
		return pb, nil
	}

	//if using different fork of protobuf, then convert it
	pt := reflect.TypeOf(pb)
	typeName := strings.Split(pt.String(), ".")[1]
	b, err := g.Marshal(pb)
	if err != nil {
		return nil, err
	}
	switch typeName {
	case "FileDescriptorSet":
		pb = &dpb.FileDescriptorSet{}
	case "FileDescriptorProto":
		pb = &dpb.FileDescriptorProto{}
	case "DescriptorProto":
		pb = &dpb.DescriptorProto{}
	case "ExtensionRangeOptions":
		pb = &dpb.ExtensionRangeOptions{}
	case "FieldDescriptorProto":
		pb = &dpb.FieldDescriptorProto{}
	case "OneofDescriptorProto":
		pb = &dpb.OneofDescriptorProto{}
	case "EnumDescriptorProto":
		pb = &dpb.EnumDescriptorProto{}
	case "EnumValueDescriptorProto":
		pb = &dpb.EnumValueDescriptorProto{}
	case "ServiceDescriptorProto":
		pb = &dpb.ServiceDescriptorProto{}
	case "MethodDescriptorProto":
		pb = &dpb.MethodDescriptorProto{}
	case "FileOptions":
		pb = &dpb.FileOptions{}
	case "MessageOptions":
		pb = &dpb.MessageOptions{}
	case "FieldOptions":
		pb = &dpb.FieldOptions{}
	case "OneofOptions":
		pb = &dpb.OneofOptions{}
	case "EnumOptions":
		pb = &dpb.EnumOptions{}
	case "EnumValueOptions":
		pb = &dpb.EnumValueOptions{}
	case "ServiceOptions":
		pb = &dpb.ServiceOptions{}
	case "MethodOptions":
		pb = &dpb.MethodOptions{}
	case "UninterpretedOption":
		pb = &dpb.UninterpretedOption{}
	case "SourceCodeInfo":
		pb = &dpb.SourceCodeInfo{}
	case "GeneratedCodeInfo":
		pb = &dpb.GeneratedCodeInfo{}
	default:
		return nil, fmt.Errorf("not proto extendableProto")
	}
	err = g.Unmarshal(b, pb)
	if err != nil {
		return nil, err
	}
	return pb, nil
}
