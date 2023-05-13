package logic

import (
	"context"
	"tiktok-plus/common/error/rpcErr"

	"tiktok-plus/service/rpc/contact/contact"
	"tiktok-plus/service/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *contact.GetMessageListRequest) (*contact.GetMessageListResponse, error) {

	DB := l.svcCtx.DBList.DB

	messageList := new([]contact.Message)
	err := DB.Table("messages").Order("create_time asc").
		Where("create_time > ? and to_user_id = ? and from_user_id = ? or "+
			"to_user_id = ? and from_user_id = ? and create_time > ? ", in.PreMsgTime, in.ToUserId, in.FromUserId,
			in.FromUserId, in.ToUserId, in.PreMsgTime).Find(messageList).Error

	if err != nil {
		return nil, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	return &contact.GetMessageListResponse{
		Messages: messageList,
	}, nil
}
