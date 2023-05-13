package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncUserTotalFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncUserTotalFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncUserTotalFavoriteLogic {
	return &SyncUserTotalFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncUserTotalFavoriteLogic) SyncUserTotalFavorite(in *user.SyncUserTotalFavoriteRequest) (*user.Empty, error) {

	DB := l.svcCtx.DB
	for i, v := range in.UserInfoIds {
		err := DB.Table("user_infos").Select("total_favorited").
			Where("id = ?", v).Update("total_favorited", in.UserTotalCounts[i]).Error
		if err != nil {
			return nil, rpcErr.DataBaseError
		}
	}
	return &user.Empty{}, nil
}
