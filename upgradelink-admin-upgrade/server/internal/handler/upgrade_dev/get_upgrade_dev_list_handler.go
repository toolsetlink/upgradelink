package upgrade_dev

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dev"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_dev/list upgrade_dev GetUpgradeDevList
//
// Get upgrade dev list | 获取UpgradeDev信息列表
//
// Get upgrade dev list | 获取UpgradeDev信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevListReq
//
// Responses:
//  200: UpgradeDevListResp

func GetUpgradeDevListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev.NewGetUpgradeDevListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDevList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
