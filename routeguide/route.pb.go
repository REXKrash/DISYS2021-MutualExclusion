// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: routeguide/route.proto

package routeguide

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LookingForLeader bool  `protobuf:"varint,1,opt,name=lookingForLeader,proto3" json:"lookingForLeader,omitempty"`
	LowestPort       int32 `protobuf:"varint,2,opt,name=lowestPort,proto3" json:"lowestPort,omitempty"`
}

func (x *LeaderRequest) Reset() {
	*x = LeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderRequest) ProtoMessage() {}

func (x *LeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderRequest.ProtoReflect.Descriptor instead.
func (*LeaderRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{0}
}

func (x *LeaderRequest) GetLookingForLeader() bool {
	if x != nil {
		return x.LookingForLeader
	}
	return false
}

func (x *LeaderRequest) GetLowestPort() int32 {
	if x != nil {
		return x.LowestPort
	}
	return 0
}

type LeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LookingForLeader bool  `protobuf:"varint,1,opt,name=lookingForLeader,proto3" json:"lookingForLeader,omitempty"`
	LowestPort       int32 `protobuf:"varint,2,opt,name=lowestPort,proto3" json:"lowestPort,omitempty"`
}

func (x *LeaderResponse) Reset() {
	*x = LeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderResponse) ProtoMessage() {}

func (x *LeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderResponse.ProtoReflect.Descriptor instead.
func (*LeaderResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{1}
}

func (x *LeaderResponse) GetLookingForLeader() bool {
	if x != nil {
		return x.LookingForLeader
	}
	return false
}

func (x *LeaderResponse) GetLowestPort() int32 {
	if x != nil {
		return x.LowestPort
	}
	return 0
}

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{2}
}

func (x *TokenRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TokenRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{3}
}

func (x *TokenResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_routeguide_route_proto protoreflect.FileDescriptor

var file_routeguide_route_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x22, 0x5b, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x22, 0x5c, 0x0a, 0x0e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6f,
	0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6c,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22,
	0x40, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x41, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xa7, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routeguide_route_proto_rawDescOnce sync.Once
	file_routeguide_route_proto_rawDescData = file_routeguide_route_proto_rawDesc
)

func file_routeguide_route_proto_rawDescGZIP() []byte {
	file_routeguide_route_proto_rawDescOnce.Do(func() {
		file_routeguide_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeguide_route_proto_rawDescData)
	})
	return file_routeguide_route_proto_rawDescData
}

var file_routeguide_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_routeguide_route_proto_goTypes = []interface{}{
	(*LeaderRequest)(nil),  // 0: routeguide.LeaderRequest
	(*LeaderResponse)(nil), // 1: routeguide.LeaderResponse
	(*TokenRequest)(nil),   // 2: routeguide.TokenRequest
	(*TokenResponse)(nil),  // 3: routeguide.TokenResponse
}
var file_routeguide_route_proto_depIdxs = []int32{
	0, // 0: routeguide.TokenService.FindLeaderRequest:input_type -> routeguide.LeaderRequest
	2, // 1: routeguide.TokenService.SendTokenRequest:input_type -> routeguide.TokenRequest
	1, // 2: routeguide.TokenService.FindLeaderRequest:output_type -> routeguide.LeaderResponse
	3, // 3: routeguide.TokenService.SendTokenRequest:output_type -> routeguide.TokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeguide_route_proto_init() }
func file_routeguide_route_proto_init() {
	if File_routeguide_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeguide_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderRequest); i {
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
		file_routeguide_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderResponse); i {
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
		file_routeguide_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_routeguide_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
			RawDescriptor: file_routeguide_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeguide_route_proto_goTypes,
		DependencyIndexes: file_routeguide_route_proto_depIdxs,
		MessageInfos:      file_routeguide_route_proto_msgTypes,
	}.Build()
	File_routeguide_route_proto = out.File
	file_routeguide_route_proto_rawDesc = nil
	file_routeguide_route_proto_goTypes = nil
	file_routeguide_route_proto_depIdxs = nil
}
