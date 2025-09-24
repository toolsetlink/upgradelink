package upgrade_apk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_apk"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_apk/list upgrade_apk GetUpgradeApkList
//
// Get upgrade apk list | 获取UpgradeApk信息列表
//
// Get upgrade apk list | 获取UpgradeApk信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkListReq
//
// Responses:
//  200: UpgradeApkListResp

func GetUpgradeApkListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk.NewGetUpgradeApkListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeApkList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
