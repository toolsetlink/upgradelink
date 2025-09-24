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
	UploadConf         UploadConf
	DatabaseConf       config.DatabaseConf
	CasbinDatabaseConf config.DatabaseConf
	RedisConf          config.RedisConf
	CasbinConf         casbin.CasbinConf
	CoreRpc            zrpc.RpcClientConf
	CROSConf           config.CROSConf
	I18nConf           i18n.Conf
}

type UploadConf struct {
	MaxImageSize        int64  `json:",default=33554432,env=MAX_IMAGE_SIZE"`
	MaxVideoSize        int64  `json:",default=1073741824,env=MAX_VIDEO_SIZE"`
	MaxAudioSize        int64  `json:",default=33554432,env=MAX_AUDIO_SIZE"`
	MaxOtherSize        int64  `json:",default=10485760,env=MAX_OTHER_SIZE"`
	PrivateStorePath    string `json:",env=PRIVATE_PATH"`
	PublicStorePath     string `json:",env=PUBLIC_PATH"`
	ServerURL           string `json:",env=SERVER_URL"`
	DeleteFileWithCloud bool   `json:",env=DELETE_FILE_WITH_CLOUD,default=true"`
}
