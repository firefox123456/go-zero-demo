package article

import (
	"context"
	"go-zero-demo/app/article/cmd/rpc/article"

	"go-zero-demo/app/article/cmd/api/internal/svc"
	"go-zero-demo/app/article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章
func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
//	// todo: add your logic here and delete this line
//
//	return
//}
func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	// 这里就是调用rpc下的AddArticle方法
	_, err = l.svcCtx.ArticleRpc.AddArticle(l.ctx, &article.AddArticleReq{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return

	return
}