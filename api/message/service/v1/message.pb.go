// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: v1/message.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid int32    `protobuf:"varint,1,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`
	RecverUid int32    `protobuf:"varint,2,opt,name=recver_uid,json=recverUid,proto3" json:"recver_uid,omitempty"`
	Content   *Content `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Ctime     int64    `protobuf:"varint,4,opt,name=ctime,proto3" json:"ctime,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSenderUid() int32 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *Message) GetRecverUid() int32 {
	if x != nil {
		return x.RecverUid
	}
	return 0
}

func (x *Message) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Content) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ChatHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid int32 `protobuf:"varint,1,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`
	RecverUid int32 `protobuf:"varint,2,opt,name=recver_uid,json=recverUid,proto3" json:"recver_uid,omitempty"`
	Page      int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Count     int32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ChatHistoryRequest) Reset() {
	*x = ChatHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHistoryRequest) ProtoMessage() {}

func (x *ChatHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHistoryRequest.ProtoReflect.Descriptor instead.
func (*ChatHistoryRequest) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *ChatHistoryRequest) GetSenderUid() int32 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *ChatHistoryRequest) GetRecverUid() int32 {
	if x != nil {
		return x.RecverUid
	}
	return 0
}

func (x *ChatHistoryRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ChatHistoryRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ChatHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs []*Message `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (x *ChatHistoryReply) Reset() {
	*x = ChatHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHistoryReply) ProtoMessage() {}

func (x *ChatHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHistoryReply.ProtoReflect.Descriptor instead.
func (*ChatHistoryReply) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *ChatHistoryReply) GetMsgs() []*Message {
	if x != nil {
		return x.Msgs
	}
	return nil
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid int32    `protobuf:"varint,1,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`
	RecverUid int32    `protobuf:"varint,2,opt,name=recver_uid,json=recverUid,proto3" json:"recver_uid,omitempty"`
	Content   *Content `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageRequest) GetSenderUid() int32 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *SendMessageRequest) GetRecverUid() int32 {
	if x != nil {
		return x.RecverUid
	}
	return 0
}

func (x *SendMessageRequest) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{5}
}

var File_v1_message_proto protoreflect.FileDescriptor

var file_v1_message_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x65,
	0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x7c, 0x0a, 0x12, 0x43, 0x68,
	0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x04,
	0x6d, 0x73, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xeb, 0x01,
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x6c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x6c, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x2c, 0x5a, 0x2a, 0x66,
	0x78, 0x6b, 0x74, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x72, 0x61, 0x69, 0x64, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_message_proto_rawDescOnce sync.Once
	file_v1_message_proto_rawDescData = file_v1_message_proto_rawDesc
)

func file_v1_message_proto_rawDescGZIP() []byte {
	file_v1_message_proto_rawDescOnce.Do(func() {
		file_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_message_proto_rawDescData)
	})
	return file_v1_message_proto_rawDescData
}

var file_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_message_proto_goTypes = []interface{}{
	(*Message)(nil),            // 0: message.v1.Message
	(*Content)(nil),            // 1: message.v1.Content
	(*ChatHistoryRequest)(nil), // 2: message.v1.ChatHistoryRequest
	(*ChatHistoryReply)(nil),   // 3: message.v1.ChatHistoryReply
	(*SendMessageRequest)(nil), // 4: message.v1.SendMessageRequest
	(*SendMessageReply)(nil),   // 5: message.v1.SendMessageReply
}
var file_v1_message_proto_depIdxs = []int32{
	1, // 0: message.v1.Message.content:type_name -> message.v1.Content
	0, // 1: message.v1.ChatHistoryReply.msgs:type_name -> message.v1.Message
	1, // 2: message.v1.SendMessageRequest.content:type_name -> message.v1.Content
	2, // 3: message.v1.MessageSystem.ChatHistory:input_type -> message.v1.ChatHistoryRequest
	4, // 4: message.v1.MessageSystem.SendMessage:input_type -> message.v1.SendMessageRequest
	3, // 5: message.v1.MessageSystem.ChatHistory:output_type -> message.v1.ChatHistoryReply
	5, // 6: message.v1.MessageSystem.SendMessage:output_type -> message.v1.SendMessageReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_message_proto_init() }
func file_v1_message_proto_init() {
	if File_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatHistoryRequest); i {
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
		file_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatHistoryReply); i {
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
		file_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
			RawDescriptor: file_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_message_proto_goTypes,
		DependencyIndexes: file_v1_message_proto_depIdxs,
		MessageInfos:      file_v1_message_proto_msgTypes,
	}.Build()
	File_v1_message_proto = out.File
	file_v1_message_proto_rawDesc = nil
	file_v1_message_proto_goTypes = nil
	file_v1_message_proto_depIdxs = nil
}