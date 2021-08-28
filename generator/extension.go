package generator

import (
	"github.com/v8platform/protoc-gen-go-ras/plugin"
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

func (e *MessageExtension) GetTypeOption() (string, string) {

	if e.EncodingMessageOptions == nil {
		return "", ""
	}

	switch e.Type.(type) {
	case *plugin.EncodingMessageOptions_PacketType:
		return e.GetPacketType(), "packet_type"
	case *plugin.EncodingMessageOptions_EndpointDataType:
		return e.GetEndpointDataType(), "endpoint_data_type"
	case *plugin.EncodingMessageOptions_MessageType:
		return e.GetMessageType(), "message_type"
	default:
		//log.Fatalln("Error type option", e)
	}
	return "", ""
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
