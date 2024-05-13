// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: contestant/v1/problem.proto

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

type QuestionType int32

const (
	QuestionType_QUESTION_TYPE_UNSPECIFIED QuestionType = 0
	QuestionType_QUESTION_TYPE_RADIO       QuestionType = 1
	QuestionType_QUESTION_TYPE_CHECKBOX    QuestionType = 2
)

// Enum value maps for QuestionType.
var (
	QuestionType_name = map[int32]string{
		0: "QUESTION_TYPE_UNSPECIFIED",
		1: "QUESTION_TYPE_RADIO",
		2: "QUESTION_TYPE_CHECKBOX",
	}
	QuestionType_value = map[string]int32{
		"QUESTION_TYPE_UNSPECIFIED": 0,
		"QUESTION_TYPE_RADIO":       1,
		"QUESTION_TYPE_CHECKBOX":    2,
	}
)

func (x QuestionType) Enum() *QuestionType {
	p := new(QuestionType)
	*p = x
	return p
}

func (x QuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_contestant_v1_problem_proto_enumTypes[0].Descriptor()
}

func (QuestionType) Type() protoreflect.EnumType {
	return &file_contestant_v1_problem_proto_enumTypes[0]
}

func (x QuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionType.Descriptor instead.
func (QuestionType) EnumDescriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{0}
}

type ProblemType int32

const (
	ProblemType_PROBLEM_TYPE_UNSPECIFIED     ProblemType = 0
	ProblemType_PROBLEM_TYPE_DESCRIPTIVE     ProblemType = 1
	ProblemType_PROBLEM_TYPE_MULTIPLE_CHOICE ProblemType = 2
)

// Enum value maps for ProblemType.
var (
	ProblemType_name = map[int32]string{
		0: "PROBLEM_TYPE_UNSPECIFIED",
		1: "PROBLEM_TYPE_DESCRIPTIVE",
		2: "PROBLEM_TYPE_MULTIPLE_CHOICE",
	}
	ProblemType_value = map[string]int32{
		"PROBLEM_TYPE_UNSPECIFIED":     0,
		"PROBLEM_TYPE_DESCRIPTIVE":     1,
		"PROBLEM_TYPE_MULTIPLE_CHOICE": 2,
	}
)

func (x ProblemType) Enum() *ProblemType {
	p := new(ProblemType)
	*p = x
	return p
}

func (x ProblemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProblemType) Descriptor() protoreflect.EnumDescriptor {
	return file_contestant_v1_problem_proto_enumTypes[1].Descriptor()
}

func (ProblemType) Type() protoreflect.EnumType {
	return &file_contestant_v1_problem_proto_enumTypes[1]
}

func (x ProblemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProblemType.Descriptor instead.
func (ProblemType) EnumDescriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{1}
}

type Choice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Choice) Reset() {
	*x = Choice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice) ProtoMessage() {}

func (x *Choice) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice.ProtoReflect.Descriptor instead.
func (*Choice) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Choice) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Choice) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body    string       `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Type    QuestionType `protobuf:"varint,3,opt,name=type,proto3,enum=contestant.v1.QuestionType" json:"type,omitempty"`
	Choices []*Choice    `protobuf:"bytes,4,rep,name=choices,proto3" json:"choices,omitempty"`
	Point   int32        `protobuf:"varint,5,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Question) GetType() QuestionType {
	if x != nil {
		return x.Type
	}
	return QuestionType_QUESTION_TYPE_UNSPECIFIED
}

func (x *Question) GetChoices() []*Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *Question) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type MultipleChoiceProblem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body      *string     `protobuf:"bytes,1,opt,name=body,proto3,oneof" json:"body,omitempty"`
	Questions []*Question `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *MultipleChoiceProblem) Reset() {
	*x = MultipleChoiceProblem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleChoiceProblem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleChoiceProblem) ProtoMessage() {}

func (x *MultipleChoiceProblem) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleChoiceProblem.ProtoReflect.Descriptor instead.
func (*MultipleChoiceProblem) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{2}
}

func (x *MultipleChoiceProblem) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *MultipleChoiceProblem) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

type ConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Command  string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ConnectionInfo) Reset() {
	*x = ConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionInfo) ProtoMessage() {}

func (x *ConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionInfo.ProtoReflect.Descriptor instead.
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ConnectionInfo) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ConnectionInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectionInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DescriptiveProblem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body            string            `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	ConnectionInfos []*ConnectionInfo `protobuf:"bytes,2,rep,name=connection_infos,json=connectionInfos,proto3" json:"connection_infos,omitempty"`
}

