package logic

import (
	"context"
	"gorm.io/gorm"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoLikeCountLogic {
	return &UpdateVideoLikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoLikeCountLogic) UpdateVideoLikeCount(in *video.UpdateVideoLikeCountRequest) (*video.Empty, error) {

	//更新操作

	db := l.svcCtx.DBList.DB
	//err := db.Table("videos").
	//Update("favorite_count", gorm.Expr("favorite + ?","xx").
	//db.Model(&User{}).Select("age", "role").
	//Updates([]map[string]interface{}{{"age": gorm.Expr("age + ?", 1), "role": "admin"}, {"age": gorm.Expr("age - ?", 2), "role": "user"}})
	maps := make([]map[string]interface{}, 30)
	for index, _ := range in.VideoID {
		maps = append(maps, map[string]interface{}{
			"favorite_count": gorm.Expr("favorite_count + ?", in.LikeCount[index]),
			"id":             in.VideoID[index],
		})
	}
	//进行批处理
	err := db.Table("videos").Select("favorite_count").
		Updates(maps).Error
	if err != nil {
		return nil, rpcErr.CacheBaseError
	}
	//if err != nil {
	//	return nil, rpcErr.CacheBaseError
	//}
	return &video.Empty{}, nil
}
