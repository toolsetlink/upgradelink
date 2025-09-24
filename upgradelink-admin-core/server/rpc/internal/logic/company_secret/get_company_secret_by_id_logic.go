package company_secret

import (
	"context"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanySecretByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompanySecretByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanySecretByIdLogic {
	return &GetCompanySecretByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompanySecretByIdLogic) GetCompanySecretById(in *core.IDReq) (*core.CompanySecretInfo, error) {
	result, err := l.svcCtx.DB.CompanySecret.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.CompanySecretInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		CompanyId: &result.CompanyID,
		AccessKey: &result.AccessKey,
		SecretKey: &result.SecretKey,
	}, nil
}
