package utils

import "fmt"

// GetPopVideoKey 欢迎视频缓存列表(Hash)
func GetPopVideoKey() string {
	return "video:pop.video"
}

// GetFavoriteKey 用于记录视频用户是否点赞(Hash)
func GetFavoriteKey(vid int64) string {
	return fmt.Sprintf("video:fav:video%v", vid)
}

// GetFavoriteCountKey 记录缓存视频点赞变化(Hash)
func GetFavoriteCountKey() string {
	return "video:fav:video.count"
}

// GetFavoriteSetKey 用于记录将要缓存更新的视频列表(Set)
func GetFavoriteSetKey() string {
	return "video:fav:video.Set"
}

// GetUserLikeSetKey 用于缓存记录用户喜欢列表(Set)
func GetUserLikeSetKey(uid int64) string {
	return fmt.Sprintf("video:user:like.set%v", uid)
}

// GetUserCancelSetKey 用户缓存记录取消点赞的列表(Set)
func GetUserCancelSetKey(uid int64) string {
	return fmt.Sprintf("video:user:cancel.set%v", uid)
}

// GetUserUpdateListKey 用来记录将要更新用户喜爱列表的用户名单
func GetUserUpdateListKey() string {
	return fmt.Sprintf("user:update.list")
}
