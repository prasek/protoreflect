// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desc_test1.proto

package testprotos // import "github.com/jhump/protoreflect/internal/testprotos"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Comment for NestedEnum
type TestMessage_NestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_VALUE1 TestMessage_NestedEnum = 1
	// Comment for VALUE2
	TestMessage_VALUE2 TestMessage_NestedEnum = 2
)

var TestMessage_NestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedEnum) Enum() *TestMessage_NestedEnum {
	p := new(TestMessage_NestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedEnum_name, int32(x))
}
func (x *TestMessage_NestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedEnum_value, data, "TestMessage_NestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedEnum(value)
	return nil
}
func (TestMessage_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0, 0}
}

// Comment for DeeplyNestedEnum
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 1
	// Comment for VALUE2
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE2 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 2
)

var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) Enum() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	p := new(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, int32(x))
}
func (x *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value, data, "TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum(value)
	return nil
}
func (TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0, 0, 0, 0, 0}
}

// Comment for TestMessage
type TestMessage struct {
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,1,opt,name=nm" json:"nm,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,2,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,3,opt,name=yanm" json:"yanm,omitempty"`
	// Comment for ne
	Ne                   []TestMessage_NestedEnum `protobuf:"varint,4,rep,name=ne,enum=testprotos.TestMessage_NestedEnum" json:"ne,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0}
}
func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (dst *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(dst, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

func (m *TestMessage) GetNe() []TestMessage_NestedEnum {
	if m != nil {
		return m.Ne
	}
	return nil
}

// Comment for NestedMessage
type TestMessage_NestedMessage struct {
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,1,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm                 *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,2,opt,name=yanm" json:"yanm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                `json:"-"`
	XXX_unrecognized     []byte                                                                  `json:"-"`
	XXX_sizecache        int32                                                                   `json:"-"`
}

func (m *TestMessage_NestedMessage) Reset()         { *m = TestMessage_NestedMessage{} }
func (m *TestMessage_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage_NestedMessage) ProtoMessage()    {}
func (*TestMessage_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0, 0}
}
func (m *TestMessage_NestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage_NestedMessage.Unmarshal(m, b)
}
func (m *TestMessage_NestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage_NestedMessage.Marshal(b, m, deterministic)
}
func (dst *TestMessage_NestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage_NestedMessage.Merge(dst, src)
}
func (m *TestMessage_NestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage_NestedMessage.Size(m)
}
func (m *TestMessage_NestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage_NestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage_NestedMessage proto.InternalMessageInfo

func (m *TestMessage_NestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

// Comment for AnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage struct {
	// Comment for yanm
	Yanm                 []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,1,rep,name=yanm" json:"yanm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                  `json:"-"`
	XXX_unrecognized     []byte                                                                    `json:"-"`
	XXX_sizecache        int32                                                                     `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0, 0, 0}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage.Unmarshal(m, b)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage.Marshal(b, m, deterministic)
}
func (dst *TestMessage_NestedMessage_AnotherNestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage.Merge(dst, src)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage.Size(m)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage proto.InternalMessageInfo

func (m *TestMessage_NestedMessage_AnotherNestedMessage) GetYanm() []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

var E_TestMessage_NestedMessage_AnotherNestedMessage_Flags = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: ([]bool)(nil),
	Field:         200,
	Name:          "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.flags",
	Tag:           "varint,200,rep,packed,name=flags",
	Filename:      "desc_test1.proto",
}

