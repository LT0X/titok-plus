package service

import (
	"bytes"
	"context"
	"douyin/database"
	"douyin/model"
	"douyin/package/cache"
	"douyin/package/constant"
	"douyin/package/util"
	"douyin/response"
	user1 "douyin/rpc/user/user"
	"errors"
	"github.com/gofrs/uuid"
	"os"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	// 密码，最长32个字符
	Password string `query:"password"`
	// 注册用户名，最长32个字符
	Username string `query:"username"`

	// 用户鉴权token
	Token string `query:"token"`
	// 用户id 注意上面token会带一个userID
	UserID uint64 `query:"user_id"`
}

func (service *UserService) RegisterService() (*response.UserRegisterOrLogin, error) {
	// 判断用户名是否合法
	if len(service.Username) <= 0 || len(service.Username) > 32 {
		return nil, errors.New(constant.BadParaRequest)
	}
	if len(service.Password) < 6 || len(service.Password) > 32 {
		return nil, errors.New(constant.SecretFormatError)
	}
	// TODO 复杂度判断 可以使用正则 记得去除常数
	if service.Password == constant.EasySecret {
		return nil, errors.New(constant.SecretFormatEasy)
	}
	//先判断用户存不存在 有唯一索引 其实可以不判断
	_, err := database.SelectUserByName(service.Username)
	//_, err := database.SelectUserByAccount(service.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error(constant.DatabaseError, zap.Error(err))
		return nil, err
	}
	if err != gorm.ErrRecordNotFound {
		return nil, errors.New(constant.UserDepulicate)
	}
	// 对密码进行加密并存储
	encryptedPassword := util.BcryptHash(service.Password)
	user := &model.User{
		Username:        service.Username,
		Password:        encryptedPassword,
		Avatar:          util.GenerateAvatar(),
		BackgroundImage: util.GenerateImage(),
		Signature:       util.GenerateSignatrue(),
	}

	//userID, err := database.CreateUser(user)
	resp, err := database.RPC.UserRpc.CreateUser(context.TODO(), &user1.CreateUserRequest{
		User: model.TransformUserInfo(user),
	})

	if err != nil {
		zap.L().Error(constant.DatabaseError, zap.Error(err))
		return nil, err
	}
	userID := resp.UserID
	// 签发token
	token, err := util.SignToken(userID)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	// 1. 缓存用户的个人信息
	// 2. 缓存关注和粉丝列表  这个刚关注肯定没有
	// 3. 缓存发布视频和喜欢的视频
	// 注册和登录之后是一样的
	go func() {
		// 将用户ID加入到布隆过滤器里  对抗缓存穿透
		cache.UserIDBloomFilter.AddString(strconv.FormatUint(userID, 10))
		err = cache.SetUserInfo(user)
		if err != nil {
			zap.L().Sugar().Error(err)
		}
		err = cache.SetFavoriteSet(userID, []uint64{})
		if err != nil {
			zap.L().Sugar().Error(err)
		}
		// 用0值维护 redis key 的存在
		err = cache.SetFollowUserIDSet(userID, []uint64{})
		if err != nil {
			zap.L().Sugar().Error(err)
		}
	}()
	return &response.UserRegisterOrLogin{
		StatusCode: response.Success,
		StatusMsg:  response.RegisterSuccess,
		Token:      &token,
		UserID:     &userID,
	}, nil
}

func (service *UserService) LoginService() (*response.UserRegisterOrLogin, error) {
	// 使用redis 限制用户一定时间的登录次数
	loginKey := constant.LoginCounterPrefix + service.Username
	logintimes, err := cache.UserRedisClient.Get(loginKey).Result()
	var logintimesInt int
	if err != nil {
		// 说明没有这个键 初始化键的登录次数
		cache.UserRedisClient.Set(loginKey, 0, constant.MaxloginInernal)
	} else {
		logintimesInt, _ = strconv.Atoi(logintimes)
		if logintimesInt >= constant.MaxLoginTime {
			return nil, errors.New(constant.FrequentLogin)
		}
	}
	// 无论登录成功还是失败 这里redis记录的数据都+1
	go cache.UserRedisClient.Set(loginKey, logintimesInt+1, constant.MaxloginInernal)
	//先判断用户存不存在

	//user, err := database.SelectUserByName(service.Username)
	resp, err := database.RPC.UserRpc.SelectUserByName(context.TODO(), &user1.SelectUserByNameRequest{
		UserName: service.Username,
	})
	user := model.TransformUser(resp.User)

	if err != nil {
		zap.L().Error(constant.DatabaseError, zap.Error(err))
		return nil, errors.New(constant.UserNoExist)
	}
	if user.ID == 0 {
		return nil, errors.New(constant.UserNoExist)
	}
	if !util.BcryptCheck(service.Password, user.Password) {
		return nil, errors.New(constant.SecretError)
	}
	// 签发token
	token, err := util.SignToken(user.ID)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	// redis预热 用户要查看个人信息 发布的视频 喜欢的视频
	go func() {
		// 个人的用户信息
		err = cache.SetUserInfo(user)
		if err != nil {
			zap.L().Sugar().Error(err)
		}
		// 喜欢的视频列表

		//favoriteIDs, err := database.SelectFavoriteVideoByUserID(user.ID)
		resp, err := database.RPC.UserRpc.SelectFavoriteVideoByUserID(context.TODO(), &user1.SelectFavoriteVideoByUserIDRequest{
			UserID: user.ID,
		})
		if err != nil {
			zap.L().Error(constant.DatabaseError, zap.Error(err))
		}
		favoriteIDs := resp.UserIDs

		if err != nil {
			zap.L().Error(constant.DatabaseError, zap.Error(err))
		} else {
			err = cache.SetFavoriteSet(user.ID, favoriteIDs)
			if err != nil {
				zap.L().Sugar().Error(err)
			}
		}
		// 关注列表
		followUserIDSet, err := database.SelectFollowingByUserID(user.ID)
		if err != nil {
			zap.L().Error(constant.DatabaseError, zap.Error(err))
			return
		}
		err = cache.SetFollowUserIDSet(user.ID, followUserIDSet)
		if err != nil {
			zap.L().Sugar().Error(err)
		}
	}()
	return &response.UserRegisterOrLogin{
		StatusCode: response.Success,
		StatusMsg:  response.LoginSucess,
		Token:      &token,
		UserID:     &user.ID,
	}, nil
}

