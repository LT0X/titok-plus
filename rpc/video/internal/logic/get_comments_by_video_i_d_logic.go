package logic

import (
	"context"
	"douyin/rpc/video/internal/model"

	"douyin/rpc/video/internal/svc"
	"douyin/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByVideoIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByVideoIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByVideoIDLogic {
	return &GetCommentsByVideoIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsByVideoIDLogic) GetCommentsByVideoID(in *video.GetCommentsByVideoIDRequest) (*video.GetCommentsByVideoIDResponse, error) {
	// todo: add your logic here and delete this line
	videos := make([]*model.Comment, 0)
	err := l.svcCtx.DBList.Mysql.Model(&model.Comment{}).Where("video_id = ?", in.VideoID).Order("created_time desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return &video.GetCommentsByVideoIDResponse{
		Comments: model.TransformRPCComments(videos),
	}, nil
}
