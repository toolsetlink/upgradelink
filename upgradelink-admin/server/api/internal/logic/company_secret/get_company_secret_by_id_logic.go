package company_secret

import (
	"context"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanySecretByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompanySecretByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanySecretByIdLogic {
	return &GetCompanySecretByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompanySecretByIdLogic) GetCompanySecretById(req *types.IDReq) (resp *types.CompanySecretInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
