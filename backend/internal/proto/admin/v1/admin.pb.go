// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: admin/v1/admin.proto

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

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Admin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMeRequest) Reset() {
	*x = GetMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeRequest) ProtoMessage() {}

func (x *GetMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeRequest.ProtoReflect.Descriptor instead.
func (*GetMeRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

type GetMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetMeResponse) Reset() {
	*x = GetMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeResponse) ProtoMessage() {}

func (x *GetMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeResponse.ProtoReflect.Descriptor instead.
func (*GetMeResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *GetMeResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type GetAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAdminRequest) Reset() {
	*x = GetAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRequest) ProtoMessage() {}

func (x *GetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRequest.ProtoReflect.Descriptor instead.
func (*GetAdminRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *GetAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetAdminResponse) Reset() {
	*x = GetAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminResponse) ProtoMessage() {}

func (x *GetAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminResponse.ProtoReflect.Descriptor instead.
func (*GetAdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{4}
}

func (x *GetAdminResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type GetAdminsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminsRequest) Reset() {
	*x = GetAdminsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminsRequest) ProtoMessage() {}

func (x *GetAdminsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminsRequest.ProtoReflect.Descriptor instead.
func (*GetAdminsRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{5}
}

type GetAdminsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admins []*Admin `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (x *GetAdminsResponse) Reset() {
	*x = GetAdminsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminsResponse) ProtoMessage() {}

func (x *GetAdminsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminsResponse.ProtoReflect.Descriptor instead.
func (*GetAdminsResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{6}
}

func (x *GetAdminsResponse) GetAdmins() []*Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

type PostAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PostAdminRequest) Reset() {
	*x = PostAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAdminRequest) ProtoMessage() {}

func (x *PostAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAdminRequest.ProtoReflect.Descriptor instead.
func (*PostAdminRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{7}
}

func (x *PostAdminRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PostAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *PostAdminResponse) Reset() {
	*x = PostAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAdminResponse) ProtoMessage() {}

func (x *PostAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAdminResponse.ProtoReflect.Descriptor instead.
func (*PostAdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{8}
}

func (x *PostAdminResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type PatchMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PatchMeRequest) Reset() {
	*x = PatchMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchMeRequest) ProtoMessage() {}

func (x *PatchMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchMeRequest.ProtoReflect.Descriptor instead.
func (*PatchMeRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{9}
}

func (x *PatchMeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PatchMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *PatchMeResponse) Reset() {
	*x = PatchMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchMeResponse) ProtoMessage() {}

func (x *PatchMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchMeResponse.ProtoReflect.Descriptor instead.
func (*PatchMeResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{10}
}

func (x *PatchMeResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

var File_admin_v1_admin_proto protoreflect.FileDescriptor

var file_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a,
	0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x22, 0x31, 0x0a, 0x10, 0x50, 0x6f, 0x73,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x2f, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x40, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x32, 0xe1, 0x02, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x16, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x09, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65,
	0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73,
	0x63, 0x2d, 0x6f, 0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_admin_v1_admin_proto_rawDescOnce sync.Once
	file_admin_v1_admin_proto_rawDescData = file_admin_v1_admin_proto_rawDesc
)

func file_admin_v1_admin_proto_rawDescGZIP() []byte {
	file_admin_v1_admin_proto_rawDescOnce.Do(func() {
		file_admin_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_admin_proto_rawDescData)
	})
	return file_admin_v1_admin_proto_rawDescData
}

var file_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_admin_v1_admin_proto_goTypes = []interface{}{
	(*Admin)(nil),             // 0: admin.v1.Admin
	(*GetMeRequest)(nil),      // 1: admin.v1.GetMeRequest
	(*GetMeResponse)(nil),     // 2: admin.v1.GetMeResponse
	(*GetAdminRequest)(nil),   // 3: admin.v1.GetAdminRequest
	(*GetAdminResponse)(nil),  // 4: admin.v1.GetAdminResponse
	(*GetAdminsRequest)(nil),  // 5: admin.v1.GetAdminsRequest
	(*GetAdminsResponse)(nil), // 6: admin.v1.GetAdminsResponse
	(*PostAdminRequest)(nil),  // 7: admin.v1.PostAdminRequest
	(*PostAdminResponse)(nil), // 8: admin.v1.PostAdminResponse
	(*PatchMeRequest)(nil),    // 9: admin.v1.PatchMeRequest
	(*PatchMeResponse)(nil),   // 10: admin.v1.PatchMeResponse
}
var file_admin_v1_admin_proto_depIdxs = []int32{
	0,  // 0: admin.v1.GetMeResponse.admin:type_name -> admin.v1.Admin
	0,  // 1: admin.v1.GetAdminResponse.admin:type_name -> admin.v1.Admin
	0,  // 2: admin.v1.GetAdminsResponse.admins:type_name -> admin.v1.Admin
	0,  // 3: admin.v1.PostAdminResponse.admin:type_name -> admin.v1.Admin
	0,  // 4: admin.v1.PatchMeResponse.admin:type_name -> admin.v1.Admin
	1,  // 5: admin.v1.AdminService.GetMe:input_type -> admin.v1.GetMeRequest
	3,  // 6: admin.v1.AdminService.GetAdmin:input_type -> admin.v1.GetAdminRequest
	5,  // 7: admin.v1.AdminService.GetAdmins:input_type -> admin.v1.GetAdminsRequest
	7,  // 8: admin.v1.AdminService.PostAdmin:input_type -> admin.v1.PostAdminRequest
	9,  // 9: admin.v1.AdminService.PatchMe:input_type -> admin.v1.PatchMeRequest
	2,  // 10: admin.v1.AdminService.GetMe:output_type -> admin.v1.GetMeResponse
	4,  // 11: admin.v1.AdminService.GetAdmin:output_type -> admin.v1.GetAdminResponse
	6,  // 12: admin.v1.AdminService.GetAdmins:output_type -> admin.v1.GetAdminsResponse
	8,  // 13: admin.v1.AdminService.PostAdmin:output_type -> admin.v1.PostAdminResponse
	10, // 14: admin.v1.AdminService.PatchMe:output_type -> admin.v1.PatchMeResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_admin_v1_admin_proto_init() }
func file_admin_v1_admin_proto_init() {
	if File_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_admin_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeResponse); i {
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
		file_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminResponse); i {
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
		file_admin_v1_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminsRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminsResponse); i {
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
		file_admin_v1_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAdminRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAdminResponse); i {
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
		file_admin_v1_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchMeRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchMeResponse); i {
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
			RawDescriptor: file_admin_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_admin_proto_goTypes,
		DependencyIndexes: file_admin_v1_admin_proto_depIdxs,
		MessageInfos:      file_admin_v1_admin_proto_msgTypes,
	}.Build()
	File_admin_v1_admin_proto = out.File
	file_admin_v1_admin_proto_rawDesc = nil
	file_admin_v1_admin_proto_goTypes = nil
	file_admin_v1_admin_proto_depIdxs = nil
}
