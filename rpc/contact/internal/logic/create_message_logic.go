package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"
	"time"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMessageLogic) CreateMessage(in *contact.CreateMessageRequest) (*contact.CreateMessageResponse, error) {
	// todo: add your logic here and delete this line
	msg := model.Message{
		Content:    in.Content,
		FromUserID: in.UserID,
		ToUserID:   in.ToUserID,
		CreateTime: time.Now(),
	}
	err := l.svcCtx.DBList.Mysql.Model(&model.Message{}).Create(&msg).Error
	if err != nil {
		return nil, err
	}
	return &contact.CreateMessageResponse{
		Reply: 1,
	}, nil
}
