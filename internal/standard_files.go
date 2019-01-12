// Package internal contains some code that should not be exported but needs to
// be shared across more than one of the protoreflect sub-packages.
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// TODO: replace this alias configuration with desc.RegisterImportPath?

// StdFileAliases are the standard protos included with protoc, but older versions of
// their respective packages registered them using incorrect paths.
var StdFileAliases = map[string]string{
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

func init() {
	// We provide aliasing in both directions, to support files with the
	// proper import path linked against older versions of the generated
	// files AND files that used the aliased import path but linked against
	// newer versions of the generated files (which register with the
	// correct path).

	// Get all files defined above
	keys := make([]string, 0, len(StdFileAliases))
	for k := range StdFileAliases {
		keys = append(keys, k)
	}
	// And add inverse mappings
	for _, k := range keys {
		alias := StdFileAliases[k]
		StdFileAliases[alias] = k
	}
}

type ErrNoSuchFile string

func (e ErrNoSuchFile) Error() string {
	return fmt.Sprintf("no such file: %q", string(e))
}

// LoadFileDescriptor loads a registered descriptor and decodes it. If the given
// name cannot be loaded but is a known standard name, an alias will be tried,
// so the standard files can be loaded even if linked against older "known bad"
// versions of packages.
func LoadFileDescriptor(file string) (*dpb.FileDescriptorProto, error) {
	fdb := proto.FileDescriptor(file)
	aliased := false
	if fdb == nil {
		var ok bool
		alias, ok := StdFileAliases[file]
		if ok {
			//fmt.Printf(" - aliased to: %v\n", alias)
			aliased = true
			if fdb = proto.FileDescriptor(alias); fdb == nil {
				return nil, ErrNoSuchFile(file)
			}
		} else {
			return nil, ErrNoSuchFile(file)
		}
	}

	fd, err := DecodeFileDescriptor(file, fdb)
	if err != nil {
		return nil, err
	}

	if aliased {
		// the file descriptor will have the alias used to load it, but
		// we need it to have the specified name in order to link it
		fd.Name = proto.String(file)
	}

	return fd, nil
}

// DecodeFileDescriptor decodes the bytes of a registered file descriptor.
// Registered file descriptors are first "proto encoded" (e.g. binary format
// for the descriptor protos) and then gzipped. So this function gunzips and
// then unmarshals into a descriptor proto.
func DecodeFileDescriptor(element string, fdb []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(fdb)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress %q descriptor: %v", element, err)
	}
	fd := dpb.FileDescriptorProto{}
	if err := proto.Unmarshal(raw, &fd); err != nil {
		return nil, fmt.Errorf("bad descriptor for %q: %v", element, err)
	}
	return &fd, nil
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}
