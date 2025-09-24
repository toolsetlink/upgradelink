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

type CreateSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsProviderLogic {
	return &CreateSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSmsProviderLogic) CreateSmsProvider(in *mcms.SmsProviderInfo) (*mcms.BaseIDResp, error) {
	result, err := l.svcCtx.DB.SmsProvider.Create().
		SetNotNilName(in.Name).
		SetNotNilSecretID(in.SecretId).
		SetNotNilSecretKey(in.SecretKey).
		SetNotNilRegion(in.Region).
		SetNotNilIsDefault(in.IsDefault).
		Save(l.ctx)

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

	return &mcms.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
