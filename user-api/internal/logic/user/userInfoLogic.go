package user

import (
	"context"
	"errors"
	"fmt"
	"zero-demo/user-api/internal/model"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {

	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	fmt.Println(user)
	logx.Error("我故意的！！！！")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return &types.UserInfoResp{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
