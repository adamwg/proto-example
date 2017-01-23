// Code generated by protoc-gen-go.
// source: model/example.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model/example.proto
	model/server.proto

It has these top-level messages:
	User
	Post
	PostContent
	Attachment
	File
	RegisterUserRequest
	RegisterUserResponse
	AuthenticateRequest
	AuthenticateResponse
	GetUserNameRequest
	GetUserNameResponse
	PostRequest
	PostResponse
	ReadRequest
	ReadResponse
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// User represents a registered user.
type User struct {
	// Id is a unique ID for this user.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Name is the name of this user.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Password is the user's hashed password.
	Password []byte `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
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

func (m *User) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

// Post is a post from a user.
type Post struct {
	// Id is the unique ID of this post.
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// UserId is the ID of the user who posted this post.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Content is the content of the post.
	Content *PostContent `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Post) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Post) GetContent() *PostContent {
	if m != nil {
		return m.Content
	}
	return nil
}

// PostContent is the content of a post.
type PostContent struct {
	// Text is the text of the post.
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// Attachments are optional attachment to the post.
	Attachments []*Attachment `protobuf:"bytes,2,rep,name=attachments" json:"attachments,omitempty"`
}

func (m *PostContent) Reset()                    { *m = PostContent{} }
func (m *PostContent) String() string            { return proto.CompactTextString(m) }
func (*PostContent) ProtoMessage()               {}
func (*PostContent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PostContent) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PostContent) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

// Attachment is a non-text attachment to a post.
type Attachment struct {
	// Types that are valid to be assigned to Attachment:
	//	*Attachment_Url
	//	*Attachment_File
	Attachment isAttachment_Attachment `protobuf_oneof:"attachment"`
}

func (m *Attachment) Reset()                    { *m = Attachment{} }
func (m *Attachment) String() string            { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()               {}
func (*Attachment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isAttachment_Attachment interface {
	isAttachment_Attachment()
}

type Attachment_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,oneof"`
}
type Attachment_File struct {
	File *File `protobuf:"bytes,2,opt,name=file,oneof"`
}

func (*Attachment_Url) isAttachment_Attachment()  {}
func (*Attachment_File) isAttachment_Attachment() {}

func (m *Attachment) GetAttachment() isAttachment_Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *Attachment) GetUrl() string {
	if x, ok := m.GetAttachment().(*Attachment_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Attachment) GetFile() *File {
	if x, ok := m.GetAttachment().(*Attachment_File); ok {
		return x.File
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Attachment) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Attachment_OneofMarshaler, _Attachment_OneofUnmarshaler, _Attachment_OneofSizer, []interface{}{
		(*Attachment_Url)(nil),
		(*Attachment_File)(nil),
	}
}

func _Attachment_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Attachment)
	// attachment
	switch x := m.Attachment.(type) {
	case *Attachment_Url:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Url)
	case *Attachment_File:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Attachment.Attachment has unexpected type %T", x)
	}
	return nil
}

func _Attachment_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Attachment)
	switch tag {
	case 1: // attachment.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Attachment = &Attachment_Url{x}
		return true, err
	case 2: // attachment.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(File)
		err := b.DecodeMessage(msg)
		m.Attachment = &Attachment_File{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Attachment_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Attachment)
	// attachment
	switch x := m.Attachment.(type) {
	case *Attachment_Url:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *Attachment_File:
		s := proto.Size(x.File)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// File is a blob.
type File struct {
	// MimeType is the MIME type of the file.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType" json:"mime_type,omitempty"`
	// Content is the contents of the file.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *File) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *File) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*Post)(nil), "model.Post")
	proto.RegisterType((*PostContent)(nil), "model.PostContent")
	proto.RegisterType((*Attachment)(nil), "model.Attachment")
	proto.RegisterType((*File)(nil), "model.File")
}

func init() { proto.RegisterFile("model/example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xdb, 0xd4, 0xf4, 0xcf, 0xb9, 0x42, 0xe2, 0x18, 0x88, 0x60, 0x29, 0x99, 0x3a, 0xa0,
	0x22, 0xa5, 0x33, 0x03, 0x20, 0x55, 0x65, 0x03, 0x0b, 0xd8, 0x50, 0x15, 0xea, 0x43, 0x58, 0x8a,
	0xe3, 0x28, 0x76, 0x45, 0xfb, 0xed, 0x91, 0x9d, 0xb4, 0xc9, 0x76, 0xef, 0xbd, 0xd3, 0x4f, 0xcf,
	0x67, 0xb8, 0xd4, 0x46, 0x52, 0x7e, 0x4f, 0xfb, 0x4c, 0x97, 0x39, 0x2d, 0xca, 0xca, 0x38, 0x83,
	0x67, 0xc1, 0x4c, 0x56, 0xc0, 0x3e, 0x2c, 0x55, 0x78, 0x0e, 0x91, 0x92, 0x71, 0x7f, 0xd6, 0x9f,
	0x4f, 0x44, 0xa4, 0x24, 0x22, 0xb0, 0x22, 0xd3, 0x14, 0x47, 0xc1, 0x09, 0x33, 0x5e, 0xc3, 0xb8,
	0xcc, 0xac, 0xfd, 0x33, 0x95, 0x8c, 0x07, 0xb3, 0xfe, 0x7c, 0x2a, 0x4e, 0x3a, 0xf9, 0x02, 0xf6,
	0x6a, 0xac, 0xeb, 0x70, 0x58, 0xe0, 0x5c, 0xc1, 0x68, 0x67, 0xa9, 0xda, 0x28, 0xd9, 0xa0, 0x86,
	0x5e, 0xbe, 0x48, 0xbc, 0x83, 0xd1, 0xd6, 0x14, 0x8e, 0x0a, 0x17, 0x58, 0x3c, 0xc5, 0x45, 0x68,
	0xb4, 0xf0, 0x98, 0xe7, 0x3a, 0x11, 0xc7, 0x95, 0xe4, 0x13, 0x78, 0xc7, 0xf7, 0xed, 0x1c, 0xed,
	0x5d, 0xd3, 0x37, 0xcc, 0xb8, 0x04, 0x9e, 0x39, 0x97, 0x6d, 0x7f, 0x35, 0x15, 0xce, 0xc6, 0xd1,
	0x6c, 0x30, 0xe7, 0xe9, 0x45, 0x03, 0x7d, 0x3c, 0x25, 0xa2, 0xbb, 0x95, 0xbc, 0x01, 0xb4, 0x11,
	0x22, 0x0c, 0x76, 0x55, 0x5e, 0x53, 0xd7, 0x3d, 0xe1, 0x05, 0xde, 0x02, 0xfb, 0x51, 0x79, 0x7d,
	0x08, 0x9e, 0xf2, 0x86, 0xb7, 0x52, 0x39, 0xad, 0x7b, 0x22, 0x44, 0x4f, 0x53, 0x80, 0x96, 0x99,
	0x3c, 0x00, 0xf3, 0x29, 0xde, 0xc0, 0x44, 0x2b, 0x4d, 0x1b, 0x77, 0x28, 0xa9, 0x29, 0x3a, 0xf6,
	0xc6, 0xfb, 0xa1, 0x24, 0x8c, 0xdb, 0xd7, 0x47, 0xe1, 0x92, 0x47, 0xf9, 0x3d, 0x0c, 0xdf, 0xb3,
	0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xab, 0x16, 0xc1, 0x71, 0xb5, 0x01, 0x00, 0x00,
}
