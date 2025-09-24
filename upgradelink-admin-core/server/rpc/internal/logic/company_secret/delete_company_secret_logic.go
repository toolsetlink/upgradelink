package company_secret

import (
	"context"

	"upgradelink-admin-core/server/rpc/ent/companysecret"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCompanySecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanySecretLogic {
	return &DeleteCompanySecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCompanySecretLogic) DeleteCompanySecret(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.CompanySecret.Delete().Where(companysecret.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
