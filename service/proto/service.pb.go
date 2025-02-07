// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.12.4
// source: google/rpc/service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContainerStatus int32

const (
	ContainerStatus_CONTAINER_UNKNOWN ContainerStatus = 0
	ContainerStatus_CONTAINER_RUNNING ContainerStatus = 1
	ContainerStatus_CONTAINER_KILLED  ContainerStatus = 2
)

// Enum value maps for ContainerStatus.
var (
	ContainerStatus_name = map[int32]string{
		0: "CONTAINER_UNKNOWN",
		1: "CONTAINER_RUNNING",
		2: "CONTAINER_KILLED",
	}
	ContainerStatus_value = map[string]int32{
		"CONTAINER_UNKNOWN": 0,
		"CONTAINER_RUNNING": 1,
		"CONTAINER_KILLED":  2,
	}
)

func (x ContainerStatus) Enum() *ContainerStatus {
	p := new(ContainerStatus)
	*p = x
	return p
}

func (x ContainerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_rpc_service_proto_enumTypes[0].Descriptor()
}

func (ContainerStatus) Type() protoreflect.EnumType {
	return &file_google_rpc_service_proto_enumTypes[0]
}

func (x ContainerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerStatus.Descriptor instead.
func (ContainerStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_rpc_service_proto_rawDescGZIP(), []int{0}
}

type ContainersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainersRequest) Reset() {
	*x = ContainersRequest{}
	mi := &file_google_rpc_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainersRequest) ProtoMessage() {}

func (x *ContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_rpc_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainersRequest.ProtoReflect.Descriptor instead.
func (*ContainersRequest) Descriptor() ([]byte, []int) {
	return file_google_rpc_service_proto_rawDescGZIP(), []int{0}
}

type ContainerInfos struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        ContainerStatus        `protobuf:"varint,2,opt,name=status,proto3,enum=daemon.ContainerStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerInfos) Reset() {
	*x = ContainerInfos{}
	mi := &file_google_rpc_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfos) ProtoMessage() {}

func (x *ContainerInfos) ProtoReflect() protoreflect.Message {
	mi := &file_google_rpc_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfos.ProtoReflect.Descriptor instead.
func (*ContainerInfos) Descriptor() ([]byte, []int) {
	return file_google_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerInfos) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerInfos) GetStatus() ContainerStatus {
	if x != nil {
		return x.Status
	}
	return ContainerStatus_CONTAINER_UNKNOWN
}

type ContainersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Containers    []*ContainerInfos      `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainersResponse) Reset() {
	*x = ContainersResponse{}
	mi := &file_google_rpc_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainersResponse) ProtoMessage() {}

func (x *ContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_rpc_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainersResponse.ProtoReflect.Descriptor instead.
func (*ContainersResponse) Descriptor() ([]byte, []int) {
	return file_google_rpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *ContainersResponse) GetContainers() []*ContainerInfos {
	if x != nil {
		return x.Containers
	}
	return nil
}

var File_google_rpc_service_proto protoreflect.FileDescriptor

var file_google_rpc_service_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2a, 0x55, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32,
	0x57, 0x0a, 0x0d, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_google_rpc_service_proto_rawDescOnce sync.Once
	file_google_rpc_service_proto_rawDescData []byte
)

func file_google_rpc_service_proto_rawDescGZIP() []byte {
	file_google_rpc_service_proto_rawDescOnce.Do(func() {
		file_google_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_rpc_service_proto_rawDesc), len(file_google_rpc_service_proto_rawDesc)))
	})
	return file_google_rpc_service_proto_rawDescData
}

var file_google_rpc_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_rpc_service_proto_goTypes = []any{
	(ContainerStatus)(0),       // 0: daemon.ContainerStatus
	(*ContainersRequest)(nil),  // 1: daemon.ContainersRequest
	(*ContainerInfos)(nil),     // 2: daemon.ContainerInfos
	(*ContainersResponse)(nil), // 3: daemon.ContainersResponse
}
var file_google_rpc_service_proto_depIdxs = []int32{
	0, // 0: daemon.ContainerInfos.status:type_name -> daemon.ContainerStatus
	2, // 1: daemon.ContainersResponse.containers:type_name -> daemon.ContainerInfos
	1, // 2: daemon.DaemonService.GetContainers:input_type -> daemon.ContainersRequest
	3, // 3: daemon.DaemonService.GetContainers:output_type -> daemon.ContainersResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_rpc_service_proto_init() }
func file_google_rpc_service_proto_init() {
	if File_google_rpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_rpc_service_proto_rawDesc), len(file_google_rpc_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_rpc_service_proto_goTypes,
		DependencyIndexes: file_google_rpc_service_proto_depIdxs,
		EnumInfos:         file_google_rpc_service_proto_enumTypes,
		MessageInfos:      file_google_rpc_service_proto_msgTypes,
	}.Build()
	File_google_rpc_service_proto = out.File
	file_google_rpc_service_proto_goTypes = nil
	file_google_rpc_service_proto_depIdxs = nil
}
