package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (gen *Generator) genParse(g *protogen.GeneratedFile, m *protogen.Message, ext *MessageExtension) {

	fields := getMessageFields(m)

	if len(fields) == 0 {
		if ext != nil && ext.GetGenerateEmpty() {
			g.P("func (x *", m.GoIdent, ") ", gen.parseFuncName, " (reader io.Reader, version int32) error {")
			g.P("return nil ")
			g.P("}")
			g.Unskip()
		}
		return
	}

	g.Unskip()

	g.P("func (x *", m.GoIdent, ") ", gen.parseFuncName, " (reader io.Reader, version int32) error {")
	g.P("if x == nil { return nil }")

	cVersion := int32(0)
	brackets := 0
	fields.Range(func(f field) bool {

		fVersion := f.Opts.GetVersion()

		if !(cVersion == fVersion) {

			if brackets > 0 {
				brackets--
				g.P("}")
			}
			g.P("if version >= ", fVersion, "{")
			brackets++
			cVersion = fVersion

		}

		gen.generateFieldParser(g, m, f)

		return true
	})

	if brackets > 0 {
		brackets--
		g.P("}")
	}

	g.P("return nil")
	g.P("}")
}

func (gen *Generator) genBytesParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	if f.Opts.GetSizeField() != 0 {
		g.P(identifier, " = make([]byte, x.Get", getMessageFields(m).FindByNumber(f.Opts.GetSizeField()).GoName, "())")
	}

	g.P("if err:= ", g.QualifiedGoIdent(parseBytes), "(reader,", identifier, "); err != nil {")
	g.P("return err")
	g.P("}")

}

func (gen *Generator) genEnumParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	enumName := gen.ObjectNamed(string(f.Descriptor.Enum().FullName())).GoIdent

	g.P("var val_", f.GoName, " int32")
	g.P("if err:= ", g.QualifiedGoIdent(f.Decoder), "(reader, ", "&val_", f.GoName, "); err != nil {")
	g.P("return err")
	g.P("}")
	g.P(identifier, " = ", enumName, "(val_", f.GoName, ")")

}

func (gen *Generator) genListParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	sizeFieldName := "size_" + f.GoName

	g.P("var ", sizeFieldName, " int")

	if f.Opts.GetSizeField() != 0 {
		g.P(sizeFieldName, " = x.Get", getMessageFields(m).FindByNumber(f.Opts.GetSizeField()).GoName, "()")
	} else {
		g.P("if err:= ", g.QualifiedGoIdent(parseSize), "(reader, &", sizeFieldName, "); err != nil {")
		g.P("return err")
		g.P("}")
	}

	g.P("for i := 0; i <", sizeFieldName, "; i++ {")

	gen.generateValueParser(g, m, f, "val", true)

	g.P(identifier, " = append("+identifier, ", val)")
	g.P("}")

}

func (gen *Generator) generateFieldParser(g *protogen.GeneratedFile, m *protogen.Message, f field) {

	identifier := "x." + f.GoName

	g.P("// decode ", identifier, " opts: "+f.Opts.String())

	switch {
	case f.Descriptor.IsList():
		gen.genListParser(g, m, f, identifier)
	case f.Oneof != nil:

		gen.generateOneOfParser(g, m, f, "x."+f.Field.Oneof.GoName)

	case f.Descriptor.IsMap():
		g.P("// TODO generate map")
	default:
		gen.generateValueParser(g, m, f, identifier, false)
	}
}

func (gen *Generator) generateValueParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	switch f.Descriptor.Kind() {
	case protoreflect.BytesKind:
		gen.genBytesParser(g, m, f, identifier)
	case protoreflect.EnumKind:
		gen.genEnumParser(g, m, f, identifier)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		gen.generateMessageParser(g, m, f, identifier, createVal)
	default:
		gen.generateSingularParser(g, f, identifier, createVal)
	}
}

func (gen *Generator) generateMessageParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	if genFunc := gen.wellKnownTypeParse(f.Descriptor.Message().FullName()); genFunc != nil {
		genFunc(g, m, f, identifier, createVal)
		return
	}

	valueName := string(f.Descriptor.Message().FullName())
	if createVal {
		g.P(identifier, " := &", gen.ObjectNamed(valueName).GoIdent, "{}")

	} else {
		g.P(identifier, " = &", gen.ObjectNamed(valueName).GoIdent, "{}")
	}
	g.P("if err:= ", identifier, ".", gen.parseFuncName, "(reader, version)", "; err != nil { return err }")
	g.P()

}

//
func (gen *Generator) generateSingularParser(g *protogen.GeneratedFile, f field, identifier string, createVal bool) {

	F := func(s protogen.GoIdent, goType string) {
		if createVal {
			g.P("var ", identifier, " "+goType)
		}
		g.P("if err:= ", g.QualifiedGoIdent(s), "(reader, &", identifier, "); err != nil { return err }")
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
		gen.Fail("Unknown type decoder")
	}

}

func (gen *Generator) generateOneOfParser(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	if f.Opts.GetTypeField() == 0 {
		return
	}

	typeField := getMessageFields(m).FindByNumber(f.Opts.GetTypeField())

	ext := GetMessageExtensionFor(f.Descriptor.Message().Options())
	enumName, _ := ext.GetTypeOption()

	valueName := string(f.Descriptor.Name())
	g.P("if x.Get", typeField.GoName, "() == ", gen.EnumNamed(enumName).GoIdent, " {")
	gen.generateMessageParser(g, m, f, valueName, true)

	g.P(identifier, " = &", f.Oneof.GoIdent, "{")
	g.P(f.Oneof.GoName, ": ", valueName, ",")
	g.P("}")

	g.P("}")

}
