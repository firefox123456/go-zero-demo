package logic

import (
	"context"

	"go-zero-demo/RPCDemo/internal/svc"
	"go-zero-demo/RPCDemo/rPCDemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rPCDemo.Request) (*rPCDemo.Response, error) {
	// todo: add your logic here and delete this line

	return &rPCDemo.Response{
		Pong: in.Ping,
	}, nil
}

