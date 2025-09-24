package configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/configuration"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /configuration/create configuration CreateConfiguration
//
// Create configuration information | 创建Configuration
//
// Create configuration information | 创建Configuration
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ConfigurationInfo
//
// Responses:
//  200: BaseMsgResp

func CreateConfigurationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConfigurationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := configuration.NewCreateConfigurationLogic(r.Context(), svcCtx)
		resp, err := l.CreateConfiguration(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
