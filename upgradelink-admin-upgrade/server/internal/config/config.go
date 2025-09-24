package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth               rest.AuthConf
	CROSConf           config.CROSConf
	CasbinDatabaseConf config.DatabaseConf
	RedisConf          config.RedisConf
	CasbinConf         casbin.CasbinConf
	DatabaseConf       config.DatabaseConf
	CoreRpc            zrpc.RpcClientConf
	I18nConf           i18n.Conf
}
