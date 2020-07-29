// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: autopeering/selection/proto/message.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/iotaledger/hive.go/autopeering/salt/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PeeringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unix time
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// salt of the requester
	Salt *proto1.Salt `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (x *PeeringRequest) Reset() {
	*x = PeeringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_selection_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeeringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeeringRequest) ProtoMessage() {}

func (x *PeeringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_selection_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeeringRequest.ProtoReflect.Descriptor instead.
func (*PeeringRequest) Descriptor() ([]byte, []int) {
	return file_autopeering_selection_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *PeeringRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PeeringRequest) GetSalt() *proto1.Salt {
	if x != nil {
		return x.Salt
	}
	return nil
}

type PeeringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash of the corresponding request
	ReqHash []byte `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	// response of a peering request
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PeeringResponse) Reset() {
	*x = PeeringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_selection_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeeringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeeringResponse) ProtoMessage() {}

func (x *PeeringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_selection_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeeringResponse.ProtoReflect.Descriptor instead.
func (*PeeringResponse) Descriptor() ([]byte, []int) {
	return file_autopeering_selection_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *PeeringResponse) GetReqHash() []byte {
	if x != nil {
		return x.ReqHash
	}
	return nil
}

func (x *PeeringResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type PeeringDrop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unix time
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PeeringDrop) Reset() {
	*x = PeeringDrop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_selection_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeeringDrop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeeringDrop) ProtoMessage() {}

func (x *PeeringDrop) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_selection_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeeringDrop.ProtoReflect.Descriptor instead.
func (*PeeringDrop) Descriptor() ([]byte, []int) {
	return file_autopeering_selection_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *PeeringDrop) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_autopeering_selection_proto_message_proto protoreflect.FileDescriptor

var file_autopeering_selection_proto_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x73, 0x61, 0x6c, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x74,
	0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x22, 0x44, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x71,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2b, 0x0a, 0x0b,
	0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x61, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x70,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autopeering_selection_proto_message_proto_rawDescOnce sync.Once
	file_autopeering_selection_proto_message_proto_rawDescData = file_autopeering_selection_proto_message_proto_rawDesc
)

func file_autopeering_selection_proto_message_proto_rawDescGZIP() []byte {
	file_autopeering_selection_proto_message_proto_rawDescOnce.Do(func() {
		file_autopeering_selection_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_autopeering_selection_proto_message_proto_rawDescData)
	})
	return file_autopeering_selection_proto_message_proto_rawDescData
}

var file_autopeering_selection_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_autopeering_selection_proto_message_proto_goTypes = []interface{}{
	(*PeeringRequest)(nil),  // 0: proto.PeeringRequest
	(*PeeringResponse)(nil), // 1: proto.PeeringResponse
	(*PeeringDrop)(nil),     // 2: proto.PeeringDrop
	(*proto1.Salt)(nil),     // 3: proto.Salt
}
var file_autopeering_selection_proto_message_proto_depIdxs = []int32{
	3, // 0: proto.PeeringRequest.salt:type_name -> proto.Salt
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autopeering_selection_proto_message_proto_init() }
func file_autopeering_selection_proto_message_proto_init() {
	if File_autopeering_selection_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autopeering_selection_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeeringRequest); i {
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
		file_autopeering_selection_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeeringResponse); i {
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
		file_autopeering_selection_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeeringDrop); i {
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
			RawDescriptor: file_autopeering_selection_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autopeering_selection_proto_message_proto_goTypes,
		DependencyIndexes: file_autopeering_selection_proto_message_proto_depIdxs,
		MessageInfos:      file_autopeering_selection_proto_message_proto_msgTypes,
	}.Build()
	File_autopeering_selection_proto_message_proto = out.File
	file_autopeering_selection_proto_message_proto_rawDesc = nil
	file_autopeering_selection_proto_message_proto_goTypes = nil
	file_autopeering_selection_proto_message_proto_depIdxs = nil
}
