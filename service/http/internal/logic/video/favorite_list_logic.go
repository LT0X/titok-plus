package video

import (
	"context"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {

	// todo: add your logic here and delete this line

	return
}
