package svc

import (
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/captcha"
	"upgradelink-admin/server/api/internal/ent"
	i18n2 "upgradelink-admin/server/api/internal/i18n"
	"upgradelink-admin/server/api/internal/middleware"

	"upgradelink-admin/server/api/internal/config"

	"github.com/casbin/casbin/v3"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ServiceContext struct {
	Config      config.Config
	Authority   rest.Middleware
	CompanyData rest.Middleware
	DB          *ent.Client
	Redis       redis.UniversalClient
	Casbin      *casbin.Enforcer
	Trans       *i18n.Translator
	Captcha     *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 1. 注册全局异常处理
	httpx.SetErrorHandlerCtx(http_error.ErrorHandler)

	// 2. 连接redis
	rds := c.RedisConf.MustNewUniversalRedis()

	// 3. 注册数据库
	db := ent.NewClient(
		ent.Log(logx.Error), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
	)

	// 4. 注册casbin
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(),
		c.RedisConf)

	// 5. 注册翻译
	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:      c,
		Captcha:     captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		DB:          db,
		Redis:       rds,
		Casbin:      cbn,
		Trans:       trans,
		Authority:   middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		CompanyData: middleware.NewCompanyDataMiddleware(db).Handle,
	}
}
