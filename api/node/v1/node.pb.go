// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.12
// source: node.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EvacRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PodInfo       *PodInfo               `protobuf:"bytes,1,opt,name=pod_info,json=podInfo,proto3" json:"pod_info,omitempty"`
	MigrationInfo *MigrationInfo         `protobuf:"bytes,2,opt,name=migration_info,json=migrationInfo,proto3" json:"migration_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvacRequest) Reset() {
	*x = EvacRequest{}
	mi := &file_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvacRequest) ProtoMessage() {}

func (x *EvacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvacRequest.ProtoReflect.Descriptor instead.
func (*EvacRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *EvacRequest) GetPodInfo() *PodInfo {
	if x != nil {
		return x.PodInfo
	}
	return nil
}

func (x *EvacRequest) GetMigrationInfo() *MigrationInfo {
	if x != nil {
		return x.MigrationInfo
	}
	return nil
}

type EvacResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Empty         *emptypb.Empty         `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvacResponse) Reset() {
	*x = EvacResponse{}
	mi := &file_node_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvacResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvacResponse) ProtoMessage() {}

func (x *EvacResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvacResponse.ProtoReflect.Descriptor instead.
func (*EvacResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *EvacResponse) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type RestoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PodInfo       *PodInfo               `protobuf:"bytes,1,opt,name=pod_info,json=podInfo,proto3" json:"pod_info,omitempty"`
	MigrationInfo *MigrationInfo         `protobuf:"bytes,2,opt,name=migration_info,json=migrationInfo,proto3" json:"migration_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RestoreRequest) Reset() {
	*x = RestoreRequest{}
	mi := &file_node_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreRequest) ProtoMessage() {}

func (x *RestoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreRequest.ProtoReflect.Descriptor instead.
func (*RestoreRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *RestoreRequest) GetPodInfo() *PodInfo {
	if x != nil {
		return x.PodInfo
	}
	return nil
}

func (x *RestoreRequest) GetMigrationInfo() *MigrationInfo {
	if x != nil {
		return x.MigrationInfo
	}
	return nil
}

type RestoreResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MigrationInfo *MigrationInfo         `protobuf:"bytes,1,opt,name=migration_info,json=migrationInfo,proto3" json:"migration_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RestoreResponse) Reset() {
	*x = RestoreResponse{}
	mi := &file_node_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreResponse) ProtoMessage() {}

func (x *RestoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreResponse.ProtoReflect.Descriptor instead.
func (*RestoreResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *RestoreResponse) GetMigrationInfo() *MigrationInfo {
	if x != nil {
		return x.MigrationInfo
	}
	return nil
}

type PodInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerName string                 `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	Ports         []int32                `protobuf:"varint,4,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	mi := &file_node_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *PodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodInfo) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *PodInfo) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

type MigrationInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageId       string                 `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	BundleDir     string                 `protobuf:"bytes,2,opt,name=bundle_dir,json=bundleDir,proto3" json:"bundle_dir,omitempty"`
	LiveMigration bool                   `protobuf:"varint,3,opt,name=live_migration,json=liveMigration,proto3" json:"live_migration,omitempty"`
	PausedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=paused_at,json=pausedAt,proto3" json:"paused_at,omitempty"`
	RestoreStart  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=restore_start,json=restoreStart,proto3" json:"restore_start,omitempty"`
	RestoreEnd    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=restore_end,json=restoreEnd,proto3" json:"restore_end,omitempty"`
	Ports         []int32                `protobuf:"varint,7,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MigrationInfo) Reset() {
	*x = MigrationInfo{}
	mi := &file_node_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MigrationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationInfo) ProtoMessage() {}

func (x *MigrationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationInfo.ProtoReflect.Descriptor instead.
func (*MigrationInfo) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *MigrationInfo) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *MigrationInfo) GetBundleDir() string {
	if x != nil {
		return x.BundleDir
	}
	return ""
}

func (x *MigrationInfo) GetLiveMigration() bool {
	if x != nil {
		return x.LiveMigration
	}
	return false
}

func (x *MigrationInfo) GetPausedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PausedAt
	}
	return nil
}

func (x *MigrationInfo) GetRestoreStart() *timestamppb.Timestamp {
	if x != nil {
		return x.RestoreStart
	}
	return nil
}

func (x *MigrationInfo) GetRestoreEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.RestoreEnd
	}
	return nil
}

func (x *MigrationInfo) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageData     []byte                 `protobuf:"bytes,1,opt,name=imageData,proto3" json:"imageData,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_node_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *Image) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

type CriuLazyPagesRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CheckpointPath string                 `protobuf:"bytes,1,opt,name=checkpoint_path,json=checkpointPath,proto3" json:"checkpoint_path,omitempty"`
	Address        string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port           int32                  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Tls            bool                   `protobuf:"varint,4,opt,name=tls,proto3" json:"tls,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CriuLazyPagesRequest) Reset() {
	*x = CriuLazyPagesRequest{}
	mi := &file_node_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CriuLazyPagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriuLazyPagesRequest) ProtoMessage() {}

func (x *CriuLazyPagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriuLazyPagesRequest.ProtoReflect.Descriptor instead.
func (*CriuLazyPagesRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *CriuLazyPagesRequest) GetCheckpointPath() string {
	if x != nil {
		return x.CheckpointPath
	}
	return ""
}

func (x *CriuLazyPagesRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CriuLazyPagesRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *CriuLazyPagesRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

type PullImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageId       string                 `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PullImageRequest) Reset() {
	*x = PullImageRequest{}
	mi := &file_node_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullImageRequest) ProtoMessage() {}

