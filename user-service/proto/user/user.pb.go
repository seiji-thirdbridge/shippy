// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	User
	Request
	Response
	Token
	Error
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Response) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x6e, 0xf3, 0xd7, 0xf6, 0x94, 0x7b, 0xe1, 0x0e, 0x57, 0x1c, 0xe2, 0xa6, 0x64, 0x25, 0x88,
	0x51, 0xea, 0xba, 0x62, 0x29, 0xa5, 0xfb, 0xf8, 0xb3, 0x8f, 0xc9, 0x41, 0x07, 0x93, 0x4c, 0x9c,
	0x99, 0x46, 0x7c, 0x13, 0x9f, 0xcd, 0xa7, 0x91, 0x39, 0x69, 0x45, 0x30, 0x55, 0xe8, 0xa6, 0x3d,
	0xdf, 0x4f, 0x32, 0xdf, 0x9c, 0x8f, 0xc0, 0x41, 0xad, 0xa4, 0x91, 0x67, 0x6b, 0x8d, 0x8a, 0x7e,
	0x62, 0xc2, 0xec, 0xdf, 0x83, 0x8c, 0x4b, 0x91, 0x29, 0x19, 0x6b, 0xd5, 0xc4, 0x56, 0x88, 0x1a,
	0xf0, 0x6e, 0x35, 0x2a, 0xf6, 0x17, 0x1c, 0x91, 0xf3, 0xfe, 0xa4, 0x7f, 0x3c, 0x4a, 0x1c, 0x91,
	0x33, 0x06, 0x5e, 0x95, 0x96, 0xc8, 0x1d, 0x62, 0x68, 0x66, 0x1c, 0x06, 0x99, 0x2c, 0xeb, 0xb4,
	0x7a, 0xe5, 0x2e, 0xd1, 0x5b, 0xc8, 0xfe, 0x83, 0x8f, 0x65, 0x2a, 0x0a, 0xee, 0x11, 0xdf, 0x02,
	0x16, 0xc2, 0xb0, 0x4e, 0xb5, 0x7e, 0x91, 0x2a, 0xe7, 0x3e, 0x09, 0x9f, 0x38, 0x1a, 0xc1, 0x20,
	0xc1, 0xe7, 0x35, 0x6a, 0x13, 0xbd, 0xf5, 0x61, 0x98, 0xa0, 0xae, 0x65, 0xa5, 0x91, 0x9d, 0x80,
	0x67, 0x73, 0x51, 0x92, 0xf1, 0xf4, 0x30, 0xfe, 0x96, 0x38, 0xb6, 0x71, 0x13, 0x32, 0xb1, 0x53,
	0xf0, 0xed, 0xbf, 0xe6, 0xce, 0xc4, 0xfd, 0xc9, 0xdd, 0xba, 0xd8, 0x39, 0x04, 0xa8, 0x94, 0x54,
	0x9a, 0xbb, 0xe4, 0xe7, 0x1d, 0xfe, 0xa5, 0x35, 0x24, 0x1b, 0x5f, 0x84, 0xe0, 0xdf, 0xc8, 0x27,
	0xac, 0xec, 0x05, 0x8d, 0x1d, 0x36, 0x1b, 0x6a, 0x81, 0x65, 0x9b, 0xb4, 0x10, 0x39, 0x6d, 0x69,
	0x98, 0xb4, 0x60, 0x8f, 0x63, 0x66, 0xe0, 0x13, 0x61, 0xb7, 0x9e, 0xc9, 0x1c, 0xe9, 0x14, 0x3f,
	0xa1, 0x99, 0x4d, 0x60, 0x9c, 0xa3, 0xce, 0x94, 0xa8, 0x8d, 0x90, 0xd5, 0xa6, 0x90, 0xaf, 0xd4,
	0xf4, 0xdd, 0x81, 0xb1, 0xbd, 0xe7, 0x35, 0xaa, 0x46, 0x64, 0xc8, 0xae, 0x20, 0x58, 0x28, 0x4c,
	0x0d, 0xb2, 0x5d, 0x1b, 0x09, 0x8f, 0x3a, 0x84, 0x6d, 0x07, 0x51, 0x8f, 0xcd, 0xc0, 0x5d, 0xa1,
	0xd9, 0xfb, 0xf1, 0x05, 0x04, 0x2b, 0x34, 0xf3, 0xa2, 0x60, 0x61, 0xa7, 0x91, 0x7a, 0xff, 0xed,
	0x25, 0x97, 0xe0, 0xcd, 0xd7, 0xe6, 0x71, 0xef, 0x10, 0x4b, 0xf8, 0x73, 0x67, 0xfb, 0x48, 0x0d,
	0xb6, 0x1d, 0x76, 0xf5, 0x40, 0x4a, 0xb8, 0x53, 0x89, 0x7a, 0xf7, 0x01, 0x7d, 0x3a, 0x17, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x8d, 0xf8, 0x32, 0x53, 0x03, 0x00, 0x00,
}
