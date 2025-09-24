package upgrade_company_income

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeCompanyIncomeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeCompanyIncomeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeCompanyIncomeByIdLogic {
	return &GetUpgradeCompanyIncomeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeCompanyIncomeByIdLogic) GetUpgradeCompanyIncomeById(req *types.IDReq) (*types.UpgradeCompanyIncomeInfoResp, error) {
	data, err := l.svcCtx.DB.UpgradeCompanyIncome.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeCompanyIncomeInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.UpgradeCompanyIncomeInfo{
			Id:           &data.ID,
			CompanyId:    &data.CompanyID,
			IncomeType:   &data.IncomeType,
			IncomeAmount: &data.IncomeAmount,
			Remark:       &data.Remark,
			Status:       &data.Status,
			IsDel:        &data.IsDel,
			IncomeTime:   pointy.GetUnixMilliPointer(data.IncomeTime.UnixMilli()),
			CreateAt:     pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:     pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
