package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"os"
	"strings"
)

// Object is an interface abstracting the abilities shared by enums, messages, extensions and imported objects.
type Object struct {
	protogen.GoIdent
	Desc      protoreflect.Descriptor
	TypeIdent protogen.GoIdent
}

type Generator struct {
	plugin           *protogen.Plugin
	typeNameToObject map[string]Object
	enumNameToObject map[string]Object
	enumToObject     map[string]Object
	parseFuncName    string
	formatFuncName   string

	idxObject             map[string]Object
	idxMessage            map[string]*protogen.Message
	idxImpl               map[string]protogen.GoIdent
	idxMessageByEnumValue map[string]*protogen.Message
	idxEnum               map[string]*protogen.Enum
	idxEnumValues         map[string]*protogen.EnumValue

	KnownTypes KnownTypes
}

type KnownTypes struct {
	EnumPacketType            *protogen.Enum
	EnumEndpointDataType      *protogen.Enum
	EnumMessageType           *protogen.Enum
	PacketEndpointMessageType *protogen.Message

	EndpointType               *protogen.Message
	EndpointImplSuffix         string
	EndpointDataMessageType    *protogen.Message
	EndpointVoidMessageType    *protogen.Message
	EndpointFailureMessageType *protogen.Message
}

type KnownTypesOptions struct {
	EnumPacketTypeName            string
	EnumEndpointDataTypeName      string
	EnumMessageTypeName           string
	EndpointMessagePacketTypeName string

	FormatterImplSuffix string
	ParseImplSuffix     string
	EndpointImplSuffix  string

	EndpointDataMessageTypeName    string
	EndpointVoidMessageTypeName    string
	EndpointFailureMessageTypeName string
}

var defaultKnownTypesOptions = KnownTypesOptions{

	EnumPacketTypeName:       "PacketType",
	EnumEndpointDataTypeName: "EndpointDataType",
	EnumMessageTypeName:      "MessageType",

	EndpointMessagePacketTypeName: "PACKET_TYPE_ENDPOINT_MESSAGE",

	FormatterImplSuffix: "Formatter",
	ParseImplSuffix:     "Parse",
	EndpointImplSuffix:  "Impl",

	EndpointDataMessageTypeName:    "ENDPOINT_DATA_TYPE_MESSAGE",
	EndpointVoidMessageTypeName:    "ENDPOINT_DATA_TYPE_VOID_MESSAGE",
	EndpointFailureMessageTypeName: "ENDPOINT_DATA_TYPE_EXCEPTION",
}

func NewGenerator(plugin *protogen.Plugin) *Generator {

	gen := &Generator{
		plugin:         plugin,
		parseFuncName:  "Parse",
		formatFuncName: "Formatter",
	}
	gen.fill()
	gen.fillKnownTypes(defaultKnownTypesOptions)

	return gen
}

func (gen *Generator) AddImpl(name string, path protogen.GoImportPath) {

	_, ok := gen.idxImpl[name]
	if !ok {
		gen.idxImpl[name] = path.Ident(name)
		return
	}

	// gen.Fail("find impl with name", name, " -> ", ident.String())

}

func (m *Generator) getIdentOrImpl(message *protogen.Message, g *protogen.GeneratedFile, prefix string) string {

	if GetClientMessageExtension(message.Desc.Options()).IsImpl {
		return prefix + g.QualifiedGoIdent(m.GetImpl(string(message.Desc.Name())))
	}

	return prefix + g.QualifiedGoIdent(message.GoIdent)
}

func (gen *Generator) GetImpl(name string) protogen.GoIdent {

	ident, ok := gen.idxImpl[name]
	if !ok {
		gen.Fail("not find impl with name", name)
	}

	return ident

}

func (t KnownTypes) EndpointImpl() protogen.GoIdent {
	return protogen.GoIdent{GoName: t.EndpointType.GoIdent.GoName + t.EndpointImplSuffix, GoImportPath: t.EndpointType.GoIdent.GoImportPath}
}

func (gen *Generator) ObjectNamed(typeName string) Object {

	if !strings.HasPrefix(typeName, ".") {
		typeName = "." + typeName
	}
	o, ok := gen.idxObject[typeName]
	if !ok {
		gen.Fail("can't find object with type", typeName)
	}
	return o
}

func (gen *Generator) EnumNamed(typeName string) *protogen.EnumValue {

	o, ok := gen.idxEnumValues[typeName]
	if !ok {
		gen.Fail("can't find enum value with enum name", typeName)
	}
	return o
}

func (gen *Generator) generateFile(file *protogen.File) {

	if !file.Generate {
		return
	}

	filename := file.GeneratedFilenamePrefix + "_ras.pb.go"
	g := gen.plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.Skip()

	gen.genHeader(g, string(file.GoPackageName))
	gen.generateRas(file, g)
	gen.generateClient(file, g)
	gen.generateEndpoint(file, g)
	gen.generateMessagesService(file, g)

}

func (gen *Generator) Generate() {

	for _, f := range gen.plugin.Files {
		if f.Generate && shouldProcess(f) {
			gen.generateFile(f)
		}
	}

}

func (gen *Generator) GenerateOnce() {

	for _, f := range gen.plugin.Files {
		if f.Generate && shouldProcess(f) {
			gen.generateRasService(f)
		}
	}
}

