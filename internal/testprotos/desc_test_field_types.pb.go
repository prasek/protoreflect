// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desc_test_field_types.proto

package testprotos

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TestEnum int32

const (
	TestEnum_INVALID TestEnum = 0
	TestEnum_FIRST   TestEnum = 1
	TestEnum_SECOND  TestEnum = 2
	TestEnum_THIRD   TestEnum = 3
)

var TestEnum_name = map[int32]string{
	0: "INVALID",
	1: "FIRST",
	2: "SECOND",
	3: "THIRD",
}

var TestEnum_value = map[string]int32{
	"INVALID": 0,
	"FIRST":   1,
	"SECOND":  2,
	"THIRD":   3,
}

func (x TestEnum) Enum() *TestEnum {
	p := new(TestEnum)
	*p = x
	return p
}

func (x TestEnum) String() string {
	return proto.EnumName(TestEnum_name, int32(x))
}

func (x *TestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestEnum_value, data, "TestEnum")
	if err != nil {
		return err
	}
	*x = TestEnum(value)
	return nil
}

func (TestEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{0}
}

type UnaryFields struct {
	I                    *int32              `protobuf:"varint,1,opt,name=i" json:"i,omitempty"`
	J                    *int64              `protobuf:"varint,2,opt,name=j" json:"j,omitempty"`
	K                    *int32              `protobuf:"zigzag32,3,opt,name=k" json:"k,omitempty"`
	L                    *int64              `protobuf:"zigzag64,4,opt,name=l" json:"l,omitempty"`
	M                    *uint32             `protobuf:"varint,5,opt,name=m" json:"m,omitempty"`
	N                    *uint64             `protobuf:"varint,6,opt,name=n" json:"n,omitempty"`
	O                    *uint32             `protobuf:"fixed32,7,opt,name=o" json:"o,omitempty"`
	P                    *uint64             `protobuf:"fixed64,8,opt,name=p" json:"p,omitempty"`
	Q                    *int32              `protobuf:"fixed32,9,opt,name=q" json:"q,omitempty"`
	R                    *int64              `protobuf:"fixed64,10,opt,name=r" json:"r,omitempty"`
	S                    *float32            `protobuf:"fixed32,11,opt,name=s" json:"s,omitempty"`
	T                    *float64            `protobuf:"fixed64,12,opt,name=t" json:"t,omitempty"`
	U                    []byte              `protobuf:"bytes,13,opt,name=u" json:"u,omitempty"`
	V                    *string             `protobuf:"bytes,14,opt,name=v" json:"v,omitempty"`
	W                    *bool               `protobuf:"varint,15,opt,name=w" json:"w,omitempty"`
	X                    *RepeatedFields     `protobuf:"bytes,16,opt,name=x" json:"x,omitempty"`
	Groupy               *UnaryFields_GroupY `protobuf:"group,17,opt,name=GroupY,json=groupy" json:"groupy,omitempty"`
	Z                    *TestEnum           `protobuf:"varint,18,opt,name=z,enum=testprotos.TestEnum" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UnaryFields) Reset()         { *m = UnaryFields{} }
func (m *UnaryFields) String() string { return proto.CompactTextString(m) }
func (*UnaryFields) ProtoMessage()    {}
func (*UnaryFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{0}
}
func (m *UnaryFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryFields.Unmarshal(m, b)
}
func (m *UnaryFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryFields.Marshal(b, m, deterministic)
}
func (m *UnaryFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryFields.Merge(m, src)
}
func (m *UnaryFields) XXX_Size() int {
	return xxx_messageInfo_UnaryFields.Size(m)
}
func (m *UnaryFields) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryFields.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryFields proto.InternalMessageInfo

func (m *UnaryFields) GetI() int32 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func (m *UnaryFields) GetJ() int64 {
	if m != nil && m.J != nil {
		return *m.J
	}
	return 0
}

func (m *UnaryFields) GetK() int32 {
	if m != nil && m.K != nil {
		return *m.K
	}
	return 0
}

func (m *UnaryFields) GetL() int64 {
	if m != nil && m.L != nil {
		return *m.L
	}
	return 0
}

func (m *UnaryFields) GetM() uint32 {
	if m != nil && m.M != nil {
		return *m.M
	}
	return 0
}

func (m *UnaryFields) GetN() uint64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *UnaryFields) GetO() uint32 {
	if m != nil && m.O != nil {
		return *m.O
	}
	return 0
}

func (m *UnaryFields) GetP() uint64 {
	if m != nil && m.P != nil {
		return *m.P
	}
	return 0
}

func (m *UnaryFields) GetQ() int32 {
	if m != nil && m.Q != nil {
		return *m.Q
	}
	return 0
}

func (m *UnaryFields) GetR() int64 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *UnaryFields) GetS() float32 {
	if m != nil && m.S != nil {
		return *m.S
	}
	return 0
}

func (m *UnaryFields) GetT() float64 {
	if m != nil && m.T != nil {
		return *m.T
	}
	return 0
}

func (m *UnaryFields) GetU() []byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *UnaryFields) GetV() string {
	if m != nil && m.V != nil {
		return *m.V
	}
	return ""
}

func (m *UnaryFields) GetW() bool {
	if m != nil && m.W != nil {
		return *m.W
	}
	return false
}

func (m *UnaryFields) GetX() *RepeatedFields {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *UnaryFields) GetGroupy() *UnaryFields_GroupY {
	if m != nil {
		return m.Groupy
	}
	return nil
}

func (m *UnaryFields) GetZ() TestEnum {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return TestEnum_INVALID
}

type UnaryFields_GroupY struct {
	Ya                   *string  `protobuf:"bytes,171,opt,name=ya" json:"ya,omitempty"`
	Yb                   *int32   `protobuf:"varint,172,opt,name=yb" json:"yb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryFields_GroupY) Reset()         { *m = UnaryFields_GroupY{} }
func (m *UnaryFields_GroupY) String() string { return proto.CompactTextString(m) }
func (*UnaryFields_GroupY) ProtoMessage()    {}
func (*UnaryFields_GroupY) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{0, 0}
}
func (m *UnaryFields_GroupY) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryFields_GroupY.Unmarshal(m, b)
}
func (m *UnaryFields_GroupY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryFields_GroupY.Marshal(b, m, deterministic)
}
func (m *UnaryFields_GroupY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryFields_GroupY.Merge(m, src)
}
func (m *UnaryFields_GroupY) XXX_Size() int {
	return xxx_messageInfo_UnaryFields_GroupY.Size(m)
}
func (m *UnaryFields_GroupY) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryFields_GroupY.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryFields_GroupY proto.InternalMessageInfo

