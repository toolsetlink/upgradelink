package upgrade_dashboard

import (
	"net/http"

	"upgradelink-admin/server/api/internal/logic/upgrade_dashboard"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTaruiAppUpgradeDashboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dashboard.NewGetTaruiAppUpgradeDashboardLogic(r.Context(), svcCtx)
		resp, err := l.GetTaruiAppUpgradeDashboard()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
