package upgrade_apk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_apk"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_apk/update upgrade_apk UpdateUpgradeApk
//
// Update upgrade apk information | 更新UpgradeApk信息
//
// Update upgrade apk information | 更新UpgradeApk信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeApkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk.NewUpdateUpgradeApkLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeApk(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
