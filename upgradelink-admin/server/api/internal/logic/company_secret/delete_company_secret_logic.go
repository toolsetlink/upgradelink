package company_secret

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/companysecret"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCompanySecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanySecretLogic {
	return &DeleteCompanySecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCompanySecretLogic) DeleteCompanySecret(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {

	_, err = l.svcCtx.DB.CompanySecret.Delete().Where(companysecret.IDIn(req.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
