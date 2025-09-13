package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectUserByNameLogic {
	return &SelectUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectUserByNameLogic) SelectUserByName(in *user.SelectUserByNameRequest) (*user.SelectUserByNameResponse, error) {
	// todo: add your logic here and delete this line
	resp := &user.SelectUserByNameResponse{}
	var user model.User
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Where("username = ? ", in.UserName).First(&user).Error
	if err != nil {
		return nil, err
	}
	info := model.TransformUserInfo(&user)
	resp.User = info
	return resp, nil
}
