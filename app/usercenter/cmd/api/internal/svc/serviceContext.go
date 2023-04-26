package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"learn-looklook/app/usercenter/cmd/api/internal/config"
	"learn-looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
