// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PeticionClient is the client API for Peticion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeticionClient interface {
	RealizarPeticion(ctx context.Context, in *Objeto, opts ...grpc.CallOption) (*Serie, error)
}

type peticionClient struct {
	cc grpc.ClientConnInterface
}

func NewPeticionClient(cc grpc.ClientConnInterface) PeticionClient {
	return &peticionClient{cc}
}

func (c *peticionClient) RealizarPeticion(ctx context.Context, in *Objeto, opts ...grpc.CallOption) (*Serie, error) {
	out := new(Serie)
	err := c.cc.Invoke(ctx, "/pb.Peticion/RealizarPeticion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeticionServer is the server API for Peticion service.
// All implementations must embed UnimplementedPeticionServer
// for forward compatibility
type PeticionServer interface {
	RealizarPeticion(context.Context, *Objeto) (*Serie, error)
	mustEmbedUnimplementedPeticionServer()
}

// UnimplementedPeticionServer must be embedded to have forward compatible implementations.
type UnimplementedPeticionServer struct {
}

func (UnimplementedPeticionServer) RealizarPeticion(context.Context, *Objeto) (*Serie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RealizarPeticion not implemented")
}
func (UnimplementedPeticionServer) mustEmbedUnimplementedPeticionServer() {}

// UnsafePeticionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeticionServer will
// result in compilation errors.
type UnsafePeticionServer interface {
	mustEmbedUnimplementedPeticionServer()
}

func RegisterPeticionServer(s *grpc.Server, srv PeticionServer) {
	s.RegisterService(&_Peticion_serviceDesc, srv)
}

func _Peticion_RealizarPeticion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Objeto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeticionServer).RealizarPeticion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Peticion/RealizarPeticion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeticionServer).RealizarPeticion(ctx, req.(*Objeto))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peticion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Peticion",
	HandlerType: (*PeticionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RealizarPeticion",
			Handler:    _Peticion_RealizarPeticion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "producto.proto",
}