// Comment for YetAnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage struct {
	// Comment for foo
	Foo *string `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	// Comment for bar
	Bar *int32 `protobuf:"varint,2,opt,name=bar" json:"bar,omitempty"`
	// Comment for baz
	Baz []byte `protobuf:"bytes,3,opt,name=baz" json:"baz,omitempty"`
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,4,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,5,opt,name=anm" json:"anm,omitempty"`
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,6,opt,name=nm" json:"nm,omitempty"`
	// Comment for tm
	Tm                   *TestMessage `protobuf:"bytes,7,opt,name=tm" json:"tm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{0, 0, 0, 0}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage.Unmarshal(m, b)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage.Marshal(b, m, deterministic)
}
func (dst *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage.Merge(dst, src)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage.Size(m)
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage proto.InternalMessageInfo

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetFoo() string {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return ""
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBar() int32 {
	if m != nil && m.Bar != nil {
		return *m.Bar
	}
	return 0
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBaz() []byte {
	if m != nil {
		return m.Baz
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetTm() *TestMessage {
	if m != nil {
		return m.Tm
	}
	return nil
}

// Comment for AnotherTestMessage
type AnotherTestMessage struct {
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,1,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for map_field1
	MapField1 map[int32]string `protobuf:"bytes,2,rep,name=map_field1,json=mapField1" json:"map_field1,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Comment for map_field2
	MapField2 map[int64]float32 `protobuf:"bytes,3,rep,name=map_field2,json=mapField2" json:"map_field2,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	// Comment for map_field3
	MapField3 map[uint32]bool `protobuf:"bytes,4,rep,name=map_field3,json=mapField3" json:"map_field3,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Comment for map_field4
	MapField4 map[string]*AnotherTestMessage `protobuf:"bytes,5,rep,name=map_field4,json=mapField4" json:"map_field4,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Rocknroll *AnotherTestMessage_RockNRoll  `protobuf:"group,6,opt,name=RockNRoll,json=rocknroll" json:"rocknroll,omitempty"`
	// Comment for atmoo
	//
	// Types that are valid to be assigned to Atmoo:
	//	*AnotherTestMessage_Str
	//	*AnotherTestMessage_Int
	Atmoo                        isAnotherTestMessage_Atmoo `protobuf_oneof:"atmoo"`
	XXX_NoUnkeyedLiteral         struct{}                   `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *AnotherTestMessage) Reset()         { *m = AnotherTestMessage{} }
func (m *AnotherTestMessage) String() string { return proto.CompactTextString(m) }
func (*AnotherTestMessage) ProtoMessage()    {}
func (*AnotherTestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{1}
}

var extRange_AnotherTestMessage = []proto.ExtensionRange{
	{Start: 100, End: 200},
}

