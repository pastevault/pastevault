package main

import (
	"context"

	"pastevault.com/internal/paste"
	pb "pastevault.com/internal/proto"
)

func (s *server) GetPasteById(_ context.Context, in *pb.Id) (*pb.Paste, error) {
	return paste.GetPasteById(in.Id)
}

func (s *server) CreatePaste(_ context.Context, in *pb.PasteRequest) (*pb.Paste, error) {
	return paste.CreatePaste(in)
}