func (m *UnaryFields_GroupY) GetYa() string {
	if m != nil && m.Ya != nil {
		return *m.Ya
	}
	return ""
}

func (m *UnaryFields_GroupY) GetYb() int32 {
	if m != nil && m.Yb != nil {
		return *m.Yb
	}
	return 0
}

type RepeatedFields struct {
	I                    []int32                  `protobuf:"varint,1,rep,name=i" json:"i,omitempty"`
	J                    []int64                  `protobuf:"varint,2,rep,name=j" json:"j,omitempty"`
	K                    []int32                  `protobuf:"zigzag32,3,rep,name=k" json:"k,omitempty"`
	L                    []int64                  `protobuf:"zigzag64,4,rep,name=l" json:"l,omitempty"`
	M                    []uint32                 `protobuf:"varint,5,rep,name=m" json:"m,omitempty"`
	N                    []uint64                 `protobuf:"varint,6,rep,name=n" json:"n,omitempty"`
	O                    []uint32                 `protobuf:"fixed32,7,rep,name=o" json:"o,omitempty"`
	P                    []uint64                 `protobuf:"fixed64,8,rep,name=p" json:"p,omitempty"`
	Q                    []int32                  `protobuf:"fixed32,9,rep,name=q" json:"q,omitempty"`
	R                    []int64                  `protobuf:"fixed64,10,rep,name=r" json:"r,omitempty"`
	S                    []float32                `protobuf:"fixed32,11,rep,name=s" json:"s,omitempty"`
	T                    []float64                `protobuf:"fixed64,12,rep,name=t" json:"t,omitempty"`
	U                    [][]byte                 `protobuf:"bytes,13,rep,name=u" json:"u,omitempty"`
	V                    []string                 `protobuf:"bytes,14,rep,name=v" json:"v,omitempty"`
	W                    []bool                   `protobuf:"varint,15,rep,name=w" json:"w,omitempty"`
	X                    []*UnaryFields           `protobuf:"bytes,16,rep,name=x" json:"x,omitempty"`
	Groupy               []*RepeatedFields_GroupY `protobuf:"group,17,rep,name=GroupY,json=groupy" json:"groupy,omitempty"`
	Z                    []TestEnum               `protobuf:"varint,18,rep,name=z,enum=testprotos.TestEnum" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RepeatedFields) Reset()         { *m = RepeatedFields{} }
func (m *RepeatedFields) String() string { return proto.CompactTextString(m) }
func (*RepeatedFields) ProtoMessage()    {}
func (*RepeatedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{1}
}
func (m *RepeatedFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedFields.Unmarshal(m, b)
}
func (m *RepeatedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedFields.Marshal(b, m, deterministic)
}
func (m *RepeatedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedFields.Merge(m, src)
}
func (m *RepeatedFields) XXX_Size() int {
	return xxx_messageInfo_RepeatedFields.Size(m)
}
func (m *RepeatedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedFields.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedFields proto.InternalMessageInfo

func (m *RepeatedFields) GetI() []int32 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *RepeatedFields) GetJ() []int64 {
	if m != nil {
		return m.J
	}
	return nil
}

func (m *RepeatedFields) GetK() []int32 {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *RepeatedFields) GetL() []int64 {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *RepeatedFields) GetM() []uint32 {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *RepeatedFields) GetN() []uint64 {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *RepeatedFields) GetO() []uint32 {
	if m != nil {
		return m.O
	}
	return nil
}

func (m *RepeatedFields) GetP() []uint64 {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *RepeatedFields) GetQ() []int32 {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *RepeatedFields) GetR() []int64 {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *RepeatedFields) GetS() []float32 {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *RepeatedFields) GetT() []float64 {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *RepeatedFields) GetU() [][]byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *RepeatedFields) GetV() []string {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *RepeatedFields) GetW() []bool {
	if m != nil {
		return m.W
	}
	return nil
}

func (m *RepeatedFields) GetX() []*UnaryFields {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *RepeatedFields) GetGroupy() []*RepeatedFields_GroupY {
	if m != nil {
		return m.Groupy
	}
	return nil
}

func (m *RepeatedFields) GetZ() []TestEnum {
	if m != nil {
		return m.Z
	}
	return nil
}

type RepeatedFields_GroupY struct {
	Ya                   *string  `protobuf:"bytes,171,opt,name=ya" json:"ya,omitempty"`
	Yb                   *int32   `protobuf:"varint,172,opt,name=yb" json:"yb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatedFields_GroupY) Reset()         { *m = RepeatedFields_GroupY{} }
func (m *RepeatedFields_GroupY) String() string { return proto.CompactTextString(m) }
func (*RepeatedFields_GroupY) ProtoMessage()    {}
func (*RepeatedFields_GroupY) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{1, 0}
}
func (m *RepeatedFields_GroupY) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedFields_GroupY.Unmarshal(m, b)
}
func (m *RepeatedFields_GroupY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedFields_GroupY.Marshal(b, m, deterministic)
}
func (m *RepeatedFields_GroupY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedFields_GroupY.Merge(m, src)
}
func (m *RepeatedFields_GroupY) XXX_Size() int {
	return xxx_messageInfo_RepeatedFields_GroupY.Size(m)
}
func (m *RepeatedFields_GroupY) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedFields_GroupY.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedFields_GroupY proto.InternalMessageInfo

