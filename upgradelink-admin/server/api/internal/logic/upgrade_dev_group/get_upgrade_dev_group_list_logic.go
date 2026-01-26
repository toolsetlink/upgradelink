package upgrade_dev_group

import (
	"context"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevgroup"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupListLogic {
	return &GetUpgradeDevGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupListLogic) GetUpgradeDevGroupList(req *types.UpgradeDevGroupListReq) (*types.UpgradeDevGroupListResp, error) {
	var predicates []predicate.UpgradeDevGroup
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradedevgroup.IsDelEQ(0))

	if req.Name != nil {
		predicates = append(predicates, upgradedevgroup.NameContains(*req.Name))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradedevgroup.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradedevgroup.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeDevGroup.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevGroupListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeDevGroupInfo{
				Id:       &v.ID,
				Name:     &v.Name,
				IsDel:    &v.IsDel,
				CreateAt: pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt: pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
