package gogo

import (
	"fmt"
	"reflect"
	"strings"

	gogo "github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/jhump/protoreflect/proto"
)

var defaultGoGoAliases = map[string]string{
	// File mappings for gogo/protobuf:
	// version: 1ef32a8b9fc3f8ec940126907cedb5998f6318e4
	// as populated by proto.RegisterFile() from the pb.go files
	// in gogo/protobuf packages: types, descriptor, and plugin.

	"google/protobuf/descriptor.proto":      "descriptor.proto",
	"google/protobuf/compiler/plugin.proto": "plugin.proto",

	/*
			// These mappings are not needed for gogo/protobuf
		  // version: 1ef32a8b9fc3f8ec940126907cedb5998f6318e4
			// which is a little after v1.0.0 which didn't name these
			// correctly in the pb.go proto.RegisterFile() but it's
			// been fixed with the commit above.

			"google/protobuf/any.proto":             "any.proto",
			"google/protobuf/api.proto":             "api.proto",
			"google/protobuf/duration.proto":        "duration.proto",
			"google/protobuf/empty.proto":           "empty.proto",
			"google/protobuf/empty.proto":           "field_mask.proto",
			"google/protobuf/struct.proto":          "source_context.proto",
			"google/protobuf/struct.proto":          "struct.proto",
			"google/protobuf/timestamp.proto":       "timestamp.proto",
			"google/protobuf/wrappers.proto":        "type.proto",
			"google/protobuf/wrappers.proto":        "wrappers.proto",
	*/
}

func NewProtoer(aliases map[string]string) proto.Protoer {
	g := &protoer{
		aliases: make(map[string]string),
	}

	if aliases == nil {
		aliases = defaultGoGoAliases
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
	return gogo.MessageName(m)
}

func (p *protoer) MessageType(name string) reflect.Type {
	return gogo.MessageType(name)
}

func (p *protoer) FileDescriptor(file string) []byte {
	fdb := gogo.FileDescriptor(file)
	if fdb == nil {
		var ok bool
		alias, ok := p.aliases[file]
		if ok {
			if fdb = gogo.FileDescriptor(alias); fdb == nil {
				return nil
			}
		} else {
			return nil
		}
	}
	return fdb
}

func (p *protoer) Unmarshal(b []byte, m proto.Message) error {
	return gogo.Unmarshal(b, m)
}

func (p *protoer) Marshal(m proto.Message) ([]byte, error) {
	return gogo.Marshal(m)
}

type extendableProto interface {
	gogo.Message
	ExtensionRangeArray() []gogo.ExtensionRange
}

func (g *protoer) GetExtension(pb proto.Message, extField int32) (interface{}, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return nil, fmt.Errorf("no pb message")
	}

	pb, err := g.EnsureNativeMessage(pb)
	if err != nil {
		return nil, err
	}

	m := gogo.RegisteredExtensions(pb)
	if m == nil {
		return nil, fmt.Errorf("no registered gogo extensions")
	}
	ext, ok := m[extField]
	if !ok {
		return nil, fmt.Errorf("Field %d not found.", extField)
	}

	return gogo.GetExtension(pb, ext)
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
		return nil, fmt.Errorf("not gogo extendableProto")
	}
	err = g.Unmarshal(b, pb)
	if err != nil {
		return nil, err
	}
	return pb, nil
}
