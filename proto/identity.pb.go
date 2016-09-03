// Code generated by protoc-gen-go.
// source: identity.proto
// DO NOT EDIT!

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	identity.proto

It has these top-level messages:
	RegisterRequest
	RegisterResponse
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterResponse struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "identity.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "identity.RegisterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Identity service

type IdentityClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/identity.Identity/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Identity service

type IdentityServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.Identity/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identity.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Identity_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x3c,
	0xb9, 0xf8, 0x83, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a,
	0x04, 0x13, 0x44, 0x0e, 0xc6, 0x57, 0xb2, 0xe3, 0x12, 0x40, 0x18, 0x55, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x2a, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x02, 0x35, 0x85, 0x29, 0x33, 0x05, 0xc5, 0x6c, 0x26,
	0x54, 0xb3, 0x8d, 0xfc, 0xb9, 0x38, 0x3c, 0xa1, 0xce, 0x12, 0x72, 0xe6, 0xe2, 0x80, 0x99, 0x25,
	0x24, 0xa9, 0x07, 0x77, 0x3d, 0x9a, 0x53, 0xa5, 0xa4, 0xb0, 0x49, 0x41, 0xac, 0x56, 0x62, 0x48,
	0x62, 0x03, 0x7b, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x21, 0x23, 0xd1, 0xfe, 0x00,
	0x00, 0x00,
}
