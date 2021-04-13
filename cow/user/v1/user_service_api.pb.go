// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: cow/user/v1/user_service_api.proto

package user

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

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *PlayerIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerRequest) GetIdentifier() *PlayerIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

type GetPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GetPlayerResponse) Reset() {
	*x = GetPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerResponse) ProtoMessage() {}

func (x *GetPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerResponse) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type GetPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiers []*PlayerIdentifier `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
}

func (x *GetPlayersRequest) Reset() {
	*x = GetPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersRequest) ProtoMessage() {}

func (x *GetPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersRequest.ProtoReflect.Descriptor instead.
func (*GetPlayersRequest) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlayersRequest) GetIdentifiers() []*PlayerIdentifier {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

type GetPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *GetPlayersResponse) Reset() {
	*x = GetPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersResponse) ProtoMessage() {}

func (x *GetPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersResponse.ProtoReflect.Descriptor instead.
func (*GetPlayersResponse) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type GetPlayerUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *PlayerIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetPlayerUserRequest) Reset() {
	*x = GetPlayerUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerUserRequest) ProtoMessage() {}

func (x *GetPlayerUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerUserRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerUserRequest) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerUserRequest) GetIdentifier() *PlayerIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

type GetPlayerUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetPlayerUserResponse) Reset() {
	*x = GetPlayerUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerUserResponse) ProtoMessage() {}

func (x *GetPlayerUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerUserResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerUserResponse) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlayerUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserPlayersRequest) Reset() {
	*x = GetUserPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPlayersRequest) ProtoMessage() {}

func (x *GetUserPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPlayersRequest.ProtoReflect.Descriptor instead.
func (*GetUserPlayersRequest) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserPlayersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *GetUserPlayersResponse) Reset() {
	*x = GetUserPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPlayersResponse) ProtoMessage() {}

func (x *GetUserPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPlayersResponse.ProtoReflect.Descriptor instead.
func (*GetUserPlayersResponse) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserPlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[10]
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
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{10}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                            // random uuid
	ReferenceId   string `protobuf:"bytes,2,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`       // e.g. minecraft uuid
	ReferenceType string `protobuf:"bytes,3,opt,name=reference_type,json=referenceType,proto3" json:"reference_type,omitempty"` // e.g. "minecraft"
	Username      string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`                                // e.g. minecraft username
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{11}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Player) GetReferenceType() string {
	if x != nil {
		return x.ReferenceType
	}
	return ""
}

func (x *Player) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// The identifier used to determine which player to return.
type PlayerIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reference id of the corresponding source (e.g. the minecraft uuid).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the reference/source id (e.g. "minecraft")
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The username assigned to the reference/source account (e.g. the minecraft username).
	Username *string `protobuf:"bytes,3,opt,name=username,proto3,oneof" json:"username,omitempty"`
}

func (x *PlayerIdentifier) Reset() {
	*x = PlayerIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_user_v1_user_service_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerIdentifier) ProtoMessage() {}

func (x *PlayerIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_cow_user_v1_user_service_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerIdentifier.ProtoReflect.Descriptor instead.
func (*PlayerIdentifier) Descriptor() ([]byte, []int) {
	return file_cow_user_v1_user_service_api_proto_rawDescGZIP(), []int{12}
}

func (x *PlayerIdentifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayerIdentifier) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PlayerIdentifier) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

var File_cow_user_v1_user_service_api_proto protoreflect.FileDescriptor

var file_cow_user_v1_user_service_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x77, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x77, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x77, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x22, 0x55, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x30, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x47, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x77,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x7e, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x64, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa1, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x0a, 0x1b, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x6d, 0x6f, 0x6f, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x77, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cow_user_v1_user_service_api_proto_rawDescOnce sync.Once
	file_cow_user_v1_user_service_api_proto_rawDescData = file_cow_user_v1_user_service_api_proto_rawDesc
)

func file_cow_user_v1_user_service_api_proto_rawDescGZIP() []byte {
	file_cow_user_v1_user_service_api_proto_rawDescOnce.Do(func() {
		file_cow_user_v1_user_service_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cow_user_v1_user_service_api_proto_rawDescData)
	})
	return file_cow_user_v1_user_service_api_proto_rawDescData
}

var file_cow_user_v1_user_service_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_cow_user_v1_user_service_api_proto_goTypes = []interface{}{
	(*GetPlayerRequest)(nil),       // 0: cow.user.v1.GetPlayerRequest
	(*GetPlayerResponse)(nil),      // 1: cow.user.v1.GetPlayerResponse
	(*GetPlayersRequest)(nil),      // 2: cow.user.v1.GetPlayersRequest
	(*GetPlayersResponse)(nil),     // 3: cow.user.v1.GetPlayersResponse
	(*GetPlayerUserRequest)(nil),   // 4: cow.user.v1.GetPlayerUserRequest
	(*GetPlayerUserResponse)(nil),  // 5: cow.user.v1.GetPlayerUserResponse
	(*GetUserRequest)(nil),         // 6: cow.user.v1.GetUserRequest
	(*GetUserResponse)(nil),        // 7: cow.user.v1.GetUserResponse
	(*GetUserPlayersRequest)(nil),  // 8: cow.user.v1.GetUserPlayersRequest
	(*GetUserPlayersResponse)(nil), // 9: cow.user.v1.GetUserPlayersResponse
	(*User)(nil),                   // 10: cow.user.v1.User
	(*Player)(nil),                 // 11: cow.user.v1.Player
	(*PlayerIdentifier)(nil),       // 12: cow.user.v1.PlayerIdentifier
}
var file_cow_user_v1_user_service_api_proto_depIdxs = []int32{
	12, // 0: cow.user.v1.GetPlayerRequest.identifier:type_name -> cow.user.v1.PlayerIdentifier
	11, // 1: cow.user.v1.GetPlayerResponse.player:type_name -> cow.user.v1.Player
	12, // 2: cow.user.v1.GetPlayersRequest.identifiers:type_name -> cow.user.v1.PlayerIdentifier
	11, // 3: cow.user.v1.GetPlayersResponse.players:type_name -> cow.user.v1.Player
	12, // 4: cow.user.v1.GetPlayerUserRequest.identifier:type_name -> cow.user.v1.PlayerIdentifier
	10, // 5: cow.user.v1.GetPlayerUserResponse.user:type_name -> cow.user.v1.User
	10, // 6: cow.user.v1.GetUserResponse.user:type_name -> cow.user.v1.User
	11, // 7: cow.user.v1.GetUserPlayersResponse.players:type_name -> cow.user.v1.Player
	0,  // 8: cow.user.v1.UserService.GetPlayer:input_type -> cow.user.v1.GetPlayerRequest
	2,  // 9: cow.user.v1.UserService.GetPlayers:input_type -> cow.user.v1.GetPlayersRequest
	4,  // 10: cow.user.v1.UserService.GetPlayerUser:input_type -> cow.user.v1.GetPlayerUserRequest
	6,  // 11: cow.user.v1.UserService.GetUser:input_type -> cow.user.v1.GetUserRequest
	8,  // 12: cow.user.v1.UserService.GetUserPlayers:input_type -> cow.user.v1.GetUserPlayersRequest
	1,  // 13: cow.user.v1.UserService.GetPlayer:output_type -> cow.user.v1.GetPlayerResponse
	3,  // 14: cow.user.v1.UserService.GetPlayers:output_type -> cow.user.v1.GetPlayersResponse
	5,  // 15: cow.user.v1.UserService.GetPlayerUser:output_type -> cow.user.v1.GetPlayerUserResponse
	7,  // 16: cow.user.v1.UserService.GetUser:output_type -> cow.user.v1.GetUserResponse
	9,  // 17: cow.user.v1.UserService.GetUserPlayers:output_type -> cow.user.v1.GetUserPlayersResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cow_user_v1_user_service_api_proto_init() }
func file_cow_user_v1_user_service_api_proto_init() {
	if File_cow_user_v1_user_service_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cow_user_v1_user_service_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerRequest); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerResponse); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersRequest); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersResponse); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerUserRequest); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerUserResponse); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPlayersRequest); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPlayersResponse); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_cow_user_v1_user_service_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerIdentifier); i {
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
	file_cow_user_v1_user_service_api_proto_msgTypes[12].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cow_user_v1_user_service_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cow_user_v1_user_service_api_proto_goTypes,
		DependencyIndexes: file_cow_user_v1_user_service_api_proto_depIdxs,
		MessageInfos:      file_cow_user_v1_user_service_api_proto_msgTypes,
	}.Build()
	File_cow_user_v1_user_service_api_proto = out.File
	file_cow_user_v1_user_service_api_proto_rawDesc = nil
	file_cow_user_v1_user_service_api_proto_goTypes = nil
	file_cow_user_v1_user_service_api_proto_depIdxs = nil
}
