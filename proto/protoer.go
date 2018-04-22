package proto

import (
	"reflect"
	"sync"
)

var (
	mu      sync.RWMutex
	protoer Protoer = NewGoGoProtoer(defaultGoGoAliases)
)

type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

type Protoer interface {
	MessageType(name string) reflect.Type
	MessageName(m Message) string
	FileDescriptor(file string) []byte
	Unmarshal(b []byte, m Message) error
	Marshal(m Message) ([]byte, error)
	GetExtension(pb Message, field int32) (interface{}, error)
}

type AliasMap map[string]string

type Aliaser interface {
	Aliases() AliasMap
}

func SetProtoer(p Protoer) {
	mu.Lock()
	defer mu.Unlock()
	protoer = p
}

func MessageType(name string) reflect.Type {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.MessageType(name)
}

func MessageName(m Message) string {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.MessageName(m)
}

func FileDescriptor(file string) []byte {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.FileDescriptor(file)
}

func Unmarshal(b []byte, m Message) error {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Unmarshal(b, m)
}

func Marshal(m Message) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Marshal(m)
}

func GetExtension(pb Message, field int32) (interface{}, error) {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.GetExtension(pb, field)
}

func Aliases() AliasMap {
	mu.RLock()
	defer mu.RUnlock()
	if a, ok := protoer.(Aliaser); ok {
		return a.Aliases()
	}
	return AliasMap{}
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it
func String(v string) *string {
	return &v
}
