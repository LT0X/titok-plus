package logic

import (
	"context"
	"gorm.io/gorm"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByNameLogic {
	return &GetAccountByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByNameLogic) GetAccountByName(in *user.GetAccountByNameRequest) (*user.GetAccountByNameResponse, error) {
	DB := l.svcCtx.DBList.DB
	account := new(user.Account)

	err := DB.Table("user_logins").
		Where("username = ?", in.UserName).
		First(account).Error

	if err == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		return nil, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	return &user.GetAccountByNameResponse{
		Account: account,
	}, nil

}
