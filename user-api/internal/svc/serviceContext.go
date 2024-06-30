package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-api/internal/model"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.Mysql)),
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}
