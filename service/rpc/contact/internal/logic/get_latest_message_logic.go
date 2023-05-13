package logic

import (
	"context"

	"tiktok-plus/service/rpc/contact/contact"
	"tiktok-plus/service/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLatestMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestMessageLogic {
	return &GetLatestMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLatestMessageLogic) GetLatestMessage(in *contact.GetLatestMessageRequest) (*contact.GetLatestMessageResponse, error) {

	DB := l.svcCtx.DBList.DB
	DB.Table("")

	return &contact.GetLatestMessageResponse{}, nil
}
