package upgrade_company_income

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeCompanyIncomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeCompanyIncomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeCompanyIncomeLogic {
	return &DeleteUpgradeCompanyIncomeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeCompanyIncomeLogic) DeleteUpgradeCompanyIncome(req *types.IDsReq) (*types.BaseMsgResp, error) {
	//_, err := l.svcCtx.DB.UpgradeCompanyIncome.Delete().Where(upgradecompanyincome.IDIn(req.Ids...)).Exec(l.ctx)
	//
	//if err != nil {
	//	return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	//}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
