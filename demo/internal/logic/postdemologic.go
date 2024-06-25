package logic

import (
	"context"

	"go-zero-demo/demo/internal/svc"
	"go-zero-demo/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDemoLogic {
	return &PostDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDemoLogic) PostDemo(req *types.PostDemoReq) (resp *types.PostDemoResp, err error) {
	// todo: add your logic here and delete this line
	return &types.PostDemoResp{
		Message: req.Message,
	}, nil
}
