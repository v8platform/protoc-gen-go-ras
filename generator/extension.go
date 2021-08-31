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
	*plugin.EncodingMessageOptions

	//PacketType       string
	//MessageType      string
	//EndpointDataType string
	//GenerateEmpty bool
}

func (e *MessageExtension) GetTypeOption(gen *Generator) *protogen.EnumValue {

	if e.EncodingMessageOptions == nil {
		return nil
	}

	switch e.Type.(type) {
	case *plugin.EncodingMessageOptions_PacketType:
		return gen.idxEnumValues[e.GetPacketType()]
	case *plugin.EncodingMessageOptions_EndpointDataType:
		return gen.idxEnumValues[e.GetEndpointDataType()]
	case *plugin.EncodingMessageOptions_MessageType:
		return gen.idxEnumValues[e.GetMessageType()]
	default:
		//log.Fatalln("Error type option", e)
	}
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

func GetMessageExtensionFor(m proto.Message) *MessageExtension {

	opts, _ := m.(*descriptorpb.MessageOptions)

	if opts == nil || !proto.HasExtension(opts, plugin.E_Options) {
		return nil
	}

	ext := proto.GetExtension(opts, plugin.E_Options).(*plugin.EncodingMessageOptions)

	return &MessageExtension{ext}
}

func GetEnumExtensionFor(m proto.Message) *EnumExtension {

	opts := m.(*descriptorpb.EnumOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_MessageOption) {
		return nil
	}

	messageOption := proto.GetExtension(opts, plugin.E_MessageOption).(string)

	return &EnumExtension{messageOption}
}

func GetIsClientExtensionFor(m proto.Message) bool {

	opts := m.(*descriptorpb.ServiceOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Client) {
		return false
	}

	ext := proto.GetExtension(opts, plugin.E_Client).(*plugin.ClientOptions)

	return ext.IsClient
}

type ClientMethodOptions struct {
	NoPacketPack bool
}

func GetClientMethodExtensionFor(m proto.Message) ClientMethodOptions {

	opts := m.(*descriptorpb.MethodOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Method) {
		return ClientMethodOptions{}
	}

	ext := proto.GetExtension(opts, plugin.E_Method).(*plugin.ClientMethodOptions)

	return ClientMethodOptions{
		NoPacketPack: ext.GetNoPacketPack(),
	}
}
