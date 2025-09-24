package upgrade_apk_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_apk_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_apk_version/update upgrade_apk_version UpdateUpgradeApkVersion
//
// Update upgrade apk version information | 更新UpgradeApkVersion信息
//
// Update upgrade apk version information | 更新UpgradeApkVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkVersionInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeApkVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk_version.NewUpdateUpgradeApkVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeApkVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
