package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/contact/contact"
	"tiktok-plus/service/rpc/contact/internal/svc"

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

func (l *CreateMessageLogic) CreateMessage(in *contact.CreateMessageRequest) (*contact.Empty, error) {

	DB := l.svcCtx.DBList.DB

	message := &contact.Message{
		Content:    in.Content,
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
	}

	err := DB.Table("messages").Create(message).Error

	if err != nil {
		return nil, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	return &contact.Empty{}, nil
}
