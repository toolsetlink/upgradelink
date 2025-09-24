package sms

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	smsprovider2 "upgradelink-admin-message/server/ent/smsprovider"
	"upgradelink-admin-message/server/internal/enum/smsprovider"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/internal/utils/smssdk"
	"upgradelink-admin-message/server/types/mcms"
)

func (l *SendSmsLogic) initProvider(in *mcms.SmsInfo) error {
	switch *in.Provider {
	case smsprovider.Tencent:
		if l.svcCtx.SmsGroup.TencentSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.TencentSmsClient, err = clientConf.NewTencentClient()
			if err != nil {
				logx.Error("failed to initialize Tencent SMS client, please check the configuration", logx.Field("detail", err))
				return errorx.NewInvalidArgumentError("failed to initialize Tencent SMS client, please check the configuration")
			}
		}
	case smsprovider.Aliyun:
		if l.svcCtx.SmsGroup.AliyunSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.AliyunSmsClient, err = clientConf.NewAliyunClient()
			if err != nil {
				logx.Error("failed to initialize Aliyun SMS client, please check the configuration", logx.Field("detail", err))
				return errorx.NewInvalidArgumentError("failed to initialize Aliyun SMS client, please check the configuration")
			}
		}
	case smsprovider.Uni:
		if l.svcCtx.SmsGroup.UniSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.UniSmsClient, err = clientConf.NewUniClient()
			if err != nil {
				logx.Error("failed to initialize Uni SMS client, please check the configuration", logx.Field("detail", err))
				return errorx.NewInvalidArgumentError("failed to initialize Uni SMS client, please check the configuration")
			}
		}
	case smsprovider.SmsBao:

	default:
		return errorx.NewInvalidArgumentError("provider not found")
	}

	return nil
}
