package app

import (
	"github.com/izumin5210-sandbox/grapi-playground/app/server"
	"github.com/izumin5210-sandbox/grapi-playground/app/server/internal"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewLibraryServiceServer(),
			internal.NewBookServiceServer(),
		),
	)
	return s.Serve()
}
