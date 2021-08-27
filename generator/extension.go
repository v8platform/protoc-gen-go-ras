package generator

import (
	"github.com/v8platfotm/protoc-gen-go-ras/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageFieldExtension struct {
	*plugin.EncodingFieldOptions
}

func GetMessageExtensionFor(m proto.Message) *MessageFieldExtension {

	opts := m.(*descriptorpb.FieldOptions)
	if opts == nil || !proto.HasExtension(opts, plugin.E_Field) {
		return nil
	}

	ext := proto.GetExtension(opts, plugin.E_Field).(*plugin.EncodingFieldOptions)

	return &MessageFieldExtension{ext}
}
