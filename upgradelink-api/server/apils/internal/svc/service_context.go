package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"upgradelink-api/server/apils/internal/config"
	"upgradelink-api/server/apils/internal/middleware"
)

type ServiceContext struct {
	Config       config.Config
	RateLimit    rest.Middleware
	ReplayAttack rest.Middleware
	Signature    rest.Middleware
	CdnRateLimit rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		RateLimit:    middleware.NewRateLimitMiddleware().Handle,
		ReplayAttack: middleware.NewReplayAttackMiddleware().Handle,
		Signature:    middleware.NewSignatureMiddleware().Handle,
		CdnRateLimit: middleware.NewCdnRateLimitMiddleware().Handle,
	}
}
