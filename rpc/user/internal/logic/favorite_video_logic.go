package logic

import (
	"context"
	"douyin/rpc/user/internal/cache"
	"douyin/rpc/user/internal/model"
	"errors"

	"douyin/package/constant"
	"gorm.io/gorm"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoLogic {
	return &FavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteVideoLogic) FavoriteVideo(in *user.FavoriteVideoRequest) (*user.FavoriteVideoResponse, error) {
	// todo: add your logic here and delete this line
	favorite := model.Favorite{
		UserID:  in.UserID,
		VideoID: in.VideoID,
	}
	// 一般输入流程 在是事务里 使用tx而不是db 返回任何错误都会回滚事务
	err := l.svcCtx.DBList.Mysql.Transaction(func(tx *gorm.DB) error {
		// 先看有没有点赞过
		var isFavorite int64
		err := tx.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ?", in.UserID, in.VideoID).Count(&isFavorite).Error
		if err != nil {
			return err
		}
		if in.Cnt == 1 && isFavorite == 0 {
			err = tx.Model(&model.Favorite{}).Create(&favorite).Error
		} else if in.Cnt == -1 && isFavorite == 1 {
			err = tx.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ? ", in.UserID, in.VideoID).Delete(&favorite).Error
		} else {
			err = errors.New(constant.BadParaRequest)
		}
		if err != nil {
			return err
		}
		// 视频表增加该视频的点赞
		video := model.Video{ID: in.VideoID} // model里需要有主键
		// 否则favorite.go:33 WHERE conditions required
		err = tx.Model(&model.Video{}).Select("favorite_count, author_id").Where("id = ?", in.VideoID).First(&video).Error
		if err != nil {
			return err
		}
		err = tx.Model(&video).Update("favorite_count", video.FavoriteCount+in.Cnt).Error
		if err != nil {
			return err
		}
		// 增加视频作者被点赞数
		author := model.User{ID: video.AuthorID} // 同理上面 可以看Model函数的说明
		err = tx.Model(&author).Select("total_favorited").First(&author).Error
		if err != nil {
			return err
		}
		err = tx.Model(&author).Update("total_favorited", author.TotalFavorited+in.Cnt).Error
		if err != nil {
			return err
		}
		// 增加用户的点赞数
		user := model.User{ID: in.UserID} // 同理上面
		err = tx.Model(&user).Select("favorite_count").First(&user).Error
		if err != nil {
			return err
		}
		err = tx.Model(&user).Update("favorite_count", user.FavoriteCount+in.Cnt).Error
		if err != nil {
			return err
		}

		return cache.FavoriteAction(in.UserID, author.ID, in.VideoID, in.Cnt)
	})

	if err != nil {
		return nil, err
	}

	return &user.FavoriteVideoResponse{
		Reply: 1,
	}, nil
}
