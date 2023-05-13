package logic

import (
	"context"
	"strconv"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncUserFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncUserFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncUserFavoriteLogic {
	return &SyncUserFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncUserFavoriteLogic) SyncUserFavorite(in *user.SyncUserFavoriteRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line
	db := l.svcCtx.DB

	//删除列表
	err := db.Table("user_favor_videos").
		Where("user_info_id = ? and video_id in (?)", in.UserId, in.CancelIds).
		Delete(nil).Error

	if err != nil {
		return nil, rpcErr.CacheBaseError
	}

	var userLike []string

	var updateLikes []struct {
		UserInfoId int64
		VideoId    int64
	}

	db.Table("user_favor_videos").Select("video_id").
		Where("user_info_id = ?", in.UserId).Find(&userLike)

	m := make(map[string]bool) // 创建一个 map
	for _, v := range in.LikeIds {
		m[v] = true // 将 LikeIds的元素作为 map 的键
	}

	j := 0 // 记录非重复元素的索引
	for _, v := range userLike {
		if m[v] { // 检查 userLike 的元素是否在 map 中
			continue // 如果在，就是相同的数字，跳过
		}
		// 如果不在，就是不同的数字，保留
		userLike[j] = v
		j++
	}
	userLike = userLike[:j] // 截取非重复元素的切片

	for _, v := range userLike {
		vid, _ := strconv.ParseInt(v, 10, 64)
		updateLikes = append(updateLikes, struct {
			UserInfoId int64
			VideoId    int64
		}{UserInfoId: in.UserId,
			VideoId: vid,
		})
	}

	err = db.Table("user_favor_videos").Create(&updateLikes).Error
	if err != nil {
		return nil, rpcErr.CacheBaseError
	}
	return &user.Empty{}, nil
}
