package configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/configuration"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /configuration/update configuration UpdateConfiguration
//
// Update configuration information | 更新Configuration
//
// Update configuration information | 更新Configuration
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ConfigurationInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateConfigurationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConfigurationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := configuration.NewUpdateConfigurationLogic(r.Context(), svcCtx)
		resp, err := l.UpdateConfiguration(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
