// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: testPB/test.proto

package testpb

import (
	studentpb "github.com/jsovalles/grpc/studentpb"
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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetTestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetTestResponse) Reset() {
	*x = SetTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTestResponse) ProtoMessage() {}

func (x *SetTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTestResponse.ProtoReflect.Descriptor instead.
func (*SetTestResponse) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{2}
}

func (x *SetTestResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetTestResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TestId   string `protobuf:"bytes,4,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[3]
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
	return file_testPB_test_proto_rawDescGZIP(), []int{3}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type SetQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SetQuestionResponse) Reset() {
	*x = SetQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetQuestionResponse) ProtoMessage() {}

func (x *SetQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetQuestionResponse.ProtoReflect.Descriptor instead.
func (*SetQuestionResponse) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{4}
}

func (x *SetQuestionResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EnrollmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TestId    string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *EnrollmentRequest) Reset() {
	*x = EnrollmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentRequest) ProtoMessage() {}

func (x *EnrollmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentRequest.ProtoReflect.Descriptor instead.
func (*EnrollmentRequest) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{5}
}

func (x *EnrollmentRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *EnrollmentRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type GetStudentsPerTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *GetStudentsPerTestRequest) Reset() {
	*x = GetStudentsPerTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentsPerTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentsPerTestRequest) ProtoMessage() {}

func (x *GetStudentsPerTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentsPerTestRequest.ProtoReflect.Descriptor instead.
func (*GetStudentsPerTestRequest) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{6}
}

func (x *GetStudentsPerTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type TakeTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *TakeTestRequest) Reset() {
	*x = TakeTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testPB_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTestRequest) ProtoMessage() {}

func (x *TakeTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testPB_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTestRequest.ProtoReflect.Descriptor instead.
func (*TakeTestRequest) Descriptor() ([]byte, []int) {
	return file_testPB_test_proto_rawDescGZIP(), []int{7}
}

func (x *TakeTestRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_testPB_test_proto protoreflect.FileDescriptor

var file_testPB_test_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x50, 0x42, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x25, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4b, 0x0a, 0x11, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x54, 0x61,
	0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0xef, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0a, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x46, 0x0a,
	0x0e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x30, 0x01,
	0x12, 0x35, 0x0a, 0x08, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x28, 0x01, 0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x6f, 0x76, 0x61, 0x6c, 0x6c, 0x65, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_testPB_test_proto_rawDescOnce sync.Once
	file_testPB_test_proto_rawDescData = file_testPB_test_proto_rawDesc
)

func file_testPB_test_proto_rawDescGZIP() []byte {
	file_testPB_test_proto_rawDescOnce.Do(func() {
		file_testPB_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testPB_test_proto_rawDescData)
	})
	return file_testPB_test_proto_rawDescData
}

var file_testPB_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_testPB_test_proto_goTypes = []interface{}{
	(*Test)(nil),                      // 0: test.Test
	(*GetTestRequest)(nil),            // 1: test.GetTestRequest
	(*SetTestResponse)(nil),           // 2: test.SetTestResponse
	(*Question)(nil),                  // 3: test.Question
	(*SetQuestionResponse)(nil),       // 4: test.SetQuestionResponse
	(*EnrollmentRequest)(nil),         // 5: test.EnrollmentRequest
	(*GetStudentsPerTestRequest)(nil), // 6: test.GetStudentsPerTestRequest
	(*TakeTestRequest)(nil),           // 7: test.TakeTestRequest
	(*studentpb.Student)(nil),         // 8: student.Student
}
var file_testPB_test_proto_depIdxs = []int32{
	1, // 0: test.TestService.GetTest:input_type -> test.GetTestRequest
	0, // 1: test.TestService.SetTest:input_type -> test.Test
	3, // 2: test.TestService.SetQuestions:input_type -> test.Question
	5, // 3: test.TestService.EnrollStudents:input_type -> test.EnrollmentRequest
	6, // 4: test.TestService.GetStudentsPerTest:input_type -> test.GetStudentsPerTestRequest
	7, // 5: test.TestService.TakeTest:input_type -> test.TakeTestRequest
	0, // 6: test.TestService.GetTest:output_type -> test.Test
	2, // 7: test.TestService.SetTest:output_type -> test.SetTestResponse
	4, // 8: test.TestService.SetQuestions:output_type -> test.SetQuestionResponse
	4, // 9: test.TestService.EnrollStudents:output_type -> test.SetQuestionResponse
	8, // 10: test.TestService.GetStudentsPerTest:output_type -> student.Student
	3, // 11: test.TestService.TakeTest:output_type -> test.Question
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testPB_test_proto_init() }
func file_testPB_test_proto_init() {
	if File_testPB_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testPB_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_testPB_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
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
		file_testPB_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTestResponse); i {
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
		file_testPB_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_testPB_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetQuestionResponse); i {
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
		file_testPB_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentRequest); i {
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
		file_testPB_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentsPerTestRequest); i {
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
		file_testPB_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTestRequest); i {
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
			RawDescriptor: file_testPB_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testPB_test_proto_goTypes,
		DependencyIndexes: file_testPB_test_proto_depIdxs,
		MessageInfos:      file_testPB_test_proto_msgTypes,
	}.Build()
	File_testPB_test_proto = out.File
	file_testPB_test_proto_rawDesc = nil
	file_testPB_test_proto_goTypes = nil
	file_testPB_test_proto_depIdxs = nil
}
