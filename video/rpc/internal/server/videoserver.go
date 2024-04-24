// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"zero/video/rpc/internal/logic"
	"zero/video/rpc/internal/svc"
	"zero/video/rpc/video"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoServer
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServer) Ping(ctx context.Context, in *video.Request) (*video.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
