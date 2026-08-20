package logic

import (
	"context"
	"douyin/rpc/user/internal/model"
	"fmt"
	"strings"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var url string
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Where("id = ?", in.UserID).
		Select("avatar").First(&url).Error

	err = l.svcCtx.DBList.Mysql.Model(&model.User{}).
		Where("id = ?", in.UserID).Updates(map[string]interface{}{
		"username":  in.Username,
		"avatar":    in.Avatar,
		"signature": in.Signature,
	}).Error

	if err != nil {
		return nil, err
	}
	fmt.Printf("url 为 %v\n", url)
	index := strings.Index(url, "static/")
	start := index + len("static/")
	url = url[start:]
	return &user.UpdateUserInfoResponse{
		AvatarUrl: url,
	}, nil
}
