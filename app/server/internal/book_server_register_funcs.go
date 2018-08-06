package internal

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	internal_pb "github.com/izumin5210-sandbox/grapi-playground/api/internal"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *bookServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	internal_pb.RegisterBookServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *bookServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return internal_pb.RegisterBookServiceHandler(ctx, mux, conn)
}
