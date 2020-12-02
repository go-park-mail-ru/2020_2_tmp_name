// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: comments.proto

package comments

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

type UserComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Comment *Comment `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *UserComment) Reset() {
	*x = UserComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserComment) ProtoMessage() {}

func (x *UserComment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserComment.ProtoReflect.Descriptor instead.
func (*UserComment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *UserComment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserComment) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Uid1         int32  `protobuf:"varint,2,opt,name=Uid1,json=uid1,proto3" json:"Uid1,omitempty"`
	Uid2         int32  `protobuf:"varint,3,opt,name=Uid2,json=uid2,proto3" json:"Uid2,omitempty"`
	TimeDelivery string `protobuf:"bytes,4,opt,name=TimeDelivery,json=timeDelivery,proto3" json:"TimeDelivery,omitempty"`
	CommentText  string `protobuf:"bytes,5,opt,name=CommentText,json=commentText,proto3" json:"CommentText,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

func (x *Comment) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Comment) GetUid1() int32 {
	if x != nil {
		return x.Uid1
	}
	return 0
}

func (x *Comment) GetUid2() int32 {
	if x != nil {
		return x.Uid2
	}
	return 0
}

func (x *Comment) GetTimeDelivery() string {
	if x != nil {
		return x.TimeDelivery
	}
	return ""
}

func (x *Comment) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32    `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Telephone  string   `protobuf:"bytes,3,opt,name=Telephone,json=telephone,proto3" json:"Telephone,omitempty"`
	Password   string   `protobuf:"bytes,4,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
	DateBirth  int32    `protobuf:"varint,5,opt,name=DateBirth,json=dateBirth,proto3" json:"DateBirth,omitempty"`
	Day        string   `protobuf:"bytes,6,opt,name=Day,json=day,proto3" json:"Day,omitempty"`
	Month      string   `protobuf:"bytes,7,opt,name=Month,json=month,proto3" json:"Month,omitempty"`
	Year       string   `protobuf:"bytes,8,opt,name=Year,json=year,proto3" json:"Year,omitempty"`
	Sex        string   `protobuf:"bytes,9,opt,name=Sex,json=sex,proto3" json:"Sex,omitempty"`
	LinkImages []string `protobuf:"bytes,10,rep,name=LinkImages,json=linkImages,proto3" json:"LinkImages,omitempty"`
	Job        string   `protobuf:"bytes,11,opt,name=Job,json=job,proto3" json:"Job,omitempty"`
	Education  string   `protobuf:"bytes,12,opt,name=Education,json=education,proto3" json:"Education,omitempty"`
	AboutMe    string   `protobuf:"bytes,13,opt,name=AboutMe,json=aboutMe,proto3" json:"AboutMe,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
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
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetDateBirth() int32 {
	if x != nil {
		return x.DateBirth
	}
	return 0
}

func (x *User) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *User) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *User) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *User) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *User) GetLinkImages() []string {
	if x != nil {
		return x.LinkImages
	}
	return nil
}

func (x *User) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *User) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *User) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

type UserFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32    `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	DateBirth   int32    `protobuf:"varint,3,opt,name=DateBirth,json=dateBirth,proto3" json:"DateBirth,omitempty"`
	LinkImages  []string `protobuf:"bytes,4,rep,name=LinkImages,json=linkImages,proto3" json:"LinkImages,omitempty"`
	Job         string   `protobuf:"bytes,5,opt,name=Job,json=job,proto3" json:"Job,omitempty"`
	Education   string   `protobuf:"bytes,6,opt,name=Education,json=education,proto3" json:"Education,omitempty"`
	AboutMe     string   `protobuf:"bytes,7,opt,name=AboutMe,json=aboutMe,proto3" json:"AboutMe,omitempty"`
	IsSuperLike bool     `protobuf:"varint,8,opt,name=IsSuperLike,json=isSuperLike,proto3" json:"IsSuperLike,omitempty"`
}

func (x *UserFeed) Reset() {
	*x = UserFeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFeed) ProtoMessage() {}

func (x *UserFeed) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFeed.ProtoReflect.Descriptor instead.
func (*UserFeed) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{5}
}

func (x *UserFeed) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserFeed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserFeed) GetDateBirth() int32 {
	if x != nil {
		return x.DateBirth
	}
	return 0
}

func (x *UserFeed) GetLinkImages() []string {
	if x != nil {
		return x.LinkImages
	}
	return nil
}

func (x *UserFeed) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *UserFeed) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *UserFeed) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *UserFeed) GetIsSuperLike() bool {
	if x != nil {
		return x.IsSuperLike
	}
	return false
}

type CommentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *UserFeed `protobuf:"bytes,1,opt,name=User,json=user,proto3" json:"User,omitempty"`
	CommentText  string    `protobuf:"bytes,2,opt,name=CommentText,json=commentText,proto3" json:"CommentText,omitempty"`
	TimeDelivery string    `protobuf:"bytes,3,opt,name=TimeDelivery,json=timeDelivery,proto3" json:"TimeDelivery,omitempty"`
}

func (x *CommentId) Reset() {
	*x = CommentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentId) ProtoMessage() {}

func (x *CommentId) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentId.ProtoReflect.Descriptor instead.
func (*CommentId) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{6}
}

