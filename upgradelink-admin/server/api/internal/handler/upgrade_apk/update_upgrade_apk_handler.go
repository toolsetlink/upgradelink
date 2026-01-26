package upgrade_apk

import (
	"net/http"

	"upgradelink-admin/server/api/internal/logic/upgrade_apk"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUpgradeApkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk.NewUpdateUpgradeApkLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeApk(&req)
		if err != nil {
			//err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
