package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowedLogic {
	return &IsFollowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowedLogic) IsFollowed(in *contact.IsFollowedRequest) (*contact.IsFollowedResponse, error) {
	// todo: add your logic here and delete this line
	var cnt int64
	err := l.svcCtx.DBList.Mysql.Model(&model.Follow{}).
		Where("user_id= ? AND to_user_id = ? ", in.UserID, in.ID).Count(&cnt).Error
	if err != nil {
		return nil, err
	} else if cnt == 0 {
		return nil, nil
	}
	return &contact.IsFollowedResponse{
		Bool: true,
	}, nil
}
