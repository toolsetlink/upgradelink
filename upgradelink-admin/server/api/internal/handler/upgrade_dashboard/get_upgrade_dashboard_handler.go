package upgrade_dashboard

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/upgrade_dashboard"
	"upgradelink-admin/server/api/internal/svc"
)

func GetUpgradeDashboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upgrade_dashboard.NewGetUpgradeDashboardLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDashboard()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
