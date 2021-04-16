// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: cow/instance/v1/instance_service_api.proto

package instance

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

type CreateInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name the instance should have
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Kubernetes instance manifest or name of
	// a manifest template configured in the service
	Manifest string `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
}

func (x *CreateInstanceRequest) Reset() {
	*x = CreateInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceRequest) ProtoMessage() {}

func (x *CreateInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceRequest.ProtoReflect.Descriptor instead.
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInstanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateInstanceRequest) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

type CreateInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instance created by the service
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *CreateInstanceResponse) Reset() {
	*x = CreateInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceResponse) ProtoMessage() {}

func (x *CreateInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceResponse.ProtoReflect.Descriptor instead.
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInstanceResponse) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type EndInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EndInstanceRequest) Reset() {
	*x = EndInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndInstanceRequest) ProtoMessage() {}

func (x *EndInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndInstanceRequest.ProtoReflect.Descriptor instead.
func (*EndInstanceRequest) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{2}
}

func (x *EndInstanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instance marked as ended by the service
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *EndInstanceResponse) Reset() {
	*x = EndInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndInstanceResponse) ProtoMessage() {}

func (x *EndInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndInstanceResponse.ProtoReflect.Descriptor instead.
func (*EndInstanceResponse) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{3}
}

func (x *EndInstanceResponse) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type GetInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInstanceRequest) Reset() {
	*x = GetInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceRequest) ProtoMessage() {}

func (x *GetInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceRequest.ProtoReflect.Descriptor instead.
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetInstanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instance requested
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *GetInstanceResponse) Reset() {
	*x = GetInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceResponse) ProtoMessage() {}

func (x *GetInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cow_instance_v1_instance_service_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceResponse.ProtoReflect.Descriptor instead.
func (*GetInstanceResponse) Descriptor() ([]byte, []int) {
	return file_cow_instance_v1_instance_service_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetInstanceResponse) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

var File_cow_instance_v1_instance_service_api_proto protoreflect.FileDescriptor

var file_cow_instance_v1_instance_service_api_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x77, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f,
	0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63,
	0x6f, 0x77, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x45, 0x6e,
	0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x90, 0x02, 0x0a,
	0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x77,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x03, 0x45,
	0x6e, 0x64, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x77, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x4f, 0x0a, 0x1f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x77, 0x2e, 0x6d,
	0x6f, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x77, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cow_instance_v1_instance_service_api_proto_rawDescOnce sync.Once
	file_cow_instance_v1_instance_service_api_proto_rawDescData = file_cow_instance_v1_instance_service_api_proto_rawDesc
)

func file_cow_instance_v1_instance_service_api_proto_rawDescGZIP() []byte {
	file_cow_instance_v1_instance_service_api_proto_rawDescOnce.Do(func() {
		file_cow_instance_v1_instance_service_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cow_instance_v1_instance_service_api_proto_rawDescData)
	})
	return file_cow_instance_v1_instance_service_api_proto_rawDescData
}

var file_cow_instance_v1_instance_service_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cow_instance_v1_instance_service_api_proto_goTypes = []interface{}{
	(*CreateInstanceRequest)(nil),  // 0: cow.instance.v1.CreateInstanceRequest
	(*CreateInstanceResponse)(nil), // 1: cow.instance.v1.CreateInstanceResponse
	(*EndInstanceRequest)(nil),     // 2: cow.instance.v1.EndInstanceRequest
	(*EndInstanceResponse)(nil),    // 3: cow.instance.v1.EndInstanceResponse
	(*GetInstanceRequest)(nil),     // 4: cow.instance.v1.GetInstanceRequest
	(*GetInstanceResponse)(nil),    // 5: cow.instance.v1.GetInstanceResponse
	(*Instance)(nil),               // 6: cow.instance.v1.Instance
}
var file_cow_instance_v1_instance_service_api_proto_depIdxs = []int32{
	6, // 0: cow.instance.v1.CreateInstanceResponse.instance:type_name -> cow.instance.v1.Instance
	6, // 1: cow.instance.v1.EndInstanceResponse.instance:type_name -> cow.instance.v1.Instance
	6, // 2: cow.instance.v1.GetInstanceResponse.instance:type_name -> cow.instance.v1.Instance
	0, // 3: cow.instance.v1.InstanceService.Create:input_type -> cow.instance.v1.CreateInstanceRequest
	2, // 4: cow.instance.v1.InstanceService.End:input_type -> cow.instance.v1.EndInstanceRequest
	4, // 5: cow.instance.v1.InstanceService.Get:input_type -> cow.instance.v1.GetInstanceRequest
	1, // 6: cow.instance.v1.InstanceService.Create:output_type -> cow.instance.v1.CreateInstanceResponse
	3, // 7: cow.instance.v1.InstanceService.End:output_type -> cow.instance.v1.EndInstanceResponse
	5, // 8: cow.instance.v1.InstanceService.Get:output_type -> cow.instance.v1.GetInstanceResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cow_instance_v1_instance_service_api_proto_init() }
func file_cow_instance_v1_instance_service_api_proto_init() {
	if File_cow_instance_v1_instance_service_api_proto != nil {
		return
	}
	file_cow_instance_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cow_instance_v1_instance_service_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstanceRequest); i {
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
		file_cow_instance_v1_instance_service_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstanceResponse); i {
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
		file_cow_instance_v1_instance_service_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndInstanceRequest); i {
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
		file_cow_instance_v1_instance_service_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndInstanceResponse); i {
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
		file_cow_instance_v1_instance_service_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstanceRequest); i {
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
		file_cow_instance_v1_instance_service_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstanceResponse); i {
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
			RawDescriptor: file_cow_instance_v1_instance_service_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cow_instance_v1_instance_service_api_proto_goTypes,
		DependencyIndexes: file_cow_instance_v1_instance_service_api_proto_depIdxs,
		MessageInfos:      file_cow_instance_v1_instance_service_api_proto_msgTypes,
	}.Build()
	File_cow_instance_v1_instance_service_api_proto = out.File
	file_cow_instance_v1_instance_service_api_proto_rawDesc = nil
	file_cow_instance_v1_instance_service_api_proto_goTypes = nil
	file_cow_instance_v1_instance_service_api_proto_depIdxs = nil
}
