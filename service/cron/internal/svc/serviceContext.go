package svc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok-plus/service/cron/internal/config"
	"tiktok-plus/service/rpc/contact/contactclient"
	"tiktok-plus/service/rpc/user/userclient"
	"tiktok-plus/service/rpc/video/videoclient"

	"time"
)

type ServiceContext struct {
	Config     config.Config
	Redis      *redis.Client
	VideoRpc   videoclient.Video
	UserRpc    userclient.User
	ContactRpc contactclient.Contact
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Redis:      initRedis(c),
		VideoRpc:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ContactRpc: contactclient.NewContact(zrpc.MustNewClient(c.ContactRpc)),
	}
}

func initRedis(c config.Config) *redis.Client {
	fmt.Println("connect Redis ...")
	db := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		//DB:       c.DBList.Redis.DB,
		//超时
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  3 * time.Second,
	})
	_, err := db.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect Redis failed")
		panic(err)
	}
	fmt.Println("connect Redis success")
	return db
}
