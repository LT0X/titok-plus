package contact

import (
	"context"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryMessageLogic {
	return &HistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryMessageLogic) HistoryMessage(req *types.HistoryMessageRequest) (resp *types.HistoryMessageReply, err error) {
	// todo: add your logic here and delete this line

	return
}
