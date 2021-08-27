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
	PacketType       string
	MessageType      string
	EndpointDataType string
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

	opts := m.(*descriptorpb.MessageOptions)

	if opts == nil {
		return nil
	}
	ext := &MessageExtension{}

	if proto.HasExtension(opts, plugin.E_PacketType) {
		packetType := proto.GetExtension(opts, plugin.E_PacketType).(string)
		ext.PacketType = packetType
	}

	if proto.HasExtension(opts, plugin.E_MessageType) {
		messageType := proto.GetExtension(opts, plugin.E_MessageType).(string)
		ext.MessageType = messageType
	}

	if proto.HasExtension(opts, plugin.E_EndpointDataType) {
		endpointDataType := proto.GetExtension(opts, plugin.E_EndpointDataType).(string)
		ext.EndpointDataType = endpointDataType
	}

	return ext
}
