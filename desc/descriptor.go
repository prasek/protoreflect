package desc

// Descriptor is the common interface implemented by all descriptor objects.
type Descriptor interface {
	// GetName returns the name of the object described by the descriptor. This will
	// be a base name that does not include enclosing message names or the package name.
	// For file descriptors, this indicates the path and name to the described file.
	GetName() string
	// GetFullyQualifiedName returns the fully-qualified name of the object described by
	// the descriptor. This will include the package name and any enclosing message names.
	// For file descriptors, this returns the path and name to the described file (same as
	// GetName).
	GetFullyQualifiedName() string
	// GetParent returns the enclosing element in a proto source file. If the described
	// object is a top-level object, this returns the file descriptor. Otherwise, it returns
	// the element in which the described object was declared. File descriptors have no
	// parent and return nil.
	GetParent() Descriptor
	// GetFile returns the file descriptor in which this element was declared. File
	// descriptors return themselves.
	GetFile() *FileDescriptor

	/*
			// GetOptions returns the options proto containing options for the described element.
			GetOptions() proto.Message

		// GetSourceInfo returns any source code information that was present in the file
		// descriptor. Source code info is optional. If no source code info is available for
		// the element (including if there is none at all in the file descriptor) then this
		// returns nil
		GetSourceInfo() *dpb.SourceCodeInfo_Location
	*/
}

// scope represents a lexical scope in a proto file in which messages and enums
// can be declared.
type scope func(string) Descriptor
