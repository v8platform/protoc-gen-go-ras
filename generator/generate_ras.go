package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

type rasGenerator struct {
	*Generator
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func (r rasGenerator) GenerateFileContent() {

	if len(r.file.Messages) == 0 {
		return
	}

	for _, message := range r.file.Messages {
		r.genMessage(message)
	}
}

func (r rasGenerator) genGetTypeMessage(message *protogen.Message, ext MessageExtension) {

	if len(ext.PacketType) > 0 {
		r.genGetTypeMessageFunc(message, ext.PacketType)
	}
	if len(ext.EndpointDataType) > 0 {
		r.genGetTypeMessageFunc(message, ext.EndpointDataType)
	}
	if len(ext.MessageType) > 0 {
		r.genGetTypeMessageFunc(message, ext.MessageType)
	}

}

func (r rasGenerator) genGetTypeMessageFunc(message *protogen.Message, typeValue string) {
	enumValue := r.idxEnumValues[typeValue]
	funcName := r.getGetTypeFuncName(enumValue)
	r.g.Annotate(message.GoIdent.GoName, message.Location)
	r.g.P("func (x *", message.GoIdent, ") ", funcName, "() ", enumValue.Parent.GoIdent, " {")
	r.g.P("return ", enumValue.GoIdent)
	r.g.P("}")
	r.g.P()
	r.g.Unskip()
}

func (r rasGenerator) getGetTypeFuncName(enum *protogen.EnumValue) string {
	return "Get" + enum.Parent.GoIdent.GoName
}
func (r rasGenerator) genMessageHelpers(message *protogen.Message, ext MessageExtension) {

	r.genGetTypeMessage(message, ext)

	r.generatePacketHelpers(message, ext)
	r.generateEndpointMessageHelpers(message, ext)
	r.generateEndpointHelpers(message, ext)

	if ext.GenerateErrorFn {
		r.generateErrorFunc(message)
	}

	if ext.GenerateIoWriteTo {
		r.generateWriteToFunc(message)
	}

	if ext.IsNegotiate {
		r.generateNegotiateHelpers(message)
	}

}

func (r rasGenerator) genMessage(message *protogen.Message) {
	ext := GetMessageExtension(message.Desc.Options())

	r.genMessageHelpers(message, ext)

	r.genParseFunc(message, ext)
	r.genFormatterFunc(message, ext)
}

func (r rasGenerator) generateEndpointMessageHelpers(m *protogen.Message, ext MessageExtension) {

	if !ext.GenerateEndpointMessageHelpers {
		return
	}

	endpointMessageDataTypeIdent := r.EnumNamed("ENDPOINT_DATA_TYPE_MESSAGE").GoIdent
	messageDataIdent := r.idxMessageByEnumValue["ENDPOINT_DATA_TYPE_MESSAGE"].GoIdent
	_ = r.idxMessageByEnumValue["ENDPOINT_DATA_TYPE_VOID_MESSAGE"].GoIdent
	_ = r.idxMessageByEnumValue["ENDPOINT_DATA_TYPE_EXCEPTION"].GoIdent

	messageDataEnumIdent := r.EnumNamed("ENDPOINT_DATA_TYPE_MESSAGE").GoIdent
	voidMessageEnumIdent := r.EnumNamed("ENDPOINT_DATA_TYPE_VOID_MESSAGE").GoIdent
	exceptionMessagEnumeIdent := r.EnumNamed("ENDPOINT_DATA_TYPE_EXCEPTION").GoIdent

	MessageTypeIdent := r.KnownTypes.EnumMessageType.GoIdent

	messageName := m.GoIdent.GoName
	messageFormatterName := messageName + "Formatter"
	messageParserName := messageName + "Parser"

	fields := getMessageFields(m)
	dataField := fields.FindByNumber(5)
	r.g.P("// Helpers generated by github.com/v8platform/protoc-r-go-ras. DO NOT EDIT")
	r.g.P()
	r.g.P("type ", messageFormatterName, " interface {")
	r.g.P(" GetMessageType() ", MessageTypeIdent)
	r.g.P(" ", r.formatFuncName, "(writer io.Writer, version int32) error")
	r.g.P("}")
	r.g.P()
	r.g.P("type ", messageParserName, " interface {")
	r.g.P(" GetMessageType() ", MessageTypeIdent)
	r.g.P(" ", r.parseFuncName, "(reader io.Reader, version int32) error")
	r.g.P("}")
	r.g.P()
	r.g.P("func New", m.GoIdent.GoName, "(endpoint ", r.KnownTypes.EndpointImpl(), ", message ", messageFormatterName, ") (*", m.GoIdent, ", error){")
	r.g.P("buf := &", bytesBuffer, "{}")
	r.g.P("if err := message.", r.formatFuncName, "(buf, endpoint.GetVersion()); err != nil { return nil, err }")
	r.g.P("return &", m.GoIdent, "{")
	r.g.P("Type: ", endpointMessageDataTypeIdent, ",")
	r.g.P("Format: endpoint.GetFormat(),")
	r.g.P("EndpointId: endpoint.GetId(),")
	r.g.P("Data: &", dataField.Oneof.GoIdent, "{")
	r.g.P(dataField.Oneof.GoName, ": &", messageDataIdent, "{")
	r.g.P("Bytes: buf.Bytes(),")
	r.g.P("Type: message.GetMessageType(),")
	r.g.P("},")
	r.g.P("},")
	r.g.P("}, nil")
	r.g.P("}")
	r.g.P()
	r.g.P("func (x *", m.GoIdent, ") Unpack(endpoint ", r.KnownTypes.EndpointImpl(), ", into ", messageParserName, ") error {")
	r.g.P("switch x.GetType() {")
	r.g.P("case ", messageDataEnumIdent, ":")
	r.g.P("buf := ", bytesNewBuffer, "(x.GetMessage().GetBytes())")
	r.g.P("if err := into.", r.parseFuncName, "(buf, endpoint.GetVersion()); err != nil { return err }")
	r.g.P("return nil ")
	r.g.P("case ", voidMessageEnumIdent, ":")
	r.g.P("return nil ")
	r.g.P("case ", exceptionMessagEnumeIdent, ":")
	r.g.P("return x.GetFailure()")
	r.g.P("default:")
	r.g.P("return ", fmtErrorf, "(\"unknown message type <%s>\", x.GetType())")
	r.g.P("}")
	r.g.P("}")
	r.g.P()

	r.AddImpl(messageFormatterName, m.GoIdent.GoImportPath)
	r.AddImpl(messageParserName, m.GoIdent.GoImportPath)

	r.g.Unskip()
}

func (r rasGenerator) generateEndpointHelpers(m *protogen.Message, ext MessageExtension) {

	if !ext.GenerateEndpointHelpers {
		return
	}

	g := r.g

	endpointMessageIdent := r.KnownTypes.PacketEndpointMessageType.GoIdent
	endpointName := m.GoIdent.GoName
	endpointImplName := endpointName + "Impl"
	messageFormatterName := endpointMessageIdent.GoName + "Formatter"
	messageParserName := endpointMessageIdent.GoName + "Parser"

	g.P("// Helpers generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT")
	g.P()
	g.P("type ", endpointImplName, " interface {")
	g.P(" GetVersion() int32 ")
	g.P(" GetId() int32 ")
	g.P(" GetService() string ")
	g.P(" GetFormat() int32 ")
	g.P(" NewMessage(data interface{}) (*", endpointMessageIdent, ", error) ")
	g.P(" UnpackMessage(data interface{}, into ", messageParserName, ") error ")
	g.P("}")
	g.P()
	g.P("func New", m.GoIdent.GoName, "(id int32, version int32) ", endpointImplName, "{")
	g.P("return &", m.GoIdent, "{")
	g.P("Service: \"v8.service.Admin.Cluster\",")
	g.P("Version: version,")
	g.P("Id: id,")
	g.P("Format: ", codecVersion, "()}")
	g.P("}")
	g.P()
	g.P("func (x *", m.GoIdent, ") NewMessage(data interface{}) (*", endpointMessageIdent, ", error) {")
	g.P("switch typed := data.(type) {")
	g.P("case ", ioReader, ":")
	g.P("packet, err := New", r.idxMessage["Packet"].GoIdent, "(data)")
	g.P("if err != nil { return nil, err }")
	g.P("var message ", endpointMessageIdent, "")
	g.P("if err := packet.Unpack(&message); err != nil { return nil, err }")
	g.P("return &message, nil")
	g.P("case ", messageFormatterName, ": ")
	g.P("return New", endpointMessageIdent.GoName, "(x, typed)")
	g.P("default:")
	g.P("return nil, ", fmtErrorf, "(\"unknown type <%T> to create new message\", typed)")
	g.P("}")
	g.P("}")
	g.P()
	g.P("func (x *", m.GoIdent, ") UnpackMessage(data interface{}, into ", messageParserName, ") error  {")
	g.P("switch typed := data.(type) {")
	g.P("case ", ioPackage.Ident("Reader"), ":")
	g.P("packet, err := New", r.idxMessage["Packet"].GoIdent, "(data)")
	g.P("if err != nil { return err }")
	g.P("var message ", endpointMessageIdent, "")
	g.P("if err := packet.Unpack(&message); err != nil { return err }")
	g.P("return message.Unpack(x, into)")
	g.P("case ", r.idxMessage["Packet"].GoIdent, ":")
	g.P("var message ", endpointMessageIdent, "")
	g.P("if err := typed.Unpack(&message); err != nil { return err }")
	g.P("return message.Unpack(x, into)")
	g.P("case *", endpointMessageIdent, ": ")
	g.P("return typed.Unpack(x, into)")
	g.P("default:")
	g.P("return ", fmtErrorf, "(\"unknown type <%T> to create unpack message\", typed)")
	g.P("}")
	g.P("}")
	g.P()
	g.Unskip()
}

func (r rasGenerator) generateNegotiateHelpers(m *protogen.Message) {

	r.g.P("const (")
	r.g.P("Magic int32 = 475223888")
	r.g.P("ProtocolVersion int32 = 256")
	r.g.P(")")
	r.g.P()
	r.g.P("func New", m.GoIdent.GoName, "() *", m.GoIdent, "{")
	r.g.P("return &", m.GoIdent, "{")
	r.g.P("Magic: Magic,")
	r.g.P("Protocol: ProtocolVersion,")
	r.g.P("Version:", codecVersion, "(),")
	r.g.P("}")
	r.g.P("}")
	r.g.P()
}
func (r rasGenerator) generatePacketHelpers(m *protogen.Message, ext MessageExtension) {

	if !ext.GeneratePacketHelpers {
		return
	}

	g := r.g
	packetTypeIdent := r.KnownTypes.EnumPacketType.GoIdent
	messageName := m.GoIdent.GoName + "Message"
	messageFormatterName := messageName + "Formatter"
	messageParserName := messageName + "Parser"

	g.P("// Helpers generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT")
	g.P("type ", messageName, " interface {")
	g.P(" ", messageFormatterName)
	g.P(" ", messageParserName)
	g.P("}")
	g.P()
	g.P("type ", messageFormatterName, " interface {")
	g.P(" GetPacketType() ", g.QualifiedGoIdent(packetTypeIdent))
	g.P(" ", r.formatFuncName, "(writer io.Writer, version int32) error")
	g.P("}")
	g.P()
	g.P("type ", messageParserName, " interface {")
	g.P(" GetPacketType() ", g.QualifiedGoIdent(packetTypeIdent))
	g.P(" ", r.parseFuncName, "(reader io.Reader, version int32) error")
	g.P("}")
	g.P()
	g.Annotate(m.GoIdent.GoName, m.Location)
	g.P("func New", m.GoIdent.GoName, "(message interface{}) (*", m.GoIdent, ", error){")
	g.P("var packet ", m.GoIdent, "")
	g.P("switch typed := message.(type) {")
	g.P("case ", ioReader, ":")
	g.P("if err := packet.", r.parseFuncName, "(typed, 0); err != nil { return nil, err }")
	g.P("case ", messageFormatterName, ": ")
	g.P("buf := &", bytesBuffer, "{}")
	g.P("if err := typed.", r.formatFuncName, "(buf, 0); err != nil { return nil, err }")
	g.P("packet.Type = typed.GetPacketType()")
	g.P("packet.Data = buf.Bytes()")
	g.P("packet.Size = int32(len(packet.Data))")
	g.P("default:")
	g.P("return nil, ", fmtErrorf, "(\"unknown type <%T> to get new packet\", typed)")
	g.P("}")
	g.P("return &packet,  nil")
	g.P("}")
	g.P()
	g.Annotate(m.GoIdent.GoName, m.Location)
	g.P("func (x *", m.GoIdent, ") Unpack(into ", messageParserName, ") error {")
	g.P("switch x.GetType() {")
	g.P("case into.GetPacketType():")
	g.P("buf := ", bytesNewBuffer, "(x.Data)")
	g.P("return into.Parse(buf, 0)")
	g.P("default:")
	g.P("if _, err := x.UnpackNew(); err != nil { return err}")
	g.P("return ", fmtErrorf, "(\"unpack type no equal packet type. Has %s want %s\", x.GetType(), into.GetPacketType())")
	g.P("}")
	g.P("}")
	g.P()
	g.Annotate(m.GoIdent.GoName, m.Location)
	g.P("func (x *", m.GoIdent, ") UnpackNew() (interface{}, error) {")
	g.P("var into interface{}")
	g.P("switch x.GetType() {")

	for enumName, objectValue := range r.idxMessageByEnumValue {
		if strings.HasPrefix(enumName, "PACKET_TYPE") {
			enumIdent := r.EnumNamed(enumName)
			g.P("// type ", enumIdent.GoIdent, " cast ", objectValue.GoIdent)
			g.P("case ", enumIdent.GoIdent, " :")
			g.P("into = &", objectValue.GoIdent, "{}")
		}
	}

	g.P("default: ")
	g.P("return nil, ", fmtErrorf, "(\"unknown unpack type %s\",  x.GetType())")
	g.P("}")
	g.P("buf := ", bytesNewBuffer, "(x.Data)")
	g.P("parser := into.(", messageParserName, ")")
	g.P("if err := parser.Parse(buf, 0); err != nil { return nil, err }")
	g.P("if err, ok := into.(error); ok { return nil, err }")
	g.P("return into, nil")
	g.P("}")
	g.P()
	g.Unskip()
}

func (r rasGenerator) genFormatterFunc(m *protogen.Message, ext MessageExtension) {

	fields := getMessageFields(m)

	if len(fields) == 0 && ext.GenerateEmpty {
		r.g.P("func (x *", m.GoIdent, ") ", r.formatFuncName, " (writer ", ioPackage.Ident("Writer"), ", version int32) error {")
		r.g.P("return nil ")
		r.g.P("}")
		r.g.Unskip()
	}

	if len(fields) == 0 {
		return
	}

	r.g.Unskip()

	r.g.P("func (x *", m.GoIdent, ") ", r.formatFuncName, " (writer ", ioPackage.Ident("Writer"), ", version int32) error {")
	r.g.P("if x == nil { return nil }")

	cVersion := int32(0)
	brackets := 0
	fields.Range(func(f field) bool {

		fVersion := f.Opts.GetVersion()

		if !(cVersion == fVersion) {

			if brackets > 0 {
				brackets--
				r.g.P("}")
			}
			r.g.P("if version >= ", fVersion, "{")
			brackets++
			cVersion = fVersion

		}

		r.generateFieldFormatter(m, f)

		return true
	})

	if brackets > 0 {
		brackets--
		r.g.P("}")
	}

	r.g.P("return nil")
	r.g.P("}")
}

func (r rasGenerator) genBytesFormatter(m *protogen.Message, f field, identifier string) {

	r.g.P("if err:= ", formatBytes, "(writer,", identifier, "); err != nil {")
	r.g.P("return err")
	r.g.P("}")

}

func (r rasGenerator) generateWriteToFunc(message *protogen.Message) {

	r.g.P("func (x *", message.GoIdent, ") WriteTo (w ", ioPackage.Ident("Writer"), ") (int64, error) {")
	r.g.P("buf := &", bytesPackage.Ident("Buffer"), "{}")
	r.g.P("if err := x.", r.formatFuncName, "(buf, 0); err != nil { return 0, err }")
	r.g.P("n, err := w.Write(buf.Bytes())")
	r.g.P("return int64(n), err")
	r.g.P("}")
	r.g.Unskip()

}

func (r rasGenerator) generateErrorFunc(message *protogen.Message) {

	r.g.P("func (x *", message.GoIdent, ") Error() string {")
	r.g.P("return x.String() ")
	r.g.P("}")
	r.g.P()
	r.g.Unskip()

}

func (r rasGenerator) genEnumFormatter(m *protogen.Message, f field, identifier string) {

	r.g.P("if err:= ", f.Encoder, "(writer, int32(", identifier, ")); err != nil {")
	r.g.P("return err")
	r.g.P("}")

}

func (r rasGenerator) genListFormatter(m *protogen.Message, f field, identifier string) {

	r.g.P("if err:= ", formatSize, "(writer, len(", identifier, ")); err != nil {")
	r.g.P("return err")
	r.g.P("}")

	r.g.P("for i := 0; i < len(", identifier, "); i++ {")

	r.generateValueFormatter(m, f, identifier+"[i]", false)

	r.g.P("}")

}

func (r rasGenerator) generateFieldFormatter(m *protogen.Message, f field) {

	identifier := "x." + f.GoName

	r.g.P("// decode ", identifier, " opts: "+f.Opts.String())

	switch {
	case f.Descriptor.IsList():
		r.genListFormatter(m, f, identifier)
	case f.Oneof != nil:
		r.generateOneOfFormatter(m, f, f.GoName)
	case f.Descriptor.IsMap():
		r.generateMapFormatter(m, f, identifier)
	default:
		r.generateValueFormatter(m, f, identifier, false)
	}
}

func (r rasGenerator) generateValueFormatter(m *protogen.Message, f field, identifier string, createVal bool) {

	switch f.Descriptor.Kind() {
	case protoreflect.BytesKind:
		r.genBytesFormatter(m, f, identifier)
	case protoreflect.EnumKind:
		r.genEnumFormatter(m, f, identifier)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		r.generateMessageFormatter(m, f, identifier, createVal)
	default:
		r.generateSingularFormatter(f, identifier, createVal)
	}
}

func (r rasGenerator) generateMapFormatter(m *protogen.Message, f field, identifier string) {

	r.g.P("if err:= ", formatSize, "(writer, len(", identifier, ")); err != nil { return err }")

}

func (r rasGenerator) generateMessageFormatter(m *protogen.Message, f field, identifier string, createVal bool) {

	if genFunc := r.wellKnownTypeFormatter(f.Descriptor.Message().FullName()); genFunc != nil {
		genFunc(r.g, m, f, identifier, createVal)
		return
	}
	r.g.P("if err:= ", identifier, ".", r.formatFuncName, "(writer, version)", "; err != nil { return err }")

}

//
func (r rasGenerator) generateSingularFormatter(f field, identifier string, createVal bool) {

	F := func(encoder protogen.GoIdent, goType string) {
		r.g.P("if err:= ", encoder, "(writer, ", identifier, "); err != nil { return err }")
	}

	encoder := f.Encoder

	switch f.Descriptor.Kind() {
	case protoreflect.BoolKind:
		F(encoder, "bool")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		F(encoder, "int32")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		F(encoder, "int64")
	case protoreflect.FloatKind:
		F(encoder, "float32")
	case protoreflect.DoubleKind:
		F(encoder, "float64")
	case protoreflect.StringKind:
		F(encoder, "string")
	case protoreflect.BytesKind:
		F(encoder, "bytes")
	default:
		r.Fail("Unknown type encoder")
	}

}

func (r rasGenerator) generateOneOfFormatter(m *protogen.Message, f field, identifier string) {

	r.g.P("if val := x.Get", identifier, "(); val != nil {")
	r.generateMessageFormatter(m, f, "val", false)
	r.g.P("}")
}

func (r rasGenerator) genParseFunc(m *protogen.Message, ext MessageExtension) {

	fields := getMessageFields(m)

	if ext.GenerateEmpty && len(fields) == 0 {
		r.g.P("func (x *", m.GoIdent, ") ", r.parseFuncName, " (reader ", ioPackage.Ident("Reader"), ", version int32) error {")
		r.g.P("return nil ")
		r.g.P("}")
		r.g.Unskip()
	}

	if len(fields) == 0 {
		return
	}

	r.g.Unskip()

	r.g.P("func (x *", m.GoIdent, ") ", r.parseFuncName, " (reader ", ioPackage.Ident("Reader"), ", version int32) error {")
	r.g.P("if x == nil { return nil }")

	cVersion := int32(0)
	brackets := 0
	fields.Range(func(f field) bool {

		fVersion := f.Opts.GetVersion()

		if !(cVersion == fVersion) {

			if brackets > 0 {
				brackets--
				r.g.P("}")
			}
			r.g.P("if version >= ", fVersion, "{")
			brackets++
			cVersion = fVersion

		}

		r.generateFieldParser(m, f)

		return true
	})

	if brackets > 0 {
		brackets--
		r.g.P("}")
	}

	r.g.P("return nil")
	r.g.P("}")
}

func (r rasGenerator) genBytesParser(m *protogen.Message, f field, identifier string) {

	if f.Opts.GetSizeField() != 0 {
		r.g.P(identifier, " = make([]byte, x.Get", getMessageFields(m).FindByNumber(f.Opts.GetSizeField()).GoName, "())")
		r.g.P("if err:= ", parseBytes, "(reader,", identifier, "); err != nil {")
		r.g.P("return err")
		r.g.P("}")
		return
	}
	r.g.P("var err error")
	r.g.P(identifier, ", err = ", ioPackage.Ident("ReadAll"), "(reader)")
	r.g.P("if err != nil {")
	r.g.P("return err")
	r.g.P("}")

}

func (r rasGenerator) genEnumParser(m *protogen.Message, f field, identifier string) {

	enumName := r.ObjectNamed(string(f.Descriptor.Enum().FullName())).GoIdent

	r.g.P("var val_", f.GoName, " int32")
	r.g.P("if err:= ", f.Decoder, "(reader, ", "&val_", f.GoName, "); err != nil {")
	r.g.P("return err")
	r.g.P("}")
	r.g.P(identifier, " = ", enumName, "(val_", f.GoName, ")")

}

func (r rasGenerator) genListParser(m *protogen.Message, f field, identifier string) {

	sizeFieldName := "size_" + f.GoName

	r.g.P("var ", sizeFieldName, " int")

	if f.Opts.GetSizeField() != 0 {
		r.g.P(sizeFieldName, " = x.Get", getMessageFields(m).FindByNumber(f.Opts.GetSizeField()).GoName, "()")
	} else {
		r.g.P("if err:= ", parseSize, "(reader, &", sizeFieldName, "); err != nil {")
		r.g.P("return err")
		r.g.P("}")
	}

	r.g.P("for i := 0; i <", sizeFieldName, "; i++ {")

	r.generateValueParser(m, f, "val", true)

	r.g.P(identifier, " = append("+identifier, ", val)")
	r.g.P("}")

}

func (r rasGenerator) generateFieldParser(m *protogen.Message, f field) {

	identifier := "x." + f.GoName

	r.g.P("// decode ", identifier, " opts: "+f.Opts.String())

	switch {
	case f.Descriptor.IsList():
		r.genListParser(m, f, identifier)
	case f.Oneof != nil:

		r.generateOneOfParser(m, f, "x."+f.Field.Oneof.GoName)

	case f.Descriptor.IsMap():
		r.generateMapParser(m, f, identifier)
	default:
		r.generateValueParser(m, f, identifier, false)
	}
}

func (r rasGenerator) generateMapParser(m *protogen.Message, f field, identifier string) {

	r.g.P("// TODO parse map")
	// g.P("if err:= ", g.QualifiedGoIdent(formatSize), "(writer, len(", identifier, ")); err != nil { return err }")

}

func (r rasGenerator) generateValueParser(m *protogen.Message, f field, identifier string, createVal bool) {

	switch f.Descriptor.Kind() {
	case protoreflect.BytesKind:
		r.genBytesParser(m, f, identifier)
	case protoreflect.EnumKind:
		r.genEnumParser(m, f, identifier)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		r.generateMessageParser(m, f, identifier, createVal)
	default:
		r.generateSingularParser(f, identifier, createVal)
	}
}

func (r rasGenerator) generateMessageParser(m *protogen.Message, f field, identifier string, createVal bool) {

	if genFunc := r.wellKnownTypeParse(f.Descriptor.Message().FullName()); genFunc != nil {
		genFunc(r.g, m, f, identifier, createVal)
		return
	}

	valueName := string(f.Descriptor.Message().FullName())
	if createVal {
		r.g.P(identifier, " := &", r.ObjectNamed(valueName).GoIdent, "{}")

	} else {
		r.g.P(identifier, " = &", r.ObjectNamed(valueName).GoIdent, "{}")
	}
	r.g.P("if err:= ", identifier, ".", r.parseFuncName, "(reader, version)", "; err != nil { return err }")
	r.g.P()

}

func (r rasGenerator) generateSingularParser(f field, identifier string, createVal bool) {

	F := func(decoder protogen.GoIdent, goType string) {
		if createVal {
			r.g.P("var ", identifier, " "+goType)
		}
		r.g.P("if err:= ", decoder, "(reader, &", identifier, "); err != nil { return err }")
	}

	decoder := f.Decoder

	switch f.Descriptor.Kind() {
	case protoreflect.BoolKind:
		F(decoder, "bool")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		F(decoder, "int32")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		F(decoder, "int64")
	case protoreflect.FloatKind:
		F(decoder, "float32")
	case protoreflect.DoubleKind:
		F(decoder, "float64")
	case protoreflect.StringKind:
		F(decoder, "string")
	default:
		r.Fail("Unknown type decoder")
	}

}

func (r rasGenerator) generateOneOfParser(m *protogen.Message, f field, identifier string) {

	if f.Opts.GetTypeField() == 0 {
		return
	}

	typeField := getMessageFields(m).FindByNumber(f.Opts.GetTypeField())

	ext := GetMessageExtension(f.Descriptor.Message().Options())
	var enumValue *protogen.EnumValue

	switch {
	case len(ext.PacketType) > 0:
		enumValue = r.idxEnumValues[ext.PacketType]
	case len(ext.EndpointDataType) > 0:
		enumValue = r.idxEnumValues[ext.EndpointDataType]
	case len(ext.MessageType) > 0:
		enumValue = r.idxEnumValues[ext.MessageType]
	}

	valueName := string(f.Descriptor.Name())
	r.g.P("if x.Get", typeField.GoName, "() == ", enumValue.GoIdent, " {")
	r.generateMessageParser(m, f, valueName, true)

	r.g.P(identifier, " = &", f.Oneof.GoIdent, "{")
	r.g.P(f.Oneof.GoName, ": ", valueName, ",")
	r.g.P("}")

	r.g.P("}")

}
