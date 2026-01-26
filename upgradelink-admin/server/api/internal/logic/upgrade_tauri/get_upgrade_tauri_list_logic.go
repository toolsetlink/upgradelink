package upgrade_tauri

import (
	"context"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradetauri"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeTauriListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeTauriListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeTauriListLogic {
	return &GetUpgradeTauriListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeTauriListLogic) GetUpgradeTauriList(req *types.UpgradeTauriListReq) (*types.UpgradeTauriListResp, error) {
	var predicates []predicate.UpgradeTauri
	predicates = append(predicates, upgradetauri.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradetauri.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradetauri.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradetauri.NameContains(*req.Name))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradetauri.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradetauri.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradetauri.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeTauri.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeTauriListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeTauriInfo{
				Id:          &v.ID,
				Key:         &v.Key,
				Name:        &v.Name,
				Description: &v.Description,
				GithubUrl:   &v.GithubURL,
				IsDel:       &v.IsDel,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
