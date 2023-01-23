// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
	ID       string `protobuf:"bytes,5,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
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
	return file_api_proto_rawDescGZIP(), []int{1}
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

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *UserDto) Reset() {
	*x = UserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDto) ProtoMessage() {}

func (x *UserDto) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDto.ProtoReflect.Descriptor instead.
func (*UserDto) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *UserDto) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AuthDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDto []*UserDto `protobuf:"bytes,1,rep,name=UserDto,proto3" json:"UserDto,omitempty"`
	Token   string     `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthDto) Reset() {
	*x = AuthDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthDto) ProtoMessage() {}

func (x *AuthDto) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthDto.ProtoReflect.Descriptor instead.
func (*AuthDto) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AuthDto) GetUserDto() []*UserDto {
	if x != nil {
		return x.UserDto
	}
	return nil
}

func (x *AuthDto) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Authentication) Reset() {
	*x = Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication) ProtoMessage() {}

func (x *Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication.ProtoReflect.Descriptor instead.
func (*Authentication) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Authentication) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Authentication) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *UserID) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res []*anypb.Any `protobuf:"bytes,1,rep,name=res,proto3" json:"res,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetRes() []*anypb.Any {
	if x != nil {
		return x.Res
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Role        string `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	TokenString string `protobuf:"bytes,4,opt,name=TokenString,proto3" json:"TokenString,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *Token) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Token) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Token) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Token) GetTokenString() string {
	if x != nil {
		return x.TokenString
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x22, 0x70, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x4c, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x44, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x40, 0x0a,
	0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x20, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0xc6, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x36, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(*Null)(nil),           // 0: apiproto.Null
	(*User)(nil),           // 1: apiproto.user
	(*UserDto)(nil),        // 2: apiproto.UserDto
	(*AuthDto)(nil),        // 3: apiproto.AuthDto
	(*Authentication)(nil), // 4: apiproto.Authentication
	(*UserID)(nil),         // 5: apiproto.userID
	(*Response)(nil),       // 6: apiproto.response
	(*Token)(nil),          // 7: apiproto.Token
	(*anypb.Any)(nil),      // 8: google.protobuf.Any
}
var file_api_proto_depIdxs = []int32{
	2, // 0: apiproto.AuthDto.UserDto:type_name -> apiproto.UserDto
	8, // 1: apiproto.response.res:type_name -> google.protobuf.Any
	1, // 2: apiproto.UserService.CreateUser:input_type -> apiproto.user
	5, // 3: apiproto.UserService.GetUser:input_type -> apiproto.userID
	5, // 4: apiproto.UserService.DeleteUserById:input_type -> apiproto.userID
	1, // 5: apiproto.UserService.UpdateUserById:input_type -> apiproto.user
	0, // 6: apiproto.UserService.GetAllUser:input_type -> apiproto.Null
	4, // 7: apiproto.UserService.SignIn:input_type -> apiproto.Authentication
	3, // 8: apiproto.UserService.CreateUser:output_type -> apiproto.AuthDto
	2, // 9: apiproto.UserService.GetUser:output_type -> apiproto.UserDto
	6, // 10: apiproto.UserService.DeleteUserById:output_type -> apiproto.response
	6, // 11: apiproto.UserService.UpdateUserById:output_type -> apiproto.response
	6, // 12: apiproto.UserService.GetAllUser:output_type -> apiproto.response
	6, // 13: apiproto.UserService.SignIn:output_type -> apiproto.response
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDto); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthDto); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
