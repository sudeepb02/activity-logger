// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: activitypb/activity.proto

package activitypb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration  string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Label     string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Activity) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Activity) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Activity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type LogActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *LogActivityRequest) Reset() {
	*x = LogActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogActivityRequest) ProtoMessage() {}

func (x *LogActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogActivityRequest.ProtoReflect.Descriptor instead.
func (*LogActivityRequest) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{2}
}

func (x *LogActivityRequest) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type LogActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LogActivityResponse) Reset() {
	*x = LogActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogActivityResponse) ProtoMessage() {}

func (x *LogActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogActivityResponse.ProtoReflect.Descriptor instead.
func (*LogActivityResponse) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{3}
}

func (x *LogActivityResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type IsDoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *IsDoneRequest) Reset() {
	*x = IsDoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsDoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsDoneRequest) ProtoMessage() {}

func (x *IsDoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsDoneRequest.ProtoReflect.Descriptor instead.
func (*IsDoneRequest) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{4}
}

func (x *IsDoneRequest) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type IsDoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IsDoneResponse) Reset() {
	*x = IsDoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsDoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsDoneResponse) ProtoMessage() {}

func (x *IsDoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsDoneResponse.ProtoReflect.Descriptor instead.
func (*IsDoneResponse) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{5}
}

func (x *IsDoneResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type IsValidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *IsValidRequest) Reset() {
	*x = IsValidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidRequest) ProtoMessage() {}

func (x *IsValidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidRequest.ProtoReflect.Descriptor instead.
func (*IsValidRequest) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{6}
}

func (x *IsValidRequest) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type IsValidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *IsValidResponse) Reset() {
	*x = IsValidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidResponse) ProtoMessage() {}

func (x *IsValidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidResponse.ProtoReflect.Descriptor instead.
func (*IsValidResponse) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{7}
}

func (x *IsValidResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{8}
}

func (x *AddUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activitypb_activity_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activitypb_activity_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_activitypb_activity_proto_rawDescGZIP(), []int{9}
}

func (x *AddUserResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_activitypb_activity_proto protoreflect.FileDescriptor

var file_activitypb_activity_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x6e, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x46, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x44, 0x0a,
	0x12, 0x4c, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x3f, 0x0a, 0x0d, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x0e, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a,
	0x0e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22,
	0x29, 0x0a, 0x0f, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x29, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa2, 0x02, 0x0a, 0x0f,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1c,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x06, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x73, 0x44, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activitypb_activity_proto_rawDescOnce sync.Once
	file_activitypb_activity_proto_rawDescData = file_activitypb_activity_proto_rawDesc
)

func file_activitypb_activity_proto_rawDescGZIP() []byte {
	file_activitypb_activity_proto_rawDescOnce.Do(func() {
		file_activitypb_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activitypb_activity_proto_rawDescData)
	})
	return file_activitypb_activity_proto_rawDescData
}

