// Code generated by protoc-gen-go. DO NOT EDIT.
// source: library.proto

package api_pb // import "github.com/izumin5210-sandbox/grapi-playground/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _type "github.com/izumin5210-sandbox/grapi-playground/api/type"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ListBooksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksRequest) Reset()         { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()    {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{0}
}
func (m *ListBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksRequest.Unmarshal(m, b)
}
func (m *ListBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksRequest.Marshal(b, m, deterministic)
}
func (dst *ListBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksRequest.Merge(dst, src)
}
func (m *ListBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBooksRequest.Size(m)
}
func (m *ListBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksRequest proto.InternalMessageInfo

type ListBooksResponse struct {
	Books                []*_type.Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{1}
}
func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (dst *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(dst, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*_type.Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	BookId               string   `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{2}
}
func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (dst *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(dst, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type CreateBookRequest struct {
	Book                 *_type.Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{3}
}
func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(dst, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

func (m *CreateBookRequest) GetBook() *_type.Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type UpdateBookRequest struct {
	Book                 *_type.Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateBookRequest) Reset()         { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()    {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{4}
}
func (m *UpdateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBookRequest.Unmarshal(m, b)
}
func (m *UpdateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBookRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBookRequest.Merge(dst, src)
}
func (m *UpdateBookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBookRequest.Size(m)
}
func (m *UpdateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBookRequest proto.InternalMessageInfo

func (m *UpdateBookRequest) GetBook() *_type.Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type DeleteBookRequest struct {
	BookId               string   `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookRequest) Reset()         { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()    {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_5c19e92b1d6c688c, []int{5}
}
func (m *DeleteBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookRequest.Unmarshal(m, b)
}
func (m *DeleteBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookRequest.Merge(dst, src)
}
func (m *DeleteBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookRequest.Size(m)
}
func (m *DeleteBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookRequest proto.InternalMessageInfo

func (m *DeleteBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

func init() {
	proto.RegisterType((*ListBooksRequest)(nil), "izumin5210.sandbox.grapi_playground.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "izumin5210.sandbox.grapi_playground.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "izumin5210.sandbox.grapi_playground.GetBookRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "izumin5210.sandbox.grapi_playground.CreateBookRequest")
	proto.RegisterType((*UpdateBookRequest)(nil), "izumin5210.sandbox.grapi_playground.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "izumin5210.sandbox.grapi_playground.DeleteBookRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LibraryService service

type LibraryServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*_type.Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*_type.Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*_type.Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type libraryServiceClient struct {
	cc *grpc.ClientConn
}

func NewLibraryServiceClient(cc *grpc.ClientConn) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/ListBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*_type.Book, error) {
	out := new(_type.Book)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*_type.Book, error) {
	out := new(_type.Book)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*_type.Book, error) {
	out := new(_type.Book)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/UpdateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/DeleteBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LibraryService service

type LibraryServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*_type.Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*_type.Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*_type.Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*empty.Empty, error)
}

func RegisterLibraryServiceServer(s *grpc.Server, srv LibraryServiceServer) {
	s.RegisterService(&_LibraryService_serviceDesc, srv)
}

func _LibraryService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.sandbox.grapi_playground.LibraryService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.sandbox.grapi_playground.LibraryService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.sandbox.grapi_playground.LibraryService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.sandbox.grapi_playground.LibraryService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.sandbox.grapi_playground.LibraryService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LibraryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "izumin5210.sandbox.grapi_playground.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _LibraryService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _LibraryService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _LibraryService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}

func init() { proto.RegisterFile("library.proto", fileDescriptor_library_5c19e92b1d6c688c) }

var fileDescriptor_library_5c19e92b1d6c688c = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0x25, 0xea, 0x76, 0xdd, 0x9f, 0x18, 0xb7, 0x83, 0x7f, 0x4a, 0x54, 0x58, 0xb2, 0x97, 0x55,
	0xdc, 0x89, 0xa6, 0x54, 0xd0, 0xde, 0x6a, 0x45, 0x84, 0x9e, 0x2a, 0x22, 0x7a, 0x29, 0x93, 0x66,
	0x8c, 0x43, 0xd3, 0xcc, 0x98, 0x99, 0x14, 0xab, 0x78, 0xf1, 0x24, 0x5e, 0x3d, 0xf9, 0x91, 0x3c,
	0xfb, 0x15, 0xfc, 0x20, 0x32, 0x93, 0xa9, 0x69, 0x0d, 0x42, 0x72, 0xd8, 0x63, 0x32, 0xf3, 0x7b,
	0xef, 0xcd, 0x7b, 0x8f, 0x1f, 0x5c, 0x4e, 0x59, 0x94, 0x93, 0x7c, 0x8d, 0x45, 0xce, 0x15, 0x47,
	0xc7, 0xec, 0x63, 0xb1, 0x64, 0xd9, 0x20, 0x7c, 0x70, 0x1f, 0x4b, 0x92, 0xc5, 0x11, 0xff, 0x80,
	0x93, 0x9c, 0x08, 0x36, 0x13, 0x29, 0x59, 0x27, 0x39, 0x2f, 0xb2, 0xd8, 0xbb, 0x95, 0x70, 0x9e,
	0xa4, 0x34, 0x20, 0x82, 0x05, 0x24, 0xcb, 0xb8, 0x22, 0x8a, 0xf1, 0x4c, 0x96, 0x10, 0xde, 0x4d,
	0x7b, 0x6a, 0xbe, 0xa2, 0xe2, 0x6d, 0x40, 0x97, 0x42, 0x59, 0x7c, 0xef, 0x8a, 0x5a, 0x0b, 0x1a,
	0x44, 0x9c, 0x2f, 0xca, 0x1f, 0x3e, 0x82, 0xc3, 0x09, 0x93, 0x6a, 0xc4, 0xf9, 0x42, 0x4e, 0xe9,
	0xfb, 0x82, 0x4a, 0xe5, 0xbf, 0x86, 0xee, 0xd6, 0x3f, 0x29, 0x78, 0x26, 0x29, 0x1a, 0xc3, 0x9e,
	0x1e, 0x93, 0x3d, 0xe7, 0xe8, 0xfc, 0xc9, 0xa5, 0x10, 0xe3, 0x06, 0x4a, 0xb1, 0x66, 0xc3, 0x1a,
	0x67, 0x5a, 0x0e, 0xfb, 0x77, 0xc0, 0x7d, 0x46, 0x0d, 0xb2, 0x25, 0x43, 0x37, 0x60, 0x5f, 0x1f,
	0xcd, 0x58, 0xdc, 0x73, 0x8e, 0x9c, 0x93, 0x83, 0x69, 0x47, 0x7f, 0x3e, 0x8f, 0xfd, 0x57, 0xd0,
	0x7d, 0x92, 0x53, 0xa2, 0xe8, 0xf6, 0xed, 0x11, 0x5c, 0xd0, 0xc7, 0xe6, 0x6a, 0x7b, 0x11, 0x66,
	0x56, 0x03, 0xbf, 0x14, 0xf1, 0x19, 0x00, 0xdf, 0x83, 0xee, 0x98, 0xa6, 0x74, 0x17, 0xf8, 0x7f,
	0xef, 0x0b, 0x7f, 0xee, 0x81, 0x3b, 0x29, 0xc3, 0x7f, 0x41, 0xf3, 0x15, 0x9b, 0x53, 0xf4, 0xcd,
	0x81, 0x83, 0xbf, 0xce, 0xa3, 0x41, 0x23, 0x11, 0xff, 0xa6, 0xe7, 0x3d, 0x6c, 0x3b, 0x56, 0x06,
	0xec, 0xbb, 0x5f, 0x7e, 0xfd, 0xfe, 0x7e, 0xee, 0x22, 0xea, 0x98, 0x7a, 0x48, 0xf4, 0xd5, 0x81,
	0x7d, 0x9b, 0x15, 0xea, 0x37, 0xc2, 0xdc, 0x4d, 0xd6, 0x6b, 0x69, 0xa2, 0xdf, 0x33, 0x02, 0x10,
	0x3a, 0x2c, 0x05, 0x04, 0x9f, 0xac, 0x6f, 0x9f, 0xb5, 0x2f, 0x50, 0x75, 0x01, 0x35, 0x7b, 0x61,
	0xad, 0x3c, 0xad, 0x05, 0x5d, 0x35, 0x82, 0x5c, 0xdf, 0x3a, 0xf2, 0xd8, 0xa4, 0x8c, 0x7e, 0x38,
	0x00, 0x55, 0x7f, 0x1a, 0x8a, 0xa9, 0x15, 0xae, 0xb5, 0x98, 0x63, 0x23, 0xe6, 0x76, 0x78, 0x6d,
	0xdb, 0x1d, 0xbc, 0xb1, 0xc8, 0x6a, 0x5b, 0x01, 0x54, 0x0d, 0x6c, 0x28, 0xad, 0x56, 0x59, 0xef,
	0x3a, 0x2e, 0x57, 0x08, 0xde, 0xac, 0x10, 0xfc, 0x54, 0xaf, 0x90, 0x4d, 0x40, 0x77, 0x6b, 0x01,
	0x8d, 0x86, 0x6f, 0x1e, 0x25, 0x4c, 0xbd, 0x2b, 0x22, 0x3c, 0xe7, 0xcb, 0xa0, 0x62, 0x3d, 0xb5,
	0xac, 0x81, 0x61, 0x3d, 0xad, 0x58, 0xf5, 0xee, 0x1a, 0x1a, 0x21, 0x51, 0xd4, 0x31, 0x34, 0xfd,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe0, 0xd0, 0xf4, 0x0b, 0x05, 0x00, 0x00,
}
