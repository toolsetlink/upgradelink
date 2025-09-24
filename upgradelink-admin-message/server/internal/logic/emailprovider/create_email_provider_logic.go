package emailprovider

import (
	"context"

	emailprovider2 "upgradelink-admin-message/server/ent/emailprovider"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailProviderLogic {
	return &CreateEmailProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEmailProviderLogic) CreateEmailProvider(in *mcms.EmailProviderInfo) (*mcms.BaseIDResp, error) {
	query := l.svcCtx.DB.EmailProvider.Create().
		SetNotNilName(in.Name).
		SetNotNilEmailAddr(in.EmailAddr).
		SetNotNilPassword(in.Password).
		SetNotNilHostName(in.HostName).
		SetNotNilIdentify(in.Identify).
		SetNotNilSecret(in.Secret).
		SetNotNilPort(in.Port).
		SetNotNilTLS(in.Tls).
		SetNotNilIsDefault(in.IsDefault)

	if in.AuthType != nil {
		query.SetNotNilAuthType(pointy.GetPointer(uint8(*in.AuthType)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// If it is default, set other default to false
	if in.IsDefault != nil && *in.IsDefault == true {
		err = l.svcCtx.DB.EmailProvider.Update().
			Where(emailprovider2.Not(emailprovider2.IDEQ(*in.Id))).
			SetIsDefault(false).
			Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
	}

	return &mcms.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
