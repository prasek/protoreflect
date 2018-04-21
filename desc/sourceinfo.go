package desc

import dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

type sourceInfoMap map[string]*dpb.SourceCodeInfo_Location

func (m sourceInfoMap) Get(path []int32) *dpb.SourceCodeInfo_Location {
	return m[asMapKey(path)]
}

func (m sourceInfoMap) Put(path []int32, loc *dpb.SourceCodeInfo_Location) {
	m[asMapKey(path)] = loc
}

func (m sourceInfoMap) PutIfAbsent(path []int32, loc *dpb.SourceCodeInfo_Location) bool {
	k := asMapKey(path)
	if _, ok := m[k]; ok {
		return false
	}
	m[k] = loc
	return true
}

func asMapKey(slice []int32) string {
	// NB: arrays should be usable as map keys, but this does not
	// work due to a bug: https://github.com/golang/go/issues/22605
	//rv := reflect.ValueOf(slice)
	//arrayType := reflect.ArrayOf(rv.Len(), rv.Type().Elem())
	//array := reflect.New(arrayType).Elem()
	//reflect.Copy(array, rv)
	//return array.Interface()

	b := make([]byte, len(slice)*4)
	for i, s := range slice {
		j := i * 4
		b[j] = byte(s)
		b[j+1] = byte(s >> 8)
		b[j+2] = byte(s >> 16)
		b[j+3] = byte(s >> 24)
	}
	return string(b)
}

func createSourceInfoMap(fd *dpb.FileDescriptorProto) sourceInfoMap {
	res := sourceInfoMap{}
	for _, l := range fd.GetSourceCodeInfo().GetLocation() {
		res.Put(l.Path, l)
	}
	return res
}
