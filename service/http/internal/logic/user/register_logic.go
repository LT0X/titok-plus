package user

import (
	"context"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/rpc/user/user"
	"unicode"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	logx.WithContext(l.ctx).Infof("注册: %v", req)

	//检查格式
	if msg, ok := CheckUserFormat(req.Username, req.Password); !ok {
		return nil, apiErr.InvalidParameter.WithDetails(msg)
	}

	//开始业务逻辑
	res, err := l.svcCtx.UserRpc.IsExitUser(l.ctx, &user.IsExitUserRequest{
		UserName: req.Username,
	})

	//用户名重复
	if res.IsExit == true {
		return nil, apiErr.UserNameConflict
	}

	//开始密码加密
	password, err := utils.EncryptPassword(req.Password)
	if err != nil {
		logx.WithContext(l.ctx).Infof("密码加密失败: %v,%v", req, err.Error())
		return nil, apiErr.EncryptionFailed
	}
	req.Password = password

	userInfo, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserRequest{
		UserName: req.Username,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Infof("用户注册失败: %v,错误err: %v", req, err.Error())
		return nil, apiErr.ServerInternal
	}

	//开始分发token
	token, err := utils.CreateToken(userInfo.UserId,
		l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)

	if err != nil {
		logx.WithContext(l.ctx).Infof("颁发token失败: %v, err:%v", req, err)
		return nil, apiErr.CreateTokenFailed
	}

	return &types.RegisterResponse{
		UserId:       userInfo.UserId,
		Token:        token,
		BaseResponse: types.BaseResponse(apiErr.Success),
	}, nil

}

func CheckUserFormat(username string, password string) (string, bool) {
	if len(username) > 20 {
		return "用户名不得超过20位", false
	} else {
		if len(password) < 20 {
			if IsChinese(password) {
				return "密码含有中文", false
			}
			return "", true
		}
		return "密码长度不得超过32位", false
	}
}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
