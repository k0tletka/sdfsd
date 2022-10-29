// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protobuf/serverapi.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PoolMode int32

const (
	PoolMode_EXTEND_MODE PoolMode = 0
	PoolMode_BACKUP_MODE PoolMode = 1
)

// Enum value maps for PoolMode.
var (
	PoolMode_name = map[int32]string{
		0: "EXTEND_MODE",
		1: "BACKUP_MODE",
	}
	PoolMode_value = map[string]int32{
		"EXTEND_MODE": 0,
		"BACKUP_MODE": 1,
	}
)

func (x PoolMode) Enum() *PoolMode {
	p := new(PoolMode)
	*p = x
	return p
}

func (x PoolMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PoolMode) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_serverapi_proto_enumTypes[0].Descriptor()
}

func (PoolMode) Type() protoreflect.EnumType {
	return &file_protobuf_serverapi_proto_enumTypes[0]
}

func (x PoolMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PoolMode.Descriptor instead.
func (PoolMode) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{0}
}

type ErrorCode int32

const (
	ErrorCode_OBJECT_NOT_FOUND ErrorCode = 0
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "OBJECT_NOT_FOUND",
	}
	ErrorCode_value = map[string]int32{
		"OBJECT_NOT_FOUND": 0,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_serverapi_proto_enumTypes[1].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_protobuf_serverapi_proto_enumTypes[1]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{1}
}

type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    ErrorCode        `protobuf:"varint,1,opt,name=errorCode,proto3,enum=ErrorCode" json:"errorCode,omitempty"`
	ErrorMessage string           `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Additional   *structpb.Struct `protobuf:"bytes,3,opt,name=additional,proto3" json:"additional,omitempty"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorInfo) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_OBJECT_NOT_FOUND
}

func (x *ErrorInfo) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ErrorInfo) GetAdditional() *structpb.Struct {
	if x != nil {
		return x.Additional
	}
	return nil
}

type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolName string   `protobuf:"bytes,1,opt,name=poolName,proto3" json:"poolName,omitempty"`
	PoolMode PoolMode `protobuf:"varint,2,opt,name=poolMode,proto3,enum=PoolMode" json:"poolMode,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{1}
}

func (x *Pool) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

func (x *Pool) GetPoolMode() PoolMode {
	if x != nil {
		return x.PoolMode
	}
	return PoolMode_EXTEND_MODE
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeName string `protobuf:"bytes,1,opt,name=volumeName,proto3" json:"volumeName,omitempty"`
	VolumeSize uint64 `protobuf:"varint,2,opt,name=volumeSize,proto3" json:"volumeSize,omitempty"`
	PoolName   string `protobuf:"bytes,3,opt,name=poolName,proto3" json:"poolName,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{2}
}

func (x *Volume) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

func (x *Volume) GetVolumeSize() uint64 {
	if x != nil {
		return x.VolumeSize
	}
	return 0
}

func (x *Volume) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

type ServerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	ApiVersion uint64 `protobuf:"varint,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
}

func (x *ServerInfoResponse) Reset() {
	*x = ServerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoResponse) ProtoMessage() {}

func (x *ServerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoResponse.ProtoReflect.Descriptor instead.
func (*ServerInfoResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{3}
}

func (x *ServerInfoResponse) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ServerInfoResponse) GetApiVersion() uint64 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

type PoolListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolList []*Pool `protobuf:"bytes,1,rep,name=poolList,proto3" json:"poolList,omitempty"`
}

func (x *PoolListResponse) Reset() {
	*x = PoolListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolListResponse) ProtoMessage() {}

func (x *PoolListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolListResponse.ProtoReflect.Descriptor instead.
func (*PoolListResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{4}
}

func (x *PoolListResponse) GetPoolList() []*Pool {
	if x != nil {
		return x.PoolList
	}
	return nil
}

type PoolInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolName string `protobuf:"bytes,1,opt,name=poolName,proto3" json:"poolName,omitempty"`
}

func (x *PoolInfoRequest) Reset() {
	*x = PoolInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolInfoRequest) ProtoMessage() {}

func (x *PoolInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolInfoRequest.ProtoReflect.Descriptor instead.
func (*PoolInfoRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{5}
}

func (x *PoolInfoRequest) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

type PoolInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	Pool      *Pool      `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (x *PoolInfoResponse) Reset() {
	*x = PoolInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolInfoResponse) ProtoMessage() {}

func (x *PoolInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolInfoResponse.ProtoReflect.Descriptor instead.
func (*PoolInfoResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{6}
}

func (x *PoolInfoResponse) GetErrorInfo() *ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *PoolInfoResponse) GetPool() *Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

type VolumeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeList []*Volume `protobuf:"bytes,1,rep,name=volumeList,proto3" json:"volumeList,omitempty"`
}

func (x *VolumeListResponse) Reset() {
	*x = VolumeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_serverapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeListResponse) ProtoMessage() {}

