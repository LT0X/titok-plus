package video

import (
	"context"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoListLogic) PublishVideoList(req *types.PublishVideoListRequest) (resp *types.PublishVideoListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
