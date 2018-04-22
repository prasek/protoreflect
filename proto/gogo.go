package proto

import (
	"fmt"
	"reflect"
	"strings"

	gogo "github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

var _ Protoer = (*protoGoGo)(nil)
var _ Aliaser = (*protoGoGo)(nil)

var defaultGoGoAliases = AliasMap{
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

func NewGoGoProtoer(aliases AliasMap) Protoer {
	g := &protoGoGo{
		aliases: make(AliasMap),
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

type protoGoGo struct {
	aliases AliasMap
}

func (p *protoGoGo) Aliases() AliasMap {
	return p.aliases
}

func (p *protoGoGo) MessageName(m Message) string {
	return gogo.MessageName(m)
}

func (p *protoGoGo) MessageType(name string) reflect.Type {
	return gogo.MessageType(name)
}

func (p *protoGoGo) FileDescriptor(file string) []byte {
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

func (p *protoGoGo) Unmarshal(b []byte, m Message) error {
	return gogo.Unmarshal(b, m)
}

func (p *protoGoGo) Marshal(m Message) ([]byte, error) {
	return gogo.Marshal(m)
}

type extendableProto interface {
	gogo.Message
	ExtensionRangeArray() []gogo.ExtensionRange
}

func (g *protoGoGo) GetExtension(pb Message, extField int32) (interface{}, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return nil, fmt.Errorf("no pb message")
	}

	pb, err := g.ensureNativeMessage(pb)
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

func (g *protoGoGo) ensureNativeMessage(pb Message) (Message, error) {
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
	case "ServiceOptions":
		pb = &dpb.ServiceOptions{}
		err := g.Unmarshal(b, pb)
		if err != nil {
			return nil, err
		}
	case "MethodOptions":
		pb = &dpb.MethodOptions{}
		err := g.Unmarshal(b, pb)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("not gogo extendableProto")
	}
	return pb, nil
}
