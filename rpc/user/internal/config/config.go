package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Address string
	}
	Email struct {
		UserName string
		Password string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
