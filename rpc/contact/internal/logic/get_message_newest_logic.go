package logic

import (
	"context"
	model "douyin/rpc/contact/internal/modle"

	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageNewestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageNewestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageNewestLogic {
	return &GetMessageNewestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageNewestLogic) GetMessageNewest(in *contact.GetMessageNewestRequest) (*contact.GetMessageNewestResponse, error) {
	// todo: add your logic here and delete this line
	msg := model.Message{}
	// 这里用 union 避免or 不走索引的情况 or两侧必须都走索引 括号也没用
	err := l.svcCtx.DBList.Mysql.Raw("? UNION ? ORDER BY create_time DESC LIMIT 1",
		l.svcCtx.DBList.Mysql.Model(&model.Message{}).Where("from_user_id = ? AND to_user_id = ?", in.UserID, in.ToUserID),
		l.svcCtx.DBList.Mysql.Model(&model.Message{}).Where("from_user_id = ? AND to_user_id = ?", in.ToUserID, in.UserID)).Scan(&msg).Error
	if err != nil {
		return nil, err
	}

	return &contact.GetMessageNewestResponse{
		Message: model.TransformRPCMessage(msg),
	}, nil
}
