package logic

import (
	"context"
	"douyin/rpc/user/internal/model"
	"math/rand"
	"time"

	"douyin/rpc/user/internal/svc"
	"douyin/rpc/user/user"

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
	// todo: add your logic here and delete this line
	resp := &user.CreateUserResponse{}
	user1 := model.TransformUser(in.User)
	//user1.Account  =  user1.Username
	//user1.Username = generateRandomUsername(8)
	err := l.svcCtx.DBList.Mysql.Model(&model.User{}).Create(user1).Error
	if err != nil {
		return nil, err
	}
	resp.UserID = user1.ID
	return resp, nil
}

func generateRandomUsername(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

//func transformUser(info *user.UserInfo) *model.User {
//	return &model.User{
//		ID:              info.ID,
//		Username:        info.Username,
//		Password:        info.Password,
//		Avatar:          info.Avatar,
//		BackgroundImage: info.BackgroundImage,
//		Signature:       info.Signature,
//		FollowCount:     info.FollowCount,
//		TotalFavorited:  info.TotalFavorited,
//		FavoriteCount:   info.FavoriteCount,
//		WorkCount:       info.WorkCount,
//	}
//}
