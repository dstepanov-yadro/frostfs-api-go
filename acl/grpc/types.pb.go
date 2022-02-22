// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: acl/grpc/types.proto

package acl

import (
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
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

// Target role of the access control rule in access control list.
type Role int32

const (
	// Unspecified  role, default value
	Role_ROLE_UNSPECIFIED Role = 0
	// User target rule is applied if sender is the owner of the container
	Role_USER Role = 1
	// System target rule is applied if sender is the storage node within the
	// container or inner ring node
	Role_SYSTEM Role = 2
	// Others target rule is applied if sender is not user nor system target
	Role_OTHERS Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "USER",
		2: "SYSTEM",
		3: "OTHERS",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"USER":             1,
		"SYSTEM":           2,
		"OTHERS":           3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_acl_grpc_types_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_acl_grpc_types_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{0}
}

// MatchType is an enumeration of match types.
type MatchType int32

const (
	// Unspecified match type, default value.
	MatchType_MATCH_TYPE_UNSPECIFIED MatchType = 0
	// Return true if strings are equal
	MatchType_STRING_EQUAL MatchType = 1
	// Return true if strings are different
	MatchType_STRING_NOT_EQUAL MatchType = 2
)

// Enum value maps for MatchType.
var (
	MatchType_name = map[int32]string{
		0: "MATCH_TYPE_UNSPECIFIED",
		1: "STRING_EQUAL",
		2: "STRING_NOT_EQUAL",
	}
	MatchType_value = map[string]int32{
		"MATCH_TYPE_UNSPECIFIED": 0,
		"STRING_EQUAL":           1,
		"STRING_NOT_EQUAL":       2,
	}
)

func (x MatchType) Enum() *MatchType {
	p := new(MatchType)
	*p = x
	return p
}

func (x MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_acl_grpc_types_proto_enumTypes[1].Descriptor()
}

func (MatchType) Type() protoreflect.EnumType {
	return &file_acl_grpc_types_proto_enumTypes[1]
}

func (x MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchType.Descriptor instead.
func (MatchType) EnumDescriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{1}
}

// Request's operation type to match if the rule is applicable to a particular
// request.
type Operation int32

const (
	// Unspecified operation, default value
	Operation_OPERATION_UNSPECIFIED Operation = 0
	// Get
	Operation_GET Operation = 1
	// Head
	Operation_HEAD Operation = 2
	// Put
	Operation_PUT Operation = 3
	// Delete
	Operation_DELETE Operation = 4
	// Search
	Operation_SEARCH Operation = 5
	// GetRange
	Operation_GETRANGE Operation = 6
	// GetRangeHash
	Operation_GETRANGEHASH Operation = 7
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "GET",
		2: "HEAD",
		3: "PUT",
		4: "DELETE",
		5: "SEARCH",
		6: "GETRANGE",
		7: "GETRANGEHASH",
	}
	Operation_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"GET":                   1,
		"HEAD":                  2,
		"PUT":                   3,
		"DELETE":                4,
		"SEARCH":                5,
		"GETRANGE":              6,
		"GETRANGEHASH":          7,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_acl_grpc_types_proto_enumTypes[2].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_acl_grpc_types_proto_enumTypes[2]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{2}
}

// Rule execution result action. Either allows or denies access if the rule's
// filters match.
type Action int32

const (
	// Unspecified action, default value
	Action_ACTION_UNSPECIFIED Action = 0
	// Allow action
	Action_ALLOW Action = 1
	// Deny action
	Action_DENY Action = 2
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ALLOW",
		2: "DENY",
	}
	Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ALLOW":              1,
		"DENY":               2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_acl_grpc_types_proto_enumTypes[3].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_acl_grpc_types_proto_enumTypes[3]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{3}
}

// Enumeration of possible sources of Headers to apply filters.
type HeaderType int32

const (
	// Unspecified header, default value.
	HeaderType_HEADER_UNSPECIFIED HeaderType = 0
	// Filter request headers
	HeaderType_REQUEST HeaderType = 1
	// Filter object headers
	HeaderType_OBJECT HeaderType = 2
	// Filter service headers. These are not processed by NeoFS nodes and
	// exist for service use only.
	HeaderType_SERVICE HeaderType = 3
)

