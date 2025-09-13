package logic

import (
	"context"
	"douyin/rpc/video/internal/model"

	"douyin/rpc/video/internal/svc"
	"douyin/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoURLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoURLLogic {
	return &UpdateVideoURLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoURLLogic) UpdateVideoURL(in *video.UpdateVideoURLRequest) (*video.UpdateVideoURLResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.DBList.Mysql.Model(&model.Video{ID: in.VideoID}).
		Updates(&model.Video{PlayURL: in.PlayURL, CoverURL: in.CoverURl}).Error
	if err != nil {
		return nil, err
	}
	return &video.UpdateVideoURLResponse{
		Reply: 1,
	}, nil
}
