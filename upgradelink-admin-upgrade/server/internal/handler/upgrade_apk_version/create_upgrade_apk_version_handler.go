package upgrade_apk_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_apk_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_apk_version/create upgrade_apk_version CreateUpgradeApkVersion
//
// Create upgrade apk version information | 创建UpgradeApkVersion信息
//
// Create upgrade apk version information | 创建UpgradeApkVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkVersionInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeApkVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk_version.NewCreateUpgradeApkVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeApkVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
