// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Impl bool `protobuf:"varint,1,opt,name=impl,proto3" json:"impl,omitempty"` // Указания на интерфейс
}

func (x *ClientMessageOptions) Reset() {
	*x = ClientMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessageOptions) ProtoMessage() {}

func (x *ClientMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessageOptions.ProtoReflect.Descriptor instead.
func (*ClientMessageOptions) Descriptor() ([]byte, []int) {
	return file_ras_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMessageOptions) GetImpl() bool {
	if x != nil {
		return x.Impl
	}
	return false
}

type EncodingMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenerateEmpty                  bool `protobuf:"varint,1,opt,name=generate_empty,json=generateEmpty,proto3" json:"generate_empty,omitempty"`
	GeneratePacketHelpers          bool `protobuf:"varint,2,opt,name=generate_packet_helpers,json=generatePacketHelpers,proto3" json:"generate_packet_helpers,omitempty"`
	GenerateEndpointMessageHelpers bool `protobuf:"varint,3,opt,name=generate_endpoint_message_helpers,json=generateEndpointMessageHelpers,proto3" json:"generate_endpoint_message_helpers,omitempty"`
	GenerateMessageHelpers         bool `protobuf:"varint,4,opt,name=generate_message_helpers,json=generateMessageHelpers,proto3" json:"generate_message_helpers,omitempty"`
	// Types that are assignable to Type:
	//	*EncodingMessageOptions_PacketType
	//	*EncodingMessageOptions_EndpointDataType
	//	*EncodingMessageOptions_MessageType
	Type                    isEncodingMessageOptions_Type `protobuf_oneof:"type"`
	GenerateErrorFn         bool                          `protobuf:"varint,8,opt,name=generate_error_fn,json=generateErrorFn,proto3" json:"generate_error_fn,omitempty"`
	GenerateEndpointHelpers bool                          `protobuf:"varint,9,opt,name=generate_endpoint_helpers,json=generateEndpointHelpers,proto3" json:"generate_endpoint_helpers,omitempty"`
	GenerateIoWriteTo       bool                          `protobuf:"varint,10,opt,name=generate_io_write_to,json=generateIoWriteTo,proto3" json:"generate_io_write_to,omitempty"`
}

func (x *EncodingMessageOptions) Reset() {
	*x = EncodingMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodingMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodingMessageOptions) ProtoMessage() {}

