package smsprovider

import (
	"context"

	smsprovider2 "upgradelink-admin-message/server/ent/smsprovider"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsProviderLogic {
	return &UpdateSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSmsProviderLogic) UpdateSmsProvider(in *mcms.SmsProviderInfo) (*mcms.BaseResp, error) {
	err := l.svcCtx.DB.SmsProvider.UpdateOneID(*in.Id).
		SetNotNilName(in.Name).
		SetNotNilSecretID(in.SecretId).
		SetNotNilSecretKey(in.SecretKey).
		SetNotNilRegion(in.Region).
		SetNotNilIsDefault(in.IsDefault).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// If it is default, set other default to false
	if in.IsDefault != nil && *in.IsDefault == true {
		err = l.svcCtx.DB.SmsProvider.Update().
			Where(smsprovider2.Not(smsprovider2.IDEQ(*in.Id))).
			SetIsDefault(false).
			Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
	}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
