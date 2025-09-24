package upgrade_company_income

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeCompanyIncomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeCompanyIncomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeCompanyIncomeLogic {
	return &CreateUpgradeCompanyIncomeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeCompanyIncomeLogic) CreateUpgradeCompanyIncome(req *types.UpgradeCompanyIncomeInfo) (*types.BaseMsgResp, error) {
	//_, err := l.svcCtx.DB.UpgradeCompanyIncome.Create().
	//		SetNotNilCompanyID(req.CompanyId).
	//		SetNotNilIncomeType(req.IncomeType).
	//		SetNotNilIncomeAmount(req.IncomeAmount).
	//		SetNotNilRemark(req.Remark).
	//		SetNotNilStatus(req.Status).
	//		SetNotNilIsDel(req.IsDel).
	//		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
	//		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
	//		Save(l.ctx)
	//
	//if err != nil {
	//	return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	//}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}
