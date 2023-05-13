package svc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok-plus/service/rpc/video/internal/config"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DBList
	AsynqClient *asynq.Client
}

type DBList struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DBList: DBList{
			DB:    initDB(c.DBList.Mysql),
			Redis: initRedis(c),
		},
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{Addr: c.DBList.Redis.Address, Password: c.DBList.Redis.Password}),
	}
}

func initDB(c config.MysqlConf) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=%v&loc=%s",
		c.User, c.PassWord, c.Host, c.Port,
		c.DataBase, c.CharSet, c.ParseTime,
		c.Loc)

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis(c config.Config) *redis.Client {
	fmt.Println("connect Redis ...")
	db := redis.NewClient(&redis.Options{
		Addr:     c.DBList.Redis.Address,
		Password: c.DBList.Redis.Password,
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
