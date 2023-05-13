package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/model"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {

	//开启事务进行多表插入
	tx := l.svcCtx.DB.Begin()

	userInfo := user.UserInfo{
		Name: in.UserName,
	}

	err := tx.Table("user_infos").Create(&userInfo).Error
	//发生错误，回滚事务
	if err != nil {
		tx.Rollback()
		return &user.CreateUserResponse{
			UserId: 0,
		}, rpcErr.DataBaseError.WithDetails(err.Error())
	}
	login := &model.UserLogin{
		Username:   in.UserName,
		Password:   in.Password,
		UserInfoId: userInfo.Id,
	}
	err = tx.Table("user_logins").Create(login).Error
	//发生错误，回滚事务
	if err != nil {
		tx.Rollback()
		return &user.CreateUserResponse{
			UserId: 0,
		}, rpcErr.DataBaseError.WithDetails(err.Error())
	}
	tx.Commit()
	return &user.CreateUserResponse{
		UserId: userInfo.Id,
	}, nil

}
