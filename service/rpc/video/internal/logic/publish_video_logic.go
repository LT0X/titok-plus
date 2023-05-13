package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *video.CreateVideoRequest) (*video.Empty, error) {

	db := l.svcCtx.DB

	videoInfo := &video.VideoInfo{
		UserInfoId:    in.UserInfoId,
		Title:         in.Title,
		CoverUrl:      in.CoverUrl,
		PlayUrl:       in.PlayUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}

	err := db.Table("videos").Create(videoInfo).Error
	if err != nil {
		return &video.Empty{}, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	return &video.Empty{}, nil
}
