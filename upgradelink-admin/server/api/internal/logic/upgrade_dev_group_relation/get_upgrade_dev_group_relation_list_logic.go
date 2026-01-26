package upgrade_dev_group_relation

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevgrouprelation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupRelationListLogic {
	return &GetUpgradeDevGroupRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupRelationListLogic) GetUpgradeDevGroupRelationList(req *types.UpgradeDevGroupRelationListReq) (*types.UpgradeDevGroupRelationListResp, error) {
	var predicates []predicate.UpgradeDevGroupRelation
	if req.DevId != nil {
		predicates = append(predicates, upgradedevgrouprelation.DevIDEQ(*req.DevId))
	}
	if req.DevGroupId != nil {
		predicates = append(predicates, upgradedevgrouprelation.DevGroupIDEQ(*req.DevGroupId))
	}
	data, err := l.svcCtx.DB.UpgradeDevGroupRelation.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevGroupRelationListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeDevGroupRelationInfo{
				Id:         &v.ID,
				DevId:      &v.DevID,
				DevGroupId: &v.DevGroupID,
			})
	}

	return resp, nil
}
