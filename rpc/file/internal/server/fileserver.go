// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package server

import (
	"context"

	"fileup/rpc/file/file"
	"fileup/rpc/file/internal/logic"
	"fileup/rpc/file/internal/svc"
)

type FileServer struct {
	svcCtx *svc.ServiceContext
	file.UnimplementedFileServer
}

func NewFileServer(svcCtx *svc.ServiceContext) *FileServer {
	return &FileServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServer) Ping(ctx context.Context, in *file.Request) (*file.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
