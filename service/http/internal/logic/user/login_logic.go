package user

import (
	"context"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/rpc/user/user"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	logx.WithContext(l.ctx).Infof("登录: %v", req)

	//检查格式
	if msg, ok := CheckUserFormat(req.Username, req.Password); !ok {
		return nil, apiErr.InvalidParameter.WithDetails(msg)
	}

	res, err := l.svcCtx.UserRpc.GetAccountByName(l.ctx, &user.GetAccountByNameRequest{
		UserName: req.Username,
	})
	if err != nil {
		return nil, apiErr.UserNotExit
	}

	//检查密码
	ok := utils.EqualsPassword(req.Password, res.Account.Password)
	if !ok {
		return nil, apiErr.InvalidAccount
	}

	//颁发token
	token, err := utils.CreateToken(res.Account.UserInfoId,
		l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)

	if err != nil {
		logx.WithContext(l.ctx).Infof("颁发token失败: %v, err:%v", req, err)
		return nil, apiErr.CreateTokenFailed
	}

	return &types.LoginResponse{
		BaseResponse: types.BaseResponse(apiErr.Success),
		Token:        token,
		UserId:       res.Account.UserInfoId,
	}, nil

}
