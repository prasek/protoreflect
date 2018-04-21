package desc

import (
	"fmt"
	"sort"

	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// EnumDescriptor describes an enum declared in a proto file.
type EnumDescriptor struct {
	proto       *dpb.EnumDescriptorProto
	parent      Descriptor
	file        *FileDescriptor
	values      []*EnumValueDescriptor
	valuesByNum sortedValues
	fqn         string
	sourceInfo  *dpb.SourceCodeInfo_Location
}

func createEnumDescriptor(fd *FileDescriptor, parent Descriptor, enclosing string, ed *dpb.EnumDescriptorProto, symbols map[string]Descriptor) (*EnumDescriptor, string) {
	enumName := merge(enclosing, ed.GetName())
	ret := &EnumDescriptor{proto: ed, parent: parent, file: fd, fqn: enumName}
	for _, ev := range ed.GetValue() {
		evd, n := createEnumValueDescriptor(fd, ret, enumName, ev)
		symbols[n] = evd
		ret.values = append(ret.values, evd)
	}
	if len(ret.values) > 0 {
		ret.valuesByNum = make(sortedValues, len(ret.values))
		copy(ret.valuesByNum, ret.values)
		sort.Stable(ret.valuesByNum)
	}
	return ret, enumName
}

type sortedValues []*EnumValueDescriptor

func (sv sortedValues) Len() int {
	return len(sv)
}

func (sv sortedValues) Less(i, j int) bool {
	return sv[i].GetNumber() < sv[j].GetNumber()
}

func (sv sortedValues) Swap(i, j int) {
	sv[i], sv[j] = sv[j], sv[i]
}

func (ed *EnumDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	ed.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Enum_valuesTag)
	for i, evd := range ed.values {
		evd.resolve(append(path, int32(i)), sourceCodeInfo)
	}
}

func (ed *EnumDescriptor) GetName() string {
	return ed.proto.GetName()
}

func (ed *EnumDescriptor) GetFullyQualifiedName() string {
	return ed.fqn
}

func (ed *EnumDescriptor) GetParent() Descriptor {
	return ed.parent
}

func (ed *EnumDescriptor) GetFile() *FileDescriptor {
	return ed.file
}

func (ed *EnumDescriptor) GetOptions() proto.Message {
	return ed.proto.GetOptions()
}

func (ed *EnumDescriptor) GetEnumOptions() *dpb.EnumOptions {
	return ed.proto.GetOptions()
}

func (ed *EnumDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return ed.sourceInfo
}

func (ed *EnumDescriptor) AsProto() proto.Message {
	return ed.proto
}

func (ed *EnumDescriptor) AsEnumDescriptorProto() *dpb.EnumDescriptorProto {
	return ed.proto
}

func (ed *EnumDescriptor) String() string {
	return ed.proto.String()
}

// GetValues returns all of the allowed values defined for this enum.
func (ed *EnumDescriptor) GetValues() []*EnumValueDescriptor {
	return ed.values
}

// FindValueByName finds the enum value with the given name. If no such value exists
// then nil is returned.
func (ed *EnumDescriptor) FindValueByName(name string) *EnumValueDescriptor {
	fqn := fmt.Sprintf("%s.%s", ed.fqn, name)
	if vd, ok := ed.file.symbols[fqn].(*EnumValueDescriptor); ok {
		return vd
	} else {
		return nil
	}
}

// FindValueByNumber finds the value with the given numeric value. If no such value
// exists then nil is returned. If aliases are allowed and multiple values have the
// given number, the first declared value is returned.
func (ed *EnumDescriptor) FindValueByNumber(num int32) *EnumValueDescriptor {
	index := sort.Search(len(ed.valuesByNum), func(i int) bool { return ed.valuesByNum[i].GetNumber() >= num })
	if index < len(ed.valuesByNum) {
		vd := ed.valuesByNum[index]
		if vd.GetNumber() == num {
			return vd
		}
	}
	return nil
}

// EnumValueDescriptor describes an allowed value of an enum declared in a proto file.
type EnumValueDescriptor struct {
	proto      *dpb.EnumValueDescriptorProto
	parent     *EnumDescriptor
	file       *FileDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

func createEnumValueDescriptor(fd *FileDescriptor, parent *EnumDescriptor, enclosing string, evd *dpb.EnumValueDescriptorProto) (*EnumValueDescriptor, string) {
	valName := merge(enclosing, evd.GetName())
	return &EnumValueDescriptor{proto: evd, parent: parent, file: fd, fqn: valName}, valName
}

func (vd *EnumValueDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	vd.sourceInfo = sourceCodeInfo.Get(path)
}

func (vd *EnumValueDescriptor) GetName() string {
	return vd.proto.GetName()
}

// GetNumber returns the numeric value associated with this enum value.
func (vd *EnumValueDescriptor) GetNumber() int32 {
	return vd.proto.GetNumber()
}

func (vd *EnumValueDescriptor) GetFullyQualifiedName() string {
	return vd.fqn
}

func (vd *EnumValueDescriptor) GetParent() Descriptor {
	return vd.parent
}

// GetEnum returns the enum in which this enum value is defined.
func (vd *EnumValueDescriptor) GetEnum() *EnumDescriptor {
	return vd.parent
}

func (vd *EnumValueDescriptor) GetFile() *FileDescriptor {
	return vd.file
}

func (vd *EnumValueDescriptor) GetOptions() proto.Message {
	return vd.proto.GetOptions()
}

func (vd *EnumValueDescriptor) GetEnumValueOptions() *dpb.EnumValueOptions {
	return vd.proto.GetOptions()
}

func (vd *EnumValueDescriptor) GetSourceInfo() *dpb.SourceCodeInfo_Location {
	return vd.sourceInfo
}

func (vd *EnumValueDescriptor) AsProto() proto.Message {
	return vd.proto
}

func (vd *EnumValueDescriptor) AsEnumValueDescriptorProto() *dpb.EnumValueDescriptorProto {
	return vd.proto
}

func (vd *EnumValueDescriptor) String() string {
	return vd.proto.String()
}
