package video

import (
	"context"
	"tiktok-plus/common/error/apiErr"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoRequest) (resp *types.PublishVideoResponse, err error) {

	//处理视频文件的妥协
	//在handler 层面处理

	return &types.PublishVideoResponse{
		BaseResponse: types.BaseResponse(apiErr.Success),
	}, nil
}
