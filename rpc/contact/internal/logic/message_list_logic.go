package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"
	"time"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageListLogic) MessageList(in *contact.MessageListRequest) (*contact.MessageListResponse, error) {
	// todo: add your logic here and delete this line
	newMsgTime := time.UnixMilli(in.MsgTime)
	msgs := make([]model.Message, 0)
	// 这里用 union 避免or 不走索引的情况 or两侧必须都走索引 括号也没用
	err := l.svcCtx.DBList.Mysql.Raw("? UNION ? ORDER BY create_time ASC",
		l.svcCtx.DBList.Mysql.Model(&model.Message{}).Where("from_user_id = ? AND to_user_id = ? AND create_time > ?",
			in.UserID, in.ToUserID, newMsgTime),
		l.svcCtx.DBList.Mysql.Model(&model.Message{}).Where("from_user_id = ? AND to_user_id = ? AND create_time > ?",
			in.ToUserID, in.UserID, newMsgTime)).Scan(&msgs).Error
	if err != nil {
		return nil, err
	}

	return &contact.MessageListResponse{
		Massages: model.TransformRPCMessages(msgs),
	}, nil
}