func (x *DescriptiveProblem) Reset() {
	*x = DescriptiveProblem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptiveProblem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptiveProblem) ProtoMessage() {}

func (x *DescriptiveProblem) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptiveProblem.ProtoReflect.Descriptor instead.
func (*DescriptiveProblem) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{4}
}

func (x *DescriptiveProblem) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *DescriptiveProblem) GetConnectionInfos() []*ConnectionInfo {
	if x != nil {
		return x.ConnectionInfos
	}
	return nil
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code  string      `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Point int32       `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
	Type  ProblemType `protobuf:"varint,5,opt,name=type,proto3,enum=contestant.v1.ProblemType" json:"type,omitempty"`
	// Types that are assignable to Body:
	//
	//	*Problem_Descriptive
	//	*Problem_MultipleChoice
	Body isProblem_Body `protobuf_oneof:"body"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{5}
}

func (x *Problem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Problem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Problem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Problem) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *Problem) GetType() ProblemType {
	if x != nil {
		return x.Type
	}
	return ProblemType_PROBLEM_TYPE_UNSPECIFIED
}

func (m *Problem) GetBody() isProblem_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Problem) GetDescriptive() *DescriptiveProblem {
	if x, ok := x.GetBody().(*Problem_Descriptive); ok {
		return x.Descriptive
	}
	return nil
}

func (x *Problem) GetMultipleChoice() *MultipleChoiceProblem {
	if x, ok := x.GetBody().(*Problem_MultipleChoice); ok {
		return x.MultipleChoice
	}
	return nil
}

type isProblem_Body interface {
	isProblem_Body()
}

type Problem_Descriptive struct {
	Descriptive *DescriptiveProblem `protobuf:"bytes,6,opt,name=descriptive,proto3,oneof"`
}

type Problem_MultipleChoice struct {
	MultipleChoice *MultipleChoiceProblem `protobuf:"bytes,7,opt,name=multiple_choice,json=multipleChoice,proto3,oneof"`
}

func (*Problem_Descriptive) isProblem_Body() {}

func (*Problem_MultipleChoice) isProblem_Body() {}

type GetProblemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProblemsRequest) Reset() {
	*x = GetProblemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemsRequest) ProtoMessage() {}

func (x *GetProblemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemsRequest.ProtoReflect.Descriptor instead.
func (*GetProblemsRequest) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{6}
}

type GetProblemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problems []*Problem `protobuf:"bytes,1,rep,name=problems,proto3" json:"problems,omitempty"`
}

func (x *GetProblemsResponse) Reset() {
	*x = GetProblemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_problem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemsResponse) ProtoMessage() {}

func (x *GetProblemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_problem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemsResponse.ProtoReflect.Descriptor instead.
func (*GetProblemsResponse) Descriptor() ([]byte, []int) {
	return file_contestant_v1_problem_proto_rawDescGZIP(), []int{7}
}

func (x *GetProblemsResponse) GetProblems() []*Problem {
	if x != nil {
		return x.Problems
	}
	return nil
}

var File_contestant_v1_problem_proto protoreflect.FileDescriptor

var file_contestant_v1_problem_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x06, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01,
	0x18, 0xe8, 0x07, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x0a, 0xba, 0x48,
	0x07, 0x92, 0x01, 0x04, 0x08, 0x02, 0x10, 0x0a, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x88, 0x01, 0x0a, 0x15, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10,
	0x01, 0x18, 0xe8, 0x07, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x41, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xba, 0x48, 0x07,
	0x92, 0x01, 0x04, 0x08, 0x01, 0x10, 0x0a, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xa2, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48,
	0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x8a, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8,
	0x07, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x54, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x92, 0x01, 0x04, 0x08, 0x00, 0x10, 0x0a, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0xe2, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x05, 0xba, 0x48, 0x02,
	0x08, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x2a, 0x62, 0x0a, 0x0c, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x44, 0x49,
	0x4f, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x42, 0x4f, 0x58, 0x10, 0x02, 0x2a,
	0x6b, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x18, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x52,
	0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49,
	0x50, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x02, 0x32, 0x66, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2d, 0x6f,
	0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contestant_v1_problem_proto_rawDescOnce sync.Once
	file_contestant_v1_problem_proto_rawDescData = file_contestant_v1_problem_proto_rawDesc
)

func file_contestant_v1_problem_proto_rawDescGZIP() []byte {
	file_contestant_v1_problem_proto_rawDescOnce.Do(func() {
		file_contestant_v1_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_contestant_v1_problem_proto_rawDescData)
	})
	return file_contestant_v1_problem_proto_rawDescData
}

var file_contestant_v1_problem_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_contestant_v1_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_contestant_v1_problem_proto_goTypes = []interface{}{
	(QuestionType)(0),             // 0: contestant.v1.QuestionType
	(ProblemType)(0),              // 1: contestant.v1.ProblemType
	(*Choice)(nil),                // 2: contestant.v1.Choice
	(*Question)(nil),              // 3: contestant.v1.Question
	(*MultipleChoiceProblem)(nil), // 4: contestant.v1.MultipleChoiceProblem
	(*ConnectionInfo)(nil),        // 5: contestant.v1.ConnectionInfo
	(*DescriptiveProblem)(nil),    // 6: contestant.v1.DescriptiveProblem
	(*Problem)(nil),               // 7: contestant.v1.Problem
	(*GetProblemsRequest)(nil),    // 8: contestant.v1.GetProblemsRequest
	(*GetProblemsResponse)(nil),   // 9: contestant.v1.GetProblemsResponse
}
var file_contestant_v1_problem_proto_depIdxs = []int32{
	0, // 0: contestant.v1.Question.type:type_name -> contestant.v1.QuestionType
	2, // 1: contestant.v1.Question.choices:type_name -> contestant.v1.Choice
	3, // 2: contestant.v1.MultipleChoiceProblem.questions:type_name -> contestant.v1.Question
	5, // 3: contestant.v1.DescriptiveProblem.connection_infos:type_name -> contestant.v1.ConnectionInfo
	1, // 4: contestant.v1.Problem.type:type_name -> contestant.v1.ProblemType
	6, // 5: contestant.v1.Problem.descriptive:type_name -> contestant.v1.DescriptiveProblem
	4, // 6: contestant.v1.Problem.multiple_choice:type_name -> contestant.v1.MultipleChoiceProblem
	7, // 7: contestant.v1.GetProblemsResponse.problems:type_name -> contestant.v1.Problem
	8, // 8: contestant.v1.ProblemService.GetProblems:input_type -> contestant.v1.GetProblemsRequest
	9, // 9: contestant.v1.ProblemService.GetProblems:output_type -> contestant.v1.GetProblemsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_contestant_v1_problem_proto_init() }
func file_contestant_v1_problem_proto_init() {
	if File_contestant_v1_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contestant_v1_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice); i {
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
		file_contestant_v1_problem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_contestant_v1_problem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleChoiceProblem); i {
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
		file_contestant_v1_problem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionInfo); i {
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
		file_contestant_v1_problem_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptiveProblem); i {
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
		file_contestant_v1_problem_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
		file_contestant_v1_problem_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProblemsRequest); i {
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
		file_contestant_v1_problem_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProblemsResponse); i {
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
	file_contestant_v1_problem_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_contestant_v1_problem_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Problem_Descriptive)(nil),
		(*Problem_MultipleChoice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contestant_v1_problem_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contestant_v1_problem_proto_goTypes,
		DependencyIndexes: file_contestant_v1_problem_proto_depIdxs,
		EnumInfos:         file_contestant_v1_problem_proto_enumTypes,
		MessageInfos:      file_contestant_v1_problem_proto_msgTypes,
	}.Build()
	File_contestant_v1_problem_proto = out.File
	file_contestant_v1_problem_proto_rawDesc = nil
	file_contestant_v1_problem_proto_goTypes = nil
	file_contestant_v1_problem_proto_depIdxs = nil
}
