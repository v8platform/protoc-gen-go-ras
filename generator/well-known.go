package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

func IsWellKnownType(fullname string) bool {
	return strings.HasPrefix(fullname, "google.protobuf.")
}

type generateParseFunc func(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool)
type generateFormatterFunc func(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool)

func (gen *Generator) wellKnownTypeFormatter(name protoreflect.FullName) generateFormatterFunc {
	switch name {
	// case genid.Any_message_name:
	// 	return decoder.unmarshalAny
	case "google.protobuf.Timestamp":
		return gen.generateTimestampFormatter
		// 	case genid.Duration_message_name:
		// 		return decoder.unmarshalDuration
		// 	case genid.BoolValue_message_name,
		// 		genid.Int32Value_message_name,
		// 		genid.Int64Value_message_name,
		// 		genid.UInt32Value_message_name,
		// 		genid.UInt64Value_message_name,
		// 		genid.FloatValue_message_name,
		// 		genid.DoubleValue_message_name,
		// 		genid.StringValue_message_name,
		// 		genid.BytesValue_message_name:
		// 		return decoder.unmarshalWrapperType
		// 	case genid.Struct_message_name:
		// 		return decoder.unmarshalStruct
		// 	case genid.ListValue_message_name:
		// 		return decoder.unmarshalListValue
		// 	case genid.Value_message_name:
		// 		return decoder.unmarshalKnownValue
		// 	case genid.FieldMask_message_name:
		// 		return decoder.unmarshalFieldMask
		// 	case genid.Empty_message_name:
		// 		return decoder.unmarshalEmpty
	}
	return nil
}

// wellKnownTypeUnmarshaler returns a unmarshal function if the message type
// has specialized serialization behavior. It returns nil otherwise.
func (gen *Generator) wellKnownTypeParse(fullname protoreflect.FullName) generateParseFunc {
	switch fullname {
	// case genid.Any_message_name:
	// 	return decoder.unmarshalAny
	case "google.protobuf.Timestamp":
		return gen.generateTimestampParse
		// 	case genid.Duration_message_name:
		// 		return decoder.unmarshalDuration
		// 	case genid.BoolValue_message_name,
		// 		genid.Int32Value_message_name,
		// 		genid.Int64Value_message_name,
		// 		genid.UInt32Value_message_name,
		// 		genid.UInt64Value_message_name,
		// 		genid.FloatValue_message_name,
		// 		genid.DoubleValue_message_name,
		// 		genid.StringValue_message_name,
		// 		genid.BytesValue_message_name:
		// 		return decoder.unmarshalWrapperType
		// 	case genid.Struct_message_name:
		// 		return decoder.unmarshalStruct
		// 	case genid.ListValue_message_name:
		// 		return decoder.unmarshalListValue
		// 	case genid.Value_message_name:
		// 		return decoder.unmarshalKnownValue
		// 	case genid.FieldMask_message_name:
		// 		return decoder.unmarshalFieldMask
		// 	case genid.Empty_message_name:
		// 		return decoder.unmarshalEmpty
	}
	return nil
}

func (gen *Generator) generateTimestampParse(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	valueName := string(f.Descriptor.Message().FullName())
	if createVal {
		g.P(identifier, " := &", gen.ObjectNamed(valueName).GoIdent, "{}")

	} else {
		g.P(identifier, " = &", gen.ObjectNamed(valueName).GoIdent, "{}")
	}

	g.P("if err:= ", f.Decoder, "(reader, ", identifier, ")", "; err != nil { return err }")

}

func (gen *Generator) generateTimestampFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	g.P("// TODO check nil")
	g.P("if err:= ", f.Encoder, "(writer, x.Get", f.GoName, "().AsTime())", "; err != nil { return err }")

}
