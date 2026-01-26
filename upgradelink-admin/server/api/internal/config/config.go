package config

import (
	"upgradelink-admin/server/api/internal/common/config"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/plugins/casbin"
	"upgradelink-admin/server/api/internal/common/utils/captcha"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CROSConf config.CROSConf
	Auth     struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
	I18nConf     i18n.Conf
	RedisConf    config.RedisConf
	Captcha      captcha.Conf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	UploadConf   UploadConf
}

type UploadConf struct {
	Bucket    string `json:",optional"`
	SecretID  string `json:",optional"`
	SecretKey string `json:",optional"`
	Endpoint  string `json:",optional"`
	Folder    string `json:",optional"`
	Region    string `json:",optional"`
	CdnUrl    string `json:",optional"`
}
