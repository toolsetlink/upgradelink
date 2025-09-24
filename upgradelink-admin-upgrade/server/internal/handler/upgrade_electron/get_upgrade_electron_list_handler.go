package upgrade_electron

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron/list upgrade_electron GetUpgradeElectronList
//
// Get upgrade electron list | 获取UpgradeElectron信息列表
//
// Get upgrade electron list | 获取UpgradeElectron信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronListReq
//
// Responses:
//  200: UpgradeElectronListResp

func GetUpgradeElectronListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron.NewGetUpgradeElectronListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeElectronList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
