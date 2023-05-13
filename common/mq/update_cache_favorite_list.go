package mq

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const TypeUpdateCacheFavoriteList = "cache:user:favorite.list"

type UpdateCacheFavoriteListPayload struct {
	LikeListKey   string
	CancelLikeKey string
	Value         int64
	ActionType    int32
}

func NewUpdateCacheFavoriteList(likeKey string, cancelKey string, value int64, ActionType int32) (*asynq.Task, error) {

	payload, err := json.Marshal(UpdateCacheFavoriteListPayload{
		LikeListKey:   likeKey,
		CancelLikeKey: cancelKey,
		Value:         value,
		ActionType:    ActionType,
	},
	)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeUpdateCacheInfo, payload), nil
}
