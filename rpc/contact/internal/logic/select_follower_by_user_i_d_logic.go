package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectFollowerByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectFollowerByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectFollowerByUserIDLogic {
	return &SelectFollowerByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectFollowerByUserIDLogic) SelectFollowerByUserID(in *contact.SelectFollowerByUserIDRequest) (*contact.SelectFollowerByUserIDResponse, error) {
	// todo: add your logic here and delete this line

	res := make([]uint64, 0)
	err := l.svcCtx.DBList.Mysql.Model(&model.Follow{}).Select("user_id").
		Where("to_user_id = ?", in.UserID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &contact.SelectFollowerByUserIDResponse{
		FollowerIDs: res,
	}, nil
}
