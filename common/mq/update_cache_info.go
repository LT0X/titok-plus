package mq

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const TypeUpdateCacheInfo = "cache:video:value"

type UpdateCacheInfoPayload struct {
	Key   string
	Field string
	Add   int64
}

func NewUpdateCacheInfoTask(key string, field string, add int64) (*asynq.Task, error) {
	payload, err := json.Marshal(UpdateCacheInfoPayload{
		Key:   key,
		Field: field,
		Add:   add,
	},
	)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeUpdateCacheInfo, payload), nil
}
