package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/izumin5210-sandbox/grapi-playground/api"
	type_pb "github.com/izumin5210-sandbox/grapi-playground/api/type"
)

// NewLibraryServiceServer creates a new LibraryServiceServer instance.
func NewLibraryServiceServer() interface {
	api_pb.LibraryServiceServer
	grapiserver.Server
} {
	return &libraryServiceServerImpl{}
}

type libraryServiceServerImpl struct {
}

func (s *libraryServiceServerImpl) ListBooks(ctx context.Context, req *api_pb.ListBooksRequest) (*api_pb.ListBooksResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) GetBook(ctx context.Context, req *api_pb.GetBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) CreateBook(ctx context.Context, req *api_pb.CreateBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) UpdateBook(ctx context.Context, req *api_pb.UpdateBookRequest) (*type_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) DeleteBook(ctx context.Context, req *api_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
