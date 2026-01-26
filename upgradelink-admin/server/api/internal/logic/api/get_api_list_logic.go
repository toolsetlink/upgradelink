package api

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/api"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	var predicates []predicate.API
	if req.Path != nil {
		predicates = append(predicates, api.PathContains(*req.Path))
	}
	if req.Description != nil {
		predicates = append(predicates, api.DescriptionContains(*req.Description))
	}
	if req.Group != nil {
		predicates = append(predicates, api.APIGroupContains(*req.Group))
	}
	if req.Method != nil {
		predicates = append(predicates, api.Method(*req.Method))
	}
	if req.ServiceName != nil {
		predicates = append(predicates, api.ServiceNameContains(*req.ServiceName))
	}
	result, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.ApiListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = result.PageDetails.Total

	// Translate api group
	var group = make(map[string]string)
	for _, v := range result.List {
		// If page size is over 1000, use api group i18n translation. Mainly use in authority list.
		if req.PageSize > 1000 {
			if _, exist := group[v.APIGroup]; !exist {
				group[v.APIGroup] = l.svcCtx.Trans.Trans(l.ctx, "apiGroup."+v.APIGroup)
			}
			v.APIGroup = group[v.APIGroup]
		}
		resp.Data.Data = append(resp.Data.Data,
			types.ApiInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Path:        &v.Path,
				Trans:       l.svcCtx.Trans.Trans(l.ctx, v.Description),
				Description: &v.Description,
				Group:       &v.APIGroup,
				Method:      &v.Method,
				IsRequired:  &v.IsRequired,
				ServiceName: &v.ServiceName,
			})
	}
	return resp, nil
}
