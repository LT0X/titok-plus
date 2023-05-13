package logic

import (
	"context"

	"tiktok-plus/service/rpc/user/internal/svc"
	"tiktok-plus/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExitUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExitUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExitUserLogic {
	return &IsExitUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExitUserLogic) IsExitUser(in *user.IsExitUserRequest) (*user.IsExitUserResponse, error) {

	//准备结果
	type exit struct {
		Id int64 `json:"id"`
	}
	judge := &exit{}

	l.svcCtx.DB.Table("user_logins").Select("id").
		Where("username = ?", in.UserName).First(judge)
	var res bool
	if judge.Id > 0 {
		res = true
	} else {
		res = false
	}

	return &user.IsExitUserResponse{
		IsExit: res,
	}, nil

}
