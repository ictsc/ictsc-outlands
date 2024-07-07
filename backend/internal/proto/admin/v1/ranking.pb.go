// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: admin/v1/ranking.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank  int32 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Team  *Team `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Point int32 `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_ranking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_ranking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_admin_v1_ranking_proto_rawDescGZIP(), []int{0}
}

func (x *Rank) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Rank) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *Rank) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type GetRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unpublished bool `protobuf:"varint,1,opt,name=unpublished,proto3" json:"unpublished,omitempty"`
}

func (x *GetRankingRequest) Reset() {
	*x = GetRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_ranking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingRequest) ProtoMessage() {}

func (x *GetRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_ranking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingRequest.ProtoReflect.Descriptor instead.
func (*GetRankingRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_ranking_proto_rawDescGZIP(), []int{1}
}

func (x *GetRankingRequest) GetUnpublished() bool {
	if x != nil {
		return x.Unpublished
	}
	return false
}

type GetRankingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking []*Rank `protobuf:"bytes,1,rep,name=ranking,proto3" json:"ranking,omitempty"`
}

func (x *GetRankingResponse) Reset() {
	*x = GetRankingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_ranking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingResponse) ProtoMessage() {}

func (x *GetRankingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_ranking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingResponse.ProtoReflect.Descriptor instead.
func (*GetRankingResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_ranking_proto_rawDescGZIP(), []int{2}
}

func (x *GetRankingResponse) GetRanking() []*Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

var File_admin_v1_ranking_proto protoreflect.FileDescriptor

var file_admin_v1_ranking_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a,
	0x02, 0x20, 0x00, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x75, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x32, 0x59, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74,
	0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2d, 0x6f, 0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64,
	0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_ranking_proto_rawDescOnce sync.Once
	file_admin_v1_ranking_proto_rawDescData = file_admin_v1_ranking_proto_rawDesc
)

func file_admin_v1_ranking_proto_rawDescGZIP() []byte {
	file_admin_v1_ranking_proto_rawDescOnce.Do(func() {
		file_admin_v1_ranking_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_ranking_proto_rawDescData)
	})
	return file_admin_v1_ranking_proto_rawDescData
}

var file_admin_v1_ranking_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_admin_v1_ranking_proto_goTypes = []any{
	(*Rank)(nil),               // 0: admin.v1.Rank
	(*GetRankingRequest)(nil),  // 1: admin.v1.GetRankingRequest
	(*GetRankingResponse)(nil), // 2: admin.v1.GetRankingResponse
	(*Team)(nil),               // 3: admin.v1.Team
}
var file_admin_v1_ranking_proto_depIdxs = []int32{
	3, // 0: admin.v1.Rank.team:type_name -> admin.v1.Team
	0, // 1: admin.v1.GetRankingResponse.ranking:type_name -> admin.v1.Rank
	1, // 2: admin.v1.RankingService.GetRanking:input_type -> admin.v1.GetRankingRequest
	2, // 3: admin.v1.RankingService.GetRanking:output_type -> admin.v1.GetRankingResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_v1_ranking_proto_init() }
func file_admin_v1_ranking_proto_init() {
	if File_admin_v1_ranking_proto != nil {
		return
	}
	file_admin_v1_team_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_ranking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Rank); i {
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
		file_admin_v1_ranking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetRankingRequest); i {
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
		file_admin_v1_ranking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetRankingResponse); i {
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
			RawDescriptor: file_admin_v1_ranking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_ranking_proto_goTypes,
		DependencyIndexes: file_admin_v1_ranking_proto_depIdxs,
		MessageInfos:      file_admin_v1_ranking_proto_msgTypes,
	}.Build()
	File_admin_v1_ranking_proto = out.File
	file_admin_v1_ranking_proto_rawDesc = nil
	file_admin_v1_ranking_proto_goTypes = nil
	file_admin_v1_ranking_proto_depIdxs = nil
}
