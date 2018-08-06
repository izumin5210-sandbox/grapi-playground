package admin

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	admin_pb "github.com/izumin5210-sandbox/grapi-playground/api/admin"
	type_pb "github.com/izumin5210-sandbox/grapi-playground/api/type"
)

// NewBookServiceServer creates a new BookServiceServer instance.
func NewBookServiceServer() interface {
	admin_pb.BookServiceServer
	grapiserver.Server
} {
	return &bookServiceServerImpl{}
}

type bookServiceServerImpl struct {
}

func (s *bookServiceServerImpl) ListBooks(ctx context.Context, req *admin_pb.ListBooksRequest) (*admin_pb.ListBooksResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) GetBook(ctx context.Context, req *admin_pb.GetBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) CreateBook(ctx context.Context, req *admin_pb.CreateBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) UpdateBook(ctx context.Context, req *admin_pb.UpdateBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) DeleteBook(ctx context.Context, req *admin_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
