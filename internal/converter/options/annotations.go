package options

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MessageAnnotator interface {
	AnnotateMessage(opts Options, schema *base.Schema, desc protoreflect.MessageDescriptor) *base.Schema
}

type FieldAnnotator interface {
	AnnotateField(opts Options, schema *base.Schema, desc protoreflect.FieldDescriptor, onlyScalar bool) *base.Schema
}

type OneofAnnotator interface {
	// Annotate the schema generated for a oneof: one titled variant per member in
	// anyOf, plus a "not" forbidding two members at once when there are several.
	AnnotateOneof(opts Options, group *base.Schema, desc protoreflect.OneofDescriptor) *base.Schema
}

type FieldReferenceAnnotator interface {
	// Annotate a field reference. This takes in the PARENT of the field, because with references
	// we can only annotate on the things on the parent like the list of required attributes.
	AnnotateFieldReference(opts Options, parent *base.Schema, desc protoreflect.FieldDescriptor) *base.Schema
}
