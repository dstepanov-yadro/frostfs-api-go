// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package object

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ObjectServiceClient is the client API for ObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectServiceClient interface {
	// Receive full object structure, including Headers and payload. Response uses
	// gRPC stream. First response message carries object with requested address.
	// Chunk messages are parts of the object's payload if it is needed. All
	// messages, except the first one, carry payload chunks. Requested object can
	// be restored by concatenation of object message payload and all chunks
	// keeping receiving order.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ObjectService_GetClient, error)
	// Put the object into container. Request uses gRPC stream. First message
	// SHOULD be of PutHeader type. `ContainerID` and `OwnerID` of an object
	// SHOULD be set. Session token SHOULD be obtained before `PUT` operation (see
	// session package). Chunk messages are considered by server as a part of an
	// object payload. All messages, except first one, SHOULD be payload chunks.
	// Chunk messages SHOULD be sent in direct order of fragmentation.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully saved in the container;
	// - Common failures (SECTION_FAILURE_COMMON).
	Put(ctx context.Context, opts ...grpc.CallOption) (ObjectService_PutClient, error)
	// Delete the object from a container. There is no immediate removal
	// guarantee. Object will be marked for removal and deleted eventually.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully marked to be removed from the container;
	// - Common failures (SECTION_FAILURE_COMMON).
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Returns the object Headers without data payload. By default full header is
	// returned. If `main_only` request field is set, the short header with only
	// the very minimal information would be returned instead.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object header has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Head(ctx context.Context, in *HeadRequest, opts ...grpc.CallOption) (*HeadResponse, error)
	// Search objects in container. Search query allows to match by Object
	// Header's filed values. Please see the corresponding NeoFS Technical
	// Specification section for more details.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// objects have been successfully selected;
	// - Common failures (SECTION_FAILURE_COMMON).
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (ObjectService_SearchClient, error)
	// Get byte range of data payload. Range is set as an (offset, length) tuple.
	// Like in `Get` method, the response uses gRPC stream. Requested range can be
	// restored by concatenation of all received payload chunks keeping receiving
	// order.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// data range of the object payload has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	GetRange(ctx context.Context, in *GetRangeRequest, opts ...grpc.CallOption) (ObjectService_GetRangeClient, error)
	// Returns homomorphic or regular hash of object's payload range after
	// applying XOR operation with the provided `salt`. Ranges are set of (offset,
	// length) tuples. Hashes order in response corresponds to ranges order in
	// request. Note that hash is calculated for XORed data.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// data range of the object payload has been successfully hashed;
	// - Common failures (SECTION_FAILURE_COMMON).
	GetRangeHash(ctx context.Context, in *GetRangeHashRequest, opts ...grpc.CallOption) (*GetRangeHashResponse, error)
}

type objectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectServiceClient(cc grpc.ClientConnInterface) ObjectServiceClient {
	return &objectServiceClient{cc}
}

func (c *objectServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ObjectService_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectService_ServiceDesc.Streams[0], "/neo.fs.v2.object.ObjectService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectService_GetClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type objectServiceGetClient struct {
	grpc.ClientStream
}

func (x *objectServiceGetClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectServiceClient) Put(ctx context.Context, opts ...grpc.CallOption) (ObjectService_PutClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectService_ServiceDesc.Streams[1], "/neo.fs.v2.object.ObjectService/Put", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectServicePutClient{stream}
	return x, nil
}

type ObjectService_PutClient interface {
	Send(*PutRequest) error
	CloseAndRecv() (*PutResponse, error)
	grpc.ClientStream
}

type objectServicePutClient struct {
	grpc.ClientStream
}

