package svc

import (
	"douyin/rpc/video/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DBList *DBList
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DBList: initDB(c),
	}
}

type DBList struct {
	Mysql *gorm.DB
}

func initDB(c config.Config) *DBList {
	dbList := new(DBList)
	dbList.Mysql = initMysql(c)

	return dbList
}

func initMysql(c config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBList.Mysql.Username,
		c.DBList.Mysql.Password,
		c.DBList.Mysql.Address,
		c.DBList.Mysql.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.DBList.Mysql.TablePrefix, // 表名前缀
			SingularTable: true,                       // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// 自动建表
	//err = db.AutoMigrate(&model.Video{}, &model.Favorite{}, &model.Comment{})
	//if err != nil {
	//	panic(err)
	//}

	return db
}
