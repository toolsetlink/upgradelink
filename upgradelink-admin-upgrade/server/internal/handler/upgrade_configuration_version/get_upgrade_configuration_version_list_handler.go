package upgrade_configuration_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_configuration_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_configuration_version/list upgrade_configuration_version GetUpgradeConfigurationVersionList
//
// Get upgrade configuration version list | 获取UpgradeConfigurationVersion信息列表
//
// Get upgrade configuration version list | 获取UpgradeConfigurationVersion信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationVersionListReq
//
// Responses:
//  200: UpgradeConfigurationVersionListResp

func GetUpgradeConfigurationVersionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationVersionListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration_version.NewGetUpgradeConfigurationVersionListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeConfigurationVersionList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
