package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"strings"
	"tiktok-plus/common/mq"
	"tiktok-plus/common/utils"
)

func (l *AsynqServer) UpdateCacheFavoriteListHandler(ctx context.Context, t *asynq.Task) error {
	var p mq.UpdateCacheFavoriteListPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		l.Logger.Errorf("json.Unmarshal failed: %v", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	//加入用户更新喜爱视频列表,利用set自动去重
	index := strings.Index(p.LikeListKey, "t") // 找到第一个 t 的位置
	userId := p.LikeListKey[index+1:]          // 从该位置的下一个字节开始截取

	err := l.svcCtx.Redis.SAdd(l.ctx, utils.GetUserUpdateListKey(), userId).Err()
	if err != nil {
		l.Logger.Errorf("mq :Sadd Error: %v", err)
	}

	if p.ActionType == 1 {
		//点赞操作

		l.svcCtx.Redis.SAdd(l.ctx, p.LikeListKey, p.Value)

		//删除取消点赞列表的缓存
		l.svcCtx.Redis.SRem(l.ctx, p.CancelLikeKey, p.Value)

	} else if p.ActionType == 2 {
		//取消点赞操作
		l.svcCtx.Redis.SAdd(l.ctx, p.CancelLikeKey, p.Value)

		//删除点赞列表的缓存
		l.svcCtx.Redis.SRem(l.ctx, p.LikeListKey, p.Value)
	}

	return nil
}
