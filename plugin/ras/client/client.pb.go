// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/client/client.proto

package client

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

type ClientOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsClient         bool `protobuf:"varint,1,opt,name=is_client,json=isClient,proto3" json:"is_client,omitempty"`
	IsEndpoint       bool `protobuf:"varint,2,opt,name=is_endpoint,json=isEndpoint,proto3" json:"is_endpoint,omitempty"`
	IsRequestService bool `protobuf:"varint,3,opt,name=is_request_service,json=isRequestService,proto3" json:"is_request_service,omitempty"`
	IsRasService     bool `protobuf:"varint,4,opt,name=is_ras_service,json=isRasService,proto3" json:"is_ras_service,omitempty"`
}

func (x *ClientOptions) Reset() {
	*x = ClientOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientOptions) ProtoMessage() {}

func (x *ClientOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientOptions.ProtoReflect.Descriptor instead.
func (*ClientOptions) Descriptor() ([]byte, []int) {
	return file_ras_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientOptions) GetIsClient() bool {
	if x != nil {
		return x.IsClient
	}
	return false
}

func (x *ClientOptions) GetIsEndpoint() bool {
	if x != nil {
		return x.IsEndpoint
	}
	return false
}

func (x *ClientOptions) GetIsRequestService() bool {
	if x != nil {
		return x.IsRequestService
	}
	return false
}

func (x *ClientOptions) GetIsRasService() bool {
	if x != nil {
		return x.IsRasService
	}
	return false
}

type ClientMethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoPacketPack    bool              `protobuf:"varint,1,opt,name=no_packet_pack,json=noPacketPack,proto3" json:"no_packet_pack,omitempty"`
	Param           map[string]string `protobuf:"bytes,2,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IgnoreEmpty     bool              `protobuf:"varint,3,opt,name=ignore_empty,json=ignoreEmpty,proto3" json:"ignore_empty,omitempty"`
	NewEndpointFunc bool              `protobuf:"varint,4,opt,name=new_endpoint_func,json=newEndpointFunc,proto3" json:"new_endpoint_func,omitempty"`
	ProxyName       string            `protobuf:"bytes,5,opt,name=proxy_name,json=proxyName,proto3" json:"proxy_name,omitempty"`
}

func (x *ClientMethodOptions) Reset() {
	*x = ClientMethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMethodOptions) ProtoMessage() {}

func (x *ClientMethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMethodOptions.ProtoReflect.Descriptor instead.
func (*ClientMethodOptions) Descriptor() ([]byte, []int) {
	return file_ras_client_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientMethodOptions) GetNoPacketPack() bool {
	if x != nil {
		return x.NoPacketPack
	}
	return false
}

func (x *ClientMethodOptions) GetParam() map[string]string {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *ClientMethodOptions) GetIgnoreEmpty() bool {
	if x != nil {
		return x.IgnoreEmpty
	}
	return false
}

func (x *ClientMethodOptions) GetNewEndpointFunc() bool {
	if x != nil {
		return x.NewEndpointFunc
	}
	return false
}

func (x *ClientMethodOptions) GetProxyName() string {
	if x != nil {
		return x.ProxyName
	}
	return ""
}

var file_ras_client_client_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ClientOptions)(nil),
		Field:         475223888,
		Name:          "ras.client.client",
		Tag:           "bytes,475223888,opt,name=client",
		Filename:      "ras/client/client.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*ClientMethodOptions)(nil),
		Field:         475223888,
		Name:          "ras.client.method",
		Tag:           "bytes,475223888,opt,name=method",
		Filename:      "ras/client/client.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional ras.client.ClientOptions client = 475223888;
	E_Client = &file_ras_client_client_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional ras.client.ClientMethodOptions method = 475223888;
	E_Method = &file_ras_client_client_proto_extTypes[1]
)

var File_ras_client_client_proto protoreflect.FileDescriptor

var file_ras_client_client_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x61, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x72, 0x61, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x52, 0x61, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x13,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x40, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e, 0x65, 0x77, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x56, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0,
	0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x5b, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0xa1, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d,
	0x72, 0x61, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x52, 0x43, 0x58, 0xaa, 0x02, 0x0a, 0x52, 0x61,
	0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x0a, 0x52, 0x61, 0x73, 0x5c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x16, 0x52, 0x61, 0x73, 0x5c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_client_client_proto_rawDescOnce sync.Once
	file_ras_client_client_proto_rawDescData = file_ras_client_client_proto_rawDesc
)

func file_ras_client_client_proto_rawDescGZIP() []byte {
	file_ras_client_client_proto_rawDescOnce.Do(func() {
		file_ras_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_client_client_proto_rawDescData)
	})
	return file_ras_client_client_proto_rawDescData
}

var file_ras_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ras_client_client_proto_goTypes = []interface{}{
	(*ClientOptions)(nil),               // 0: ras.client.ClientOptions
	(*ClientMethodOptions)(nil),         // 1: ras.client.ClientMethodOptions
	nil,                                 // 2: ras.client.ClientMethodOptions.ParamEntry
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 4: google.protobuf.MethodOptions
}
var file_ras_client_client_proto_depIdxs = []int32{
	2, // 0: ras.client.ClientMethodOptions.param:type_name -> ras.client.ClientMethodOptions.ParamEntry
	3, // 1: ras.client.client:extendee -> google.protobuf.ServiceOptions
	4, // 2: ras.client.method:extendee -> google.protobuf.MethodOptions
	0, // 3: ras.client.client:type_name -> ras.client.ClientOptions
	1, // 4: ras.client.method:type_name -> ras.client.ClientMethodOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ras_client_client_proto_init() }
func file_ras_client_client_proto_init() {
	if File_ras_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientOptions); i {
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
		file_ras_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMethodOptions); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ras_client_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_ras_client_client_proto_goTypes,
		DependencyIndexes: file_ras_client_client_proto_depIdxs,
		MessageInfos:      file_ras_client_client_proto_msgTypes,
		ExtensionInfos:    file_ras_client_client_proto_extTypes,
	}.Build()
	File_ras_client_client_proto = out.File
	file_ras_client_client_proto_rawDesc = nil
	file_ras_client_client_proto_goTypes = nil
	file_ras_client_client_proto_depIdxs = nil
}
