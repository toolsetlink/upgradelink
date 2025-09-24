package upgrade_configuration_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration_version upgrade_configuration_version GetUpgradeConfigurationVersionById
//
// Get upgrade configuration version by ID | 通过ID获取UpgradeConfigurationVersion信息
//
// Get upgrade configuration version by ID | 通过ID获取UpgradeConfigurationVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeConfigurationVersionInfoResp

func GetUpgradeConfigurationVersionByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration_version.NewGetUpgradeConfigurationVersionByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeConfigurationVersionById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