func (x *objectServicePutClient) Send(m *PutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *objectServicePutClient) CloseAndRecv() (*PutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.object.ObjectService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) Head(ctx context.Context, in *HeadRequest, opts ...grpc.CallOption) (*HeadResponse, error) {
	out := new(HeadResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.object.ObjectService/Head", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (ObjectService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectService_ServiceDesc.Streams[2], "/neo.fs.v2.object.ObjectService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectService_SearchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type objectServiceSearchClient struct {
	grpc.ClientStream
}

func (x *objectServiceSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectServiceClient) GetRange(ctx context.Context, in *GetRangeRequest, opts ...grpc.CallOption) (ObjectService_GetRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectService_ServiceDesc.Streams[3], "/neo.fs.v2.object.ObjectService/GetRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectServiceGetRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectService_GetRangeClient interface {
	Recv() (*GetRangeResponse, error)
	grpc.ClientStream
}

type objectServiceGetRangeClient struct {
	grpc.ClientStream
}

func (x *objectServiceGetRangeClient) Recv() (*GetRangeResponse, error) {
	m := new(GetRangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectServiceClient) GetRangeHash(ctx context.Context, in *GetRangeHashRequest, opts ...grpc.CallOption) (*GetRangeHashResponse, error) {
	out := new(GetRangeHashResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.object.ObjectService/GetRangeHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServiceServer is the server API for ObjectService service.
// All implementations should embed UnimplementedObjectServiceServer
// for forward compatibility
type ObjectServiceServer interface {
	// Receive full object structure, including Headers and payload. Response uses
	// gRPC stream. First response message carries object with requested address.
	// Chunk messages are parts of the object's payload if it is needed. All
	// messages, except the first one, carry payload chunks. Requested object can
	// be restored by concatenation of object message payload and all chunks
	// keeping receiving order.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Get(*GetRequest, ObjectService_GetServer) error
	// Put the object into container. Request uses gRPC stream. First message
	// SHOULD be of PutHeader type. `ContainerID` and `OwnerID` of an object
	// SHOULD be set. Session token SHOULD be obtained before `PUT` operation (see
	// session package). Chunk messages are considered by server as a part of an
	// object payload. All messages, except first one, SHOULD be payload chunks.
	// Chunk messages SHOULD be sent in direct order of fragmentation.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully saved in the container;
	// - Common failures (SECTION_FAILURE_COMMON).
	Put(ObjectService_PutServer) error
	// Delete the object from a container. There is no immediate removal
	// guarantee. Object will be marked for removal and deleted eventually.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object has been successfully marked to be removed from the container;
	// - Common failures (SECTION_FAILURE_COMMON).
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Returns the object Headers without data payload. By default full header is
	// returned. If `main_only` request field is set, the short header with only
	// the very minimal information would be returned instead.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// object header has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Head(context.Context, *HeadRequest) (*HeadResponse, error)
	// Search objects in container. Search query allows to match by Object
	// Header's filed values. Please see the corresponding NeoFS Technical
	// Specification section for more details.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// objects have been successfully selected;
	// - Common failures (SECTION_FAILURE_COMMON).
	Search(*SearchRequest, ObjectService_SearchServer) error
	// Get byte range of data payload. Range is set as an (offset, length) tuple.
	// Like in `Get` method, the response uses gRPC stream. Requested range can be
	// restored by concatenation of all received payload chunks keeping receiving
	// order.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// data range of the object payload has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	GetRange(*GetRangeRequest, ObjectService_GetRangeServer) error
	// Returns homomorphic or regular hash of object's payload range after
	// applying XOR operation with the provided `salt`. Ranges are set of (offset,
	// length) tuples. Hashes order in response corresponds to ranges order in
	// request. Note that hash is calculated for XORed data.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// data range of the object payload has been successfully hashed;
	// - Common failures (SECTION_FAILURE_COMMON).
	GetRangeHash(context.Context, *GetRangeHashRequest) (*GetRangeHashResponse, error)
}

// UnimplementedObjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedObjectServiceServer struct {
}

func (UnimplementedObjectServiceServer) Get(*GetRequest, ObjectService_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedObjectServiceServer) Put(ObjectService_PutServer) error {
	return status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedObjectServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedObjectServiceServer) Head(context.Context, *HeadRequest) (*HeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Head not implemented")
}
func (UnimplementedObjectServiceServer) Search(*SearchRequest, ObjectService_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedObjectServiceServer) GetRange(*GetRangeRequest, ObjectService_GetRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRange not implemented")
}
func (UnimplementedObjectServiceServer) GetRangeHash(context.Context, *GetRangeHashRequest) (*GetRangeHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRangeHash not implemented")
}

// UnsafeObjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectServiceServer will
// result in compilation errors.
type UnsafeObjectServiceServer interface {
	mustEmbedUnimplementedObjectServiceServer()
}

func RegisterObjectServiceServer(s grpc.ServiceRegistrar, srv ObjectServiceServer) {
	s.RegisterService(&ObjectService_ServiceDesc, srv)
}

func _ObjectService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectServiceServer).Get(m, &objectServiceGetServer{stream})
}

type ObjectService_GetServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type objectServiceGetServer struct {
	grpc.ServerStream
}

func (x *objectServiceGetServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectService_Put_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ObjectServiceServer).Put(&objectServicePutServer{stream})
}

type ObjectService_PutServer interface {
	SendAndClose(*PutResponse) error
	Recv() (*PutRequest, error)
	grpc.ServerStream
}

type objectServicePutServer struct {
	grpc.ServerStream
}

func (x *objectServicePutServer) SendAndClose(m *PutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *objectServicePutServer) Recv() (*PutRequest, error) {
	m := new(PutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ObjectService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.object.ObjectService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_Head_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).Head(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.object.ObjectService/Head",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).Head(ctx, req.(*HeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectServiceServer).Search(m, &objectServiceSearchServer{stream})
}

type ObjectService_SearchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type objectServiceSearchServer struct {
	grpc.ServerStream
}

func (x *objectServiceSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectService_GetRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectServiceServer).GetRange(m, &objectServiceGetRangeServer{stream})
}

type ObjectService_GetRangeServer interface {
	Send(*GetRangeResponse) error
	grpc.ServerStream
}

type objectServiceGetRangeServer struct {
	grpc.ServerStream
}

func (x *objectServiceGetRangeServer) Send(m *GetRangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectService_GetRangeHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRangeHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetRangeHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.object.ObjectService/GetRangeHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetRangeHash(ctx, req.(*GetRangeHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectService_ServiceDesc is the grpc.ServiceDesc for ObjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neo.fs.v2.object.ObjectService",
	HandlerType: (*ObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _ObjectService_Delete_Handler,
		},
		{
			MethodName: "Head",
			Handler:    _ObjectService_Head_Handler,
		},
		{
			MethodName: "GetRangeHash",
			Handler:    _ObjectService_GetRangeHash_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _ObjectService_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Put",
			Handler:       _ObjectService_Put_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Search",
			Handler:       _ObjectService_Search_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRange",
			Handler:       _ObjectService_GetRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "object/grpc/service.proto",
}