// Enum value maps for HeaderType.
var (
	HeaderType_name = map[int32]string{
		0: "HEADER_UNSPECIFIED",
		1: "REQUEST",
		2: "OBJECT",
		3: "SERVICE",
	}
	HeaderType_value = map[string]int32{
		"HEADER_UNSPECIFIED": 0,
		"REQUEST":            1,
		"OBJECT":             2,
		"SERVICE":            3,
	}
)

func (x HeaderType) Enum() *HeaderType {
	p := new(HeaderType)
	*p = x
	return p
}

func (x HeaderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeaderType) Descriptor() protoreflect.EnumDescriptor {
	return file_acl_grpc_types_proto_enumTypes[4].Descriptor()
}

func (HeaderType) Type() protoreflect.EnumType {
	return &file_acl_grpc_types_proto_enumTypes[4]
}

func (x HeaderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeaderType.Descriptor instead.
func (HeaderType) EnumDescriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{4}
}

// Describes a single eACL rule.
type EACLRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NeoFS request Verb to match
	Operation Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=neo.fs.v2.acl.Operation" json:"operation,omitempty"`
	// Rule execution result. Either allows or denies access if filters match.
	Action Action `protobuf:"varint,2,opt,name=action,proto3,enum=neo.fs.v2.acl.Action" json:"action,omitempty"`
	// List of filters to match and see if rule is applicable
	Filters []*EACLRecord_Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// List of target subjects to apply ACL rule to
	Targets []*EACLRecord_Target `protobuf:"bytes,4,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *EACLRecord) Reset() {
	*x = EACLRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACLRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACLRecord) ProtoMessage() {}

func (x *EACLRecord) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACLRecord.ProtoReflect.Descriptor instead.
func (*EACLRecord) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *EACLRecord) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_OPERATION_UNSPECIFIED
}

func (x *EACLRecord) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *EACLRecord) GetFilters() []*EACLRecord_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *EACLRecord) GetTargets() []*EACLRecord_Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

// Extended ACL rules table. Defined a list of ACL rules additionally to Basic
// ACL. Extended ACL rules can be attached to the container and can be updated
// or may be defined in `BearerToken` structure. Please see the corresponding
// NeoFS Technical Specification's section for detailed description.
type EACLTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// eACL format version. Effectively the version of API library used to create
	// eACL Table.
	Version *grpc.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Identifier of the container that should use given access control rules
	ContainerId *grpc.ContainerID `protobuf:"bytes,2,opt,name=container_id,json=containerID,proto3" json:"container_id,omitempty"`
	// List of Extended ACL rules
	Records []*EACLRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *EACLTable) Reset() {
	*x = EACLTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACLTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACLTable) ProtoMessage() {}

func (x *EACLTable) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACLTable.ProtoReflect.Descriptor instead.
func (*EACLTable) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{1}
}

func (x *EACLTable) GetVersion() *grpc.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *EACLTable) GetContainerId() *grpc.ContainerID {
	if x != nil {
		return x.ContainerId
	}
	return nil
}

func (x *EACLTable) GetRecords() []*EACLRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// BearerToken allows to attach signed Extended ACL rules to the request in
// `RequestMetaHeader`. If container's Basic ACL rules allow, the attached rule
// set will be checked instead of one attached to the container itself. Just
// like [JWT](https://jwt.io), it has a limited lifetime and scope, hence can be
// used in the similar use cases, like providing authorisation to externally
// authenticated party.
//
// BearerToken can be issued only by container's owner and must be signed using
// the key associated with container's `OwnerID`.
type BearerToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bearer Token body
	Body *BearerToken_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Signature of BearerToken body
	Signature *grpc.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *BearerToken) Reset() {
	*x = BearerToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerToken) ProtoMessage() {}

func (x *BearerToken) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerToken.ProtoReflect.Descriptor instead.
func (*BearerToken) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{2}
}

func (x *BearerToken) GetBody() *BearerToken_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BearerToken) GetSignature() *grpc.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Filter to check particular properties of the request or object.
//
// By default `key` field refers to the corresponding object's `Attribute`.
// Some Object's header fields can also be accessed by adding `$Object:`
// prefix to the name. Here is the list of fields available via this prefix:
//
// * $Object:version \
//   version
// * $Object:objectID \
//   object_id
// * $Object:containerID \
//   container_id
// * $Object:ownerID \
//   owner_id
// * $Object:creationEpoch \
//   creation_epoch
// * $Object:payloadLength \
//   payload_length
// * $Object:payloadHash \
//   payload_hash
// * $Object:objectType \
//   object_type
// * $Object:homomorphicHash \
//   homomorphic_hash
//
// Please note, that if request or response does not have object's headers or
// full object (Range, RangeHash, Search, Delete), it will not be possible to
// filter by object header fields or user attributes. From the well-known list
// only `$Object:objectID` and `$Object:containerID` will be available, as
// it's possible to take that information from the requested address.
type EACLRecord_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define if Object or Request header will be used
	HeaderType HeaderType `protobuf:"varint,1,opt,name=header_type,json=headerType,proto3,enum=neo.fs.v2.acl.HeaderType" json:"header_type,omitempty"`
	// Match operation type
	MatchType MatchType `protobuf:"varint,2,opt,name=match_type,json=matchType,proto3,enum=neo.fs.v2.acl.MatchType" json:"match_type,omitempty"`
	// Name of the Header to use
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// Expected Header Value or pattern to match
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EACLRecord_Filter) Reset() {
	*x = EACLRecord_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACLRecord_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACLRecord_Filter) ProtoMessage() {}

func (x *EACLRecord_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACLRecord_Filter.ProtoReflect.Descriptor instead.
func (*EACLRecord_Filter) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EACLRecord_Filter) GetHeaderType() HeaderType {
	if x != nil {
		return x.HeaderType
	}
	return HeaderType_HEADER_UNSPECIFIED
}

func (x *EACLRecord_Filter) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_UNSPECIFIED
}

func (x *EACLRecord_Filter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EACLRecord_Filter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Target to apply ACL rule. Can be a subject's role class or a list of public
// keys to match.
type EACLRecord_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target subject's role class
	Role Role `protobuf:"varint,1,opt,name=role,proto3,enum=neo.fs.v2.acl.Role" json:"role,omitempty"`
	// List of public keys to identify target subject
	Keys [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *EACLRecord_Target) Reset() {
	*x = EACLRecord_Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACLRecord_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACLRecord_Target) ProtoMessage() {}

func (x *EACLRecord_Target) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACLRecord_Target.ProtoReflect.Descriptor instead.
func (*EACLRecord_Target) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EACLRecord_Target) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *EACLRecord_Target) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}

// Bearer Token body structure contains Extended ACL table issued by container
// owner with additional information preventing token's abuse.
type BearerToken_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Table of Extended ACL rules to use instead of the ones attached to the
	// container
	EaclTable *EACLTable `protobuf:"bytes,1,opt,name=eacl_table,json=eaclTable,proto3" json:"eacl_table,omitempty"`
	// `OwnerID` to whom the token was issued. Must match the request
	// originator's `OwnerID`. If empty, any token bearer will be accepted.
	OwnerId *grpc.OwnerID `protobuf:"bytes,2,opt,name=owner_id,json=ownerID,proto3" json:"owner_id,omitempty"`
	// Token expiration and valid time period parameters
	Lifetime *BearerToken_Body_TokenLifetime `protobuf:"bytes,3,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
}

