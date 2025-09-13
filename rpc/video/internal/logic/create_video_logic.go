package logic

import (
	"context"

	"douyin/rpc/video/internal/svc"
	"douyin/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoLogic) CreateVideo(in *video.CreateVideoRequest) (*video.CreateVideoResponse, error) {
	// todo: add your logic here and delete this line
	//video  := model.TransformVideo(in.VideoID)
	//err := l.svcCtx.DBList.Mysql.Transaction(func(tx *gorm.DB) error {
	//	// If value doesn't contain a matching primary key, value is inserted.
	//	err := tx.Create(video).Error
	//	if err != nil {
	//		return err
	//	}
	//	cnt, err := SelectWorkCount(video.AuthorID)
	//	if err != nil {
	//		return err
	//	}
	//	err = tx.Model(&model.User{}).Where("id = ?", video.AuthorID).Update("work_count", cnt+1).Error
	//	if err != nil {
	//		return err
	//	}
	//	return cache.PublishVideo(video.AuthorID, video.ID)
	//})
	//if err != nil {
	//	return 0, err
	//}
	//return video.ID, nil
	return &video.CreateVideoResponse{}, nil
}
