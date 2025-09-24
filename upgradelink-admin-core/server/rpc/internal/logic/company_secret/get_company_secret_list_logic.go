package company_secret

import (
	"context"
	"time"

	"upgradelink-admin-core/server/rpc/ent/companysecret"
	"upgradelink-admin-core/server/rpc/ent/predicate"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanySecretListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompanySecretListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanySecretListLogic {
	return &GetCompanySecretListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompanySecretListLogic) GetCompanySecretList(in *core.CompanySecretListReq) (*core.CompanySecretListResp, error) {
	var predicates []predicate.CompanySecret
	if in.CreatedAt != nil {
		predicates = append(predicates, companysecret.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, companysecret.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.CompanyId != nil {
		predicates = append(predicates, companysecret.CompanyIDEQ(*in.CompanyId))
	}
	if in.AccessKey != nil {
		predicates = append(predicates, companysecret.AccessKeyContains(*in.AccessKey))
	}
	if in.SecretKey != nil {
		predicates = append(predicates, companysecret.SecretKeyContains(*in.SecretKey))
	}
	result, err := l.svcCtx.DB.CompanySecret.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.CompanySecretListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.CompanySecretInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			CompanyId: &v.CompanyID,
			AccessKey: &v.AccessKey,
			SecretKey: &v.SecretKey,
		})
	}

	return resp, nil
}