func (x *EncodingMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodingMessageOptions.ProtoReflect.Descriptor instead.
func (*EncodingMessageOptions) Descriptor() ([]byte, []int) {
	return file_ras_proto_rawDescGZIP(), []int{1}
}

func (x *EncodingMessageOptions) GetGenerateEmpty() bool {
	if x != nil {
		return x.GenerateEmpty
	}
	return false
}

func (x *EncodingMessageOptions) GetGeneratePacketHelpers() bool {
	if x != nil {
		return x.GeneratePacketHelpers
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateEndpointMessageHelpers() bool {
	if x != nil {
		return x.GenerateEndpointMessageHelpers
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateMessageHelpers() bool {
	if x != nil {
		return x.GenerateMessageHelpers
	}
	return false
}

func (m *EncodingMessageOptions) GetType() isEncodingMessageOptions_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *EncodingMessageOptions) GetPacketType() string {
	if x, ok := x.GetType().(*EncodingMessageOptions_PacketType); ok {
		return x.PacketType
	}
	return ""
}

func (x *EncodingMessageOptions) GetEndpointDataType() string {
	if x, ok := x.GetType().(*EncodingMessageOptions_EndpointDataType); ok {
		return x.EndpointDataType
	}
	return ""
}

func (x *EncodingMessageOptions) GetMessageType() string {
	if x, ok := x.GetType().(*EncodingMessageOptions_MessageType); ok {
		return x.MessageType
	}
	return ""
}

func (x *EncodingMessageOptions) GetGenerateErrorFn() bool {
	if x != nil {
		return x.GenerateErrorFn
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateEndpointHelpers() bool {
	if x != nil {
		return x.GenerateEndpointHelpers
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateIoWriteTo() bool {
	if x != nil {
		return x.GenerateIoWriteTo
	}
	return false
}

type isEncodingMessageOptions_Type interface {
	isEncodingMessageOptions_Type()
}

type EncodingMessageOptions_PacketType struct {
	PacketType string `protobuf:"bytes,5,opt,name=packet_type,json=packetType,proto3,oneof"`
}

type EncodingMessageOptions_EndpointDataType struct {
	EndpointDataType string `protobuf:"bytes,6,opt,name=endpoint_data_type,json=endpointDataType,proto3,oneof"`
}

type EncodingMessageOptions_MessageType struct {
	MessageType string `protobuf:"bytes,7,opt,name=message_type,json=messageType,proto3,oneof"`
}

func (*EncodingMessageOptions_PacketType) isEncodingMessageOptions_Type() {}

func (*EncodingMessageOptions_EndpointDataType) isEncodingMessageOptions_Type() {}

func (*EncodingMessageOptions_MessageType) isEncodingMessageOptions_Type() {}

type EncodingFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoder   *string `protobuf:"bytes,1,opt,name=encoder,proto3,oneof" json:"encoder,omitempty"`
	Order     *int32  `protobuf:"varint,2,opt,name=order,proto3,oneof" json:"order,omitempty"`
	Version   *int32  `protobuf:"varint,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
	TypeField *int32  `protobuf:"varint,7,opt,name=type_field,json=typeField,proto3,oneof" json:"type_field,omitempty"`
	TypeUrl   string  `protobuf:"bytes,8,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Ignore    bool    `protobuf:"varint,9,opt,name=ignore,proto3" json:"ignore,omitempty"`
	SizeField *int32  `protobuf:"varint,10,opt,name=size_field,json=sizeField,proto3,oneof" json:"size_field,omitempty"`
}

func (x *EncodingFieldOptions) Reset() {
	*x = EncodingFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodingFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodingFieldOptions) ProtoMessage() {}

func (x *EncodingFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodingFieldOptions.ProtoReflect.Descriptor instead.
func (*EncodingFieldOptions) Descriptor() ([]byte, []int) {
	return file_ras_proto_rawDescGZIP(), []int{2}
}

func (x *EncodingFieldOptions) GetEncoder() string {
	if x != nil && x.Encoder != nil {
		return *x.Encoder
	}
	return ""
}

func (x *EncodingFieldOptions) GetOrder() int32 {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return 0
}

func (x *EncodingFieldOptions) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *EncodingFieldOptions) GetTypeField() int32 {
	if x != nil && x.TypeField != nil {
		return *x.TypeField
	}
	return 0
}

func (x *EncodingFieldOptions) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *EncodingFieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *EncodingFieldOptions) GetSizeField() int32 {
	if x != nil && x.SizeField != nil {
		return *x.SizeField
	}
	return 0
}

var file_ras_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*EncodingFieldOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.field",
		Tag:           "bytes,475223888,opt,name=field",
		Filename:      "ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*EncodingFieldOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.oneof_field",
		Tag:           "bytes,475223888,opt,name=oneof_field",
		Filename:      "ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*EncodingMessageOptions)(nil),
		Field:         475223889,
		Name:          "ras.encoding.options",
		Tag:           "bytes,475223889,opt,name=options",
		Filename:      "ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ClientMessageOptions)(nil),
		Field:         475223890,
		Name:          "ras.encoding.client_options",
		Tag:           "bytes,475223890,opt,name=client_options",
		Filename:      "ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         475223891,
		Name:          "ras.encoding.message_option",
		Tag:           "bytes,475223891,opt,name=message_option",
		Filename:      "ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         475223891,
		Name:          "ras.encoding.type_url",
		Tag:           "bytes,475223891,opt,name=type_url",
		Filename:      "ras.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional ras.encoding.EncodingFieldOptions field = 475223888;
	E_Field = &file_ras_proto_extTypes[0]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional ras.encoding.EncodingFieldOptions oneof_field = 475223888;
	E_OneofField = &file_ras_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ras.encoding.EncodingMessageOptions options = 475223889;
	E_Options = &file_ras_proto_extTypes[2]
	// optional ras.encoding.ClientMessageOptions client_options = 475223890;
	E_ClientOptions = &file_ras_proto_extTypes[3]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// Название опции в message
	//
	// optional string message_option = 475223891;
	E_MessageOption = &file_ras_proto_extTypes[4]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string type_url = 475223891;
	E_TypeUrl = &file_ras_proto_extTypes[5]
)

var File_ras_proto protoreflect.FileDescriptor

var file_ras_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x61, 0x73,
	0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x14, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x22, 0x95, 0x04, 0x0a, 0x16, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x17, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x68, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x73, 0x12, 0x49, 0x0a, 0x21, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x18,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x66, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x6f, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6f,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xaa, 0x02, 0x0a, 0x14, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x09, 0x74, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x09,
	0x73, 0x69, 0x7a, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x5b, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x66, 0x0a, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x3a, 0x63, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0xae,
	0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x6e, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0xae, 0xcd, 0xe2, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x47, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x40, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3,
	0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72,
	0x6c, 0x42, 0x9d, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x52, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x58, 0xaa, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x2e, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x5c, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x18, 0x52, 0x61, 0x73, 0x5c, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_proto_rawDescOnce sync.Once
	file_ras_proto_rawDescData = file_ras_proto_rawDesc
)

func file_ras_proto_rawDescGZIP() []byte {
	file_ras_proto_rawDescOnce.Do(func() {
		file_ras_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_proto_rawDescData)
	})
	return file_ras_proto_rawDescData
}

var file_ras_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ras_proto_goTypes = []interface{}{
	(*ClientMessageOptions)(nil),          // 0: ras.encoding.ClientMessageOptions
	(*EncodingMessageOptions)(nil),        // 1: ras.encoding.EncodingMessageOptions
	(*EncodingFieldOptions)(nil),          // 2: ras.encoding.EncodingFieldOptions
	(*descriptorpb.FieldOptions)(nil),     // 3: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),     // 4: google.protobuf.OneofOptions
	(*descriptorpb.MessageOptions)(nil),   // 5: google.protobuf.MessageOptions
	(*descriptorpb.EnumOptions)(nil),      // 6: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 7: google.protobuf.EnumValueOptions
}
var file_ras_proto_depIdxs = []int32{
	3,  // 0: ras.encoding.field:extendee -> google.protobuf.FieldOptions
	4,  // 1: ras.encoding.oneof_field:extendee -> google.protobuf.OneofOptions
	5,  // 2: ras.encoding.options:extendee -> google.protobuf.MessageOptions
	5,  // 3: ras.encoding.client_options:extendee -> google.protobuf.MessageOptions
	6,  // 4: ras.encoding.message_option:extendee -> google.protobuf.EnumOptions
	7,  // 5: ras.encoding.type_url:extendee -> google.protobuf.EnumValueOptions
	2,  // 6: ras.encoding.field:type_name -> ras.encoding.EncodingFieldOptions
	2,  // 7: ras.encoding.oneof_field:type_name -> ras.encoding.EncodingFieldOptions
	1,  // 8: ras.encoding.options:type_name -> ras.encoding.EncodingMessageOptions
	0,  // 9: ras.encoding.client_options:type_name -> ras.encoding.ClientMessageOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	6,  // [6:10] is the sub-list for extension type_name
	0,  // [0:6] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ras_proto_init() }
func file_ras_proto_init() {
	if File_ras_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessageOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ras_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodingMessageOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ras_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodingFieldOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_ras_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EncodingMessageOptions_PacketType)(nil),
		(*EncodingMessageOptions_EndpointDataType)(nil),
		(*EncodingMessageOptions_MessageType)(nil),
	}
	file_ras_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ras_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_ras_proto_goTypes,
		DependencyIndexes: file_ras_proto_depIdxs,
		MessageInfos:      file_ras_proto_msgTypes,
		ExtensionInfos:    file_ras_proto_extTypes,
	}.Build()
	File_ras_proto = out.File
	file_ras_proto_rawDesc = nil
	file_ras_proto_goTypes = nil
	file_ras_proto_depIdxs = nil
}
