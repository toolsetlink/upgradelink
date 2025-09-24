package upgrade_url_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_url_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_url_version/update upgrade_url_version UpdateUpgradeUrlVersion
//
// Update upgrade url version information | 更新UpgradeUrlVersion信息
//
// Update upgrade url version information | 更新UpgradeUrlVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlVersionInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeUrlVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url_version.NewUpdateUpgradeUrlVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeUrlVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