func (*AnotherTestMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_AnotherTestMessage
}
func (m *AnotherTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnotherTestMessage.Unmarshal(m, b)
}
func (m *AnotherTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnotherTestMessage.Marshal(b, m, deterministic)
}
func (dst *AnotherTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnotherTestMessage.Merge(dst, src)
}
func (m *AnotherTestMessage) XXX_Size() int {
	return xxx_messageInfo_AnotherTestMessage.Size(m)
}
func (m *AnotherTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AnotherTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AnotherTestMessage proto.InternalMessageInfo

type isAnotherTestMessage_Atmoo interface {
	isAnotherTestMessage_Atmoo()
}

type AnotherTestMessage_Str struct {
	Str string `protobuf:"bytes,7,opt,name=str,oneof"`
}
type AnotherTestMessage_Int struct {
	Int int64 `protobuf:"varint,8,opt,name=int,oneof"`
}

func (*AnotherTestMessage_Str) isAnotherTestMessage_Atmoo() {}
func (*AnotherTestMessage_Int) isAnotherTestMessage_Atmoo() {}

func (m *AnotherTestMessage) GetAtmoo() isAnotherTestMessage_Atmoo {
	if m != nil {
		return m.Atmoo
	}
	return nil
}

func (m *AnotherTestMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *AnotherTestMessage) GetMapField1() map[int32]string {
	if m != nil {
		return m.MapField1
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField2() map[int64]float32 {
	if m != nil {
		return m.MapField2
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField3() map[uint32]bool {
	if m != nil {
		return m.MapField3
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField4() map[string]*AnotherTestMessage {
	if m != nil {
		return m.MapField4
	}
	return nil
}

func (m *AnotherTestMessage) GetRocknroll() *AnotherTestMessage_RockNRoll {
	if m != nil {
		return m.Rocknroll
	}
	return nil
}

func (m *AnotherTestMessage) GetStr() string {
	if x, ok := m.GetAtmoo().(*AnotherTestMessage_Str); ok {
		return x.Str
	}
	return ""
}

func (m *AnotherTestMessage) GetInt() int64 {
	if x, ok := m.GetAtmoo().(*AnotherTestMessage_Int); ok {
		return x.Int
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnotherTestMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnotherTestMessage_OneofMarshaler, _AnotherTestMessage_OneofUnmarshaler, _AnotherTestMessage_OneofSizer, []interface{}{
		(*AnotherTestMessage_Str)(nil),
		(*AnotherTestMessage_Int)(nil),
	}
}

func _AnotherTestMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnotherTestMessage)
	// atmoo
	switch x := m.Atmoo.(type) {
	case *AnotherTestMessage_Str:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Str)
	case *AnotherTestMessage_Int:
		_ = b.EncodeVarint(8<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Int))
	case nil:
	default:
		return fmt.Errorf("AnotherTestMessage.Atmoo has unexpected type %T", x)
	}
	return nil
}

func _AnotherTestMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnotherTestMessage)
	switch tag {
	case 7: // atmoo.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Atmoo = &AnotherTestMessage_Str{x}
		return true, err
	case 8: // atmoo.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Atmoo = &AnotherTestMessage_Int{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _AnotherTestMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnotherTestMessage)
	// atmoo
	switch x := m.Atmoo.(type) {
	case *AnotherTestMessage_Str:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *AnotherTestMessage_Int:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Comment for RockNRoll
type AnotherTestMessage_RockNRoll struct {
	// Comment for beatles
	Beatles *string `protobuf:"bytes,1,opt,name=beatles" json:"beatles,omitempty"`
	// Comment for stones
	Stones *string `protobuf:"bytes,2,opt,name=stones" json:"stones,omitempty"`
	// Comment for doors
	Doors                *string  `protobuf:"bytes,3,opt,name=doors" json:"doors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnotherTestMessage_RockNRoll) Reset()         { *m = AnotherTestMessage_RockNRoll{} }
func (m *AnotherTestMessage_RockNRoll) String() string { return proto.CompactTextString(m) }
func (*AnotherTestMessage_RockNRoll) ProtoMessage()    {}
func (*AnotherTestMessage_RockNRoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test1_976c7515ea918599, []int{1, 4}
}
func (m *AnotherTestMessage_RockNRoll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnotherTestMessage_RockNRoll.Unmarshal(m, b)
}
func (m *AnotherTestMessage_RockNRoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnotherTestMessage_RockNRoll.Marshal(b, m, deterministic)
}
func (dst *AnotherTestMessage_RockNRoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnotherTestMessage_RockNRoll.Merge(dst, src)
}
func (m *AnotherTestMessage_RockNRoll) XXX_Size() int {
	return xxx_messageInfo_AnotherTestMessage_RockNRoll.Size(m)
}
func (m *AnotherTestMessage_RockNRoll) XXX_DiscardUnknown() {
	xxx_messageInfo_AnotherTestMessage_RockNRoll.DiscardUnknown(m)
}

var xxx_messageInfo_AnotherTestMessage_RockNRoll proto.InternalMessageInfo

func (m *AnotherTestMessage_RockNRoll) GetBeatles() string {
	if m != nil && m.Beatles != nil {
		return *m.Beatles
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetStones() string {
	if m != nil && m.Stones != nil {
		return *m.Stones
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetDoors() string {
	if m != nil && m.Doors != nil {
		return *m.Doors
	}
	return ""
}

var E_Xtm = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*TestMessage)(nil),
	Field:         100,
	Name:          "testprotos.xtm",
	Tag:           "bytes,100,opt,name=xtm",
	Filename:      "desc_test1.proto",
}

var E_Xs = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         101,
	Name:          "testprotos.xs",
	Tag:           "bytes,101,opt,name=xs",
	Filename:      "desc_test1.proto",
}

var E_Xi = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         102,
	Name:          "testprotos.xi",
	Tag:           "varint,102,opt,name=xi",
	Filename:      "desc_test1.proto",
}

var E_Xui = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         103,
	Name:          "testprotos.xui",
	Tag:           "varint,103,opt,name=xui",
	Filename:      "desc_test1.proto",
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "testprotos.TestMessage")
	proto.RegisterType((*TestMessage_NestedMessage)(nil), "testprotos.TestMessage.NestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")
	proto.RegisterType((*AnotherTestMessage)(nil), "testprotos.AnotherTestMessage")
	proto.RegisterMapType((map[int32]string)(nil), "testprotos.AnotherTestMessage.MapField1Entry")
	proto.RegisterMapType((map[int64]float32)(nil), "testprotos.AnotherTestMessage.MapField2Entry")
	proto.RegisterMapType((map[uint32]bool)(nil), "testprotos.AnotherTestMessage.MapField3Entry")
	proto.RegisterMapType((map[string]*AnotherTestMessage)(nil), "testprotos.AnotherTestMessage.MapField4Entry")
	proto.RegisterType((*AnotherTestMessage_RockNRoll)(nil), "testprotos.AnotherTestMessage.RockNRoll")
	proto.RegisterEnum("testprotos.TestMessage_NestedEnum", TestMessage_NestedEnum_name, TestMessage_NestedEnum_value)
	proto.RegisterEnum("testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum", TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value)
	proto.RegisterExtension(E_TestMessage_NestedMessage_AnotherNestedMessage_Flags)
	proto.RegisterExtension(E_Xtm)
	proto.RegisterExtension(E_Xs)
	proto.RegisterExtension(E_Xi)
	proto.RegisterExtension(E_Xui)
}

func init() { proto.RegisterFile("desc_test1.proto", fileDescriptor_desc_test1_976c7515ea918599) }

var fileDescriptor_desc_test1_976c7515ea918599 = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcd, 0x6a, 0xdb, 0x4a,
	0x14, 0xc7, 0xaf, 0x46, 0x96, 0x6d, 0x9d, 0xdc, 0x04, 0x31, 0x84, 0x1b, 0x91, 0xc5, 0xc5, 0x98,
	0x7b, 0xa9, 0x09, 0x54, 0xae, 0x65, 0x17, 0x5a, 0xd3, 0x4d, 0x42, 0x13, 0x5a, 0x48, 0xb2, 0x98,
	0x7e, 0x40, 0x4b, 0x21, 0x28, 0xd6, 0xd8, 0x51, 0x23, 0xcd, 0x18, 0xcd, 0xb8, 0xd8, 0x79, 0x9a,
	0x2e, 0xfa, 0x00, 0x7d, 0x8c, 0xf6, 0x1d, 0xfa, 0x26, 0xdd, 0x94, 0x19, 0xdb, 0xb1, 0xfc, 0x91,
	0x28, 0x29, 0x49, 0x77, 0x73, 0x46, 0xe7, 0xff, 0x9b, 0xf3, 0xa5, 0x03, 0x4e, 0x48, 0x45, 0xe7,
	0x44, 0x52, 0x21, 0x1b, 0x5e, 0x3f, 0xe5, 0x92, 0x63, 0x50, 0x86, 0x3e, 0x8a, 0xea, 0xcf, 0x32,
	0xac, 0xbd, 0xa6, 0x42, 0x1e, 0x51, 0x21, 0x82, 0x1e, 0xc5, 0x8f, 0x01, 0xb1, 0xc4, 0x35, 0x2a,
	0x46, 0x6d, 0xcd, 0xff, 0xdf, 0x9b, 0x39, 0x7a, 0x19, 0x27, 0xef, 0x98, 0x0a, 0x49, 0xc3, 0x89,
	0x45, 0x10, 0x4b, 0xf0, 0x21, 0x98, 0x01, 0x4b, 0x5c, 0xa4, 0x75, 0xed, 0x1b, 0xe9, 0xbc, 0x5d,
	0xc6, 0xe5, 0x19, 0x4d, 0xe7, 0x61, 0x0a, 0x83, 0xbb, 0x50, 0x18, 0x29, 0x9c, 0xa9, 0x71, 0xe4,
	0xf7, 0x71, 0xde, 0x3b, 0x2a, 0x57, 0x3e, 0xa3, 0xf9, 0xd8, 0x07, 0xc4, 0xa8, 0x5b, 0xa8, 0x98,
	0xb5, 0x0d, 0xbf, 0x7a, 0xfd, 0x2b, 0xfb, 0x6c, 0x90, 0x10, 0xc4, 0xe8, 0xf6, 0x97, 0x22, 0xac,
	0xcf, 0xb1, 0xa6, 0xb9, 0x1b, 0x77, 0x9b, 0x3b, 0xba, 0xdf, 0xdc, 0xb7, 0x7f, 0x14, 0x60, 0x73,
	0xd5, 0xe7, 0xcb, 0x00, 0x8c, 0x8a, 0x79, 0xaf, 0x01, 0x7c, 0x36, 0x61, 0xeb, 0x0a, 0x0f, 0xec,
	0x80, 0xd9, 0xe5, 0x5c, 0x97, 0xd4, 0x26, 0xea, 0xa8, 0x6e, 0x4e, 0x83, 0x54, 0x57, 0xc5, 0x22,
	0xea, 0x38, 0xbe, 0xb9, 0xd0, 0x33, 0xf2, 0xb7, 0xba, 0xb9, 0xc0, 0x03, 0x30, 0x43, 0xdd, 0x4f,
	0xa3, 0xb6, 0xe1, 0x77, 0xee, 0x3e, 0x70, 0xef, 0x39, 0xa5, 0xfd, 0x78, 0x94, 0x19, 0x08, 0xf5,
	0xde, 0xb4, 0xff, 0xd6, 0xdd, 0xf4, 0x7f, 0xfc, 0x03, 0x16, 0x6f, 0xfb, 0x03, 0x3e, 0x00, 0x24,
	0x13, 0xb7, 0xa4, 0x65, 0x5b, 0x57, 0xc8, 0x08, 0x92, 0x49, 0x75, 0x07, 0x9c, 0xc5, 0x34, 0x30,
	0x40, 0xf1, 0xed, 0xee, 0xe1, 0x9b, 0xfd, 0x86, 0x63, 0x5c, 0x9e, 0x7d, 0x07, 0xf9, 0x4f, 0xc1,
	0xea, 0xc6, 0x41, 0x4f, 0xe0, 0x7f, 0xb3, 0xc4, 0x49, 0xec, 0x19, 0xb0, 0xfb, 0x4d, 0x0d, 0x4b,
	0x79, 0x0f, 0x39, 0x06, 0x19, 0x2b, 0xaa, 0xff, 0x01, 0xe4, 0x3f, 0x50, 0xfd, 0x5a, 0x02, 0xbc,
	0x8c, 0x9b, 0x36, 0xd2, 0xf8, 0xe3, 0x8d, 0x84, 0x24, 0xe8, 0x9f, 0x74, 0x23, 0x1a, 0x87, 0x0d,
	0x17, 0xe9, 0xf9, 0x7f, 0x78, 0x7d, 0xe6, 0xde, 0x51, 0xd0, 0x3f, 0xd0, 0xfe, 0xfb, 0x4c, 0xa6,
	0x23, 0x62, 0x27, 0x53, 0x7b, 0x8e, 0xe6, 0xbb, 0xe6, 0xad, 0x68, 0xfe, 0x02, 0xcd, 0x9f, 0xa3,
	0x35, 0xf5, 0xca, 0xba, 0x39, 0xad, 0xb9, 0x40, 0x6b, 0xce, 0xd1, 0x5a, 0xae, 0x75, 0x2b, 0x5a,
	0x6b, 0x81, 0xd6, 0xc2, 0x07, 0x60, 0xa7, 0xbc, 0x73, 0xce, 0x52, 0x1e, 0xc7, 0x7a, 0x72, 0xc1,
	0xaf, 0xe5, 0xc0, 0x08, 0xef, 0x9c, 0x1f, 0x13, 0x1e, 0xc7, 0x64, 0x26, 0xc5, 0x18, 0x4c, 0x21,
	0x53, 0x3d, 0xc4, 0xf6, 0x8b, 0xbf, 0x88, 0x32, 0xd4, 0x5d, 0xc4, 0xa4, 0x5b, 0xae, 0x18, 0x35,
	0x53, 0xdd, 0x45, 0x4c, 0x6e, 0x3f, 0x83, 0x8d, 0xf9, 0xb2, 0xab, 0x5d, 0x70, 0x4e, 0x47, 0x7a,
	0x60, 0x2c, 0xa2, 0x8e, 0x78, 0x13, 0xac, 0x4f, 0x41, 0x3c, 0xa0, 0x7a, 0x63, 0xd8, 0x64, 0x6c,
	0xb4, 0xd1, 0x13, 0x23, 0xab, 0xf6, 0x97, 0xd4, 0xe6, 0x0a, 0x35, 0xba, 0x42, 0xdd, 0x5c, 0x52,
	0xaf, 0xaf, 0x50, 0x97, 0xb3, 0xea, 0x0f, 0x33, 0x75, 0x6b, 0x49, 0x6d, 0x8f, 0xd5, 0xad, 0xac,
	0x7a, 0xcd, 0xcf, 0xf9, 0xf5, 0xb2, 0xf4, 0x57, 0x60, 0x5f, 0xd6, 0x15, 0xbb, 0x50, 0x3a, 0xa5,
	0x81, 0x8c, 0xa9, 0x98, 0xc0, 0xa7, 0x26, 0xfe, 0x07, 0x8a, 0x42, 0x72, 0x46, 0xc5, 0xa4, 0x36,
	0x13, 0x4b, 0x85, 0x1d, 0x72, 0x9e, 0x0a, 0xbd, 0x52, 0x6d, 0x32, 0x36, 0x76, 0xac, 0x72, 0xe8,
	0x7c, 0x37, 0xf6, 0x4a, 0x60, 0x05, 0x32, 0xe1, 0xbc, 0xfd, 0x12, 0xcc, 0xa1, 0x4c, 0x72, 0x37,
	0x42, 0x78, 0xfd, 0x26, 0x52, 0x8c, 0xb6, 0x07, 0x68, 0x98, 0xbf, 0x5b, 0xa8, 0x8e, 0x06, 0x0d,
	0x85, 0xf6, 0x8f, 0x72, 0xfd, 0xbb, 0x7a, 0x08, 0xd0, 0x30, 0x6a, 0x3f, 0x02, 0x73, 0x38, 0xc8,
	0x17, 0xf4, 0x2a, 0x46, 0xad, 0x40, 0x94, 0xeb, 0x5e, 0xf3, 0x7d, 0xa3, 0x17, 0xc9, 0xb3, 0xc1,
	0xa9, 0xd7, 0xe1, 0x49, 0xfd, 0xe3, 0xd9, 0x20, 0xe9, 0xd7, 0xb5, 0x30, 0xa5, 0xdd, 0x98, 0x76,
	0x64, 0x3d, 0x62, 0x92, 0xa6, 0x2c, 0x88, 0xeb, 0x33, 0xe4, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9b, 0x1d, 0x1b, 0x3b, 0x61, 0x09, 0x00, 0x00,
}
