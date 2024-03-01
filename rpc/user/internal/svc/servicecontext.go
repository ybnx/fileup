package svc

import (
	"fileup/rpc/user/internal/config"
	"fileup/util"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     util.InitDB(c.Mysql.DataSource),
		RDB:    util.InitRDB(c.Redis.Address), // TODO redis并发问题
	}
}