var file_activitypb_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_activitypb_activity_proto_goTypes = []interface{}{
	(*Activity)(nil),            // 0: activity.Activity
	(*User)(nil),                // 1: activity.User
	(*LogActivityRequest)(nil),  // 2: activity.LogActivityRequest
	(*LogActivityResponse)(nil), // 3: activity.LogActivityResponse
	(*IsDoneRequest)(nil),       // 4: activity.IsDoneRequest
	(*IsDoneResponse)(nil),      // 5: activity.IsDoneResponse
	(*IsValidRequest)(nil),      // 6: activity.IsValidRequest
	(*IsValidResponse)(nil),     // 7: activity.IsValidResponse
	(*AddUserRequest)(nil),      // 8: activity.AddUserRequest
	(*AddUserResponse)(nil),     // 9: activity.AddUserResponse
}
var file_activitypb_activity_proto_depIdxs = []int32{
	0, // 0: activity.LogActivityRequest.activity:type_name -> activity.Activity
	0, // 1: activity.IsDoneRequest.activity:type_name -> activity.Activity
	0, // 2: activity.IsValidRequest.activity:type_name -> activity.Activity
	1, // 3: activity.AddUserRequest.user:type_name -> activity.User
	2, // 4: activity.ActivityService.LogActivity:input_type -> activity.LogActivityRequest
	4, // 5: activity.ActivityService.IsDone:input_type -> activity.IsDoneRequest
	6, // 6: activity.ActivityService.IsValid:input_type -> activity.IsValidRequest
	8, // 7: activity.ActivityService.AddUser:input_type -> activity.AddUserRequest
	3, // 8: activity.ActivityService.LogActivity:output_type -> activity.LogActivityResponse
	5, // 9: activity.ActivityService.IsDone:output_type -> activity.IsDoneResponse
	7, // 10: activity.ActivityService.IsValid:output_type -> activity.IsValidResponse
	9, // 11: activity.ActivityService.AddUser:output_type -> activity.AddUserResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_activitypb_activity_proto_init() }
func file_activitypb_activity_proto_init() {
	if File_activitypb_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activitypb_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_activitypb_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_activitypb_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogActivityRequest); i {
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
		file_activitypb_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogActivityResponse); i {
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
		file_activitypb_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsDoneRequest); i {
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
		file_activitypb_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsDoneResponse); i {
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
		file_activitypb_activity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidRequest); i {
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
		file_activitypb_activity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidResponse); i {
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
		file_activitypb_activity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_activitypb_activity_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResponse); i {
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
			RawDescriptor: file_activitypb_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activitypb_activity_proto_goTypes,
		DependencyIndexes: file_activitypb_activity_proto_depIdxs,
		MessageInfos:      file_activitypb_activity_proto_msgTypes,
	}.Build()
	File_activitypb_activity_proto = out.File
	file_activitypb_activity_proto_rawDesc = nil
	file_activitypb_activity_proto_goTypes = nil
	file_activitypb_activity_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityServiceClient interface {
	LogActivity(ctx context.Context, in *LogActivityRequest, opts ...grpc.CallOption) (*LogActivityResponse, error)
	IsDone(ctx context.Context, in *IsDoneRequest, opts ...grpc.CallOption) (*IsDoneResponse, error)
	IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (*IsValidResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) LogActivity(ctx context.Context, in *LogActivityRequest, opts ...grpc.CallOption) (*LogActivityResponse, error) {
	out := new(LogActivityResponse)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/LogActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) IsDone(ctx context.Context, in *IsDoneRequest, opts ...grpc.CallOption) (*IsDoneResponse, error) {
	out := new(IsDoneResponse)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/IsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (*IsValidResponse, error) {
	out := new(IsValidResponse)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/IsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
type ActivityServiceServer interface {
	LogActivity(context.Context, *LogActivityRequest) (*LogActivityResponse, error)
	IsDone(context.Context, *IsDoneRequest) (*IsDoneResponse, error)
	IsValid(context.Context, *IsValidRequest) (*IsValidResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
}

// UnimplementedActivityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (*UnimplementedActivityServiceServer) LogActivity(context.Context, *LogActivityRequest) (*LogActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogActivity not implemented")
}
func (*UnimplementedActivityServiceServer) IsDone(context.Context, *IsDoneRequest) (*IsDoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsDone not implemented")
}
func (*UnimplementedActivityServiceServer) IsValid(context.Context, *IsValidRequest) (*IsValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (*UnimplementedActivityServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_LogActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).LogActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/LogActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).LogActivity(ctx, req.(*LogActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_IsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).IsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/IsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).IsDone(ctx, req.(*IsDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/IsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).IsValid(ctx, req.(*IsValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activity.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogActivity",
			Handler:    _ActivityService_LogActivity_Handler,
		},
		{
			MethodName: "IsDone",
			Handler:    _ActivityService_IsDone_Handler,
		},
		{
			MethodName: "IsValid",
			Handler:    _ActivityService_IsValid_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ActivityService_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activitypb/activity.proto",
}
