package desc

// These are standard protos included with protoc, but older versions of their
// respective packages registered them using incorrect paths.
var stdFileAliases = map[string]string{
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
	keys := make([]string, 0, len(stdFileAliases))
	for k := range stdFileAliases {
		keys = append(keys, k)
	}
	// And add inverse mappings
	for _, k := range keys {
		alias := stdFileAliases[k]
		stdFileAliases[alias] = k
	}
}
