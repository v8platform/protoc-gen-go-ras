package generator

import (
	"bytes"
	"google.golang.org/protobuf/compiler/protogen"
	"html/template"
)

// Object is an interface abstracting the abilities shared by enums, messages, extensions and imported objects.
type MapObject struct {
	protogen.GoIdent
	File    *protogen.File
	Message *protogen.Message
	Type    protogen.GoIdent
}

type GenerateHelpers struct {
	file                  *protogen.File
	idxMessageByEnumValue map[string]*protogen.Message
	idxEnum               map[string]*protogen.Enum
	idxEnumValues         map[string]*protogen.EnumValue

	formatSuffix       string
	parseSuffix        string
	endpointMessage    *protogen.Message
	endpointImplSuffix string

	serviceName string // v8.service.Admin.Cluster

}

func (h GenerateHelpers) EndpointImpl() string {
	return h.endpointMessage.GoIdent.GoName + h.endpointImplSuffix
}
func (h GenerateHelpers) ServiceName() string {
	return "v8.service.Admin.Cluster"
}

type EndpointHelpers struct {
	GenerateHelpers
}

func (h EndpointHelpers) Execute(g *protogen.GeneratedFile) error {

	w := bytes.NewBuffer(nil)

	if err := messageTemplate.Execute(w, tplMessage{
		// File: h.file,
	}); err != nil {
		return err
	}
	_, err := g.Write(w.Bytes())

	return err
}

type tplMessage struct {
}

var (
	messageTemplate = template.Must(template.New("message").Parse(`

type  {{ .EndpointImpl}} interface {
	GetVersion() int32
	GetId() int32
	GetService() string
	GetFormat() int32
}

type EndpointMessageFormatter interface {
	GetMessageType() MessageType
	Formatter(writer io.Writer, version int32) error
}

type EndpointMessageParser interface {
	GetMessageType() MessageType
	Parse(reader io.Reader, version int32) error
}

func NewEndpoint(id int32, version int32) *{{ .Endpoint.GoIdent.GoName}} {
	return &{{ .Endpoint.GoIdent.GoName}} {
		Service: "{{ .ServiceName}}",
		Version: version,
		Id:      id,
		Format:  codec256.Version()}
}

func NewEndpointMessage(endpoint EndpointImpl, message EndpointMessageFormatter) (*EndpointMessage, error) {
	buf := &bytes.Buffer{}
	if err := message.Formatter(buf, endpoint.GetVersion()); err != nil {
		return nil, err
	}
	return &EndpointMessage{
		Type:       EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
		Format:     endpoint.GetFormat(),
		EndpointId: endpoint.GetId(),
		Data: &EndpointMessage_Message{
			Message: &EndpointDataMessage{
				Bytes: buf.Bytes(),
				Type:  message.GetMessageType(),
			},
		},
	}, nil
}

func (x *Endpoint) NewMessage(message EndpointMessageFormatter) (*EndpointMessage, error) {
	return NewEndpointMessage(x, message)
}
`))
)
