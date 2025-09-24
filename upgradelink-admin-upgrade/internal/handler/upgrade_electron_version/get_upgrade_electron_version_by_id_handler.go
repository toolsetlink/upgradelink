package upgrade_electron_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_electron_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_electron_version upgrade_electron_version GetUpgradeElectronVersionById
//
// Get upgrade electron version by ID | 通过ID获取UpgradeElectronVersion信息
//
// Get upgrade electron version by ID | 通过ID获取UpgradeElectronVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeElectronVersionInfoResp

func GetUpgradeElectronVersionByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_version.NewGetUpgradeElectronVersionByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeElectronVersionById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
