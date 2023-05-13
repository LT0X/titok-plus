package user

import (
	"context"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/service/rpc/user/user"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	logx.WithContext(l.ctx).Infof("用户信息: %v", req)

	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
		UserId: req.UserId,
	})
	if err == rpcErr.UserNotExit {
		return nil, apiErr.UserNotExit
	} else if err != nil {
		return nil, apiErr.ServerInternal
	}

	return &types.UserInfoResponse{
		BaseResponse: types.BaseResponse(apiErr.Success),
		User: types.UserInfo{
			Id:              res.UserInfo.Id,
			Name:            res.UserInfo.Name,
			FollowCount:     res.UserInfo.FollowCount,
			FollowerCount:   res.UserInfo.FollowerCount,
			Avatar:          res.UserInfo.Avatar,
			BackgroundImage: res.UserInfo.BackgroundImage,
			Signature:       res.UserInfo.Signature,
			TotalFavorited:  res.UserInfo.TotalFavorited,
			WorkCount:       res.UserInfo.WorkCount,
			FavoriteCount:   res.UserInfo.FavoriteCount,
		},
	}, nil
}
