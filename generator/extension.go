package generator

import (
	"github.com/v8platform/protoc-gen-go-ras/plugin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageFieldExtension struct {
	*plugin.EncodingFieldOptions
}

type MessageExtension struct {
	GenerateEmpty                  bool
	GeneratePacketHelpers          bool
	GenerateEndpointMessageHelpers bool
	GenerateMessageHelpers         bool
	PacketType                     string
	EndpointDataType               string
	MessageType                    string
	GenerateErrorFn                bool
	GenerateEndpointHelpers        bool
	GenerateIoWriteTo              bool
	IsNegotiate                    bool
}

func (e *MessageExtension) GetTypeOption(gen *Generator) *protogen.EnumValue {

	return nil
}

type EnumExtension struct {
	MessageOption string
}

func GetMessageFieldExtensionFor(m proto.Message) *MessageFieldExtension {

	opts := m.(*descriptorpb.FieldOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Field) {
		return nil
	}

	ext := proto.GetExtension(opts, plugin.E_Field).(*plugin.EncodingFieldOptions)

	return &MessageFieldExtension{ext}
}

func GetMessageExtension(m proto.Message) MessageExtension {

	opts, _ := m.(*descriptorpb.MessageOptions)

	if opts == nil || !proto.HasExtension(opts, plugin.E_Options) {
		return MessageExtension{}
	}

	ext := proto.GetExtension(opts, plugin.E_Options).(*plugin.EncodingMessageOptions)

	return MessageExtension{
		GenerateEmpty:                  ext.GetGenerateEmpty(),
		GeneratePacketHelpers:          ext.GetGeneratePacketHelpers(),
		GenerateEndpointMessageHelpers: ext.GetGenerateEndpointMessageHelpers(),
		GenerateEndpointHelpers:        ext.GetGenerateEndpointHelpers(),
		GenerateMessageHelpers:         ext.GetGenerateMessageHelpers(),
		PacketType:                     ext.GetPacketType(),
		EndpointDataType:               ext.GetEndpointDataType(),
		MessageType:                    ext.GetMessageType(),
		GenerateErrorFn:                ext.GetGenerateErrorFn(),
		GenerateIoWriteTo:              ext.GetGenerateIoWriteTo(),
		IsNegotiate:                    ext.GetIsNegotiate(),
	}
}

func GetEnumExtension(m proto.Message) *EnumExtension {

	opts := m.(*descriptorpb.EnumOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_MessageOption) {
		return nil
	}

	messageOption := proto.GetExtension(opts, plugin.E_MessageOption).(string)

	return &EnumExtension{messageOption}
}

func GetIsClientExtension(m proto.Message) bool {

	opts := m.(*descriptorpb.ServiceOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Client) {
		return false
	}

	ext := proto.GetExtension(opts, plugin.E_Client).(*plugin.ClientOptions)

	return ext.GetIsClient()
}
func GetIsEndpointExtension(m proto.Message) bool {

	opts := m.(*descriptorpb.ServiceOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Client) {
		return false
	}

	ext := proto.GetExtension(opts, plugin.E_Client).(*plugin.ClientOptions)

	return ext.GetIsEndpoint()
}

func GetIsRequestServiceExtension(m proto.Message) bool {

	opts := m.(*descriptorpb.ServiceOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Client) {
		return false
	}

	ext := proto.GetExtension(opts, plugin.E_Client).(*plugin.ClientOptions)
	return ext.GetIsRequestService()
}

func GetIsRASServiceExtension(m proto.Message) bool {

	opts := m.(*descriptorpb.ServiceOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Client) {
		return false
	}

	ext := proto.GetExtension(opts, plugin.E_Client).(*plugin.ClientOptions)
	return ext.GetIsRasService()
}

type ClientMethodOptions struct {
	MethodParams    map[string]string
	IgnoreEmpty     bool
	NoPacketPack    bool
	NewEndpointFunc bool
	ProxyName       string
}

func GetClientMethodExtension(m proto.Message) ClientMethodOptions {

	opts := m.(*descriptorpb.MethodOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Method) {
		return ClientMethodOptions{}
	}

	ext := proto.GetExtension(opts, plugin.E_Method).(*plugin.ClientMethodOptions)

	return ClientMethodOptions{
		NoPacketPack:    ext.GetNoPacketPack(),
		MethodParams:    ext.GetParam(),
		IgnoreEmpty:     ext.GetIgnoreEmpty(),
		NewEndpointFunc: ext.GetNewEndpointFunc(),
		ProxyName:       ext.GetProxyName(),
	}
}

type ClientMessageOptions struct {
	IsImpl bool
}

func GetClientMessageExtension(m proto.Message) ClientMessageOptions {

	opts := m.(*descriptorpb.MessageOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_ClientOptions) {
		return ClientMessageOptions{}
	}

	ext := proto.GetExtension(opts, plugin.E_ClientOptions).(*plugin.ClientMessageOptions)

	return ClientMessageOptions{
		IsImpl: ext.GetImpl(),
	}
}

type FileImplOptions struct {
	impl []string
}

func GetFileImplExtension(m proto.Message) FileImplOptions {

	opts := m.(*descriptorpb.FileOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Impl) {
		return FileImplOptions{}
	}

	ext := proto.GetExtension(opts, plugin.E_Impl).(*plugin.FileImplOptions)

	return FileImplOptions{
		impl: ext.GetInterface(),
	}
}
