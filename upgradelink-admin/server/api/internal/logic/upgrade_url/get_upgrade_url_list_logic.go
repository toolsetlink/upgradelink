package upgrade_url

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeurl"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeUrlListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeUrlListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeUrlListLogic {
	return &GetUpgradeUrlListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeUrlListLogic) GetUpgradeUrlList(req *types.UpgradeUrlListReq) (*types.UpgradeUrlListResp, error) {
	var predicates []predicate.UpgradeUrl
	predicates = append(predicates, upgradeurl.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradeurl.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradeurl.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeurl.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeUrlListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeUrlInfo{
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
