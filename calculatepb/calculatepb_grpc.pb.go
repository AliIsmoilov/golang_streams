// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: calculatepb/calculatepb.proto

package calculatepb

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server-side stream
	PrimeDecomposition(ctx context.Context, in *PrimeDecompositionRequest, opts ...grpc.CallOption) (CalculateService_PrimeDecompositionClient, error)
	// Client-side stream
	ComputeAvarage(ctx context.Context, opts ...grpc.CallOption) (CalculateService_ComputeAvarageClient, error)
	// Bi-Directional stream
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaximumClient, error)
	SquareRoot(ctx context.Context, in *Squareroot, opts ...grpc.CallOption) (*Squareroot, error)
	PerfectNumber(ctx context.Context, in *PerfectNumberRequest, opts ...grpc.CallOption) (CalculateService_PerfectNumberClient, error)
	TotalNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_TotalNumberClient, error)
	FindMinimum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMinimumClient, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculatepb.CalculateService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) PrimeDecomposition(ctx context.Context, in *PrimeDecompositionRequest, opts ...grpc.CallOption) (CalculateService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[0], "/calculatepb.CalculateService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_PrimeDecompositionClient interface {
	Recv() (*PrimeDecompositionResponse, error)
	grpc.ClientStream
}

type calculateServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculateServicePrimeDecompositionClient) Recv() (*PrimeDecompositionResponse, error) {
	m := new(PrimeDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) ComputeAvarage(ctx context.Context, opts ...grpc.CallOption) (CalculateService_ComputeAvarageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[1], "/calculatepb.CalculateService/ComputeAvarage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceComputeAvarageClient{stream}
	return x, nil
}

type CalculateService_ComputeAvarageClient interface {
	Send(*ComputeAvarageRequest) error
	CloseAndRecv() (*ComputeAvarageResponse, error)
	grpc.ClientStream
}

type calculateServiceComputeAvarageClient struct {
	grpc.ClientStream
}

func (x *calculateServiceComputeAvarageClient) Send(m *ComputeAvarageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceComputeAvarageClient) CloseAndRecv() (*ComputeAvarageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAvarageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[2], "/calculatepb.CalculateService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceFindMaximumClient{stream}
	return x, nil
}

type CalculateService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculateServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculateServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *Squareroot, opts ...grpc.CallOption) (*Squareroot, error) {
	out := new(Squareroot)
	err := c.cc.Invoke(ctx, "/calculatepb.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) PerfectNumber(ctx context.Context, in *PerfectNumberRequest, opts ...grpc.CallOption) (CalculateService_PerfectNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[3], "/calculatepb.CalculateService/PerfectNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServicePerfectNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_PerfectNumberClient interface {
	Recv() (*PerfectNumberResponse, error)
	grpc.ClientStream
}

type calculateServicePerfectNumberClient struct {
	grpc.ClientStream
}

func (x *calculateServicePerfectNumberClient) Recv() (*PerfectNumberResponse, error) {
	m := new(PerfectNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) TotalNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_TotalNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[4], "/calculatepb.CalculateService/TotalNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceTotalNumberClient{stream}
	return x, nil
}

type CalculateService_TotalNumberClient interface {
	Send(*TotalNumberRequest) error
	CloseAndRecv() (*TotalNumberResponse, error)
	grpc.ClientStream
}

type calculateServiceTotalNumberClient struct {
	grpc.ClientStream
}

func (x *calculateServiceTotalNumberClient) Send(m *TotalNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceTotalNumberClient) CloseAndRecv() (*TotalNumberResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TotalNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) FindMinimum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMinimumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[5], "/calculatepb.CalculateService/FindMinimum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceFindMinimumClient{stream}
	return x, nil
}

type CalculateService_FindMinimumClient interface {
	Send(*FindMinimumRequest) error
	Recv() (*FindMinimumResponse, error)
	grpc.ClientStream
}

type calculateServiceFindMinimumClient struct {
	grpc.ClientStream
}

func (x *calculateServiceFindMinimumClient) Send(m *FindMinimumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceFindMinimumClient) Recv() (*FindMinimumResponse, error) {
	m := new(FindMinimumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations must embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server-side stream
	PrimeDecomposition(*PrimeDecompositionRequest, CalculateService_PrimeDecompositionServer) error
	// Client-side stream
	ComputeAvarage(CalculateService_ComputeAvarageServer) error
	// Bi-Directional stream
	FindMaximum(CalculateService_FindMaximumServer) error
	SquareRoot(context.Context, *Squareroot) (*Squareroot, error)
	PerfectNumber(*PerfectNumberRequest, CalculateService_PerfectNumberServer) error
	TotalNumber(CalculateService_TotalNumberServer) error
	FindMinimum(CalculateService_FindMinimumServer) error
	mustEmbedUnimplementedCalculateServiceServer()
}

// UnimplementedCalculateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculateServiceServer) PrimeDecomposition(*PrimeDecompositionRequest, CalculateService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (UnimplementedCalculateServiceServer) ComputeAvarage(CalculateService_ComputeAvarageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvarage not implemented")
}
func (UnimplementedCalculateServiceServer) FindMaximum(CalculateService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (UnimplementedCalculateServiceServer) SquareRoot(context.Context, *Squareroot) (*Squareroot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}
func (UnimplementedCalculateServiceServer) PerfectNumber(*PerfectNumberRequest, CalculateService_PerfectNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method PerfectNumber not implemented")
}
func (UnimplementedCalculateServiceServer) TotalNumber(CalculateService_TotalNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method TotalNumber not implemented")
}
func (UnimplementedCalculateServiceServer) FindMinimum(CalculateService_FindMinimumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMinimum not implemented")
}
func (UnimplementedCalculateServiceServer) mustEmbedUnimplementedCalculateServiceServer() {}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculatepb.CalculateService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).PrimeDecomposition(m, &calculateServicePrimeDecompositionServer{stream})
}

