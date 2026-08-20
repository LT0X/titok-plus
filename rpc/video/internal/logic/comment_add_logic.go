package logic

import (
	"context"
	"douyin/rpc/video/internal/cache"

	"douyin/rpc/video/internal/model"
	"gorm.io/gorm"

	"douyin/rpc/video/internal/svc"
	"douyin/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentAddLogic {
	return &CommentAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentAddLogic) CommentAdd(in *video.CommentAddRequest) (*video.CommentAddResponse, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.DBList.Mysql.Transaction(func(tx *gorm.DB) error {
		// 先查询videoID是否存在
		com := model.TransformComment(in.Comment)
		video := model.Video{ID: com.VideoID}

		err := tx.First(&video).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.Comment{}).Create(&com).Error
		if err != nil {
			return err
		}
		err = tx.Model(&video).Update("comment_count", video.CommentCount+1).Error
		if err != nil {
			return err
		}
		return cache.CommentAdd(com)
	})
	if err != nil {
		return nil, err
	}

	return &video.CommentAddResponse{
		Reply: 1,
	}, nil
}
