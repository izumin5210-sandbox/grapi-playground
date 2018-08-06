// Code generated by protoc-gen-go. DO NOT EDIT.
// source: library.proto

package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import type_pb "github.com/izumin5210/grapi-playground/api/type_pb"
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
	return fileDescriptor_library_0491b63951b8e202, []int{0}
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
	Books                []*type_pb.Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_0491b63951b8e202, []int{1}
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

func (m *ListBooksResponse) GetBooks() []*type_pb.Book {
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
	return fileDescriptor_library_0491b63951b8e202, []int{2}
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
	Book                 *type_pb.Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_0491b63951b8e202, []int{3}
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

func (m *CreateBookRequest) GetBook() *type_pb.Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type UpdateBookRequest struct {
	Book                 *type_pb.Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateBookRequest) Reset()         { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()    {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_library_0491b63951b8e202, []int{4}
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

func (m *UpdateBookRequest) GetBook() *type_pb.Book {
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
	return fileDescriptor_library_0491b63951b8e202, []int{5}
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
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error)
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

func (c *libraryServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error) {
	out := new(type_pb.Book)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error) {
	out := new(type_pb.Book)
	err := grpc.Invoke(ctx, "/izumin5210.sandbox.grapi_playground.LibraryService/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*type_pb.Book, error) {
	out := new(type_pb.Book)
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
	GetBook(context.Context, *GetBookRequest) (*type_pb.Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*type_pb.Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*type_pb.Book, error)
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

func init() { proto.RegisterFile("library.proto", fileDescriptor_library_0491b63951b8e202) }

var fileDescriptor_library_0491b63951b8e202 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0xeb, 0xd3, 0x30,
	0x1c, 0xc6, 0xa9, 0xba, 0x6e, 0xfb, 0x8a, 0x75, 0x0b, 0xfe, 0x19, 0x55, 0x61, 0x74, 0x97, 0x29,
	0x92, 0x6a, 0xc7, 0x3c, 0x78, 0x9c, 0x13, 0x11, 0x76, 0x9a, 0x88, 0xe8, 0x65, 0xa4, 0x36, 0x96,
	0xb0, 0xae, 0x89, 0x4d, 0x3a, 0x9c, 0xe2, 0xc5, 0x93, 0x78, 0xf5, 0xe4, 0x4b, 0xf2, 0xec, 0x5b,
	0xf0, 0x85, 0x48, 0xd2, 0xce, 0x6d, 0xbf, 0xf2, 0x83, 0xf6, 0xf0, 0x3b, 0xb6, 0xc9, 0xf7, 0xc9,
	0x27, 0xcf, 0xf3, 0x10, 0xb8, 0x96, 0xb0, 0x30, 0x23, 0xd9, 0x0e, 0x8b, 0x8c, 0x2b, 0x8e, 0x46,
	0xec, 0x73, 0xbe, 0x61, 0xe9, 0x34, 0x78, 0xfc, 0x08, 0x4b, 0x92, 0x46, 0x21, 0xff, 0x84, 0xe3,
	0x8c, 0x08, 0xb6, 0x12, 0x09, 0xd9, 0xc5, 0x19, 0xcf, 0xd3, 0xc8, 0xbd, 0x1b, 0x73, 0x1e, 0x27,
	0xd4, 0x27, 0x82, 0xf9, 0x24, 0x4d, 0xb9, 0x22, 0x8a, 0xf1, 0x54, 0x16, 0x12, 0xee, 0x9d, 0x72,
	0xd5, 0x7c, 0x85, 0xf9, 0x07, 0x9f, 0x6e, 0x84, 0x2a, 0xf5, 0xdd, 0xeb, 0x6a, 0x27, 0xa8, 0x1f,
	0x72, 0xbe, 0x2e, 0x7e, 0x78, 0x08, 0x7a, 0x0b, 0x26, 0xd5, 0x8c, 0xf3, 0xb5, 0x5c, 0xd2, 0x8f,
	0x39, 0x95, 0xca, 0x7b, 0x0b, 0xfd, 0xa3, 0x7f, 0x52, 0xf0, 0x54, 0x52, 0x34, 0x87, 0x96, 0x1e,
	0x93, 0x03, 0x6b, 0x78, 0x79, 0x7c, 0x35, 0xc0, 0xb8, 0x06, 0x29, 0xd6, 0xa7, 0x61, 0xad, 0xb3,
	0x2c, 0x86, 0xbd, 0xfb, 0xe0, 0xbc, 0xa0, 0x46, 0xb9, 0x3c, 0x0c, 0xdd, 0x86, 0xb6, 0x5e, 0x5a,
	0xb1, 0x68, 0x60, 0x0d, 0xad, 0x71, 0x77, 0x69, 0xeb, 0xcf, 0x97, 0x91, 0xf7, 0x06, 0xfa, 0xcf,
	0x32, 0x4a, 0x14, 0x3d, 0xde, 0x3d, 0x83, 0x2b, 0x7a, 0xd9, 0x6c, 0x6d, 0x0e, 0x61, 0x66, 0xb5,
	0xf0, 0x6b, 0x11, 0x5d, 0x80, 0xf0, 0x43, 0xe8, 0xcf, 0x69, 0x42, 0x4f, 0x85, 0xcf, 0xbb, 0x5f,
	0xf0, 0xbb, 0x05, 0xce, 0xa2, 0x08, 0xff, 0x15, 0xcd, 0xb6, 0xec, 0x3d, 0x45, 0x3f, 0x2c, 0xe8,
	0xfe, 0x77, 0x1e, 0x4d, 0x6b, 0x41, 0x9c, 0x4d, 0xcf, 0x7d, 0xd2, 0x74, 0xac, 0x08, 0xd8, 0x73,
	0xbe, 0xfd, 0xf9, 0xfb, 0xf3, 0x52, 0x07, 0xd9, 0xa6, 0x1e, 0x12, 0x7d, 0xb7, 0xa0, 0x5d, 0x66,
	0x85, 0x26, 0xb5, 0x34, 0x4f, 0x93, 0x75, 0x1b, 0x9a, 0xe8, 0x0d, 0x0c, 0x00, 0x42, 0xbd, 0x02,
	0xc0, 0xff, 0x52, 0xfa, 0xf6, 0x55, 0xfb, 0x02, 0x87, 0x2e, 0xa0, 0x7a, 0x37, 0xac, 0x94, 0xa7,
	0x31, 0xd0, 0x0d, 0x03, 0xe4, 0x78, 0xa5, 0x23, 0x4f, 0x4d, 0xca, 0xe8, 0x97, 0x05, 0x70, 0xe8,
	0x4f, 0x4d, 0x98, 0x4a, 0xe1, 0x1a, 0xc3, 0x8c, 0x0c, 0xcc, 0xbd, 0xe0, 0xe6, 0xb1, 0x3b, 0x78,
	0x6f, 0x51, 0xc9, 0xb6, 0x05, 0x38, 0x34, 0xb0, 0x26, 0x5a, 0xa5, 0xb2, 0xee, 0x2d, 0x5c, 0x3c,
	0x21, 0x78, 0xff, 0x84, 0xe0, 0xe7, 0xfa, 0x09, 0xd9, 0x07, 0xf4, 0xa0, 0x12, 0xd0, 0xac, 0xf3,
	0xce, 0x36, 0xaa, 0x61, 0x68, 0x9b, 0x99, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x0e,
	0x54, 0x38, 0xd8, 0x04, 0x00, 0x00,
}
