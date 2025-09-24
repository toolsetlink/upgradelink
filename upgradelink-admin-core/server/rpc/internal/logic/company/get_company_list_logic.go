package company

import (
	"context"
	"time"

	"upgradelink-admin-core/server/rpc/ent/company"
	"upgradelink-admin-core/server/rpc/ent/predicate"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompanyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanyListLogic {
	return &GetCompanyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompanyListLogic) GetCompanyList(in *core.CompanyListReq) (*core.CompanyListResp, error) {
	var predicates []predicate.Company
	if in.CreatedAt != nil {
		predicates = append(predicates, company.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, company.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.Name != nil {
		predicates = append(predicates, company.NameContains(*in.Name))
	}
	result, err := l.svcCtx.DB.Company.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.CompanyListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.CompanyInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:      &v.Name,
		})
	}

	return resp, nil
}
