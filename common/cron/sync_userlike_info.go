package cron

import (
	"github.com/hibiken/asynq"
)

const TypeSyncUserLikeInfoCache = "sync:video.Info"

func NewSyncUserLikeInfoTask() (*asynq.Task, error) {
	return asynq.NewTask(TypeSyncUserLikeInfoCache, nil), nil
}
