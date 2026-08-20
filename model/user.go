package model

import (
	"douyin/rpc/user/user"
)

// 这些count是不是可以考虑去除 要的时候再去表里计算TODO
// 这里的头像和背景存储文件名而不是URL 方便更换OSS域名
type User struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
	//Account         string `gorm:"uniqueIndex;type:varchar(255);not null" json:"account"`
	Username        string `gorm:"uniqueIndex;type:varchar(63);not null" json:"username"`
	Password        string `gorm:"type:varchar(255);not null" json:"password"`
	Avatar          string `gorm:"type:varchar(255);not null" json:"avatar"`
	BackgroundImage string `gorm:"type:varchar(255);not null" json:"background_image"`
	Signature       string `gorm:"type:varchar(255);not null" json:"signature"`
	FollowCount     int64  `gorm:"default:0;not null" json:"follow_count"`
	FollowerCount   int64  `gorm:"default:0;not null" json:"follower_count"`
	TotalFavorited  int64  `gorm:"default:0;not null" json:"total_favorited"`
	FavoriteCount   int64  `gorm:"default:0;not null" json:"favorite_count"`
	WorkCount       int64  `gorm:"default:0;not null" json:"work_count"`
}

func TransformUserInfo(user1 *User) *user.UserInfo {
	return &user.UserInfo{
		ID:              user1.ID,
		Username:        user1.Username,
		Password:        user1.Password,
		Avatar:          user1.Avatar,
		BackgroundImage: user1.BackgroundImage,
		Signature:       user1.Signature,
		FollowCount:     user1.FollowCount,
		TotalFavorited:  user1.TotalFavorited,
		FavoriteCount:   user1.FavoriteCount,
		WorkCount:       user1.WorkCount,
	}
}

func TransformUser(user1 *user.UserInfo) *User {
	return &User{
		ID:              user1.ID,
		Username:        user1.Username,
		Password:        user1.Password,
		Avatar:          user1.Avatar,
		BackgroundImage: user1.BackgroundImage,
		Signature:       user1.Signature,
		FollowCount:     user1.FollowCount,
		TotalFavorited:  user1.TotalFavorited,
		FavoriteCount:   user1.FavoriteCount,
		WorkCount:       user1.WorkCount,
	}
}
