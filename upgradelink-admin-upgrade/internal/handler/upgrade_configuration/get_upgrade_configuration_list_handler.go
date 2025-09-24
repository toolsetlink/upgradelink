package upgrade_configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration/list upgrade_configuration GetUpgradeConfigurationList
//
// Get upgrade configuration list | 获取UpgradeConfiguration信息列表
//
// Get upgrade configuration list | 获取UpgradeConfiguration信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationListReq
//
// Responses:
//  200: UpgradeConfigurationListResp

func GetUpgradeConfigurationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration.NewGetUpgradeConfigurationListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeConfigurationList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