func (m *RepeatedFields_GroupY) GetYa() string {
	if m != nil && m.Ya != nil {
		return *m.Ya
	}
	return ""
}

func (m *RepeatedFields_GroupY) GetYb() int32 {
	if m != nil && m.Yb != nil {
		return *m.Yb
	}
	return 0
}

type RepeatedPackedFields struct {
	I                    []int32                        `protobuf:"varint,1,rep,packed,name=i" json:"i,omitempty"`
	J                    []int64                        `protobuf:"varint,2,rep,packed,name=j" json:"j,omitempty"`
	K                    []int32                        `protobuf:"zigzag32,3,rep,packed,name=k" json:"k,omitempty"`
	L                    []int64                        `protobuf:"zigzag64,4,rep,packed,name=l" json:"l,omitempty"`
	M                    []uint32                       `protobuf:"varint,5,rep,packed,name=m" json:"m,omitempty"`
	N                    []uint64                       `protobuf:"varint,6,rep,packed,name=n" json:"n,omitempty"`
	O                    []uint32                       `protobuf:"fixed32,7,rep,packed,name=o" json:"o,omitempty"`
	P                    []uint64                       `protobuf:"fixed64,8,rep,packed,name=p" json:"p,omitempty"`
	Q                    []int32                        `protobuf:"fixed32,9,rep,packed,name=q" json:"q,omitempty"`
	R                    []int64                        `protobuf:"fixed64,10,rep,packed,name=r" json:"r,omitempty"`
	S                    []float32                      `protobuf:"fixed32,11,rep,packed,name=s" json:"s,omitempty"`
	T                    []float64                      `protobuf:"fixed64,12,rep,packed,name=t" json:"t,omitempty"`
	U                    []bool                         `protobuf:"varint,13,rep,packed,name=u" json:"u,omitempty"`
	Groupy               []*RepeatedPackedFields_GroupY `protobuf:"group,14,rep,name=GroupY,json=groupy" json:"groupy,omitempty"`
	V                    []TestEnum                     `protobuf:"varint,15,rep,packed,name=v,enum=testprotos.TestEnum" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RepeatedPackedFields) Reset()         { *m = RepeatedPackedFields{} }
func (m *RepeatedPackedFields) String() string { return proto.CompactTextString(m) }
func (*RepeatedPackedFields) ProtoMessage()    {}
func (*RepeatedPackedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{2}
}
func (m *RepeatedPackedFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedPackedFields.Unmarshal(m, b)
}
func (m *RepeatedPackedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedPackedFields.Marshal(b, m, deterministic)
}
func (m *RepeatedPackedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedPackedFields.Merge(m, src)
}
func (m *RepeatedPackedFields) XXX_Size() int {
	return xxx_messageInfo_RepeatedPackedFields.Size(m)
}
func (m *RepeatedPackedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedPackedFields.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedPackedFields proto.InternalMessageInfo

func (m *RepeatedPackedFields) GetI() []int32 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *RepeatedPackedFields) GetJ() []int64 {
	if m != nil {
		return m.J
	}
	return nil
}

func (m *RepeatedPackedFields) GetK() []int32 {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *RepeatedPackedFields) GetL() []int64 {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *RepeatedPackedFields) GetM() []uint32 {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *RepeatedPackedFields) GetN() []uint64 {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *RepeatedPackedFields) GetO() []uint32 {
	if m != nil {
		return m.O
	}
	return nil
}

func (m *RepeatedPackedFields) GetP() []uint64 {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *RepeatedPackedFields) GetQ() []int32 {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *RepeatedPackedFields) GetR() []int64 {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *RepeatedPackedFields) GetS() []float32 {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *RepeatedPackedFields) GetT() []float64 {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *RepeatedPackedFields) GetU() []bool {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *RepeatedPackedFields) GetGroupy() []*RepeatedPackedFields_GroupY {
	if m != nil {
		return m.Groupy
	}
	return nil
}

func (m *RepeatedPackedFields) GetV() []TestEnum {
	if m != nil {
		return m.V
	}
	return nil
}

type RepeatedPackedFields_GroupY struct {
	Yb                   []int32  `protobuf:"varint,141,rep,packed,name=yb" json:"yb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatedPackedFields_GroupY) Reset()         { *m = RepeatedPackedFields_GroupY{} }
func (m *RepeatedPackedFields_GroupY) String() string { return proto.CompactTextString(m) }
func (*RepeatedPackedFields_GroupY) ProtoMessage()    {}
func (*RepeatedPackedFields_GroupY) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{2, 0}
}
func (m *RepeatedPackedFields_GroupY) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedPackedFields_GroupY.Unmarshal(m, b)
}
func (m *RepeatedPackedFields_GroupY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedPackedFields_GroupY.Marshal(b, m, deterministic)
}
func (m *RepeatedPackedFields_GroupY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedPackedFields_GroupY.Merge(m, src)
}
func (m *RepeatedPackedFields_GroupY) XXX_Size() int {
	return xxx_messageInfo_RepeatedPackedFields_GroupY.Size(m)
}
func (m *RepeatedPackedFields_GroupY) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedPackedFields_GroupY.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedPackedFields_GroupY proto.InternalMessageInfo

