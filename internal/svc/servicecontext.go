package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl.com/shorturl/internal/config"
	"shorturl.com/shorturl/rpc/transform/transformer"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // manual code
	}
}
