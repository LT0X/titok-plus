package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoIdsByVideoIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoIdsByVideoIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoIdsByVideoIdsLogic {
	return &GetUserInfoIdsByVideoIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoIdsByVideoIdsLogic) GetUserInfoIdsByVideoIds(in *video.GetUserInfoIdsByVideoIdsRequest) (*video.GetUserInfoIdsByVideoIdsResponse, error) {
	// todo: add your logic here and delete this line

	var userInfoIds []int64
	var totalCount []int64
	DB := l.svcCtx.DB
	err := DB.Table("videos").Select("user_Info_id").
		Where("id in (?)", in.VideoIds).Find(userInfoIds).Error
	if err != nil {
		return nil, rpcErr.CacheBaseError
	}
	for _, v := range userInfoIds {
		var count struct {
			TotalCount int64
		}
		DB.Table("videos").Select("count(favorite_count) as total_count").
			Where("user_info_id = ?", v).First(&count)
		totalCount = append(totalCount, count.TotalCount)
	}
	return &video.GetUserInfoIdsByVideoIdsResponse{
		UserInfoIds: userInfoIds,
		TotalCount:  totalCount,
	}, nil

}
