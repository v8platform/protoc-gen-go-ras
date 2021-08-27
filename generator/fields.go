package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"sort"
)

type Fields []field

func (f Fields) Range(fn func(f field) bool) {

	for _, f2 := range f {
		if !fn(f2) {
			break
		}
	}
}

func (f Fields) FindByNumber(n int32) field {

	for _, f2 := range f {

		if int32(f2.Descriptor.Number()) == n {
			return f2
		}

	}

	panic("no found field by number")
}

type field struct {
	GoName     string
	GoIdent    protogen.GoIdent
	Message    *protogen.Field
	Descriptor protoreflect.FieldDescriptor

	Opts *MessageFieldExtension

	Encoder protogen.GoIdent
	Decoder protogen.GoIdent
}

func getMessageFields(m *protogen.Message) Fields {

	var fields Fields

	for _, mfield := range m.Fields {

		opts := GetMessageExtensionFor(mfield.Desc.Options())

		if opts == nil {
			continue
		}

		fields = append(fields, field{
			Opts:       opts,
			GoName:     mfield.GoName,
			GoIdent:    mfield.GoIdent,
			Message:    mfield,
			Descriptor: mfield.Desc,
			Encoder:    getEncoder(opts.GetEncoder(), mfield.Desc),
			Decoder:    getDecoder(opts.GetEncoder(), mfield.Desc),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Opts.GetOrder() < fields[j].Opts.GetOrder()
	})

	return fields

}

func getDecoder(name string, desc protoreflect.FieldDescriptor) protogen.GoIdent {
	encoder, ok := getEncoderByName(name)

	if !ok {
		encoder = getEncoderByKind(desc)
	}

	return encoder
}

func getEncoder(name string, desc protoreflect.FieldDescriptor) protogen.GoIdent {
	encoder, ok := getEncoderByName(name)

	if !ok {
		encoder = getEncoderByKind(desc)
	}

	return encoder
}

func getEncoderByName(encoder string) (protogen.GoIdent, bool) {

	ident, ok := encoders[encoder]

	if ok {
		return ident, ok
	}

	return protogen.GoIdent{}, false
}

func getEncoderByKind(desc protoreflect.FieldDescriptor) protogen.GoIdent {

	switch desc.Kind() {
	case protoreflect.BoolKind:
		return encoders["bool"]
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return encoders["int32"]
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return encoders["int64"]
	case protoreflect.FloatKind:
		return encoders["float32"]
	case protoreflect.DoubleKind:
		return encoders["float64"]
	case protoreflect.StringKind:
		return encoders["string"]
	default:
		return protogen.GoIdent{}
	}
}
