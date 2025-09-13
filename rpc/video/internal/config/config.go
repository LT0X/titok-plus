package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DBList DBListConf
}

type DBListConf struct {
	Mysql MysqlConf
	Redis RedisConf
}

type MysqlConf struct {
	Address     string
	Username    string
	Password    string
	DBName      string
	TablePrefix string
}

type RedisConf struct {
	Host     string
	Port     string
	Password string
	PoolSize int
}