func (m *RepeatedPackedFields_GroupY) GetYb() []int32 {
	if m != nil {
		return m.Yb
	}
	return nil
}

type MapKeyFields struct {
	I                    map[int32]string  `protobuf:"bytes,1,rep,name=i" json:"i,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	J                    map[int64]string  `protobuf:"bytes,2,rep,name=j" json:"j,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	K                    map[int32]string  `protobuf:"bytes,3,rep,name=k" json:"k,omitempty" protobuf_key:"zigzag32,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	L                    map[int64]string  `protobuf:"bytes,4,rep,name=l" json:"l,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	M                    map[uint32]string `protobuf:"bytes,5,rep,name=m" json:"m,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	N                    map[uint64]string `protobuf:"bytes,6,rep,name=n" json:"n,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	O                    map[uint32]string `protobuf:"bytes,7,rep,name=o" json:"o,omitempty" protobuf_key:"fixed32,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	P                    map[uint64]string `protobuf:"bytes,8,rep,name=p" json:"p,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Q                    map[int32]string  `protobuf:"bytes,9,rep,name=q" json:"q,omitempty" protobuf_key:"fixed32,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	R                    map[int64]string  `protobuf:"bytes,10,rep,name=r" json:"r,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	S                    map[string]string `protobuf:"bytes,11,rep,name=s" json:"s,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	T                    map[bool]string   `protobuf:"bytes,12,rep,name=t" json:"t,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapKeyFields) Reset()         { *m = MapKeyFields{} }
func (m *MapKeyFields) String() string { return proto.CompactTextString(m) }
func (*MapKeyFields) ProtoMessage()    {}
func (*MapKeyFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{3}
}
func (m *MapKeyFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapKeyFields.Unmarshal(m, b)
}
func (m *MapKeyFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapKeyFields.Marshal(b, m, deterministic)
}
func (m *MapKeyFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapKeyFields.Merge(m, src)
}
func (m *MapKeyFields) XXX_Size() int {
	return xxx_messageInfo_MapKeyFields.Size(m)
}
func (m *MapKeyFields) XXX_DiscardUnknown() {
	xxx_messageInfo_MapKeyFields.DiscardUnknown(m)
}

var xxx_messageInfo_MapKeyFields proto.InternalMessageInfo

func (m *MapKeyFields) GetI() map[int32]string {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *MapKeyFields) GetJ() map[int64]string {
	if m != nil {
		return m.J
	}
	return nil
}

func (m *MapKeyFields) GetK() map[int32]string {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *MapKeyFields) GetL() map[int64]string {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *MapKeyFields) GetM() map[uint32]string {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *MapKeyFields) GetN() map[uint64]string {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *MapKeyFields) GetO() map[uint32]string {
	if m != nil {
		return m.O
	}
	return nil
}

func (m *MapKeyFields) GetP() map[uint64]string {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *MapKeyFields) GetQ() map[int32]string {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *MapKeyFields) GetR() map[int64]string {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *MapKeyFields) GetS() map[string]string {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *MapKeyFields) GetT() map[bool]string {
	if m != nil {
		return m.T
	}
	return nil
}

