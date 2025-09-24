package company_secret

import (
	"context"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompanySecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanySecretLogic {
	return &UpdateCompanySecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCompanySecretLogic) UpdateCompanySecret(in *core.CompanySecretInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.CompanySecret.UpdateOneID(*in.Id).
		SetNotNilCompanyID(in.CompanyId).
		SetNotNilAccessKey(in.AccessKey).
		SetNotNilSecretKey(in.SecretKey).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
