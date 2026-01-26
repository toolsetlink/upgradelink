package upgrade_url_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeurlversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeUrlVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeUrlVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeUrlVersionListLogic {
	return &GetUpgradeUrlVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeUrlVersionListLogic) GetUpgradeUrlVersionList(req *types.UpgradeUrlVersionListReq) (*types.UpgradeUrlVersionListResp, error) {
	var predicates []predicate.UpgradeUrlVersion
	predicates = append(predicates, upgradeurlversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradeurlversion.IsDelEQ(0))

	if req.UrlId != nil {
		predicates = append(predicates, upgradeurlversion.URLIDEQ(*req.UrlId))
	}
	if req.VersionName != nil {
		predicates = append(predicates, upgradeurlversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradeurlversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.UrlPath != nil {
		predicates = append(predicates, upgradeurlversion.URLPathContains(*req.UrlPath))
	}

	data, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeUrlVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		urlData, err := l.svcCtx.DB.UpgradeUrl.Get(l.ctx, v.URLID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeUrlVersionInfo{
				Id:          &v.ID,
				UrlId:       &v.URLID,
				UrlName:     &urlData.Name,
				UrlPath:     &v.URLPath,
				VersionName: &v.VersionName,
				VersionCode: &v.VersionCode,
				Description: &v.Description,
				IsDel:       &v.IsDel,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
