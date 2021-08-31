package generator

import (
	"fmt"
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

func (gen *Generator) generateClient(file *protogen.File) {
	// GenerateFile generates a .mock.pb.go file containing gRPC service definitions.
	if !file.Generate {
		return
	}
	if len(file.Services) == 0 {
		return
	}
	// fmt.Println("FILENAME ", file.GeneratedFilenamePrefix)
	filename := file.GeneratedFilenamePrefix + "_client.pb.go"
	g := gen.plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.Skip()
	mockGenerator := clientGenerator{
		Generator: gen,
		gen:       gen.plugin,
		file:      file,
		g:         g,
	}
	mockGenerator.genHeader(string(file.GoPackageName))
	mockGenerator.GenerateFileContent()
	return

}

func (gen *Generator) fill() {

	gen.idxObject = make(map[string]Object)
	gen.idxMessageByEnumValue = make(map[string]*protogen.Message)
	gen.idxEnum = make(map[string]*protogen.Enum)
	gen.idxEnumValues = make(map[string]*protogen.EnumValue)
	gen.idxMessage = make(map[string]*protogen.Message)

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
	}

	for _, message := range gen.idxMessage {
		messageExtension := GetMessageExtensionFor(message.Desc.Options())
		if messageExtension != nil {
			enumValue := messageExtension.GetTypeOption(gen)
			if enumValue != nil {
				gen.idxMessageByEnumValue[string(enumValue.Desc.Name())] = message
			}
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

func (t KnownTypes) EndpointImpl() protogen.GoIdent {
	return protogen.GoIdent{GoName: t.EndpointType.GoIdent.GoName + t.EndpointImplSuffix, GoImportPath: t.EndpointType.GoIdent.GoImportPath}
}

// dottedSlice turns a sliced name into a dotted name.
func dottedSlice(elem []string) string { return strings.Join(elem, ".") }

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

const (
	// contextPkg         = protogen.GoImportPath("context")
	// encoderPackage    = protogen.GoImportPath("github.com/v8platform/encoder/ras")
	// errPkg             = protogen.GoImportPath("errors")

	deprecationComment = "// Deprecated: Do not use."
)

func (gen *Generator) GenerateFile(plugin *protogen.Plugin, file *protogen.File) {

	filename := file.GeneratedFilenamePrefix + "_ras.pb.go"

	// log.Printf("Processing %s", file.na)
	log.Printf("Generating %s\n", fmt.Sprintf("%s_ras.pb.go", file.GeneratedFilenamePrefix))

	g := plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.Skip()
	g.P("// Code generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.")
	// g.P("// ", contextPkg.Ident(""),  errPkg.Ident(""))
	g.P("// ", encoderPackage.Ident(""), ioPackage.Ident(""))
	g.P()

	for _, message := range file.Messages {
		//log.Println(message)
		gen.genMessage(g, message)
	}

	gen.generateClient(file)

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

func (gen *Generator) genMessage(g *protogen.GeneratedFile, m *protogen.Message) {

	ext := GetMessageExtensionFor(m.Desc.Options())

	if ext != nil {

		if ext.GetType() != nil {
			enumValue := ext.GetTypeOption(gen)
			funcName := "Get" + enumValue.Parent.GoIdent.GoName
			g.Annotate(m.GoIdent.GoName, m.Location)
			g.P("func (x *", m.GoIdent, ") ", funcName, "() ", enumValue.Parent.GoIdent, " {")
			g.P("return ", enumValue.GoIdent)
			g.P("}")
			g.P()
			g.Unskip()
		}

		if ext.GetGeneratePacketHelpers() {
			gen.generatePacketHelpers(g, m)
		}

		if ext.GetGenerateEndpointMessageHelpers() {
			gen.generateEndpointMessageHelpers(g, m)
		}
		if ext.GetGenerateEndpointHelpers() {
			gen.generateEndpointHelpers(g, m)
		}

		if ext.GetGenerateErrorFn() {
			gen.generateErrorHelpers(g, m)
		}
	}

	gen.genParse(g, m, ext)
	gen.genFormatter(g, m, ext)
}
