package generator

import "google.golang.org/protobuf/compiler/protogen"

func (gen *Generator) generateErrorHelpers(g *protogen.GeneratedFile, m *protogen.Message) {

	g.P("func (x *", m.GoIdent, ") Error() string {")
	g.P("return x.String() ")
	g.P("}")
	g.P()
	g.Unskip()

}
