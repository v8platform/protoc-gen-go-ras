package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (gen *Generator) genFormatter(g *protogen.GeneratedFile, m *protogen.Message, ext *MessageExtension) {

	fields := getMessageFields(m)

	if len(fields) == 0 {
		if ext != nil && ext.GetGenerateEmpty() {
			g.P("func (x *", m.GoIdent, ") ", gen.formatFuncName, " (writer io.Writer, version int32) error {")
			g.P("return nil ")
			g.P("}")
			g.Unskip()
		}
		return
	}

	if ext != nil && ext.GetGenerateIoWriteTo() {
		g.P("func (x *", m.GoIdent, ") WriteTo (w io.Writer) (int64, error) {")
		g.P("buf := &", bytesBuffer, "{}")
		g.P("if err := x.", gen.formatFuncName, "(buf, 0); err != nil { return 0, err }")
		g.P("n, err := w.Write(buf.Bytes())")
		g.P("return int64(n), err")
		g.P("}")
		g.Unskip()
	}

	g.Unskip()

	g.P("func (x *", m.GoIdent, ") ", gen.formatFuncName, " (writer io.Writer, version int32) error {")
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

		gen.generateFieldFormatter(g, m, f)

		return true
	})

	if brackets > 0 {
		brackets--
		g.P("}")
	}

	g.P("return nil")
	g.P("}")
}

func (gen *Generator) genBytesFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	g.P("if err:= ", g.QualifiedGoIdent(formatBytes), "(writer,", identifier, "); err != nil {")
	g.P("return err")
	g.P("}")

}

func (gen *Generator) genEnumFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	g.P("if err:= ", g.QualifiedGoIdent(f.Encoder), "(writer, int32(", identifier, ")); err != nil {")
	g.P("return err")
	g.P("}")

}

func (gen *Generator) genListFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	g.P("if err:= ", g.QualifiedGoIdent(formatSize), "(writer, len(", identifier, ")); err != nil {")
	g.P("return err")
	g.P("}")

	g.P("for i := 0; i < len(", identifier, "); i++ {")

	gen.generateValueFormatter(g, m, f, identifier+"[i]", false)

	g.P("}")

}

func (gen *Generator) generateFieldFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field) {

	identifier := "x." + f.GoName

	g.P("// decode ", identifier, " opts: "+f.Opts.String())

	switch {
	case f.Descriptor.IsList():
		gen.genListFormatter(g, m, f, identifier)
	case f.Oneof != nil:

		gen.generateOneOfFormatter(g, m, f, f.GoName)

	case f.Descriptor.IsMap():
		gen.generateMapFormatter(g, m, f, identifier)
	default:
		gen.generateValueFormatter(g, m, f, identifier, false)
	}
}

func (gen *Generator) generateValueFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	switch f.Descriptor.Kind() {
	case protoreflect.BytesKind:
		gen.genBytesFormatter(g, m, f, identifier)
	case protoreflect.EnumKind:
		gen.genEnumFormatter(g, m, f, identifier)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		gen.generateMessageFormatter(g, m, f, identifier, createVal)
	default:
		gen.generateSingularFormatter(g, f, identifier, createVal)
	}
}

func (gen *Generator) generateMapFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	g.P("if err:= ", g.QualifiedGoIdent(formatSize), "(writer, len(", identifier, ")); err != nil { return err }")

}

func (gen *Generator) generateMessageFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string, createVal bool) {

	if genFunc := gen.wellKnownTypeFormatter(f.Descriptor.Message().FullName()); genFunc != nil {
		genFunc(g, m, f, identifier, createVal)
		return
	}
	g.P("if err:= ", identifier, ".", gen.formatFuncName, "(writer, version)", "; err != nil { return err }")

}

//
func (gen *Generator) generateSingularFormatter(g *protogen.GeneratedFile, f field, identifier string, createVal bool) {

	F := func(s protogen.GoIdent, goType string) {
		g.P("if err:= ", g.QualifiedGoIdent(s), "(writer, ", identifier, "); err != nil { return err }")
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
	default:
		gen.Fail("Unknown type encoder")
	}

}

func (gen *Generator) generateOneOfFormatter(g *protogen.GeneratedFile, m *protogen.Message, f field, identifier string) {

	//g.P("// TODO oneof formatter")
	g.P("if val := x.Get", identifier, "(); val != nil {")
	gen.generateMessageFormatter(g, m, f, "val", false)
	g.P("}")

}
