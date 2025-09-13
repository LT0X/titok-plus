package model

import (
	"douyin/rpc/contact/contact"
	"time"
)

type Message struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	Content    string    `gorm:"not null;type:text" json:"content"`
	CreateTime time.Time `gorm:"not null;index" json:"create_time"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID uint64    `gorm:"not null;index:idx_user_touser" json:"from_user_id"`
	ToUserID   uint64    `gorm:"not null;index:idx_user_touser" json:"to_user_id"`
}

func TransformMessages(msg []*contact.Message) []Message {
	res := make([]Message, len(msg))
	for i, v := range msg {
		res[i] = Message{
			ID:         v.ID,
			Content:    v.Content,
			CreateTime: time.UnixMilli(v.CreateTime),
			FromUserID: v.FromUserID,
			ToUserID:   uint64(v.ToUserID),
		}
	}
	return res
}

func TransformMessage(msg *contact.Message) *Message {
	return &Message{
		ID:         msg.ID,
		Content:    msg.Content,
		CreateTime: time.UnixMilli(msg.CreateTime),
		FromUserID: msg.FromUserID,
		ToUserID:   uint64(msg.ToUserID),
	}
}
