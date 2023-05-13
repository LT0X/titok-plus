package video

import (
	"context"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/rpc/video/video"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionRequest) (resp *types.FavoriteActionResponse, err error) {
	logx.WithContext(l.ctx).Infof("点赞功能 %v", req)
	uid, err := utils.GetUserIDFormToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, apiErr.InvalidToken
	}

	_, err = l.svcCtx.VideoRpc.FavoriteAction(l.ctx, &video.FavoriteActionRequest{
		UserId:     uid,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		logx.WithContext(l.ctx).Infof("FavoriteAction error : %v", err)
		return nil, apiErr.ServerInternal
	}

	return &types.FavoriteActionResponse{
		BaseResponse: types.BaseResponse(apiErr.Success),
	}, nil
}