func (service *UserService) InfoService(loginUserID uint64) (*response.InfoResponse, error) {
	// 使用布隆过滤器判断用户ID是否存在
	if !cache.UserIDBloomFilter.TestString(strconv.FormatUint(service.UserID, 10)) {
		err := errors.New(constant.BloomFilterRejected)
		zap.L().Sugar().Error(err)
		return nil, err
	}
	// 去redis里查询用户信息 这是热点数据 redis缓存确实快了很多
	user, err := cache.GetUserInfo(service.UserID)
	// 缓存未命中再去查数据库
	if err != nil {
		zap.L().Warn(constant.CacheMiss, zap.Error(err))
		user, err = database.SelectUserByID(service.UserID)
		if err != nil {
			zap.L().Error(constant.DatabaseError, zap.Error(err))
			return nil, err
		}
		// 设置缓存
		go func() {
			err = cache.SetUserInfo(user)
			if err != nil {
				zap.L().Error(constant.SetCacheError, zap.Error(err))
			}
		}()
	}
	// 判断是否是关注用户
	var isFollow bool
	// 用户未登录
	if loginUserID == 0 {
		isFollow = false
	} else if loginUserID == service.UserID { // 自己查自己 当然是关注了的
		isFollow = true
	} else {
		isFollow, err = cache.IsFollow(loginUserID, service.UserID)
		// 缓存未命中 查询数据库
		if err != nil {
			zap.L().Warn(constant.CacheMiss, zap.Error(err))
			isFollow, err = database.IsFollowed(loginUserID, service.UserID)
			if err != nil {
				zap.L().Error(constant.DatabaseError, zap.Error(err))
				return nil, err
			}
			go func() {
				// 关注列表
				followUserIDSet, err := database.SelectFollowingByUserID(loginUserID)
				if err != nil {
					zap.L().Error(constant.DatabaseError, zap.Error(err))
					return
				}
				err = cache.SetFollowUserIDSet(user.ID, followUserIDSet)
				if err != nil {
					zap.L().Sugar().Error(err)
				}
			}()
		}
	}
	return &response.InfoResponse{
		StatusCode: response.Success,
		User:       response.UserInfo(user, isFollow),
	}, nil
}

func (service *UserService) UpdateInfoService(uid uint64, username string, signature string, fileFormat string, buf *bytes.Buffer) (*response.UpdateInfoResponse, error) {

	u1, err := uuid.NewV4()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	fileName := u1.String() + fileFormat

	// 感觉应该使用消息队列 来异步上传到OSS FIXME
	err = util.UploadAvater(buf.Bytes(), fileName)
	if err != nil {
		return nil, err
	}
	fileName = "http://127.0.0.1:8000/static/avater/" + fileName
	resp, err := database.RPC.UserRpc.UpdateUserInfo(context.TODO(), &user1.UpdateUserInfoRequest{
		UserID:    uid,
		Username:  username,
		Avatar:    fileName,
		Signature: signature,
	})
	if err != nil {
		return nil, err
	}
	err = os.Remove("C:\\work\\GoWord\\v1-clip\\DikTok\\douyinVideo\\" + resp.AvatarUrl)
	if err != nil {
		// 如果删除过程中出现错误，返回错误信息
		return nil, err
	}

	user, err := database.SelectUserByID(uid)
	err = cache.SetUserInfo(user)
	if err != nil {
		// 如果删除过程中出现错误，返回错误信息
		return nil, err
	}

	return &response.UpdateInfoResponse{
		StatusCode: response.Success,
		StatusMsg:  "成功",
	}, nil

}

func (service *UserService) SearchUserService(uid int64, username string) (*response.SearchUserResponse, error) {

	resp, err := database.RPC.UserRpc.SearchUser(context.TODO(), &user1.SearchUserRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	userInfo := make([]*response.SearchUser, len(resp.UserInfos))
	for i, user := range resp.UserInfos {
		userInfo[i] = &response.SearchUser{
			Username: user.Username,
			Avatar:   user.Avatar,
			UId:      user.ID,
		}
	}
	return &response.SearchUserResponse{

		UserInfos: userInfo,
	}, err
}
