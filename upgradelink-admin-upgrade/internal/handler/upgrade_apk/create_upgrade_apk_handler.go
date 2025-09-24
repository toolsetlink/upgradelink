package upgrade_apk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_apk"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_apk/create upgrade_apk CreateUpgradeApk
//
// Create upgrade apk information | 创建UpgradeApk信息
//
// Create upgrade apk information | 创建UpgradeApk信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeApkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk.NewCreateUpgradeApkLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeApk(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
