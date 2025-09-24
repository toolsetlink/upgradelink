package company_secret

import (
	"context"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCompanySecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanySecretLogic {
	return &CreateCompanySecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCompanySecretLogic) CreateCompanySecret(in *core.CompanySecretInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.CompanySecret.Create().
		SetNotNilCompanyID(in.CompanyId).
		SetNotNilAccessKey(in.AccessKey).
		SetNotNilSecretKey(in.SecretKey).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
