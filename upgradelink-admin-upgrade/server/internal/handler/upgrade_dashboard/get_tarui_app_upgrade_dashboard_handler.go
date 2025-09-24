package upgrade_dashboard

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dashboard"
	"upgradelink-admin-upgrade/server/internal/svc"
)

// swagger:route post /tauri_app_upgrade_dashboard upgrade_dashboard GetTaruiAppUpgradeDashboard
//
// 获取应用级别tauri应用统计数据
//
// 获取应用级别tauri应用统计数据
//
// Responses:
//  200: TaruiAppUpgradeDashboardResp

func GetTaruiAppUpgradeDashboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upgrade_dashboard.NewGetTaruiAppUpgradeDashboardLogic(r.Context(), svcCtx)
		resp, err := l.GetTaruiAppUpgradeDashboard()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
