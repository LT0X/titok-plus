package handler

import (
	"context"
	"github.com/hibiken/asynq"
	"strconv"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/rpc/user/user"
)

func (l *AsynqServer) SyncUserLikeINfoHandler(ctx context.Context, t *asynq.Task) error {

	//开始更新用户喜爱列表
	userIds := l.svcCtx.Redis.SMembers(l.ctx, utils.GetUserUpdateListKey()).Val()

	for _, v := range userIds {

		vid, _ := strconv.ParseInt(v, 10, 64)
		likeKey := utils.GetUserLikeSetKey(vid)
		cancelKey := utils.GetUserCancelSetKey(vid)

		likes := l.svcCtx.Redis.SMembers(l.ctx, likeKey).Val()
		cancels := l.svcCtx.Redis.SMembers(l.ctx, cancelKey).Val()
		_, err := l.svcCtx.UserRpc.SyncUserFavorite(l.ctx, &user.SyncUserFavoriteRequest{
			LikeIds:   likes,
			CancelIds: cancels,
		})
		if err != nil {
			l.Logger.Errorf("mq SyncUserLikeINfo failed: %v", err)
		}
		l.svcCtx.Redis.Del(l.ctx, likeKey, cancelKey)

	}

	l.svcCtx.Redis.Del(l.ctx, utils.GetUserUpdateListKey())
	return nil
}