func (x *BearerToken_Body) Reset() {
	*x = BearerToken_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerToken_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerToken_Body) ProtoMessage() {}

func (x *BearerToken_Body) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerToken_Body.ProtoReflect.Descriptor instead.
func (*BearerToken_Body) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{2, 0}
}

func (x *BearerToken_Body) GetEaclTable() *EACLTable {
	if x != nil {
		return x.EaclTable
	}
	return nil
}

func (x *BearerToken_Body) GetOwnerId() *grpc.OwnerID {
	if x != nil {
		return x.OwnerId
	}
	return nil
}

func (x *BearerToken_Body) GetLifetime() *BearerToken_Body_TokenLifetime {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

// Lifetime parameters of the token. Field names taken from
// [rfc7519](https://tools.ietf.org/html/rfc7519).
type BearerToken_Body_TokenLifetime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expiration Epoch
	Exp uint64 `protobuf:"varint,1,opt,name=exp,proto3" json:"exp,omitempty"`
	// Not valid before Epoch
	Nbf uint64 `protobuf:"varint,2,opt,name=nbf,proto3" json:"nbf,omitempty"`
	// Issued at Epoch
	Iat uint64 `protobuf:"varint,3,opt,name=iat,proto3" json:"iat,omitempty"`
}

func (x *BearerToken_Body_TokenLifetime) Reset() {
	*x = BearerToken_Body_TokenLifetime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_grpc_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BearerToken_Body_TokenLifetime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerToken_Body_TokenLifetime) ProtoMessage() {}

