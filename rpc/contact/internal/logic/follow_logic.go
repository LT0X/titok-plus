package logic

import (
	"context"
	"douyin/rpc/contact/internal/cache"
	model "douyin/rpc/contact/internal/modle"

	"gorm.io/gorm"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}

}

func (l *FollowLogic) Follow(in *contact.FollowRequest) (*contact.FollowResponse, error) {
	// todo: add your logic here and delete this line
	follow := model.Follow{
		UserID:   in.UserID,
		ToUserID: in.ToUserID,
	}
	err := l.svcCtx.DBList.Mysql.Transaction(func(tx *gorm.DB) error {
		// 关注表里更新
		var err error
		var ff model.Follow
		if in.Cnt == 1 {
			//  这里设置联合唯一索引 应该不需要检查了
			err = tx.Model(&model.Follow{}).Create(&follow).Error
		} else if in.Cnt == -1 {
			err = tx.Model(&model.Follow{}).Where("user_id = ? AND to_user_id = ?", in.UserID, in.ToUserID).Delete(&ff).Error
		}
		if err != nil {
			return err
		}
		// 然后更新用户的关注数
		user := model.User{ID: in.UserID}
		// Model会检查主键
		err = tx.Model(&user).First(&user).Error
		if err != nil {
			return err
		}
		err = tx.Model(&user).Update("follow_count", user.FollowCount+in.Cnt).Error
		if err != nil {
			return err
		}
		// 被关注用户的被关注数+1
		toUser := model.User{ID: in.ToUserID}
		err = tx.Model(&toUser).First(&toUser).Error
		if err != nil {
			return err
		}
		err = tx.Model(&toUser).Update("follower_count", toUser.FollowerCount+in.Cnt).Error
		if err != nil {
			return err
		}
		return cache.FollowAction(in.UserID, in.ToUserID, in.Cnt)
	})

	if err != nil {
		return nil, err
	}

	return &contact.FollowResponse{
		Reply: 1,
	}, nil
}
