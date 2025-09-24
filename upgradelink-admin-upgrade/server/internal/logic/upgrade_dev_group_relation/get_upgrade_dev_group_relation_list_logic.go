package upgrade_dev_group_relation

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedevgrouprelation"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevGroupRelationListResp{}
	resp.Msg = errormsg.Success
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
