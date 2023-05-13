package cron

import (
	"github.com/hibiken/asynq"
)

const TypeSyncVideoInfoCache = "sync:video.Info"

func NewSyncVideoInfoTask() (*asynq.Task, error) {
	return asynq.NewTask(TypeSyncVideoInfoCache, nil), nil
}
