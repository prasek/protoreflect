# Protocol Buffer and GRPC Reflection
[![Build Status](https://travis-ci.org/prasek/protoreflect.svg?branch=master)](https://travis-ci.org/prasek/protoreflect/branches)
[![codecov](https://codecov.io/gh/prasek/protoreflect/branch/master/graph/badge.svg)](https://codecov.io/gh/prasek/protoreflect)
----
## Forked to work with gogo/protobuf

Sample glide.yaml
```yaml
 - package: github.com/jhump/protoreflect
   vcs: git
   repo: https://github.com/prasek/protoreflect.git
   version: v0.1.0
```
----

This repo provides reflection APIs for [protocol buffers](https://developers.google.com/protocol-buffers/) (also known as "protobufs" for short)
and [gRPC](https://grpc.io/). The core of reflection in protobufs is the
[descriptor](https://github.com/google/protobuf/blob/199d82fde1734ab5bc931cd0de93309e50cd7ab9/src/google/protobuf/descriptor.proto).
A descriptor is itself a protobuf message that describes a `.proto` source file or any element
therein. So a collection of descriptors can describe an entire schema of protobuf types, including
RPC services.

[![GoDoc](https://godoc.org/github.com/jhump/protoreflect?status.svg)](https://godoc.org/github.com/jhump/protoreflect)

----
## Descriptors: The Language Model of Protocol Buffers

```go
import "github.com/jhump/protoreflect/desc"
```

The `desc` package herein introduces a `Descriptor` interface and implementations of it that
correspond to each of the descriptor types. These new types are effectively smart wrappers around
the [generated protobuf types](https://github.com/gogo/protobuf/blob/master/protoc-gen-gogo/descriptor/descriptor.pb.go)
that make them *much* more useful and easier to use.

You can construct descriptors from file descriptor sets (which can be generated by `protoc`), and
you can also load descriptors for messages and services that are linked into the current binary.
"What does it mean for messages and services to be linked in?" you may ask. It means your binary
imports a package that was generated by `protoc`. When you generate Go code from your `.proto`
sources, the resulting package has descriptor information embedded in it. The `desc` package allows
you to easily extract those embedded descriptors.

Descriptors can also be acquired directly from `.proto` source files (using the `protoparse` sub-package)
or by programmatically constructing them (using the `builder` sub-package).

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/desc)*

```go
import "github.com/jhump/protoreflect/desc/protoparse"
```

The `protoparse` package allows for parsing of `.proto` source files into rich descriptors. Without
this package, you must invoke `protoc` to either generate a file descriptor set file or to generate
Go code (which has descriptor information embedded in it). This package allows reading the source
directly without having to invoke `protoc`.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/desc/protoparse)*

```go
import "github.com/jhump/protoreflect/desc/protoprint"
```

The `protoprint` package allows for printing of descriptors to `.proto` source files. This is
effectively the inverse of the `protoparse` package. Combined with the `builder` package, this
is a useful tool for programmatically generating protocol buffer sources.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/desc/protoprint)*

```go
import "github.com/jhump/protoreflect/desc/builder"
```

The `builder` package allows for programmatic construction of rich descriptors. Descriptors can
be constructed programmatically by creating trees of descriptor protos and using the `desc` package
to link those into rich descriptors. But constructing a valid tree of descriptor protos is far from
trivial.

So this package provides generous API to greatly simplify that task. It also allows for converting
rich descriptors into builders, which means you can programmatically modify/tweak existing
descriptors.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/desc/builder)*

----
## Dynamic Messages and Stubs

```go
import "github.com/jhump/protoreflect/dynamic"
```

The `dynamic` package provides a dynamic message implementation. It implements `proto.Message` but
is backed by a message descriptor and a map of fields->values, instead of a generated struct. This
is useful for acting generically with protocol buffer messages, without having to generate and link
in Go code for every kind of message. This is particularly useful for general-purpose tools that
need to operate on arbitrary protocol buffer schemas. This is made possible by having the tools load
descriptors at runtime.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/dynamic)*

```go
import "github.com/jhump/protoreflect/dynamic/grpcdynamic"
```

There is also sub-package named `grpcdynamic`, which provides a dynamic stub implementation. The stub can
be used to issue RPC methods using method descriptors instead of generated client interfaces.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/dynamic/grpcdynamic)*

----
## gRPC Server Reflection

```go
import "github.com/jhump/protoreflect/grpcreflect"
```

The `grpcreflect` package provides an easy-to-use client for the
[gRPC reflection service](https://github.com/grpc/grpc-go/blob/6bd4f6eb1ea9d81d1209494242554dcde44429a4/reflection/grpc_reflection_v1alpha/reflection.proto#L36),
making it much easier to query for and work with the schemas of remote services.

It also provides some helper methods for querying for rich service descriptors for the
services registered in a gRPC server.

*[Read more ≫](https://godoc.org/github.com/jhump/protoreflect/grpcreflect)*
