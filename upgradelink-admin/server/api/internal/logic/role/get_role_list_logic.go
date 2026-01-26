package role

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/role"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	var predicates []predicate.Role
	if req.Name != nil {
		predicates = append(predicates, role.NameContains(*req.Name))
	}

	result, err := l.svcCtx.DB.Role.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize, func(pager *ent.RolePager) {
		pager.Order = ent.Asc(role.FieldSort)
	})
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.RoleListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data.Data = append(resp.Data.Data,
			types.RoleInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Trans:         l.svcCtx.Trans.Trans(l.ctx, v.Name),
				Status:        pointy.GetPointer(uint32(v.Status)),
				Name:          &v.Name,
				Code:          &v.Code,
				DefaultRouter: &v.DefaultRouter,
				Remark:        &v.Remark,
				Sort:          &v.Sort,
			})
	}
	return resp, nil
}
