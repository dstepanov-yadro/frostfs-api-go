// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: v2/accounting/grpc/service.proto

package accounting

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

// BalanceRequest message
type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the balance request message.
	Body *BalanceRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_accounting_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_accounting_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_v2_accounting_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceRequest) GetBody() *BalanceRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BalanceRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *BalanceRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// BalanceResponse message
type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the balance response message.
	Body *BalanceResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_accounting_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_accounting_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_v2_accounting_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceResponse) GetBody() *BalanceResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BalanceResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *BalanceResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// To indicate the account for which the balance is requested, it's identifier
// is used. It can be any existing account in NeoFS sidechain `Balance` smart
// contract. If omitted, client implementation MUST set it to the request's
// signer `OwnerID`.
type BalanceRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid user identifier in `OwnerID` format for which the balance is
	// requested. Required field.
	OwnerId *grpc1.OwnerID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *BalanceRequest_Body) Reset() {
	*x = BalanceRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_accounting_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest_Body) ProtoMessage() {}

func (x *BalanceRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_v2_accounting_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest_Body.ProtoReflect.Descriptor instead.
func (*BalanceRequest_Body) Descriptor() ([]byte, []int) {
	return file_v2_accounting_grpc_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BalanceRequest_Body) GetOwnerId() *grpc1.OwnerID {
	if x != nil {
		return x.OwnerId
	}
	return nil
}

// The amount of funds in GAS token for the `OwnerID`'s account requested.
// Balance is `Decimal` format to avoid precision issues with rounding.
type BalanceResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount of funds in GAS token for the requested account.
	Balance *Decimal `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceResponse_Body) Reset() {
	*x = BalanceResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_accounting_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse_Body) ProtoMessage() {}

func (x *BalanceResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_v2_accounting_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse_Body.ProtoReflect.Descriptor instead.
func (*BalanceResponse_Body) Descriptor() ([]byte, []int) {
	return file_v2_accounting_grpc_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BalanceResponse_Body) GetBalance() *Decimal {
	if x != nil {
		return x.Balance
	}
	return nil
}

var File_v2_accounting_grpc_service_proto protoreflect.FileDescriptor

var file_v2_accounting_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1e, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa5, 0x02, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
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
	0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x3a, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x52, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0xae, 0x02, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x3f, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x37, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x6b, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x62, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0xaa, 0x02, 0x1e, 0x4e, 0x65, 0x6f, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2_accounting_grpc_service_proto_rawDescOnce sync.Once
	file_v2_accounting_grpc_service_proto_rawDescData = file_v2_accounting_grpc_service_proto_rawDesc
)

func file_v2_accounting_grpc_service_proto_rawDescGZIP() []byte {
	file_v2_accounting_grpc_service_proto_rawDescOnce.Do(func() {
		file_v2_accounting_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_accounting_grpc_service_proto_rawDescData)
	})
	return file_v2_accounting_grpc_service_proto_rawDescData
}

var file_v2_accounting_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_accounting_grpc_service_proto_goTypes = []interface{}{
	(*BalanceRequest)(nil),                  // 0: neo.fs.v2.accounting.BalanceRequest
	(*BalanceResponse)(nil),                 // 1: neo.fs.v2.accounting.BalanceResponse
	(*BalanceRequest_Body)(nil),             // 2: neo.fs.v2.accounting.BalanceRequest.Body
	(*BalanceResponse_Body)(nil),            // 3: neo.fs.v2.accounting.BalanceResponse.Body
	(*grpc.RequestMetaHeader)(nil),          // 4: neo.fs.v2.session.RequestMetaHeader
	(*grpc.RequestVerificationHeader)(nil),  // 5: neo.fs.v2.session.RequestVerificationHeader
	(*grpc.ResponseMetaHeader)(nil),         // 6: neo.fs.v2.session.ResponseMetaHeader
	(*grpc.ResponseVerificationHeader)(nil), // 7: neo.fs.v2.session.ResponseVerificationHeader
	(*grpc1.OwnerID)(nil),                   // 8: neo.fs.v2.refs.OwnerID
	(*Decimal)(nil),                         // 9: neo.fs.v2.accounting.Decimal
}
var file_v2_accounting_grpc_service_proto_depIdxs = []int32{
	2, // 0: neo.fs.v2.accounting.BalanceRequest.body:type_name -> neo.fs.v2.accounting.BalanceRequest.Body
	4, // 1: neo.fs.v2.accounting.BalanceRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	5, // 2: neo.fs.v2.accounting.BalanceRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	3, // 3: neo.fs.v2.accounting.BalanceResponse.body:type_name -> neo.fs.v2.accounting.BalanceResponse.Body
	6, // 4: neo.fs.v2.accounting.BalanceResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	7, // 5: neo.fs.v2.accounting.BalanceResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	8, // 6: neo.fs.v2.accounting.BalanceRequest.Body.owner_id:type_name -> neo.fs.v2.refs.OwnerID
	9, // 7: neo.fs.v2.accounting.BalanceResponse.Body.balance:type_name -> neo.fs.v2.accounting.Decimal
	0, // 8: neo.fs.v2.accounting.AccountingService.Balance:input_type -> neo.fs.v2.accounting.BalanceRequest
	1, // 9: neo.fs.v2.accounting.AccountingService.Balance:output_type -> neo.fs.v2.accounting.BalanceResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_v2_accounting_grpc_service_proto_init() }
func file_v2_accounting_grpc_service_proto_init() {
	if File_v2_accounting_grpc_service_proto != nil {
		return
	}
	file_v2_accounting_grpc_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v2_accounting_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest); i {
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
		file_v2_accounting_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
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
		file_v2_accounting_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest_Body); i {
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
		file_v2_accounting_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse_Body); i {
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
			RawDescriptor: file_v2_accounting_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_accounting_grpc_service_proto_goTypes,
		DependencyIndexes: file_v2_accounting_grpc_service_proto_depIdxs,
		MessageInfos:      file_v2_accounting_grpc_service_proto_msgTypes,
	}.Build()
	File_v2_accounting_grpc_service_proto = out.File
	file_v2_accounting_grpc_service_proto_rawDesc = nil
	file_v2_accounting_grpc_service_proto_goTypes = nil
	file_v2_accounting_grpc_service_proto_depIdxs = nil
}
