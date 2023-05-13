package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok-plus/service/rpc/contact/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DBList
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DBList: DBList{
			DB: initDB(c.DBList.Mysql),
		},
	}
}

type DBList struct {
	DB *gorm.DB
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