type MapValFields struct {
	I                    map[string]int32        `protobuf:"bytes,1,rep,name=i" json:"i,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	J                    map[string]int64        `protobuf:"bytes,2,rep,name=j" json:"j,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	K                    map[string]int32        `protobuf:"bytes,3,rep,name=k" json:"k,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"zigzag32,2,opt,name=value"`
	L                    map[string]int64        `protobuf:"bytes,4,rep,name=l" json:"l,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"zigzag64,2,opt,name=value"`
	M                    map[string]uint32       `protobuf:"bytes,5,rep,name=m" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	N                    map[string]uint64       `protobuf:"bytes,6,rep,name=n" json:"n,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	O                    map[string]uint32       `protobuf:"bytes,7,rep,name=o" json:"o,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	P                    map[string]uint64       `protobuf:"bytes,8,rep,name=p" json:"p,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	Q                    map[string]int32        `protobuf:"bytes,9,rep,name=q" json:"q,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	R                    map[string]int64        `protobuf:"bytes,10,rep,name=r" json:"r,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	S                    map[string]float32      `protobuf:"bytes,11,rep,name=s" json:"s,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	T                    map[string]float64      `protobuf:"bytes,12,rep,name=t" json:"t,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	U                    map[string][]byte       `protobuf:"bytes,13,rep,name=u" json:"u,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	V                    map[string]string       `protobuf:"bytes,14,rep,name=v" json:"v,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	W                    map[string]bool         `protobuf:"bytes,15,rep,name=w" json:"w,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	X                    map[string]*UnaryFields `protobuf:"bytes,16,rep,name=x" json:"x,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Y                    map[string]TestEnum     `protobuf:"bytes,17,rep,name=y" json:"y,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=testprotos.TestEnum"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MapValFields) Reset()         { *m = MapValFields{} }
func (m *MapValFields) String() string { return proto.CompactTextString(m) }
func (*MapValFields) ProtoMessage()    {}
func (*MapValFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154c715c2e2b2ea, []int{4}
}
func (m *MapValFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValFields.Unmarshal(m, b)
}
func (m *MapValFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValFields.Marshal(b, m, deterministic)
}
func (m *MapValFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValFields.Merge(m, src)
}
func (m *MapValFields) XXX_Size() int {
	return xxx_messageInfo_MapValFields.Size(m)
}
func (m *MapValFields) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValFields.DiscardUnknown(m)
}

var xxx_messageInfo_MapValFields proto.InternalMessageInfo

func (m *MapValFields) GetI() map[string]int32 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *MapValFields) GetJ() map[string]int64 {
	if m != nil {
		return m.J
	}
	return nil
}

func (m *MapValFields) GetK() map[string]int32 {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *MapValFields) GetL() map[string]int64 {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *MapValFields) GetM() map[string]uint32 {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *MapValFields) GetN() map[string]uint64 {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *MapValFields) GetO() map[string]uint32 {
	if m != nil {
		return m.O
	}
	return nil
}

func (m *MapValFields) GetP() map[string]uint64 {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *MapValFields) GetQ() map[string]int32 {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *MapValFields) GetR() map[string]int64 {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *MapValFields) GetS() map[string]float32 {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *MapValFields) GetT() map[string]float64 {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *MapValFields) GetU() map[string][]byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *MapValFields) GetV() map[string]string {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *MapValFields) GetW() map[string]bool {
	if m != nil {
		return m.W
	}
	return nil
}

func (m *MapValFields) GetX() map[string]*UnaryFields {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *MapValFields) GetY() map[string]TestEnum {
	if m != nil {
		return m.Y
	}
	return nil
}

func init() {
	proto.RegisterEnum("testprotos.TestEnum", TestEnum_name, TestEnum_value)
	proto.RegisterType((*UnaryFields)(nil), "testprotos.UnaryFields")
	proto.RegisterType((*UnaryFields_GroupY)(nil), "testprotos.UnaryFields.GroupY")
	proto.RegisterType((*RepeatedFields)(nil), "testprotos.RepeatedFields")
	proto.RegisterType((*RepeatedFields_GroupY)(nil), "testprotos.RepeatedFields.GroupY")
	proto.RegisterType((*RepeatedPackedFields)(nil), "testprotos.RepeatedPackedFields")
	proto.RegisterType((*RepeatedPackedFields_GroupY)(nil), "testprotos.RepeatedPackedFields.GroupY")
	proto.RegisterType((*MapKeyFields)(nil), "testprotos.MapKeyFields")
	proto.RegisterMapType((map[int32]string)(nil), "testprotos.MapKeyFields.IEntry")
	proto.RegisterMapType((map[int64]string)(nil), "testprotos.MapKeyFields.JEntry")
	proto.RegisterMapType((map[int32]string)(nil), "testprotos.MapKeyFields.KEntry")
	proto.RegisterMapType((map[int64]string)(nil), "testprotos.MapKeyFields.LEntry")
	proto.RegisterMapType((map[uint32]string)(nil), "testprotos.MapKeyFields.MEntry")
	proto.RegisterMapType((map[uint64]string)(nil), "testprotos.MapKeyFields.NEntry")
	proto.RegisterMapType((map[uint32]string)(nil), "testprotos.MapKeyFields.OEntry")
	proto.RegisterMapType((map[uint64]string)(nil), "testprotos.MapKeyFields.PEntry")
	proto.RegisterMapType((map[int32]string)(nil), "testprotos.MapKeyFields.QEntry")
	proto.RegisterMapType((map[int64]string)(nil), "testprotos.MapKeyFields.REntry")
	proto.RegisterMapType((map[string]string)(nil), "testprotos.MapKeyFields.SEntry")
	proto.RegisterMapType((map[bool]string)(nil), "testprotos.MapKeyFields.TEntry")
	proto.RegisterType((*MapValFields)(nil), "testprotos.MapValFields")
	proto.RegisterMapType((map[string]int32)(nil), "testprotos.MapValFields.IEntry")
	proto.RegisterMapType((map[string]int64)(nil), "testprotos.MapValFields.JEntry")
	proto.RegisterMapType((map[string]int32)(nil), "testprotos.MapValFields.KEntry")
	proto.RegisterMapType((map[string]int64)(nil), "testprotos.MapValFields.LEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "testprotos.MapValFields.MEntry")
	proto.RegisterMapType((map[string]uint64)(nil), "testprotos.MapValFields.NEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "testprotos.MapValFields.OEntry")
	proto.RegisterMapType((map[string]uint64)(nil), "testprotos.MapValFields.PEntry")
	proto.RegisterMapType((map[string]int32)(nil), "testprotos.MapValFields.QEntry")
	proto.RegisterMapType((map[string]int64)(nil), "testprotos.MapValFields.REntry")
	proto.RegisterMapType((map[string]float32)(nil), "testprotos.MapValFields.SEntry")
	proto.RegisterMapType((map[string]float64)(nil), "testprotos.MapValFields.TEntry")
	proto.RegisterMapType((map[string][]byte)(nil), "testprotos.MapValFields.UEntry")
	proto.RegisterMapType((map[string]string)(nil), "testprotos.MapValFields.VEntry")
	proto.RegisterMapType((map[string]bool)(nil), "testprotos.MapValFields.WEntry")
	proto.RegisterMapType((map[string]*UnaryFields)(nil), "testprotos.MapValFields.XEntry")
	proto.RegisterMapType((map[string]TestEnum)(nil), "testprotos.MapValFields.YEntry")
}

func init() { proto.RegisterFile("desc_test_field_types.proto", fileDescriptor_2154c715c2e2b2ea) }

var fileDescriptor_2154c715c2e2b2ea = []byte{
	// 1080 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x97, 0xdd, 0x6e, 0xe3, 0x44,
	0x1b, 0xc7, 0xdf, 0x27, 0xd3, 0xe6, 0x63, 0x92, 0x34, 0xae, 0xdf, 0x4a, 0x8c, 0x8a, 0x00, 0x53,
	0x09, 0x18, 0x55, 0x6a, 0x2a, 0x02, 0x42, 0x2c, 0x27, 0x88, 0xa5, 0xbb, 0x90, 0x7e, 0xef, 0x34,
	0xed, 0x6e, 0x39, 0xa9, 0xd2, 0xd6, 0xbb, 0x9b, 0x6f, 0xd7, 0x71, 0xd2, 0x66, 0xaf, 0x80, 0x13,
	0x0e, 0x90, 0xb8, 0x0b, 0xb8, 0x12, 0x2e, 0x87, 0x2b, 0x40, 0x8f, 0x67, 0x12, 0xdb, 0xb5, 0x63,
	0x8f, 0x56, 0x9c, 0xf5, 0x79, 0xf4, 0xfb, 0x3b, 0x89, 0x7f, 0xd3, 0x7f, 0x1c, 0xfa, 0xe1, 0xad,
	0x3d, 0xbe, 0xb9, 0xf2, 0xec, 0xb1, 0x77, 0xf5, 0xba, 0x63, 0xf7, 0x6f, 0xaf, 0xbc, 0x99, 0x63,
	0x8f, 0xeb, 0x8e, 0x3b, 0xf2, 0x46, 0x26, 0xc5, 0xbd, 0xff, 0xe7, 0x78, 0xeb, 0x77, 0x42, 0xcb,
	0xe7, 0xc3, 0xb6, 0x3b, 0x7b, 0x8e, 0xd8, 0xd8, 0xac, 0x50, 0xe8, 0x30, 0xb0, 0x80, 0xaf, 0x0a,
	0xe8, 0xe0, 0xd4, 0x65, 0x39, 0x0b, 0x38, 0x11, 0xd0, 0xc5, 0xa9, 0xc7, 0x88, 0x05, 0x7c, 0x5d,
	0x40, 0x0f, 0xa7, 0x3e, 0x5b, 0xb1, 0x80, 0x9b, 0x02, 0xfa, 0x38, 0x0d, 0xd8, 0xaa, 0x05, 0xbc,
	0x2a, 0x60, 0x80, 0xd3, 0x90, 0xe5, 0x2d, 0xe0, 0x2b, 0x02, 0x86, 0x38, 0x8d, 0x58, 0xc1, 0x02,
	0x5e, 0x10, 0x30, 0xc2, 0xc9, 0x61, 0x45, 0x0b, 0x78, 0x5e, 0x80, 0x83, 0xd3, 0x1d, 0x2b, 0x59,
	0xc0, 0x6b, 0x02, 0xee, 0x70, 0x72, 0x19, 0xb5, 0x80, 0x1b, 0x02, 0x5c, 0x9c, 0xc6, 0xac, 0x6c,
	0x01, 0xcf, 0x09, 0xf0, 0xdf, 0x99, 0xc7, 0x2a, 0x16, 0x70, 0x10, 0xe0, 0xe1, 0x34, 0x61, 0x55,
	0x0b, 0x78, 0x45, 0xc0, 0x04, 0xa7, 0x29, 0x5b, 0xb3, 0x80, 0x97, 0x04, 0x4c, 0x71, 0xba, 0x67,
	0x35, 0x0b, 0x78, 0x51, 0xc0, 0xbd, 0xc9, 0x29, 0x3c, 0x30, 0xc3, 0x02, 0x5e, 0x6e, 0x6c, 0xd6,
	0x83, 0x4f, 0x5e, 0x17, 0xb6, 0x63, 0xb7, 0x3d, 0xfb, 0x56, 0x7e, 0x70, 0x01, 0x0f, 0xe6, 0x37,
	0x34, 0xff, 0xc6, 0x1d, 0x4d, 0x9c, 0x19, 0x5b, 0xb7, 0x80, 0xd3, 0xc6, 0xc7, 0x61, 0x3c, 0x74,
	0x93, 0xea, 0x3f, 0x21, 0x75, 0x29, 0x14, 0x6d, 0x6e, 0x51, 0x78, 0xc7, 0x4c, 0x0b, 0xf8, 0x5a,
	0x63, 0x23, 0x1c, 0x69, 0xd9, 0x63, 0xef, 0xd9, 0x70, 0x32, 0x10, 0xf0, 0x6e, 0x73, 0x9b, 0xe6,
	0x65, 0xca, 0xac, 0xd1, 0xdc, 0xac, 0xcd, 0xfe, 0x04, 0xff, 0xdd, 0xe6, 0x66, 0x6d, 0x7f, 0x71,
	0xcd, 0xfe, 0x92, 0x37, 0x3d, 0x37, 0xbb, 0xde, 0xfa, 0x83, 0xd0, 0xb5, 0xe8, 0xbb, 0x9b, 0x6b,
	0x21, 0x11, 0x2d, 0x24, 0xa2, 0x85, 0x44, 0xb4, 0x90, 0x88, 0x16, 0x12, 0xd1, 0x42, 0x22, 0x5a,
	0x48, 0x44, 0x0b, 0x89, 0x68, 0x21, 0x11, 0x2d, 0x24, 0xa2, 0x85, 0x44, 0xb4, 0x90, 0x88, 0x16,
	0x12, 0xd1, 0x42, 0x22, 0x5a, 0x88, 0xd4, 0xf2, 0x99, 0xd4, 0x42, 0x78, 0xb9, 0xf1, 0xc1, 0x92,
	0xfb, 0x8c, 0x4e, 0x9e, 0x84, 0x9c, 0x10, 0x4e, 0x1b, 0x9f, 0x2e, 0x57, 0xb8, 0x44, 0x0b, 0xf9,
	0xaf, 0xb4, 0xfc, 0x4a, 0xe8, 0xc6, 0xfc, 0x15, 0x4f, 0xdb, 0x37, 0xbd, 0x85, 0x1c, 0x63, 0x21,
	0xe7, 0x69, 0xce, 0x00, 0x14, 0x64, 0x2c, 0x04, 0xc9, 0x4d, 0x17, 0x37, 0x4a, 0x92, 0xdc, 0xf4,
	0x70, 0xa3, 0x44, 0xc9, 0x4d, 0x1f, 0x37, 0x4a, 0x96, 0xdc, 0x0c, 0x70, 0xa3, 0x84, 0xc9, 0xcd,
	0x10, 0x37, 0x4a, 0x9a, 0xdc, 0x8c, 0x70, 0xa3, 0xc4, 0xc9, 0x8d, 0x83, 0x1b, 0x25, 0x4f, 0x6e,
	0xee, 0x70, 0xa3, 0x04, 0xca, 0x8d, 0x8b, 0x1b, 0x25, 0x51, 0x6e, 0xfc, 0x4f, 0xa1, 0x44, 0xca,
	0x8d, 0x87, 0x1b, 0x29, 0xb3, 0x28, 0x37, 0x13, 0xf3, 0xfb, 0x85, 0x8d, 0x35, 0xdf, 0xc6, 0x17,
	0x49, 0x36, 0xc2, 0xf7, 0xe6, 0xb1, 0x93, 0xcf, 0xf1, 0x44, 0xd4, 0x96, 0x3b, 0x91, 0x2f, 0x34,
	0xdd, 0xfc, 0x68, 0xe1, 0xe5, 0xff, 0xbe, 0x86, 0xdf, 0x82, 0xdb, 0x8b, 0x2a, 0xfe, 0x2e, 0xd1,
	0xca, 0x51, 0xdb, 0x39, 0xb0, 0xe7, 0xb5, 0xb5, 0x33, 0x57, 0x50, 0x6e, 0x7c, 0x12, 0xbe, 0x6e,
	0x18, 0xaa, 0x37, 0x9f, 0x0d, 0x3d, 0x77, 0x86, 0x7e, 0x76, 0xe6, 0x7e, 0xd2, 0xf0, 0x7d, 0x85,
	0x77, 0x11, 0x97, 0xf2, 0xd2, 0xf0, 0x03, 0x85, 0xf7, 0x10, 0x97, 0x66, 0xd3, 0xf0, 0x43, 0x85,
	0xf7, 0x11, 0x97, 0xda, 0xd3, 0xf0, 0x23, 0x85, 0x0f, 0x10, 0x97, 0x67, 0x22, 0x0d, 0x3f, 0x56,
	0xf8, 0x10, 0x71, 0x79, 0x60, 0xd2, 0xf0, 0x13, 0x85, 0x8f, 0x10, 0x97, 0xa7, 0x29, 0x0d, 0x3f,
	0x55, 0xb8, 0x83, 0xb8, 0x3c, 0x6a, 0x69, 0xf8, 0x0b, 0x85, 0xdf, 0x21, 0x2e, 0xcf, 0x61, 0x1a,
	0x2e, 0x14, 0xee, 0x22, 0x2e, 0x0f, 0x69, 0x1a, 0x7e, 0xa6, 0x70, 0xff, 0x10, 0xc8, 0x13, 0x9c,
	0x86, 0xb7, 0x14, 0xee, 0x6d, 0x7e, 0x4d, 0xf3, 0xf2, 0x44, 0x98, 0x06, 0x25, 0x3d, 0x7b, 0xa6,
	0xbe, 0xf6, 0xf0, 0x4f, 0x73, 0x83, 0xae, 0x4e, 0xdb, 0xfd, 0x89, 0xed, 0x7f, 0xf9, 0x95, 0x84,
	0x1c, 0xbe, 0xcb, 0x7d, 0x0b, 0x98, 0xda, 0x8f, 0xa5, 0x88, 0x46, 0xea, 0x20, 0x96, 0x5a, 0xd7,
	0x48, 0x1d, 0xc6, 0x52, 0xa6, 0x46, 0xea, 0x28, 0x96, 0xaa, 0x6a, 0xa4, 0x8e, 0x63, 0xa9, 0x15,
	0x8d, 0xd4, 0x49, 0x2c, 0x55, 0xd0, 0x48, 0x9d, 0xc6, 0x52, 0x79, 0x8d, 0xd4, 0x8b, 0x58, 0xaa,
	0xa6, 0x91, 0x12, 0xb1, 0x94, 0xa1, 0x91, 0x3a, 0x8b, 0xa5, 0x4a, 0x1a, 0xa9, 0x56, 0x2c, 0x55,
	0xcc, 0x48, 0x6d, 0xfd, 0x53, 0xf5, 0xcb, 0xec, 0xa2, 0xdd, 0xcf, 0x2e, 0xb3, 0x05, 0xa4, 0x57,
	0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55,
	0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55,
	0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xae, 0x55, 0x66, 0x01, 0xbe, 0x28, 0x33, 0xc4, 0xe5, 0x77,
	0x75, 0x1a, 0x7e, 0xae, 0xf0, 0x09, 0xe2, 0xf2, 0xc9, 0x2c, 0x0d, 0xbf, 0x50, 0xf8, 0x14, 0x71,
	0xf9, 0xe8, 0x96, 0x86, 0xbf, 0x54, 0xf8, 0x3d, 0xe2, 0xf3, 0x67, 0xbb, 0xe5, 0xf8, 0x2b, 0x85,
	0x3f, 0x20, 0x2e, 0x1f, 0xef, 0xd2, 0xf0, 0x4b, 0x85, 0xcf, 0x92, 0x7b, 0x3b, 0xe9, 0x7f, 0x73,
	0x35, 0xb3, 0xb7, 0x93, 0x52, 0x24, 0xb3, 0xb7, 0x93, 0x52, 0xeb, 0x99, 0xbd, 0x9d, 0x94, 0x32,
	0x33, 0x7b, 0x3b, 0x29, 0x55, 0xcd, 0xec, 0xed, 0xa4, 0xd4, 0x4a, 0x66, 0x6f, 0x27, 0xa5, 0x0a,
	0x99, 0xbd, 0x9d, 0x94, 0xca, 0x67, 0xf6, 0x76, 0x52, 0xaa, 0x96, 0xd9, 0xdb, 0x49, 0x29, 0xe3,
	0xbd, 0x7a, 0x3b, 0x97, 0xd9, 0xdb, 0x49, 0x29, 0x78, 0x94, 0x3a, 0xd7, 0x4a, 0x55, 0x1e, 0xa5,
	0x2e, 0xde, 0xeb, 0x9b, 0xe5, 0xa5, 0x56, 0xaa, 0x18, 0x4e, 0x1d, 0xd1, 0xfc, 0xab, 0x65, 0xa9,
	0x9d, 0x70, 0x2a, 0xe5, 0x37, 0x58, 0xe8, 0x72, 0xfb, 0x34, 0x7f, 0xb9, 0xec, 0x72, 0xdb, 0xe1,
	0xcb, 0x2d, 0xfb, 0xc1, 0x15, 0x5c, 0x6b, 0xfb, 0x09, 0x2d, 0xce, 0xd7, 0x66, 0x99, 0x16, 0x9a,
	0xc7, 0x17, 0x3f, 0x1c, 0x36, 0xf7, 0x8c, 0xff, 0x99, 0x25, 0xba, 0xfa, 0xbc, 0x29, 0xce, 0x5a,
	0x06, 0x98, 0x14, 0x65, 0xfe, 0x78, 0x72, 0xbc, 0x67, 0xe4, 0x70, 0xdd, 0xfa, 0xb9, 0x29, 0xf6,
	0x0c, 0xf2, 0xf4, 0xab, 0x5f, 0xbe, 0x7c, 0xd3, 0xf1, 0xde, 0x4e, 0xae, 0xeb, 0x37, 0xa3, 0xc1,
	0x6e, 0xf7, 0xed, 0x64, 0xe0, 0xec, 0xfa, 0xaf, 0xe4, 0xda, 0xaf, 0xfb, 0xf6, 0x8d, 0xb7, 0xdb,
	0x19, 0x7a, 0xb6, 0x3b, 0x6c, 0xf7, 0x77, 0x83, 0xf7, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x51, 0x5d, 0x5c, 0x11, 0x11, 0x00, 0x00,
}
