package upgrade_configuration_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeconfigurationversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeConfigurationVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeConfigurationVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeConfigurationVersionListLogic {
	return &GetUpgradeConfigurationVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeConfigurationVersionListLogic) GetUpgradeConfigurationVersionList(req *types.UpgradeConfigurationVersionListReq) (*types.UpgradeConfigurationVersionListResp, error) {
	var predicates []predicate.UpgradeConfigurationVersion
	predicates = append(predicates, upgradeconfigurationversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradeconfigurationversion.IsDelEQ(0))

	if req.ConfigurationId != nil {
		predicates = append(predicates, upgradeconfigurationversion.ConfigurationIDEQ(*req.ConfigurationId))
	}

	if req.VersionName != nil {
		predicates = append(predicates, upgradeconfigurationversion.VersionNameContains(*req.VersionName))
	}

	if req.VersionCode != nil {
		predicates = append(predicates, upgradeconfigurationversion.VersionCodeEQ(*req.VersionCode))
	}

	data, err := l.svcCtx.DB.UpgradeConfigurationVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeConfigurationVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		configurationData, err := l.svcCtx.DB.UpgradeConfiguration.Get(l.ctx, v.ConfigurationID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeConfigurationVersionInfo{
				Id:                &v.ID,
				ConfigurationId:   &v.ConfigurationID,
				ConfigurationName: &configurationData.Name,
				Content:           &v.Content,
				VersionName:       &v.VersionName,
				VersionCode:       &v.VersionCode,
				Description:       &v.Description,
				IsDel:             &v.IsDel,
				CreateAt:          pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:          pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
