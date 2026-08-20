package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *user.SearchUserRequest) (*user.SearchUserResponse, error) {
	// todo: add your logic here and delete this line

	var users []model.User
	// (?)  ( ? )会多加一个括号
	queryValue := "%" + in.Username + "%"
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).
		Where("username like  (?)  ", queryValue).Find(&users).Error
	if err != nil {
		return nil, err
	}
	infos := model.TransformUsers(users)
	return &user.SearchUserResponse{
		UserInfos: infos,
	}, nil
}
