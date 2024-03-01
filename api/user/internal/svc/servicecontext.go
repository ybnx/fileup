package svc

import (
	"fileup/api/user/internal/config"
	"fileup/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Rpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