func (x *BearerToken_Body_TokenLifetime) ProtoReflect() protoreflect.Message {
	mi := &file_acl_grpc_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerToken_Body_TokenLifetime.ProtoReflect.Descriptor instead.
func (*BearerToken_Body_TokenLifetime) Descriptor() ([]byte, []int) {
	return file_acl_grpc_types_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *BearerToken_Body_TokenLifetime) GetExp() uint64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *BearerToken_Body_TokenLifetime) GetNbf() uint64 {
	if x != nil {
		return x.Nbf
	}
	return 0
}

func (x *BearerToken_Body_TokenLifetime) GetIat() uint64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

var File_acl_grpc_types_proto protoreflect.FileDescriptor

var file_acl_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x63, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x63, 0x6c, 0x1a, 0x15, 0x72, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a,
	0x0a, 0x45, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x61, 0x63, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x61, 0x63, 0x6c, 0x2e, 0x45, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3a,
	0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e,
	0x45, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0xa5, 0x01, 0x0a, 0x06, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x45, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x09, 0x45, 0x41,
	0x43, 0x4c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65,
	0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x45, 0x41, 0x43, 0x4c,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x83, 0x03, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x33, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x85, 0x02,
	0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x61, 0x63, 0x6c, 0x5f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x45, 0x41, 0x43, 0x4c, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x65, 0x61, 0x63, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x66, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x49, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x66, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x45,
	0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x78,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x62, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6e, 0x62, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x69, 0x61, 0x74, 0x2a, 0x3e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x53, 0x10, 0x03, 0x2a, 0x4f, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x7a, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x06,
	0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x48, 0x41, 0x53, 0x48,
	0x10, 0x07, 0x2a, 0x35, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02, 0x2a, 0x4a, 0x0a, 0x0a, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x45, 0x41, 0x44, 0x45,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x10, 0x03, 0x42, 0x4d, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63,
	0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x61, 0x63, 0x6c, 0xaa, 0x02, 0x17, 0x4e, 0x65, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49,
	0x2e, 0x41, 0x63, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acl_grpc_types_proto_rawDescOnce sync.Once
	file_acl_grpc_types_proto_rawDescData = file_acl_grpc_types_proto_rawDesc
)

func file_acl_grpc_types_proto_rawDescGZIP() []byte {
	file_acl_grpc_types_proto_rawDescOnce.Do(func() {
		file_acl_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_acl_grpc_types_proto_rawDescData)
	})
	return file_acl_grpc_types_proto_rawDescData
}

