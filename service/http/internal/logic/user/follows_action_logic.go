package user

import (
	"context"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowsActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowsActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowsActionLogic {
	return &FollowsActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowsActionLogic) FollowsAction(req *types.FollowsActionRequest) (resp *types.FollowsActionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
