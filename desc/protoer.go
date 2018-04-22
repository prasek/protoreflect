package desc

import (
	"fmt"
	"reflect"
	"strings"

	gogo "github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

type ExtensionDesc struct {
	ExtendedType  Message
	ExtensionType interface{}
	Field         int32
	Name          string
	Tag           string
	Filename      string
}

type Protoer interface {
	MessageType(name string) reflect.Type
	MessageName(m Message) string
	FileDescriptor(name string) []byte
	Unmarshal(b []byte, m Message) error
	Marshal(m Message) ([]byte, error)
	RegisteredExtensions(m Message) map[int32]*ExtensionDesc
	GetExtension(pb Message, extField int32) (interface{}, error)
}

var protoer Protoer = &GoGoProto{}

func SetProtoer(p Protoer) {
	protoer = p
}

type GoGoProto struct {
}

func (p *GoGoProto) MessageName(m Message) string {
	return gogo.MessageName(m)
}

func (p *GoGoProto) MessageType(name string) reflect.Type {
	return gogo.MessageType(name)
}

func (p *GoGoProto) FileDescriptor(name string) []byte {
	return gogo.FileDescriptor(name)
}

func (p *GoGoProto) Unmarshal(b []byte, m Message) error {
	return gogo.Unmarshal(b, m)
}

func (p *GoGoProto) Marshal(m Message) ([]byte, error) {
	return gogo.Marshal(m)
}

type extendableProto interface {
	gogo.Message
	ExtensionRangeArray() []gogo.ExtensionRange
}

func (g *GoGoProto) GetExtension(pb Message, extField int32) (interface{}, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return nil, fmt.Errorf("no pb message")
	}
	if _, ok := pb.(extendableProto); !ok {
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

func (p *GoGoProto) RegisteredExtensions(m Message) map[int32]*ExtensionDesc {
	res := make(map[int32]*ExtensionDesc)
	extensions := gogo.RegisteredExtensions(m)
	for k, v := range extensions {
		res[k] = &ExtensionDesc{
			ExtendedType:  v.ExtendedType,
			ExtensionType: v.ExtensionType,
			Field:         v.Field,
			Name:          v.Name,
			Tag:           v.Tag,
			Filename:      v.Filename,
		}
	}
	return res
}

func GetExtension(pb Message, field int32) (interface{}, error) {
	return protoer.GetExtension(pb, field)
}

func GetBoolExtension(pb Message, field int32, def bool) bool {
	v, err := protoer.GetExtension(pb, field)
	if err != nil {
		return def
	}
	if v == nil {
		return def
	}
	return *v.(*bool)
}
