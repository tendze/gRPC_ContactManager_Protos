// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: cm/cm.proto

package cmv1

import (
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

type CreateContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreateContactRequest) Reset() {
	*x = CreateContactRequest{}
	mi := &file_cm_cm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactRequest) ProtoMessage() {}

func (x *CreateContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactRequest.ProtoReflect.Descriptor instead.
func (*CreateContactRequest) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateContactRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateContactRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CreateContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateContactResponse) Reset() {
	*x = CreateContactResponse{}
	mi := &file_cm_cm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactResponse) ProtoMessage() {}

func (x *CreateContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactResponse.ProtoReflect.Descriptor instead.
func (*CreateContactResponse) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContactResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateContactResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetContactByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetContactByNameRequest) Reset() {
	*x = GetContactByNameRequest{}
	mi := &file_cm_cm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByNameRequest) ProtoMessage() {}

func (x *GetContactByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByNameRequest.ProtoReflect.Descriptor instead.
func (*GetContactByNameRequest) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{2}
}

func (x *GetContactByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetContactByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetContactByEmailRequest) Reset() {
	*x = GetContactByEmailRequest{}
	mi := &file_cm_cm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByEmailRequest) ProtoMessage() {}

func (x *GetContactByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetContactByEmailRequest) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{3}
}

func (x *GetContactByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetContactByPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetContactByPhoneRequest) Reset() {
	*x = GetContactByPhoneRequest{}
	mi := &file_cm_cm_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactByPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByPhoneRequest) ProtoMessage() {}

func (x *GetContactByPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByPhoneRequest.ProtoReflect.Descriptor instead.
func (*GetContactByPhoneRequest) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{4}
}

func (x *GetContactByPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetContactResponse) Reset() {
	*x = GetContactResponse{}
	mi := &file_cm_cm_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactResponse) ProtoMessage() {}

func (x *GetContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactResponse.ProtoReflect.Descriptor instead.
func (*GetContactResponse) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{5}
}

func (x *GetContactResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetContactResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetContactResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetContactResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type DeleteContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteContactRequest) Reset() {
	*x = DeleteContactRequest{}
	mi := &file_cm_cm_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContactRequest) ProtoMessage() {}

func (x *DeleteContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContactRequest.ProtoReflect.Descriptor instead.
func (*DeleteContactRequest) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteContactRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteContactResponse) Reset() {
	*x = DeleteContactResponse{}
	mi := &file_cm_cm_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContactResponse) ProtoMessage() {}

func (x *DeleteContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cm_cm_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContactResponse.ProtoReflect.Descriptor instead.
func (*DeleteContactResponse) Descriptor() ([]byte, []int) {
	return file_cm_cm_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteContactResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_cm_cm_proto protoreflect.FileDescriptor

var file_cm_cm_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6d, 0x2f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x56, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x64, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xf3, 0x03, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x5c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x24, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x28, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x61, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42,
	0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x28, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x64, 0x61, 0x6e, 0x67, 0x2e, 0x63, 0x6d, 0x2e, 0x76, 0x31,
	0x3b, 0x63, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cm_cm_proto_rawDescOnce sync.Once
	file_cm_cm_proto_rawDescData = file_cm_cm_proto_rawDesc
)

func file_cm_cm_proto_rawDescGZIP() []byte {
	file_cm_cm_proto_rawDescOnce.Do(func() {
		file_cm_cm_proto_rawDescData = protoimpl.X.CompressGZIP(file_cm_cm_proto_rawDescData)
	})
	return file_cm_cm_proto_rawDescData
}

var file_cm_cm_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cm_cm_proto_goTypes = []any{
	(*CreateContactRequest)(nil),     // 0: ContactManager.CreateContactRequest
	(*CreateContactResponse)(nil),    // 1: ContactManager.CreateContactResponse
	(*GetContactByNameRequest)(nil),  // 2: ContactManager.GetContactByNameRequest
	(*GetContactByEmailRequest)(nil), // 3: ContactManager.GetContactByEmailRequest
	(*GetContactByPhoneRequest)(nil), // 4: ContactManager.GetContactByPhoneRequest
	(*GetContactResponse)(nil),       // 5: ContactManager.GetContactResponse
	(*DeleteContactRequest)(nil),     // 6: ContactManager.DeleteContactRequest
	(*DeleteContactResponse)(nil),    // 7: ContactManager.DeleteContactResponse
}
var file_cm_cm_proto_depIdxs = []int32{
	0, // 0: ContactManager.ContactManager.CreateContact:input_type -> ContactManager.CreateContactRequest
	2, // 1: ContactManager.ContactManager.GetContactByName:input_type -> ContactManager.GetContactByNameRequest
	3, // 2: ContactManager.ContactManager.GetContactByEmail:input_type -> ContactManager.GetContactByEmailRequest
	4, // 3: ContactManager.ContactManager.GetContactByPhone:input_type -> ContactManager.GetContactByPhoneRequest
	6, // 4: ContactManager.ContactManager.DeleteContact:input_type -> ContactManager.DeleteContactRequest
	1, // 5: ContactManager.ContactManager.CreateContact:output_type -> ContactManager.CreateContactResponse
	5, // 6: ContactManager.ContactManager.GetContactByName:output_type -> ContactManager.GetContactResponse
	5, // 7: ContactManager.ContactManager.GetContactByEmail:output_type -> ContactManager.GetContactResponse
	5, // 8: ContactManager.ContactManager.GetContactByPhone:output_type -> ContactManager.GetContactResponse
	7, // 9: ContactManager.ContactManager.DeleteContact:output_type -> ContactManager.DeleteContactResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cm_cm_proto_init() }
func file_cm_cm_proto_init() {
	if File_cm_cm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cm_cm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cm_cm_proto_goTypes,
		DependencyIndexes: file_cm_cm_proto_depIdxs,
		MessageInfos:      file_cm_cm_proto_msgTypes,
	}.Build()
	File_cm_cm_proto = out.File
	file_cm_cm_proto_rawDesc = nil
	file_cm_cm_proto_goTypes = nil
	file_cm_cm_proto_depIdxs = nil
}