func (x *PullImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullImageRequest.ProtoReflect.Descriptor instead.
func (*PullImageRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *PullImageRequest) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x7a, 0x65,
	0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0b,
	0x45, 0x76, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70,
	0x6f, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x45, 0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70,
	0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3c, 0x0a, 0x0c, 0x45, 0x76, 0x61, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x7a, 0x65, 0x72,
	0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a,
	0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x58, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0d, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x78,
	0x0a, 0x07, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f,
	0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x44, 0x69, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x69,
	0x76, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x45,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x7f, 0x0a, 0x14, 0x43, 0x72, 0x69, 0x75, 0x4c, 0x61, 0x7a, 0x79, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c, 0x73,
	0x22, 0x2d, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32,
	0xd6, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x45, 0x76, 0x61, 0x63,
	0x12, 0x1c, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a,
	0x0b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x76, 0x61, 0x63, 0x12, 0x1c, 0x2e, 0x7a,
	0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x7a, 0x65, 0x72,
	0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x61,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70,
	0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x4e,
	0x65, 0x77, 0x43, 0x72, 0x69, 0x75, 0x4c, 0x61, 0x7a, 0x79, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x25, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x69, 0x75, 0x4c, 0x61, 0x7a, 0x79, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x09, 0x50, 0x75, 0x6c, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x7a, 0x65,
	0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x6f, 0x64, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x74, 0x72, 0x6f, 0x78, 0x2f, 0x7a, 0x65, 0x72,
	0x6f, 0x70, 0x6f, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_node_proto_goTypes = []any{
	(*EvacRequest)(nil),           // 0: zeropod.node.v1.EvacRequest
	(*EvacResponse)(nil),          // 1: zeropod.node.v1.EvacResponse
	(*RestoreRequest)(nil),        // 2: zeropod.node.v1.RestoreRequest
	(*RestoreResponse)(nil),       // 3: zeropod.node.v1.RestoreResponse
	(*PodInfo)(nil),               // 4: zeropod.node.v1.PodInfo
	(*MigrationInfo)(nil),         // 5: zeropod.node.v1.MigrationInfo
	(*Image)(nil),                 // 6: zeropod.node.v1.Image
	(*CriuLazyPagesRequest)(nil),  // 7: zeropod.node.v1.CriuLazyPagesRequest
	(*PullImageRequest)(nil),      // 8: zeropod.node.v1.PullImageRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_node_proto_depIdxs = []int32{
	4,  // 0: zeropod.node.v1.EvacRequest.pod_info:type_name -> zeropod.node.v1.PodInfo
	5,  // 1: zeropod.node.v1.EvacRequest.migration_info:type_name -> zeropod.node.v1.MigrationInfo
	9,  // 2: zeropod.node.v1.EvacResponse.empty:type_name -> google.protobuf.Empty
	4,  // 3: zeropod.node.v1.RestoreRequest.pod_info:type_name -> zeropod.node.v1.PodInfo
	5,  // 4: zeropod.node.v1.RestoreRequest.migration_info:type_name -> zeropod.node.v1.MigrationInfo
	5,  // 5: zeropod.node.v1.RestoreResponse.migration_info:type_name -> zeropod.node.v1.MigrationInfo
	10, // 6: zeropod.node.v1.MigrationInfo.paused_at:type_name -> google.protobuf.Timestamp
	10, // 7: zeropod.node.v1.MigrationInfo.restore_start:type_name -> google.protobuf.Timestamp
	10, // 8: zeropod.node.v1.MigrationInfo.restore_end:type_name -> google.protobuf.Timestamp
	0,  // 9: zeropod.node.v1.Node.Evac:input_type -> zeropod.node.v1.EvacRequest
	0,  // 10: zeropod.node.v1.Node.PrepareEvac:input_type -> zeropod.node.v1.EvacRequest
	2,  // 11: zeropod.node.v1.Node.Restore:input_type -> zeropod.node.v1.RestoreRequest
	2,  // 12: zeropod.node.v1.Node.FinishRestore:input_type -> zeropod.node.v1.RestoreRequest
	7,  // 13: zeropod.node.v1.Node.NewCriuLazyPages:input_type -> zeropod.node.v1.CriuLazyPagesRequest
	8,  // 14: zeropod.node.v1.Node.PullImage:input_type -> zeropod.node.v1.PullImageRequest
	1,  // 15: zeropod.node.v1.Node.Evac:output_type -> zeropod.node.v1.EvacResponse
	1,  // 16: zeropod.node.v1.Node.PrepareEvac:output_type -> zeropod.node.v1.EvacResponse
	3,  // 17: zeropod.node.v1.Node.Restore:output_type -> zeropod.node.v1.RestoreResponse
	3,  // 18: zeropod.node.v1.Node.FinishRestore:output_type -> zeropod.node.v1.RestoreResponse
	9,  // 19: zeropod.node.v1.Node.NewCriuLazyPages:output_type -> google.protobuf.Empty
	6,  // 20: zeropod.node.v1.Node.PullImage:output_type -> zeropod.node.v1.Image
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