type CalculateService_PrimeDecompositionServer interface {
	Send(*PrimeDecompositionResponse) error
	grpc.ServerStream
}

type calculateServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculateServicePrimeDecompositionServer) Send(m *PrimeDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_ComputeAvarage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).ComputeAvarage(&calculateServiceComputeAvarageServer{stream})
}

type CalculateService_ComputeAvarageServer interface {
	SendAndClose(*ComputeAvarageResponse) error
	Recv() (*ComputeAvarageRequest, error)
	grpc.ServerStream
}

type calculateServiceComputeAvarageServer struct {
	grpc.ServerStream
}

func (x *calculateServiceComputeAvarageServer) SendAndClose(m *ComputeAvarageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceComputeAvarageServer) Recv() (*ComputeAvarageRequest, error) {
	m := new(ComputeAvarageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).FindMaximum(&calculateServiceFindMaximumServer{stream})
}

type CalculateService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculateServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculateServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Squareroot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculatepb.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*Squareroot))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_PerfectNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PerfectNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).PerfectNumber(m, &calculateServicePerfectNumberServer{stream})
}

type CalculateService_PerfectNumberServer interface {
	Send(*PerfectNumberResponse) error
	grpc.ServerStream
}

type calculateServicePerfectNumberServer struct {
	grpc.ServerStream
}

func (x *calculateServicePerfectNumberServer) Send(m *PerfectNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_TotalNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).TotalNumber(&calculateServiceTotalNumberServer{stream})
}

type CalculateService_TotalNumberServer interface {
	SendAndClose(*TotalNumberResponse) error
	Recv() (*TotalNumberRequest, error)
	grpc.ServerStream
}

type calculateServiceTotalNumberServer struct {
	grpc.ServerStream
}

func (x *calculateServiceTotalNumberServer) SendAndClose(m *TotalNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceTotalNumberServer) Recv() (*TotalNumberRequest, error) {
	m := new(TotalNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_FindMinimum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).FindMinimum(&calculateServiceFindMinimumServer{stream})
}

type CalculateService_FindMinimumServer interface {
	Send(*FindMinimumResponse) error
	Recv() (*FindMinimumRequest, error)
	grpc.ServerStream
}

type calculateServiceFindMinimumServer struct {
	grpc.ServerStream
}

func (x *calculateServiceFindMinimumServer) Send(m *FindMinimumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceFindMinimumServer) Recv() (*FindMinimumRequest, error) {
	m := new(FindMinimumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculatepb.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculateService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _CalculateService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvarage",
			Handler:       _CalculateService_ComputeAvarage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculateService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PerfectNumber",
			Handler:       _CalculateService_PerfectNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TotalNumber",
			Handler:       _CalculateService_TotalNumber_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMinimum",
			Handler:       _CalculateService_FindMinimum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculatepb/calculatepb.proto",
}