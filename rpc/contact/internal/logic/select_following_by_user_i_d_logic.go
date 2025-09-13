package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectFollowingByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectFollowingByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectFollowingByUserIDLogic {
	return &SelectFollowingByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectFollowingByUserIDLogic) SelectFollowingByUserID(in *contact.SelectFollowingByUserIDRequest) (*contact.SelectFollowingByUserIDResponse, error) {
	// todo: add your logic here and delete this line
	res := make([]uint64, 0)
	err := l.svcCtx.DBList.Mysql.Model(&model.Follow{}).Select("to_user_id").Where("user_id = ?", in.UserID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &contact.SelectFollowingByUserIDResponse{
		UserID: res,
	}, nil
}
