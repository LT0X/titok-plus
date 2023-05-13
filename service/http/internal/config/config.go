package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	//Rpc配置
	UserRPC  zrpc.RpcClientConf
	VideoRPC zrpc.RpcClientConf

	//jwt配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Path struct {
		StaticPath string
	}
}
