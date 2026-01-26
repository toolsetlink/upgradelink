package upgrade_dev_model

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevmodel"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevModelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevModelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevModelListLogic {
	return &GetUpgradeDevModelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevModelListLogic) GetUpgradeDevModelList(req *types.UpgradeDevModelListReq) (*types.UpgradeDevModelListResp, error) {
	var predicates []predicate.UpgradeDevModel
	predicates = append(predicates, upgradedevmodel.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradedevmodel.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradedevmodel.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradedevmodel.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeDevModel.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevModelListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeDevModelInfo{
				Id:       &v.ID,
				Key:      &v.Key,
				Name:     &v.Name,
				IsDel:    &v.IsDel,
				CreateAt: pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt: pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