var file_acl_grpc_types_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_acl_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_acl_grpc_types_proto_goTypes = []interface{}{
	(Role)(0),                              // 0: neo.fs.v2.acl.Role
	(MatchType)(0),                         // 1: neo.fs.v2.acl.MatchType
	(Operation)(0),                         // 2: neo.fs.v2.acl.Operation
	(Action)(0),                            // 3: neo.fs.v2.acl.Action
	(HeaderType)(0),                        // 4: neo.fs.v2.acl.HeaderType
	(*EACLRecord)(nil),                     // 5: neo.fs.v2.acl.EACLRecord
	(*EACLTable)(nil),                      // 6: neo.fs.v2.acl.EACLTable
	(*BearerToken)(nil),                    // 7: neo.fs.v2.acl.BearerToken
	(*EACLRecord_Filter)(nil),              // 8: neo.fs.v2.acl.EACLRecord.Filter
	(*EACLRecord_Target)(nil),              // 9: neo.fs.v2.acl.EACLRecord.Target
	(*BearerToken_Body)(nil),               // 10: neo.fs.v2.acl.BearerToken.Body
	(*BearerToken_Body_TokenLifetime)(nil), // 11: neo.fs.v2.acl.BearerToken.Body.TokenLifetime
	(*grpc.Version)(nil),                   // 12: neo.fs.v2.refs.Version
	(*grpc.ContainerID)(nil),               // 13: neo.fs.v2.refs.ContainerID
	(*grpc.Signature)(nil),                 // 14: neo.fs.v2.refs.Signature
	(*grpc.OwnerID)(nil),                   // 15: neo.fs.v2.refs.OwnerID
}
var file_acl_grpc_types_proto_depIdxs = []int32{
	2,  // 0: neo.fs.v2.acl.EACLRecord.operation:type_name -> neo.fs.v2.acl.Operation
	3,  // 1: neo.fs.v2.acl.EACLRecord.action:type_name -> neo.fs.v2.acl.Action
	8,  // 2: neo.fs.v2.acl.EACLRecord.filters:type_name -> neo.fs.v2.acl.EACLRecord.Filter
	9,  // 3: neo.fs.v2.acl.EACLRecord.targets:type_name -> neo.fs.v2.acl.EACLRecord.Target
	12, // 4: neo.fs.v2.acl.EACLTable.version:type_name -> neo.fs.v2.refs.Version
	13, // 5: neo.fs.v2.acl.EACLTable.container_id:type_name -> neo.fs.v2.refs.ContainerID
	5,  // 6: neo.fs.v2.acl.EACLTable.records:type_name -> neo.fs.v2.acl.EACLRecord
	10, // 7: neo.fs.v2.acl.BearerToken.body:type_name -> neo.fs.v2.acl.BearerToken.Body
	14, // 8: neo.fs.v2.acl.BearerToken.signature:type_name -> neo.fs.v2.refs.Signature
	4,  // 9: neo.fs.v2.acl.EACLRecord.Filter.header_type:type_name -> neo.fs.v2.acl.HeaderType
	1,  // 10: neo.fs.v2.acl.EACLRecord.Filter.match_type:type_name -> neo.fs.v2.acl.MatchType
	0,  // 11: neo.fs.v2.acl.EACLRecord.Target.role:type_name -> neo.fs.v2.acl.Role
	6,  // 12: neo.fs.v2.acl.BearerToken.Body.eacl_table:type_name -> neo.fs.v2.acl.EACLTable
	15, // 13: neo.fs.v2.acl.BearerToken.Body.owner_id:type_name -> neo.fs.v2.refs.OwnerID
	11, // 14: neo.fs.v2.acl.BearerToken.Body.lifetime:type_name -> neo.fs.v2.acl.BearerToken.Body.TokenLifetime
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_acl_grpc_types_proto_init() }
func file_acl_grpc_types_proto_init() {
	if File_acl_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acl_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACLRecord); i {
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
		file_acl_grpc_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACLTable); i {
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
		file_acl_grpc_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerToken); i {
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
		file_acl_grpc_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACLRecord_Filter); i {
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
		file_acl_grpc_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACLRecord_Target); i {
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
		file_acl_grpc_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerToken_Body); i {
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
		file_acl_grpc_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BearerToken_Body_TokenLifetime); i {
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
			RawDescriptor: file_acl_grpc_types_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acl_grpc_types_proto_goTypes,
		DependencyIndexes: file_acl_grpc_types_proto_depIdxs,
		EnumInfos:         file_acl_grpc_types_proto_enumTypes,
		MessageInfos:      file_acl_grpc_types_proto_msgTypes,
	}.Build()
	File_acl_grpc_types_proto = out.File
	file_acl_grpc_types_proto_rawDesc = nil
	file_acl_grpc_types_proto_goTypes = nil
	file_acl_grpc_types_proto_depIdxs = nil
}