func (x *VolumeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_serverapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeListResponse.ProtoReflect.Descriptor instead.
func (*VolumeListResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_serverapi_proto_rawDescGZIP(), []int{7}
}

func (x *VolumeListResponse) GetVolumeList() []*Volume {
	if x != nil {
		return x.VolumeList
	}
	return nil
}

var File_protobuf_serverapi_proto protoreflect.FileDescriptor

var file_protobuf_serverapi_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x49, 0x0a, 0x04, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x6f,
	0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x64, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x35, 0x0a, 0x10, 0x50, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x08,
	0x70, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x50, 0x6f, 0x6f, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x10, 0x50, 0x6f, 0x6f, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c,
	0x22, 0x3d, 0x0a, 0x12, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2a,
	0x2c, 0x0a, 0x08, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x58, 0x54, 0x45, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x01, 0x2a, 0x21, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00,
	0x32, 0xef, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x3c,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x11, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x30, 0x74, 0x6c, 0x65, 0x74, 0x6b, 0x61, 0x2f, 0x73, 0x64, 0x66, 0x73, 0x64, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_serverapi_proto_rawDescOnce sync.Once
	file_protobuf_serverapi_proto_rawDescData = file_protobuf_serverapi_proto_rawDesc
)

func file_protobuf_serverapi_proto_rawDescGZIP() []byte {
	file_protobuf_serverapi_proto_rawDescOnce.Do(func() {
		file_protobuf_serverapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_serverapi_proto_rawDescData)
	})
	return file_protobuf_serverapi_proto_rawDescData
}

var file_protobuf_serverapi_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobuf_serverapi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_serverapi_proto_goTypes = []interface{}{
	(PoolMode)(0),              // 0: PoolMode
	(ErrorCode)(0),             // 1: ErrorCode
	(*ErrorInfo)(nil),          // 2: ErrorInfo
	(*Pool)(nil),               // 3: Pool
	(*Volume)(nil),             // 4: Volume
	(*ServerInfoResponse)(nil), // 5: ServerInfoResponse
	(*PoolListResponse)(nil),   // 6: PoolListResponse
	(*PoolInfoRequest)(nil),    // 7: PoolInfoRequest
	(*PoolInfoResponse)(nil),   // 8: PoolInfoResponse
	(*VolumeListResponse)(nil), // 9: VolumeListResponse
	(*structpb.Struct)(nil),    // 10: google.protobuf.Struct
	(*emptypb.Empty)(nil),      // 11: google.protobuf.Empty
}
var file_protobuf_serverapi_proto_depIdxs = []int32{
	1,  // 0: ErrorInfo.errorCode:type_name -> ErrorCode
	10, // 1: ErrorInfo.additional:type_name -> google.protobuf.Struct
	0,  // 2: Pool.poolMode:type_name -> PoolMode
	3,  // 3: PoolListResponse.poolList:type_name -> Pool
	2,  // 4: PoolInfoResponse.errorInfo:type_name -> ErrorInfo
	3,  // 5: PoolInfoResponse.pool:type_name -> Pool
	4,  // 6: VolumeListResponse.volumeList:type_name -> Volume
	11, // 7: ServerAPI.GetServerInfo:input_type -> google.protobuf.Empty
	11, // 8: ServerAPI.GetPools:input_type -> google.protobuf.Empty
	7,  // 9: ServerAPI.GetPoolInfo:input_type -> PoolInfoRequest
	11, // 10: ServerAPI.GetVolumes:input_type -> google.protobuf.Empty
	5,  // 11: ServerAPI.GetServerInfo:output_type -> ServerInfoResponse
	6,  // 12: ServerAPI.GetPools:output_type -> PoolListResponse
	8,  // 13: ServerAPI.GetPoolInfo:output_type -> PoolInfoResponse
	9,  // 14: ServerAPI.GetVolumes:output_type -> VolumeListResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protobuf_serverapi_proto_init() }
func file_protobuf_serverapi_proto_init() {
	if File_protobuf_serverapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_serverapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
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
		file_protobuf_serverapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
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
		file_protobuf_serverapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_protobuf_serverapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoResponse); i {
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
		file_protobuf_serverapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolListResponse); i {
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
		file_protobuf_serverapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolInfoRequest); i {
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
		file_protobuf_serverapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolInfoResponse); i {
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
		file_protobuf_serverapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeListResponse); i {
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
			RawDescriptor: file_protobuf_serverapi_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_serverapi_proto_goTypes,
		DependencyIndexes: file_protobuf_serverapi_proto_depIdxs,
		EnumInfos:         file_protobuf_serverapi_proto_enumTypes,
		MessageInfos:      file_protobuf_serverapi_proto_msgTypes,
	}.Build()
	File_protobuf_serverapi_proto = out.File
	file_protobuf_serverapi_proto_rawDesc = nil
	file_protobuf_serverapi_proto_goTypes = nil
	file_protobuf_serverapi_proto_depIdxs = nil
}
