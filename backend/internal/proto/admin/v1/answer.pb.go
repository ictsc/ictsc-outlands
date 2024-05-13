// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: admin/v1/answer.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QuestionAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer []int32 `protobuf:"varint,6,rep,packed,name=answer,proto3" json:"answer,omitempty"`
}

func (x *QuestionAnswer) Reset() {
	*x = QuestionAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionAnswer) ProtoMessage() {}

func (x *QuestionAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionAnswer.ProtoReflect.Descriptor instead.
func (*QuestionAnswer) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionAnswer) GetAnswer() []int32 {
	if x != nil {
		return x.Answer
	}
	return nil
}

type MultipleChoiceAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionAnswer []*QuestionAnswer `protobuf:"bytes,1,rep,name=question_answer,json=questionAnswer,proto3" json:"question_answer,omitempty"`
}

func (x *MultipleChoiceAnswer) Reset() {
	*x = MultipleChoiceAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleChoiceAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleChoiceAnswer) ProtoMessage() {}

func (x *MultipleChoiceAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleChoiceAnswer.ProtoReflect.Descriptor instead.
func (*MultipleChoiceAnswer) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{1}
}

func (x *MultipleChoiceAnswer) GetQuestionAnswer() []*QuestionAnswer {
	if x != nil {
		return x.QuestionAnswer
	}
	return nil
}

type DescriptiveAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DescriptiveAnswer) Reset() {
	*x = DescriptiveAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptiveAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptiveAnswer) ProtoMessage() {}

func (x *DescriptiveAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptiveAnswer.ProtoReflect.Descriptor instead.
func (*DescriptiveAnswer) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{2}
}

func (x *DescriptiveAnswer) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProblemId   string      `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	TeamId      string      `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ProblemType ProblemType `protobuf:"varint,4,opt,name=problem_type,json=problemType,proto3,enum=admin.v1.ProblemType" json:"problem_type,omitempty"`
	// Types that are assignable to Body:
	//
	//	*Answer_MultipleChoice
	//	*Answer_Descriptive
	Body          isAnswer_Body          `protobuf_oneof:"body"`
	Point         *int32                 `protobuf:"varint,7,opt,name=point,proto3,oneof" json:"point,omitempty"`
	MarkPublished bool                   `protobuf:"varint,8,opt,name=mark_published,json=markPublished,proto3" json:"mark_published,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{3}
}

func (x *Answer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Answer) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Answer) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Answer) GetProblemType() ProblemType {
	if x != nil {
		return x.ProblemType
	}
	return ProblemType_PROBLEM_TYPE_UNSPECIFIED
}

func (m *Answer) GetBody() isAnswer_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Answer) GetMultipleChoice() *MultipleChoiceAnswer {
	if x, ok := x.GetBody().(*Answer_MultipleChoice); ok {
		return x.MultipleChoice
	}
	return nil
}

func (x *Answer) GetDescriptive() *DescriptiveAnswer {
	if x, ok := x.GetBody().(*Answer_Descriptive); ok {
		return x.Descriptive
	}
	return nil
}

func (x *Answer) GetPoint() int32 {
	if x != nil && x.Point != nil {
		return *x.Point
	}
	return 0
}

func (x *Answer) GetMarkPublished() bool {
	if x != nil {
		return x.MarkPublished
	}
	return false
}

func (x *Answer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type isAnswer_Body interface {
	isAnswer_Body()
}

type Answer_MultipleChoice struct {
	MultipleChoice *MultipleChoiceAnswer `protobuf:"bytes,5,opt,name=multiple_choice,json=multipleChoice,proto3,oneof"`
}

type Answer_Descriptive struct {
	Descriptive *DescriptiveAnswer `protobuf:"bytes,6,opt,name=descriptive,proto3,oneof"`
}

func (*Answer_MultipleChoice) isAnswer_Body() {}

func (*Answer_Descriptive) isAnswer_Body() {}

type GetAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAnswerRequest) Reset() {
	*x = GetAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswerRequest) ProtoMessage() {}

func (x *GetAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswerRequest.ProtoReflect.Descriptor instead.
func (*GetAnswerRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{4}
}

func (x *GetAnswerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer *Answer `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *GetAnswerResponse) Reset() {
	*x = GetAnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswerResponse) ProtoMessage() {}

func (x *GetAnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswerResponse.ProtoReflect.Descriptor instead.
func (*GetAnswerResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{5}
}

func (x *GetAnswerResponse) GetAnswer() *Answer {
	if x != nil {
		return x.Answer
	}
	return nil
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit             int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	LastItemCreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_item_created_at,json=lastItemCreatedAt,proto3" json:"last_item_created_at,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{6}
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetLastItemCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastItemCreatedAt
	}
	return nil
}

type GetAnswersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string      `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	TeamId     *string     `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
	Pagination *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
}

func (x *GetAnswersRequest) Reset() {
	*x = GetAnswersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswersRequest) ProtoMessage() {}

func (x *GetAnswersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswersRequest.ProtoReflect.Descriptor instead.
func (*GetAnswersRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{7}
}

func (x *GetAnswersRequest) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *GetAnswersRequest) GetTeamId() string {
	if x != nil && x.TeamId != nil {
		return *x.TeamId
	}
	return ""
}

func (x *GetAnswersRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetAnswersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*Answer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *GetAnswersResponse) Reset() {
	*x = GetAnswersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswersResponse) ProtoMessage() {}

