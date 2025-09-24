package upgrade_dev_group

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradedevgroup"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
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

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(companyID))

	// 删除状态
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevGroupListResp{}
	resp.Msg = errormsg.Success
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
