package logic

import (
	"context"
	"douyin/rpc/user/internal/model"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectWorkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectWorkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectWorkCountLogic {
	return &SelectWorkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectWorkCountLogic) SelectWorkCount(in *user.SelectWorkCountRequest) (*user.SelectWorkCountResponse, error) {
	// todo: add your logic here and delete this line
	var cnt int64
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Select("work_count").Where("id = ? ", in.UserID).First(&cnt).Error
	if err != nil {
		return nil, err
	}

	return &user.SelectWorkCountResponse{
		Count: cnt,
	}, nil
}