func (x *GetAnswersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswersResponse.ProtoReflect.Descriptor instead.
func (*GetAnswersResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{8}
}

func (x *GetAnswersResponse) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type PutPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Point int32  `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *PutPointRequest) Reset() {
	*x = PutPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPointRequest) ProtoMessage() {}

func (x *PutPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPointRequest.ProtoReflect.Descriptor instead.
func (*PutPointRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{9}
}

func (x *PutPointRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutPointRequest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type PutPointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutPointResponse) Reset() {
	*x = PutPointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_answer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPointResponse) ProtoMessage() {}

func (x *PutPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_answer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPointResponse.ProtoReflect.Descriptor instead.
func (*PutPointResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_answer_proto_rawDescGZIP(), []int{10}
}

var File_admin_v1_answer_proto protoreflect.FileDescriptor

var file_admin_v1_answer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x42, 0x12, 0xba, 0x48, 0x0f, 0x92, 0x01,
	0x0c, 0x08, 0x01, 0x10, 0x0a, 0x18, 0x01, 0x22, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x4d, 0x0a,
	0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01, 0x10, 0x0a, 0x52, 0x0e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x11,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0xe5, 0x03, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98,
	0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x1a, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x48, 0x01, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x6d, 0x61, 0x72, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x41, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x0d, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x80,
	0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x53, 0x0a, 0x14,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x48, 0x00, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x22, 0x4a, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x50, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xe1, 0x01, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2d, 0x6f,
	0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_answer_proto_rawDescOnce sync.Once
	file_admin_v1_answer_proto_rawDescData = file_admin_v1_answer_proto_rawDesc
)

func file_admin_v1_answer_proto_rawDescGZIP() []byte {
	file_admin_v1_answer_proto_rawDescOnce.Do(func() {
		file_admin_v1_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_answer_proto_rawDescData)
	})
	return file_admin_v1_answer_proto_rawDescData
}

var file_admin_v1_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_admin_v1_answer_proto_goTypes = []interface{}{
	(*QuestionAnswer)(nil),        // 0: admin.v1.QuestionAnswer
	(*MultipleChoiceAnswer)(nil),  // 1: admin.v1.MultipleChoiceAnswer
	(*DescriptiveAnswer)(nil),     // 2: admin.v1.DescriptiveAnswer
	(*Answer)(nil),                // 3: admin.v1.Answer
	(*GetAnswerRequest)(nil),      // 4: admin.v1.GetAnswerRequest
	(*GetAnswerResponse)(nil),     // 5: admin.v1.GetAnswerResponse
	(*Pagination)(nil),            // 6: admin.v1.Pagination
	(*GetAnswersRequest)(nil),     // 7: admin.v1.GetAnswersRequest
	(*GetAnswersResponse)(nil),    // 8: admin.v1.GetAnswersResponse
	(*PutPointRequest)(nil),       // 9: admin.v1.PutPointRequest
	(*PutPointResponse)(nil),      // 10: admin.v1.PutPointResponse
	(ProblemType)(0),              // 11: admin.v1.ProblemType
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_admin_v1_answer_proto_depIdxs = []int32{
	0,  // 0: admin.v1.MultipleChoiceAnswer.question_answer:type_name -> admin.v1.QuestionAnswer
	11, // 1: admin.v1.Answer.problem_type:type_name -> admin.v1.ProblemType
	1,  // 2: admin.v1.Answer.multiple_choice:type_name -> admin.v1.MultipleChoiceAnswer
	2,  // 3: admin.v1.Answer.descriptive:type_name -> admin.v1.DescriptiveAnswer
	12, // 4: admin.v1.Answer.created_at:type_name -> google.protobuf.Timestamp
	3,  // 5: admin.v1.GetAnswerResponse.answer:type_name -> admin.v1.Answer
	12, // 6: admin.v1.Pagination.last_item_created_at:type_name -> google.protobuf.Timestamp
	6,  // 7: admin.v1.GetAnswersRequest.pagination:type_name -> admin.v1.Pagination
	3,  // 8: admin.v1.GetAnswersResponse.answers:type_name -> admin.v1.Answer
	4,  // 9: admin.v1.AnswerService.GetAnswer:input_type -> admin.v1.GetAnswerRequest
	7,  // 10: admin.v1.AnswerService.GetAnswers:input_type -> admin.v1.GetAnswersRequest
	9,  // 11: admin.v1.AnswerService.PutPoint:input_type -> admin.v1.PutPointRequest
	5,  // 12: admin.v1.AnswerService.GetAnswer:output_type -> admin.v1.GetAnswerResponse
	8,  // 13: admin.v1.AnswerService.GetAnswers:output_type -> admin.v1.GetAnswersResponse
	10, // 14: admin.v1.AnswerService.PutPoint:output_type -> admin.v1.PutPointResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_admin_v1_answer_proto_init() }
func file_admin_v1_answer_proto_init() {
	if File_admin_v1_answer_proto != nil {
		return
	}
	file_admin_v1_problem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionAnswer); i {
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
		file_admin_v1_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleChoiceAnswer); i {
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
		file_admin_v1_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptiveAnswer); i {
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
		file_admin_v1_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_admin_v1_answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnswerRequest); i {
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
		file_admin_v1_answer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnswerResponse); i {
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
		file_admin_v1_answer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_admin_v1_answer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnswersRequest); i {
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
		file_admin_v1_answer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnswersResponse); i {
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
		file_admin_v1_answer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPointRequest); i {
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
		file_admin_v1_answer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPointResponse); i {
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
	file_admin_v1_answer_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Answer_MultipleChoice)(nil),
		(*Answer_Descriptive)(nil),
	}
	file_admin_v1_answer_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_v1_answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_answer_proto_goTypes,
		DependencyIndexes: file_admin_v1_answer_proto_depIdxs,
		MessageInfos:      file_admin_v1_answer_proto_msgTypes,
	}.Build()
	File_admin_v1_answer_proto = out.File
	file_admin_v1_answer_proto_rawDesc = nil
	file_admin_v1_answer_proto_goTypes = nil
	file_admin_v1_answer_proto_depIdxs = nil
}
