package svc

import (
	"upgradelink-admin-upgrade/server/ent"
	_ "upgradelink-admin-upgrade/server/ent/runtime"
	"upgradelink-admin-upgrade/server/internal/config"
	"upgradelink-admin-upgrade/server/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	DB        *ent.Client
	Trans     *i18n.Translator
	CoreRpc   coreclient.Core
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type, c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(c.I18nConf, i18n.LocaleFS)

	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		DB:        db,
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		Trans:     trans,
	}
}
