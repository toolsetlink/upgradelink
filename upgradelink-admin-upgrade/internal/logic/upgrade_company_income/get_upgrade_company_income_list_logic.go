package upgrade_company_income

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradecompanyincome"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeCompanyIncomeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeCompanyIncomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeCompanyIncomeListLogic {
	return &GetUpgradeCompanyIncomeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeCompanyIncomeListLogic) GetUpgradeCompanyIncomeList(req *types.UpgradeCompanyIncomeListReq) (*types.UpgradeCompanyIncomeListResp, error) {
	var predicates []predicate.UpgradeCompanyIncome

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradecompanyincome.CompanyIDEQ(companyID))

	if req.IncomeType != nil {
		predicates = append(predicates, upgradecompanyincome.IncomeTypeEQ(*req.IncomeType))
	}
	if req.IncomeAmount != nil {
		predicates = append(predicates, upgradecompanyincome.IncomeAmountEQ(*req.IncomeAmount))
	}
	if req.Remark != nil {
		predicates = append(predicates, upgradecompanyincome.RemarkContains(*req.Remark))
	}
	if req.Status != nil {
		predicates = append(predicates, upgradecompanyincome.StatusEQ(*req.Status))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradecompanyincome.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradecompanyincome.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeCompanyIncome.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeCompanyIncomeListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total
	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeCompanyIncomeInfo{
				Id:           &v.ID,
				CompanyId:    &v.CompanyID,
				IncomeType:   &v.IncomeType,
				IncomeAmount: &v.IncomeAmount,
				Remark:       &v.Remark,
				Status:       &v.Status,
				IsDel:        &v.IsDel,
				IncomeTime:   pointy.GetUnixMilliPointer(v.IncomeTime.UnixMilli()),
				CreateAt:     pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:     pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	//// 统计出未结算金额
	//incomeAmount, err := l.svcCtx.ResourceCtx.GetIncomeAmount(l.ctx, companyID)
	//if err != nil {
	//	return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	//}
	//resp.Data.IncomeAmount = &incomeAmount
	//
	//// 统计已结算金额
	//settleAmount, err := l.svcCtx.ResourceCtx.GetSettleAmount(l.ctx, companyID)
	//if err != nil {
	//	return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	//}
	//resp.Data.SettleAmount = &settleAmount

	return resp, nil
}