func shouldProcess(file *protogen.File) bool {
	ignoredFiles := []string{
		"graphql/graphql.proto",
		"graphql.proto",
		"ras/encoding/ras.proto",
		"ras/encoding/client.proto",
		"ras/encoding/file.proto"}
	for _, ignored := range ignoredFiles {
		if *file.Proto.Name == ignored {

			log.Println("ignore file", *file.Proto.Name)
			return false
		}
	}
	// if proto.HasExtension(file.Proto.Options, graphql.E_Disabled) {
	// 	return !proto.GetExtension(file.Proto.Options, graphql.E_Disabled).(bool)
	// }
	return true
}

// Fail reports a problem and exits the program.
func (gen *Generator) Fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("protoc-gen-go-ras: error:", s)
	os.Exit(1)
}

// Error reports a problem, including an error, and exits the program.
func (gen *Generator) Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("protoc-gen-go-ras: error:", s)
	os.Exit(1)
}

func (gen *Generator) generateRasService(file *protogen.File) {

	if len(file.Services) == 0 {
		return
	}

	filename := file.GeneratedFilenamePrefix + "_ras.pb.go"
	g := gen.plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.Skip()

	gen.genHeader(g, string(file.GoPackageName))

	generator := rasServiceGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	generator.init()
	generator.GenerateFileContent()

}

func (gen *Generator) generateRas(file *protogen.File, g *protogen.GeneratedFile) {

	if len(file.Messages) == 0 {
		return
	}
	generator := rasGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	generator.GenerateFileContent()

}
func (gen *Generator) genHeader(g *protogen.GeneratedFile, packageName string) {
	g.P("// Code generated by protoc-gen-go-ras. DO NOT EDIT.")
	g.P()
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.")
	g.P()
	g.P("package ", packageName)
	g.P()

}

func (gen *Generator) generateClient(file *protogen.File, g *protogen.GeneratedFile) {

	if len(file.Services) == 0 {
		return
	}
	generator := clientGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	generator.GenerateFileContent()
}

func (gen *Generator) generateEndpoint(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	generator := endpointGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	generator.GenerateFileContent()
}

func (gen *Generator) generateMessagesService(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	generator := messageServiceGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	generator.GenerateFileContent()
}

func (gen *Generator) fill() {

	gen.idxObject = make(map[string]Object)
	gen.idxMessageByEnumValue = make(map[string]*protogen.Message)
	gen.idxEnum = make(map[string]*protogen.Enum)
	gen.idxEnumValues = make(map[string]*protogen.EnumValue)
	gen.idxMessage = make(map[string]*protogen.Message)
	gen.idxImpl = make(map[string]protogen.GoIdent)

	for _, f := range gen.plugin.Files {
		dottedPkg := "." + string(f.Proto.GetPackage())

		if dottedPkg != "." {
			dottedPkg += "."
		}

		for _, enum := range f.Enums {
			name := dottedPkg + enum.GoIdent.GoName

			gen.idxObject[name] = Object{
				GoIdent: enum.GoIdent,
				Desc:    enum.Desc,
			}

			gen.idxEnum[enum.GoIdent.GoName] = enum

			for _, value := range enum.Values {
				gen.idxEnumValues[string(value.Desc.Name())] = value
			}
		}
		for _, message := range f.Messages {
			name := dottedPkg + message.GoIdent.GoName
			gen.idxObject[name] = Object{
				GoIdent: message.GoIdent,
				Desc:    message.Desc,
			}

			gen.idxMessage[string(message.Desc.Name())] = message
		}

		extImpl := GetFileImplExtension(f.Desc.Options())
		for _, impl := range extImpl.impl {
			gen.idxImpl[impl] = f.GoImportPath.Ident(impl)
		}
	}

	for _, message := range gen.idxMessage {
		ext := GetMessageExtension(message.Desc.Options())
		switch {
		case len(ext.PacketType) > 0:
			gen.idxMessageByEnumValue[ext.PacketType] = message
		case len(ext.EndpointDataType) > 0:
			gen.idxMessageByEnumValue[ext.EndpointDataType] = message
		case len(ext.MessageType) > 0:
			gen.idxMessageByEnumValue[ext.MessageType] = message
		}
	}

}

func (gen *Generator) fillKnownTypes(options KnownTypesOptions) {

	gen.KnownTypes = KnownTypes{
		EnumPacketType:       gen.idxEnum[options.EnumPacketTypeName],
		EnumEndpointDataType: gen.idxEnum[options.EnumEndpointDataTypeName],
		EnumMessageType:      gen.idxEnum[options.EnumMessageTypeName],

		EndpointType:               gen.idxMessage["Endpoint"],
		EndpointImplSuffix:         options.EndpointImplSuffix,
		EndpointDataMessageType:    gen.idxMessageByEnumValue[options.EndpointDataMessageTypeName],
		EndpointVoidMessageType:    gen.idxMessageByEnumValue[options.EndpointVoidMessageTypeName],
		EndpointFailureMessageType: gen.idxMessageByEnumValue[options.EndpointFailureMessageTypeName],
		PacketEndpointMessageType:  gen.idxMessageByEnumValue[options.EndpointMessagePacketTypeName],
	}

}
