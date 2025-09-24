package upgrade_dashboard

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dashboard"
	"upgradelink-admin-upgrade/server/internal/svc"
)

// swagger:route post /upgrade_dashboard upgrade_dashboard GetUpgradeDashboard
//
// Get upgrade url by ID | 通过ID获取UpgradeUrl信息
//
// Get upgrade url by ID | 通过ID获取UpgradeUrl信息
//
// Responses:
//  200: UpgradeDashboardResp

func GetUpgradeDashboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upgrade_dashboard.NewGetUpgradeDashboardLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDashboard()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
