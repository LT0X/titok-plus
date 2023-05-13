package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"tiktok-plus/common/mq"
	"tiktok-plus/common/utils"
)

func (l *AsynqServer) UpdateCacheInfoHandler(ctx context.Context, t *asynq.Task) error {
	var p mq.UpdateCacheInfoPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		l.Logger.Errorf("json.Unmarshal failed: %v", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := l.svcCtx.Redis.HIncrBy(ctx, p.Key, p.Field, p.Add).Err()
	if err != nil {
		l.Logger.Errorf("redis.HIncrBy failed: %v", err)
		return err
	}
	
	err = l.svcCtx.Redis.SAdd(ctx, utils.GetFavoriteSetKey(), p.Field).Err()
	if err != nil {
		l.Logger.Errorf("redis.SAdd failed: %v", err)
		return err
	}

	return nil
}
