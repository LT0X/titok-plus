package handler

import (
	"context"
	"strconv"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/rpc/user/user"
	"tiktok-plus/service/rpc/video/video"

	"github.com/hibiken/asynq"
)

func (l *AsynqServer) SyncVideoInfoHandler(ctx context.Context, t *asynq.Task) error {

	//缓存同步到数据库并删除缓存
	result, err := l.svcCtx.Redis.SMembers(ctx, utils.GetFavoriteSetKey()).Result()
	if err != nil {
		l.Logger.Errorf("mq redis.SMembers failed: %v", err)
		return err
	}

	//对视频点赞信息进行更新
	vids := make([]string, 30)
	likeCounts := make([]string, 30)
	for _, v := range result {

		count, err := l.svcCtx.Redis.HGet(l.ctx, utils.GetFavoriteCountKey(), v).Result()
		if err != nil {
			l.Logger.Errorf("mq redis.HGet failed: %v", err)
			return err
		}
		vids = append(vids, v)
		likeCounts = append(likeCounts, count)
	}
	_, err = l.svcCtx.VideoRpc.UpdateVideoLikeCount(l.ctx, &video.UpdateVideoLikeCountRequest{
		LikeCount: likeCounts,
		VideoID:   vids,
	})

	//把string 数组转为int64
	videoIds := make([]int64, len(vids))
	for i, v := range vids {
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		videoIds[i] = int64(num)
	}

	res, err := l.svcCtx.VideoRpc.GetUserInfoIdsByVideoIds(l.ctx, &video.GetUserInfoIdsByVideoIdsRequest{
		VideoIds: videoIds,
	})
	if err != nil {
		l.Logger.Errorf("mq.UpdateVideoLikeCount failed: %v", err)
		return err
	}
	_, err = l.svcCtx.UserRpc.SyncUserTotalFavorite(l.ctx, &user.SyncUserTotalFavoriteRequest{
		UserInfoIds:     res.UserInfoIds,
		UserTotalCounts: res.TotalCount,
	})
	if err != nil {
		return err
	}

	//删除点赞视频和视频播放数量变化缓存
	err = l.svcCtx.Redis.
		Del(l.ctx, utils.GetFavoriteSetKey(), utils.GetFavoriteCountKey()).Err()
	if err != nil {
		l.Logger.Errorf("mq.UpdateVideoLikeCount failed: %v", err)
		return err
	}
	return nil
}
