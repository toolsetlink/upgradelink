package upgrade_dev_model

import (
	"context"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradedevmodel"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
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

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradedevmodel.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradedevmodel.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradedevmodel.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradedevmodel.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeDevModel.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevModelListResp{}
	resp.Msg = errormsg.Success
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
