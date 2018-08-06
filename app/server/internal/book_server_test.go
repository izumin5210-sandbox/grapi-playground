package internal

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	internal_pb "github.com/izumin5210-sandbox/grapi-playground/api/internal"
)

func Test_BookServiceServer_ListBooks(t *testing.T) {
	svr := NewBookServiceServer()

	ctx := context.Background()
	req := &internal_pb.ListBooksRequest{}

	resp, err := svr.ListBooks(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_BookServiceServer_GetBook(t *testing.T) {
	svr := NewBookServiceServer()

	ctx := context.Background()
	req := &internal_pb.GetBookRequest{}

	resp, err := svr.GetBook(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_BookServiceServer_CreateBook(t *testing.T) {
	svr := NewBookServiceServer()

	ctx := context.Background()
	req := &internal_pb.CreateBookRequest{}

	resp, err := svr.CreateBook(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_BookServiceServer_UpdateBook(t *testing.T) {
	svr := NewBookServiceServer()

	ctx := context.Background()
	req := &internal_pb.UpdateBookRequest{}

	resp, err := svr.UpdateBook(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_BookServiceServer_DeleteBook(t *testing.T) {
	svr := NewBookServiceServer()

	ctx := context.Background()
	req := &internal_pb.DeleteBookRequest{}

	resp, err := svr.DeleteBook(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}
