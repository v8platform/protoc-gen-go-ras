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

func (f Fields) FindByName(name string) field {

	for _, f2 := range f {

		if f2.GoName == name {
			return f2
		}

	}

	panic("no found field by name")
}

type field struct {
	GoName     string
	GoIdent    protogen.GoIdent
	Message    *protogen.Field
	Descriptor protoreflect.FieldDescriptor

	Opts    *MessageFieldExtension
	Field   *protogen.Field
	Oneof   *protogen.Field
	Encoder protogen.GoIdent
	Decoder protogen.GoIdent
}

func getMessageFields(m *protogen.Message) Fields {

	var fields Fields

	findOneOf := func(f *protogen.Field) *protogen.Field {
		if f.Oneof == nil {
			return nil
		}

		for _, OneofField := range f.Oneof.Fields {
			if OneofField.GoName == f.GoName {
				return OneofField
			}
		}

		return nil
	}

	for _, mfield := range m.Fields {

		opts := GetMessageFieldExtensionFor(mfield.Desc.Options())

		if opts == nil {
			continue
		}

		fields = append(fields, field{
			Opts:       opts,
			GoName:     mfield.GoName,
			GoIdent:    mfield.GoIdent,
			Message:    mfield,
			Descriptor: mfield.Desc,
			Field:      mfield,
			Oneof:      findOneOf(mfield),
			Encoder:    getEncoder(opts.GetEncoder(), mfield.Desc),
			Decoder:    getDecoder(opts.GetEncoder(), mfield.Desc),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Opts.GetOrder() < fields[j].Opts.GetOrder()
	})

	return fields

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
		return encoders["int"]
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return encoders["long"]
	case protoreflect.FloatKind:
		return encoders["float"]
	case protoreflect.DoubleKind:
		return encoders["double"]
	case protoreflect.StringKind:
		return encoders["string"]
	case protoreflect.EnumKind:
		return encoders["byte"]
	default:
		return protogen.GoIdent{}
	}
}

func getDecoder(name string, desc protoreflect.FieldDescriptor) protogen.GoIdent {
	encoder, ok := getDecoderByName(name)

	if !ok {
		encoder = getDecoderByKind(desc)
	}

	return encoder
}

func getDecoderByName(encoder string) (protogen.GoIdent, bool) {

	ident, ok := decoders[encoder]

	if ok {
		return ident, ok
	}

	return protogen.GoIdent{}, false
}

func getDecoderByKind(desc protoreflect.FieldDescriptor) protogen.GoIdent {

	switch desc.Kind() {
	case protoreflect.BoolKind:
		return decoders["bool"]
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return decoders["int"]
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return decoders["long"]
	case protoreflect.FloatKind:
		return decoders["float"]
	case protoreflect.DoubleKind:
		return decoders["double"]
	case protoreflect.StringKind:
		return decoders["string"]
	case protoreflect.EnumKind:
		return decoders["byte"]
	default:
		return protogen.GoIdent{}
	}
}
