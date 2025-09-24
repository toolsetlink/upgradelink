package upgrade_company_income

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeCompanyIncomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeCompanyIncomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeCompanyIncomeLogic {
	return &UpdateUpgradeCompanyIncomeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeCompanyIncomeLogic) UpdateUpgradeCompanyIncome(req *types.UpgradeCompanyIncomeInfo) (*types.BaseMsgResp, error) {
	//err := l.svcCtx.DB.UpgradeCompanyIncome.UpdateOneID(*req.Id).
	//		SetNotNilCompanyID(req.CompanyId).
	//		SetNotNilIncomeType(req.IncomeType).
	//		SetNotNilIncomeAmount(req.IncomeAmount).
	//		SetNotNilRemark(req.Remark).
	//		SetNotNilStatus(req.Status).
	//		SetNotNilIsDel(req.IsDel).
	//		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
	//		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
	//		Exec(l.ctx)
	//
	//if err != nil {
	//	return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	//}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}
