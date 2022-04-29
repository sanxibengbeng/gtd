package svc

import (
	"shorturl.com/shorturl/rpc/transform/internal/config"
	"shorturl.com/shorturl/rpc/transform/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
