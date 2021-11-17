// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: netmap/grpc/service.proto

package netmap

import (
	grpc1 "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
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

// Get NodeInfo structure from the particular node directly
type LocalNodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the LocalNodeInfo request message
	Body *LocalNodeInfoRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *LocalNodeInfoRequest) Reset() {
	*x = LocalNodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalNodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalNodeInfoRequest) ProtoMessage() {}

func (x *LocalNodeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalNodeInfoRequest.ProtoReflect.Descriptor instead.
func (*LocalNodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *LocalNodeInfoRequest) GetBody() *LocalNodeInfoRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LocalNodeInfoRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *LocalNodeInfoRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Local Node Info, including API Version in use
type LocalNodeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the balance response message.
	Body *LocalNodeInfoResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect response execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *LocalNodeInfoResponse) Reset() {
	*x = LocalNodeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalNodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalNodeInfoResponse) ProtoMessage() {}

func (x *LocalNodeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalNodeInfoResponse.ProtoReflect.Descriptor instead.
func (*LocalNodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *LocalNodeInfoResponse) GetBody() *LocalNodeInfoResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LocalNodeInfoResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *LocalNodeInfoResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Get NetworkInfo structure with the network view from particular node.
type NetworkInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the NetworkInfo request message
	Body *NetworkInfoRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *NetworkInfoRequest) Reset() {
	*x = NetworkInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoRequest) ProtoMessage() {}

func (x *NetworkInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoRequest.ProtoReflect.Descriptor instead.
func (*NetworkInfoRequest) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInfoRequest) GetBody() *NetworkInfoRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *NetworkInfoRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *NetworkInfoRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Response with NetworkInfo structure including current epoch and
// sidechain magic number.
type NetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the NetworkInfo response message.
	Body *NetworkInfoResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect response execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *NetworkInfoResponse) Reset() {
	*x = NetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoResponse) ProtoMessage() {}

func (x *NetworkInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkInfoResponse) GetBody() *NetworkInfoResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *NetworkInfoResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *NetworkInfoResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// LocalNodeInfo request body is empty.
type LocalNodeInfoRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalNodeInfoRequest_Body) Reset() {
	*x = LocalNodeInfoRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalNodeInfoRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalNodeInfoRequest_Body) ProtoMessage() {}

func (x *LocalNodeInfoRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalNodeInfoRequest_Body.ProtoReflect.Descriptor instead.
func (*LocalNodeInfoRequest_Body) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{0, 0}
}

