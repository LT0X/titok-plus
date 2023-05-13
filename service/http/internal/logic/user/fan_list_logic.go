package user

import (
	"context"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FanListLogic {
	return &FanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FanListLogic) FanList(req *types.FanListRequest) (resp *types.FanListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
