package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectUserListByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectUserListByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectUserListByIDsLogic {
	return &SelectUserListByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectUserListByIDsLogic) SelectUserListByIDs(in *user.SelectUserListByIDsRequest) (*user.SelectUserListByIDsResponse, error) {
	// todo: add your logic here and delete this line
	var users []model.User
	// (?)  ( ? )会多加一个括号
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Where("id IN (?)  ", in.UserIDs).Find(&users).Error
	if err != nil {
		return nil, err
	}
	transformUsers := model.TransformUsers(users)
	return &user.SelectUserListByIDsResponse{
		Users: transformUsers,
	}, nil
}
