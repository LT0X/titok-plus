package model

import (
	"douyin/rpc/video/video"
	"time"
)

type Comment struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	VideoID     uint64    `gorm:"not null;index" json:"video_id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	Content     string    `gorm:"not null;type:varchar(255)" json:"content"`
	CreatedTime time.Time `gorm:"not null" json:"created_time"`
	//IsDeleted   uint8     `gorm:"default:0;not null" json:"is_deleted"`
}

func TransformComment(comment *video.Comment) *Comment {
	return &Comment{
		ID:          comment.ID,
		VideoID:     comment.VideoID,
		UserID:      comment.UserID,
		Content:     comment.Content,
		CreatedTime: time.UnixMilli(comment.CreatedTime),
	}
}

func TransformRPCComment(comment *Comment) *video.Comment {
	return &video.Comment{
		ID:          comment.ID,
		VideoID:     comment.VideoID,
		UserID:      comment.UserID,
		Content:     comment.Content,
		CreatedTime: comment.CreatedTime.UnixMilli(),
	}
}

func TransformComments(comment []*video.Comment) []*Comment {
	res := make([]*Comment, len(comment))
	for i, v := range comment {
		res[i] = TransformComment(v)
	}
	return res
}
