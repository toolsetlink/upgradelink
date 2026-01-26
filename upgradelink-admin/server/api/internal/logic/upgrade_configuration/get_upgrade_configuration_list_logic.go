package upgrade_configuration

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeconfiguration"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeConfigurationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeConfigurationListLogic {
	return &GetUpgradeConfigurationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeConfigurationListLogic) GetUpgradeConfigurationList(req *types.UpgradeConfigurationListReq) (*types.UpgradeConfigurationListResp, error) {
	var predicates []predicate.UpgradeConfiguration
	predicates = append(predicates, upgradeconfiguration.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradeconfiguration.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradeconfiguration.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeconfiguration.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeConfiguration.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeConfigurationListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeConfigurationInfo{
				Id:          &v.ID,
				Key:         &v.Key,
				Name:        &v.Name,
				Description: &v.Description,
				IsDel:       &v.IsDel,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
