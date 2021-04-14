// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: activity.proto

package activity

import (
	context "context"
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

type Activity_ActivityType int32

const (
	Activity_Play  Activity_ActivityType = 0
	Activity_Sleep Activity_ActivityType = 1
	Activity_Eat   Activity_ActivityType = 2
	Activity_Read  Activity_ActivityType = 3
)

// Enum value maps for Activity_ActivityType.
var (
	Activity_ActivityType_name = map[int32]string{
		0: "Play",
		1: "Sleep",
		2: "Eat",
		3: "Read",
	}
	Activity_ActivityType_value = map[string]int32{
		"Play":  0,
		"Sleep": 1,
		"Eat":   2,
		"Read":  3,
	}
)

func (x Activity_ActivityType) Enum() *Activity_ActivityType {
	p := new(Activity_ActivityType)
	*p = x
	return p
}

func (x Activity_ActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_ActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_proto_enumTypes[0].Descriptor()
}

func (Activity_ActivityType) Type() protoreflect.EnumType {
	return &file_activity_proto_enumTypes[0]
}

func (x Activity_ActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_ActivityType.Descriptor instead.
func (Activity_ActivityType) EnumDescriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{1, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
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
	return file_activity_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User         *User                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	ActivityType Activity_ActivityType `protobuf:"varint,3,opt,name=activityType,proto3,enum=activity.Activity_ActivityType" json:"activityType,omitempty"`
	Timestamp    int64                 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration     int64                 `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Label        string                `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
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
	return file_activity_proto_rawDescGZIP(), []int{1}
}

func (x *Activity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Activity) GetActivityType() Activity_ActivityType {
	if x != nil {
		return x.ActivityType
	}
	return Activity_Play
}

func (x *Activity) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Activity) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Activity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *CreateActivityReq) Reset() {
	*x = CreateActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActivityReq) ProtoMessage() {}

func (x *CreateActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActivityReq.ProtoReflect.Descriptor instead.
func (*CreateActivityReq) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{2}
}

func (x *CreateActivityReq) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type OneActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OneActivityReq) Reset() {
	*x = OneActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneActivityReq) ProtoMessage() {}

func (x *OneActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneActivityReq.ProtoReflect.Descriptor instead.
func (*OneActivityReq) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{3}
}

func (x *OneActivityReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListUserActivitiesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ListUserActivitiesReq) Reset() {
	*x = ListUserActivitiesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserActivitiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserActivitiesReq) ProtoMessage() {}

func (x *ListUserActivitiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserActivitiesReq.ProtoReflect.Descriptor instead.
func (*ListUserActivitiesReq) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserActivitiesReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x56, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x8b, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x61, 0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x10, 0x03,
	0x22, 0x43, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xde, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1b,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x3b, 0x0a, 0x0b, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_activity_proto_goTypes = []interface{}{
	(Activity_ActivityType)(0),    // 0: activity.Activity.ActivityType
	(*User)(nil),                  // 1: activity.User
	(*Activity)(nil),              // 2: activity.Activity
	(*CreateActivityReq)(nil),     // 3: activity.CreateActivityReq
	(*OneActivityReq)(nil),        // 4: activity.OneActivityReq
	(*ListUserActivitiesReq)(nil), // 5: activity.ListUserActivitiesReq
}
var file_activity_proto_depIdxs = []int32{
	1, // 0: activity.Activity.user:type_name -> activity.User
	0, // 1: activity.Activity.activityType:type_name -> activity.Activity.ActivityType
	2, // 2: activity.CreateActivityReq.activity:type_name -> activity.Activity
	3, // 3: activity.ActivityService.CreateActivity:input_type -> activity.CreateActivityReq
	4, // 4: activity.ActivityService.OneActivity:input_type -> activity.OneActivityReq
	5, // 5: activity.ActivityService.ListUserActivities:input_type -> activity.ListUserActivitiesReq
	2, // 6: activity.ActivityService.CreateActivity:output_type -> activity.Activity
	2, // 7: activity.ActivityService.OneActivity:output_type -> activity.Activity
	2, // 8: activity.ActivityService.ListUserActivities:output_type -> activity.Activity
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActivityReq); i {
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
		file_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneActivityReq); i {
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
		file_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserActivitiesReq); i {
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
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		EnumInfos:         file_activity_proto_enumTypes,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
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
	CreateActivity(ctx context.Context, in *CreateActivityReq, opts ...grpc.CallOption) (*Activity, error)
	OneActivity(ctx context.Context, in *OneActivityReq, opts ...grpc.CallOption) (*Activity, error)
	ListUserActivities(ctx context.Context, in *ListUserActivitiesReq, opts ...grpc.CallOption) (ActivityService_ListUserActivitiesClient, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) CreateActivity(ctx context.Context, in *CreateActivityReq, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) OneActivity(ctx context.Context, in *OneActivityReq, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/OneActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) ListUserActivities(ctx context.Context, in *ListUserActivitiesReq, opts ...grpc.CallOption) (ActivityService_ListUserActivitiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActivityService_serviceDesc.Streams[0], "/activity.ActivityService/ListUserActivities", opts...)
	if err != nil {
		return nil, err
	}
	x := &activityServiceListUserActivitiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActivityService_ListUserActivitiesClient interface {
	Recv() (*Activity, error)
	grpc.ClientStream
}

type activityServiceListUserActivitiesClient struct {
	grpc.ClientStream
}

func (x *activityServiceListUserActivitiesClient) Recv() (*Activity, error) {
	m := new(Activity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActivityServiceServer is the server API for ActivityService service.
type ActivityServiceServer interface {
	CreateActivity(context.Context, *CreateActivityReq) (*Activity, error)
	OneActivity(context.Context, *OneActivityReq) (*Activity, error)
	ListUserActivities(*ListUserActivitiesReq, ActivityService_ListUserActivitiesServer) error
}

// UnimplementedActivityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (*UnimplementedActivityServiceServer) CreateActivity(context.Context, *CreateActivityReq) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (*UnimplementedActivityServiceServer) OneActivity(context.Context, *OneActivityReq) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneActivity not implemented")
}
func (*UnimplementedActivityServiceServer) ListUserActivities(*ListUserActivitiesReq, ActivityService_ListUserActivitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUserActivities not implemented")
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).CreateActivity(ctx, req.(*CreateActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_OneActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).OneActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/OneActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).OneActivity(ctx, req.(*OneActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_ListUserActivities_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUserActivitiesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivityServiceServer).ListUserActivities(m, &activityServiceListUserActivitiesServer{stream})
}

type ActivityService_ListUserActivitiesServer interface {
	Send(*Activity) error
	grpc.ServerStream
}

type activityServiceListUserActivitiesServer struct {
	grpc.ServerStream
}

func (x *activityServiceListUserActivitiesServer) Send(m *Activity) error {
	return x.ServerStream.SendMsg(m)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activity.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _ActivityService_CreateActivity_Handler,
		},
		{
			MethodName: "OneActivity",
			Handler:    _ActivityService_OneActivity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUserActivities",
			Handler:       _ActivityService_ListUserActivities_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "activity.proto",
}
