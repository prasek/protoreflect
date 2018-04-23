package desc

import (
	"os"
	"testing"

	"github.com/jhump/protoreflect/proto"
	"github.com/jhump/protoreflect/proto/golang"
)

func TestMain(m *testing.M) {
	proto.SetProtoer(golang.NewProtoer(nil))
	code := m.Run()
	os.Exit(code)
}
