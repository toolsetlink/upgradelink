package upgrade_dashboard

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaruiAppUpgradeDashboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaruiAppUpgradeDashboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaruiAppUpgradeDashboardLogic {
	return &GetTaruiAppUpgradeDashboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaruiAppUpgradeDashboardLogic) GetTaruiAppUpgradeDashboard() (resp *types.TaruiAppUpgradeDashboardResp, err error) {
	// todo: add your logic here and delete this line

	return
}