// Local Node Info, including API Version in use.
type LocalNodeInfoResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latest NeoFS API version in use
	Version *grpc1.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// NodeInfo structure with recent information from node itself
	NodeInfo *NodeInfo `protobuf:"bytes,2,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
}

func (x *LocalNodeInfoResponse_Body) Reset() {
	*x = LocalNodeInfoResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalNodeInfoResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalNodeInfoResponse_Body) ProtoMessage() {}

func (x *LocalNodeInfoResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalNodeInfoResponse_Body.ProtoReflect.Descriptor instead.
func (*LocalNodeInfoResponse_Body) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LocalNodeInfoResponse_Body) GetVersion() *grpc1.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *LocalNodeInfoResponse_Body) GetNodeInfo() *NodeInfo {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

// NetworkInfo request body is empty.
type NetworkInfoRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkInfoRequest_Body) Reset() {
	*x = NetworkInfoRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoRequest_Body) ProtoMessage() {}

func (x *NetworkInfoRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoRequest_Body.ProtoReflect.Descriptor instead.
func (*NetworkInfoRequest_Body) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{2, 0}
}

// Information about the network.
type NetworkInfoResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NetworkInfo structure with recent information.
	NetworkInfo *NetworkInfo `protobuf:"bytes,1,opt,name=network_info,json=networkInfo,proto3" json:"network_info,omitempty"`
}

func (x *NetworkInfoResponse_Body) Reset() {
	*x = NetworkInfoResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netmap_grpc_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoResponse_Body) ProtoMessage() {}

func (x *NetworkInfoResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_netmap_grpc_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoResponse_Body.ProtoReflect.Descriptor instead.
func (*NetworkInfoResponse_Body) Descriptor() ([]byte, []int) {
	return file_netmap_grpc_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *NetworkInfoResponse_Body) GetNetworkInfo() *NetworkInfo {
	if x != nil {
		return x.NetworkInfo
	}
	return nil
}

var File_netmap_grpc_service_proto protoreflect.FileDescriptor

var file_netmap_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x1a, 0x17, 0x6e,
	0x65, 0x74, 0x6d, 0x61, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61,
	0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x06, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0xe9, 0x02, 0x0a, 0x15, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65,
	0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x46, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x72, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0xf5, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0d,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a,
	0x06, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xbb, 0x02, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x46, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x48, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x40, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xcd, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x6d, 0x61, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74,
	0x6d, 0x61, 0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x61,
	0x70, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6e, 0x65,
	0x74, 0x6d, 0x61, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x70,
	0xaa, 0x02, 0x1a, 0x4e, 0x65, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x4e, 0x65, 0x74, 0x6d, 0x61, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_netmap_grpc_service_proto_rawDescOnce sync.Once
	file_netmap_grpc_service_proto_rawDescData = file_netmap_grpc_service_proto_rawDesc
)

func file_netmap_grpc_service_proto_rawDescGZIP() []byte {
	file_netmap_grpc_service_proto_rawDescOnce.Do(func() {
		file_netmap_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_netmap_grpc_service_proto_rawDescData)
	})
	return file_netmap_grpc_service_proto_rawDescData
}

var file_netmap_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_netmap_grpc_service_proto_goTypes = []interface{}{
	(*LocalNodeInfoRequest)(nil),            // 0: neo.fs.v2.netmap.LocalNodeInfoRequest
	(*LocalNodeInfoResponse)(nil),           // 1: neo.fs.v2.netmap.LocalNodeInfoResponse
	(*NetworkInfoRequest)(nil),              // 2: neo.fs.v2.netmap.NetworkInfoRequest
	(*NetworkInfoResponse)(nil),             // 3: neo.fs.v2.netmap.NetworkInfoResponse
	(*LocalNodeInfoRequest_Body)(nil),       // 4: neo.fs.v2.netmap.LocalNodeInfoRequest.Body
	(*LocalNodeInfoResponse_Body)(nil),      // 5: neo.fs.v2.netmap.LocalNodeInfoResponse.Body
	(*NetworkInfoRequest_Body)(nil),         // 6: neo.fs.v2.netmap.NetworkInfoRequest.Body
	(*NetworkInfoResponse_Body)(nil),        // 7: neo.fs.v2.netmap.NetworkInfoResponse.Body
	(*grpc.RequestMetaHeader)(nil),          // 8: neo.fs.v2.session.RequestMetaHeader
	(*grpc.RequestVerificationHeader)(nil),  // 9: neo.fs.v2.session.RequestVerificationHeader
	(*grpc.ResponseMetaHeader)(nil),         // 10: neo.fs.v2.session.ResponseMetaHeader
	(*grpc.ResponseVerificationHeader)(nil), // 11: neo.fs.v2.session.ResponseVerificationHeader
	(*grpc1.Version)(nil),                   // 12: neo.fs.v2.refs.Version
	(*NodeInfo)(nil),                        // 13: neo.fs.v2.netmap.NodeInfo
	(*NetworkInfo)(nil),                     // 14: neo.fs.v2.netmap.NetworkInfo
}
var file_netmap_grpc_service_proto_depIdxs = []int32{
	4,  // 0: neo.fs.v2.netmap.LocalNodeInfoRequest.body:type_name -> neo.fs.v2.netmap.LocalNodeInfoRequest.Body
	8,  // 1: neo.fs.v2.netmap.LocalNodeInfoRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	9,  // 2: neo.fs.v2.netmap.LocalNodeInfoRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	5,  // 3: neo.fs.v2.netmap.LocalNodeInfoResponse.body:type_name -> neo.fs.v2.netmap.LocalNodeInfoResponse.Body
	10, // 4: neo.fs.v2.netmap.LocalNodeInfoResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	11, // 5: neo.fs.v2.netmap.LocalNodeInfoResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	6,  // 6: neo.fs.v2.netmap.NetworkInfoRequest.body:type_name -> neo.fs.v2.netmap.NetworkInfoRequest.Body
	8,  // 7: neo.fs.v2.netmap.NetworkInfoRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	9,  // 8: neo.fs.v2.netmap.NetworkInfoRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	7,  // 9: neo.fs.v2.netmap.NetworkInfoResponse.body:type_name -> neo.fs.v2.netmap.NetworkInfoResponse.Body
	10, // 10: neo.fs.v2.netmap.NetworkInfoResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	11, // 11: neo.fs.v2.netmap.NetworkInfoResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	12, // 12: neo.fs.v2.netmap.LocalNodeInfoResponse.Body.version:type_name -> neo.fs.v2.refs.Version
	13, // 13: neo.fs.v2.netmap.LocalNodeInfoResponse.Body.node_info:type_name -> neo.fs.v2.netmap.NodeInfo
	14, // 14: neo.fs.v2.netmap.NetworkInfoResponse.Body.network_info:type_name -> neo.fs.v2.netmap.NetworkInfo
	0,  // 15: neo.fs.v2.netmap.NetmapService.LocalNodeInfo:input_type -> neo.fs.v2.netmap.LocalNodeInfoRequest
	2,  // 16: neo.fs.v2.netmap.NetmapService.NetworkInfo:input_type -> neo.fs.v2.netmap.NetworkInfoRequest
	1,  // 17: neo.fs.v2.netmap.NetmapService.LocalNodeInfo:output_type -> neo.fs.v2.netmap.LocalNodeInfoResponse
	3,  // 18: neo.fs.v2.netmap.NetmapService.NetworkInfo:output_type -> neo.fs.v2.netmap.NetworkInfoResponse
	17, // [17:19] is the sub-list for method output_type
	15, // [15:17] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_netmap_grpc_service_proto_init() }
func file_netmap_grpc_service_proto_init() {
	if File_netmap_grpc_service_proto != nil {
		return
	}
	file_netmap_grpc_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_netmap_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalNodeInfoRequest); i {
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
		file_netmap_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalNodeInfoResponse); i {
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
		file_netmap_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoRequest); i {
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
		file_netmap_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoResponse); i {
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
		file_netmap_grpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalNodeInfoRequest_Body); i {
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
		file_netmap_grpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalNodeInfoResponse_Body); i {
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
		file_netmap_grpc_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoRequest_Body); i {
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
		file_netmap_grpc_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoResponse_Body); i {
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
			RawDescriptor: file_netmap_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_netmap_grpc_service_proto_goTypes,
		DependencyIndexes: file_netmap_grpc_service_proto_depIdxs,
		MessageInfos:      file_netmap_grpc_service_proto_msgTypes,
	}.Build()
	File_netmap_grpc_service_proto = out.File
	file_netmap_grpc_service_proto_rawDesc = nil
	file_netmap_grpc_service_proto_goTypes = nil
	file_netmap_grpc_service_proto_depIdxs = nil
}
