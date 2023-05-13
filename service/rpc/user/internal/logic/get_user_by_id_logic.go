package logic

import (
	"context"
	"gorm.io/gorm"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {

	DB := l.svcCtx.DBList.DB

	userInfo := new(user.UserInfo)

	err := DB.Table("user_infos").
		Where("id =  ?", in.UserId).First(userInfo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, rpcErr.UserNotExit
	} else if err != nil {
		return nil, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	return &user.GetUserByIdResponse{
		UserInfo: userInfo,
	}, nil
}
