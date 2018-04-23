package desc

import (
	"os"
	"testing"

	"github.com/jhump/protoreflect/proto"
	"github.com/jhump/protoreflect/proto/gogo"
)

func TestMain(m *testing.M) {
	proto.SetProtoer(gogo.NewProtoer(nil))
	code := m.Run()
	os.Exit(code)
}
