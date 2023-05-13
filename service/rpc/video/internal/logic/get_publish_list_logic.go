package logic

import (
	"context"

	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishListLogic {
	return &GetPublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublishListLogic) GetPublishList(in *video.GetPublishListRequest) (*video.GetPublishListResponse, error) {
	// todo: add your logic here and delete this line
	
	return &video.GetPublishListResponse{}, nil
}
