package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectUserByIDLogic {
	return &SelectUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectUserByIDLogic) SelectUserByID(in *user.SelectUserByIDRequest) (*user.SelectUserByIDResponse, error) {
	// todo: add your logic here and delete this line
	resp := &user.SelectUserByIDResponse{}
	var user model.User
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Where("id = ? ", in.UserID).First(&user).Error
	if err != nil {
		return nil, err
	}
	resp.User = model.TransformUserInfo(&user)
	return resp, nil
}
