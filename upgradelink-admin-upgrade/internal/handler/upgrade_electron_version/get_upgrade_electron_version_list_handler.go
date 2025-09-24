package upgrade_electron_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_electron_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_electron_version/list upgrade_electron_version GetUpgradeElectronVersionList
//
// Get upgrade electron version list | 获取UpgradeElectronVersion信息列表
//
// Get upgrade electron version list | 获取UpgradeElectronVersion信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronVersionListReq
//
// Responses:
//  200: UpgradeElectronVersionListResp

func GetUpgradeElectronVersionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronVersionListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_version.NewGetUpgradeElectronVersionListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeElectronVersionList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
