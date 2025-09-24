package upgrade_url_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_url_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_url_version/create upgrade_url_version CreateUpgradeUrlVersion
//
// Create upgrade url version information | 创建UpgradeUrlVersion信息
//
// Create upgrade url version information | 创建UpgradeUrlVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlVersionInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeUrlVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url_version.NewCreateUpgradeUrlVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeUrlVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
