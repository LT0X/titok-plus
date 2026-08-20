package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectFavoriteVideoByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectFavoriteVideoByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectFavoriteVideoByUserIDLogic {
	return &SelectFavoriteVideoByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectFavoriteVideoByUserIDLogic) SelectFavoriteVideoByUserID(in *user.SelectFavoriteVideoByUserIDRequest) (*user.SelectFavoriteVideoByUserIDResponse, error) {
	// todo: add your logic here and delete this line
	res := make([]uint64, 0)
	err := l.svcCtx.DBList.Mysql.Model(&model.Favorite{}).Select("video_id").Where("user_id = ?", in.UserID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &user.SelectFavoriteVideoByUserIDResponse{
		UserIDs: res,
	}, nil
}
