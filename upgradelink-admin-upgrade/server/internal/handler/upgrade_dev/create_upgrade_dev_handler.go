package upgrade_dev

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dev"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_dev/create upgrade_dev CreateUpgradeDev
//
// Create upgrade dev information | 创建UpgradeDev信息
//
// Create upgrade dev information | 创建UpgradeDev信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeDevHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev.NewCreateUpgradeDevLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeDev(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