func (x *CommentId) GetUser() *UserFeed {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CommentId) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *CommentId) GetTimeDelivery() string {
	if x != nil {
		return x.TimeDelivery
	}
	return ""
}

type CommentsById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentById []*CommentId `protobuf:"bytes,1,rep,name=CommentById,json=commentById,proto3" json:"CommentById,omitempty"`
}

func (x *CommentsById) Reset() {
	*x = CommentsById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsById) ProtoMessage() {}

func (x *CommentsById) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsById.ProtoReflect.Descriptor instead.
func (*CommentsById) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{7}
}

func (x *CommentsById) GetCommentById() []*CommentId {
	if x != nil {
		return x.CommentById
	}
	return nil
}

type CommentsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CommentsById `protobuf:"bytes,1,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (x *CommentsData) Reset() {
	*x = CommentsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsData) ProtoMessage() {}

func (x *CommentsData) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsData.ProtoReflect.Descriptor instead.
func (*CommentsData) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{8}
}

func (x *CommentsData) GetData() *CommentsById {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5e, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x69, 0x64, 0x31, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x69, 0x64, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x69, 0x64,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x69, 0x64, 0x32, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x22, 0xba, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x4d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65,
	0x22, 0xd8, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f,
	0x62, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x79, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x22, 0x45, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x22, 0x3a, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x7e, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x52, 0x50, 0x43, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_comments_proto_goTypes = []interface{}{
	(*UserComment)(nil),  // 0: comments.UserComment
	(*Empty)(nil),        // 1: comments.Empty
	(*Id)(nil),           // 2: comments.Id
	(*Comment)(nil),      // 3: comments.Comment
	(*User)(nil),         // 4: comments.User
	(*UserFeed)(nil),     // 5: comments.UserFeed
	(*CommentId)(nil),    // 6: comments.CommentId
	(*CommentsById)(nil), // 7: comments.CommentsById
	(*CommentsData)(nil), // 8: comments.CommentsData
}
var file_comments_proto_depIdxs = []int32{
	4, // 0: comments.UserComment.user:type_name -> comments.User
	3, // 1: comments.UserComment.comment:type_name -> comments.Comment
	5, // 2: comments.CommentId.User:type_name -> comments.UserFeed
	6, // 3: comments.CommentsById.CommentById:type_name -> comments.CommentId
	7, // 4: comments.CommentsData.Data:type_name -> comments.CommentsById
	0, // 5: comments.CommentsGRPCHandler.Comment:input_type -> comments.UserComment
	2, // 6: comments.CommentsGRPCHandler.CommentsById:input_type -> comments.Id
	1, // 7: comments.CommentsGRPCHandler.Comment:output_type -> comments.Empty
	8, // 8: comments.CommentsGRPCHandler.CommentsById:output_type -> comments.CommentsData
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserComment); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFeed); i {
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
		file_comments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentId); i {
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
		file_comments_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsById); i {
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
		file_comments_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsData); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommentsGRPCHandlerClient is the client API for CommentsGRPCHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentsGRPCHandlerClient interface {
	Comment(ctx context.Context, in *UserComment, opts ...grpc.CallOption) (*Empty, error)
	CommentsById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CommentsData, error)
}

type commentsGRPCHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsGRPCHandlerClient(cc grpc.ClientConnInterface) CommentsGRPCHandlerClient {
	return &commentsGRPCHandlerClient{cc}
}

func (c *commentsGRPCHandlerClient) Comment(ctx context.Context, in *UserComment, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/comments.CommentsGRPCHandler/Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsGRPCHandlerClient) CommentsById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CommentsData, error) {
	out := new(CommentsData)
	err := c.cc.Invoke(ctx, "/comments.CommentsGRPCHandler/CommentsById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsGRPCHandlerServer is the server API for CommentsGRPCHandler service.
type CommentsGRPCHandlerServer interface {
	Comment(context.Context, *UserComment) (*Empty, error)
	CommentsById(context.Context, *Id) (*CommentsData, error)
}

// UnimplementedCommentsGRPCHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCommentsGRPCHandlerServer struct {
}

func (*UnimplementedCommentsGRPCHandlerServer) Comment(context.Context, *UserComment) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (*UnimplementedCommentsGRPCHandlerServer) CommentsById(context.Context, *Id) (*CommentsData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentsById not implemented")
}

func RegisterCommentsGRPCHandlerServer(s *grpc.Server, srv CommentsGRPCHandlerServer) {
	s.RegisterService(&_CommentsGRPCHandler_serviceDesc, srv)
}

func _CommentsGRPCHandler_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsGRPCHandlerServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comments.CommentsGRPCHandler/Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsGRPCHandlerServer).Comment(ctx, req.(*UserComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsGRPCHandler_CommentsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsGRPCHandlerServer).CommentsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comments.CommentsGRPCHandler/CommentsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsGRPCHandlerServer).CommentsById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentsGRPCHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comments.CommentsGRPCHandler",
	HandlerType: (*CommentsGRPCHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Comment",
			Handler:    _CommentsGRPCHandler_Comment_Handler,
		},
		{
			MethodName: "CommentsById",
			Handler:    _CommentsGRPCHandler_CommentsById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments.proto",
}