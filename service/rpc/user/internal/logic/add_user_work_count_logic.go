package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserWorkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserWorkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserWorkCountLogic {
	return &AddUserWorkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserWorkCountLogic) AddUserWorkCount(in *user.AddUserWorkCountRequest) (*user.Empty, error) {

	db := l.svcCtx.DBList.DB

	err := db.Table("user_infos").
		Exec("update user_infos set work_count = work_count + 1 where id = ?", in.UserInfoId).Error
	if err != nil {
		return &user.Empty{}, rpcErr.DataBaseError.WithDetails(err.Error())
	}
	return &user.Empty{}, nil
}
