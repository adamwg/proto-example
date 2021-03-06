// Code generated by protoc-gen-go.
// source: model/server.proto
// DO NOT EDIT!

package model

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

type RegisterUserRequest struct {
	// User is the user's info. It is an error to specify a user ID.
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *RegisterUserRequest) Reset()                    { *m = RegisterUserRequest{} }
func (m *RegisterUserRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterUserRequest) ProtoMessage()               {}
func (*RegisterUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RegisterUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type RegisterUserResponse struct {
	// Id is the ID assigned to the user.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RegisterUserResponse) Reset()                    { *m = RegisterUserResponse{} }
func (m *RegisterUserResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterUserResponse) ProtoMessage()               {}
func (*RegisterUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *RegisterUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AuthenticateRequest struct {
	// UserId is the ID of the user authenticating.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Password is the user's hashed password.
	Password []byte `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *AuthenticateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type AuthenticateResponse struct {
	// Token is an authentication token that can be used to post.
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *AuthenticateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetUserNameRequest struct {
	// UserId is the ID of the user whose name is requested.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserNameRequest) Reset()                    { *m = GetUserNameRequest{} }
func (m *GetUserNameRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserNameRequest) ProtoMessage()               {}
func (*GetUserNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetUserNameRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserNameResponse struct {
	// UserName is the name of the user.
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *GetUserNameResponse) Reset()                    { *m = GetUserNameResponse{} }
func (m *GetUserNameResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserNameResponse) ProtoMessage()               {}
func (*GetUserNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetUserNameResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type PostRequest struct {
	// Token is the authentication token for this post request. The user ID in
	// the post must match the user that was granted this token.
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Post is the post. If an ID is specified it will be ignored.
	Post *Post `protobuf:"bytes,2,opt,name=post" json:"post,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type PostResponse struct {
	// Id is the ID assigned to the post.
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PostResponse) Reset()                    { *m = PostResponse{} }
func (m *PostResponse) String() string            { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()               {}
func (*PostResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *PostResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadRequest struct {
	// PostIds is the set of post IDs requested. If empty, all posts will be
	// returned.
	PostIds []uint64 `protobuf:"varint,1,rep,packed,name=post_ids,json=postIds" json:"post_ids,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ReadRequest) GetPostIds() []uint64 {
	if m != nil {
		return m.PostIds
	}
	return nil
}

type ReadResponse struct {
	// Post is a single post. ReadResponses will be streamed to the client if
	// multiple posts are requested.
	Post *Post `protobuf:"bytes,1,opt,name=post" json:"post,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ReadResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterUserRequest)(nil), "model.RegisterUserRequest")
	proto.RegisterType((*RegisterUserResponse)(nil), "model.RegisterUserResponse")
	proto.RegisterType((*AuthenticateRequest)(nil), "model.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "model.AuthenticateResponse")
	proto.RegisterType((*GetUserNameRequest)(nil), "model.GetUserNameRequest")
	proto.RegisterType((*GetUserNameResponse)(nil), "model.GetUserNameResponse")
	proto.RegisterType((*PostRequest)(nil), "model.PostRequest")
	proto.RegisterType((*PostResponse)(nil), "model.PostResponse")
	proto.RegisterType((*ReadRequest)(nil), "model.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "model.ReadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Posts service

type PostsClient interface {
	// RegisterUser registers a user.
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	// Authenticate authenticates a user.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// GetUserName gets a user's name.
	GetUserName(ctx context.Context, in *GetUserNameRequest, opts ...grpc.CallOption) (*GetUserNameResponse, error)
	// Post creates a post.
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	// Read gets posts.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (Posts_ReadClient, error)
}

type postsClient struct {
	cc *grpc.ClientConn
}

func NewPostsClient(cc *grpc.ClientConn) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := grpc.Invoke(ctx, "/model.Posts/RegisterUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/model.Posts/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetUserName(ctx context.Context, in *GetUserNameRequest, opts ...grpc.CallOption) (*GetUserNameResponse, error) {
	out := new(GetUserNameResponse)
	err := grpc.Invoke(ctx, "/model.Posts/GetUserName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := grpc.Invoke(ctx, "/model.Posts/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (Posts_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Posts_serviceDesc.Streams[0], c.cc, "/model.Posts/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &postsReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Posts_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type postsReadClient struct {
	grpc.ClientStream
}

func (x *postsReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Posts service

type PostsServer interface {
	// RegisterUser registers a user.
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	// Authenticate authenticates a user.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// GetUserName gets a user's name.
	GetUserName(context.Context, *GetUserNameRequest) (*GetUserNameResponse, error)
	// Post creates a post.
	Post(context.Context, *PostRequest) (*PostResponse, error)
	// Read gets posts.
	Read(*ReadRequest, Posts_ReadServer) error
}

func RegisterPostsServer(s *grpc.Server, srv PostsServer) {
	s.RegisterService(&_Posts_serviceDesc, srv)
}

func _Posts_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Posts/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Posts/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Posts/GetUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetUserName(ctx, req.(*GetUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Posts/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostsServer).Read(m, &postsReadServer{stream})
}

type Posts_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type postsReadServer struct {
	grpc.ServerStream
}

func (x *postsReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Posts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Posts_RegisterUser_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Posts_Authenticate_Handler,
		},
		{
			MethodName: "GetUserName",
			Handler:    _Posts_GetUserName_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Posts_Post_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _Posts_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "model/server.proto",
}

func init() { proto.RegisterFile("model/server.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x63, 0xd7, 0xf9, 0x37, 0x0e, 0x3d, 0xc8, 0x81, 0x26, 0x0e, 0xb4, 0x41, 0x87, 0xe2,
	0x43, 0x9b, 0xb4, 0x09, 0xf4, 0x5e, 0x28, 0x2d, 0xe9, 0xa1, 0x14, 0x43, 0xcf, 0xc1, 0x5d, 0x0d,
	0xbb, 0x66, 0x63, 0xcb, 0x6b, 0xc9, 0xbb, 0xfb, 0x25, 0xf7, 0x3b, 0x2d, 0x92, 0x65, 0xc5, 0x26,
	0x5e, 0xf6, 0x38, 0x9e, 0xf7, 0x7e, 0x23, 0xbd, 0x91, 0x81, 0x64, 0x9c, 0xe1, 0x69, 0x2b, 0xb0,
	0xbc, 0xc7, 0x72, 0x53, 0x94, 0x5c, 0x72, 0x32, 0xd4, 0xdf, 0xc2, 0xa0, 0x6e, 0xe1, 0x63, 0x92,
	0x15, 0x27, 0xac, 0x7b, 0xf4, 0x1b, 0x04, 0x31, 0x5e, 0xa7, 0x42, 0x62, 0xf9, 0x4f, 0x60, 0x19,
	0xe3, 0x5d, 0x85, 0x42, 0x92, 0x0f, 0xe0, 0x55, 0x02, 0xcb, 0x85, 0xb3, 0x76, 0x22, 0x7f, 0xe7,
	0x6f, 0xb4, 0x75, 0xa3, 0x15, 0xba, 0x41, 0x3f, 0xc2, 0xbc, 0xeb, 0x13, 0x05, 0xcf, 0x05, 0x92,
	0xb7, 0xe0, 0xa6, 0x4c, 0xdb, 0xa6, 0xb1, 0x9b, 0x32, 0xfa, 0x1b, 0x82, 0xef, 0x95, 0xbc, 0xc1,
	0x5c, 0xa6, 0x57, 0x89, 0xc4, 0x86, 0xff, 0x0e, 0xc6, 0x0a, 0x73, 0xb4, 0xda, 0x91, 0x2a, 0x0f,
	0x8c, 0x84, 0x30, 0x29, 0x12, 0x21, 0x1e, 0x78, 0xc9, 0x16, 0xee, 0xda, 0x89, 0x66, 0xb1, 0xad,
	0xe9, 0x27, 0x98, 0x77, 0x59, 0x66, 0xe6, 0x1c, 0x86, 0x92, 0xdf, 0x62, 0x6e, 0x50, 0x75, 0x41,
	0x3f, 0x03, 0xf9, 0x85, 0x52, 0x1d, 0xee, 0x4f, 0x92, 0xbd, 0x3a, 0x98, 0xee, 0x20, 0xe8, 0xc8,
	0x0d, 0x7b, 0x05, 0x53, 0xad, 0xcf, 0x93, 0x0c, 0x8d, 0x63, 0x52, 0x19, 0x11, 0xfd, 0x01, 0xfe,
	0x5f, 0x2e, 0x64, 0xc3, 0xee, 0x3d, 0x87, 0x8a, 0xb2, 0xe0, 0x42, 0xea, 0xdb, 0x9c, 0xa3, 0xd4,
	0x3e, 0xdd, 0xa0, 0xef, 0x61, 0x56, 0x53, 0x2e, 0x22, 0xf4, 0x74, 0x84, 0x11, 0xf8, 0x31, 0x26,
	0xac, 0x99, 0xb2, 0x84, 0x89, 0xb2, 0x1d, 0x53, 0x26, 0x16, 0xce, 0xfa, 0x4d, 0xe4, 0xc5, 0x63,
	0x55, 0x1f, 0x98, 0xa0, 0x5b, 0x98, 0xd5, 0x4a, 0x43, 0x6a, 0x46, 0x3b, 0x2f, 0x8c, 0xde, 0x3d,
	0xb9, 0x30, 0x54, 0xa5, 0x20, 0x07, 0x65, 0x3d, 0xef, 0x93, 0x84, 0x46, 0xdc, 0xf3, 0x38, 0xc2,
	0x55, 0x6f, 0xaf, 0x9e, 0x49, 0x07, 0x0a, 0xd5, 0x5e, 0x93, 0x45, 0xf5, 0xbc, 0x03, 0x8b, 0xea,
	0xdb, 0x2b, 0x1d, 0x90, 0x9f, 0xe0, 0xb7, 0x96, 0x42, 0x96, 0x46, 0x7d, 0xb9, 0xd7, 0x30, 0xec,
	0x6b, 0x59, 0xce, 0x57, 0xf0, 0xd4, 0x35, 0x09, 0x69, 0x47, 0x60, 0x9c, 0x41, 0xe7, 0x9b, 0xb5,
	0xec, 0xc1, 0x53, 0x59, 0x5a, 0x4b, 0x6b, 0x05, 0xd6, 0xd2, 0x0e, 0x9b, 0x0e, 0xbe, 0x38, 0xff,
	0x47, 0xfa, 0xa7, 0xda, 0x3f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x52, 0xe9, 0x4a, 0x00, 0x86, 0x03,
	0x00, 0x00,
}
