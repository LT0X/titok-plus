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
	Host      string
	Port      int64
	User      string
	PassWord  string
	DataBase  string
	CharSet   string
	ParseTime bool
	Loc       string
}

type RedisConf struct {
	Address  string
	Password string
	//DB       int
}
