package model

// UserLogin 用户登录信息，
type UserLogin struct {
	Id         int64  `gorm:"id"`
	UserInfoId int64  `gorm:"user_info_id"`
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
}
