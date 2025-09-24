package publicapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/publicapi"
	"upgradelink-admin-core/server/api/internal/svc"
)

// swagger:route get /configuration/system/list publicapi GetPublicSystemConfigurationList
//
// Get public system configuration list | 获取公开系统参数列表
//
// Get public system configuration list | 获取公开系统参数列表
//
// Responses:
//  200: ConfigurationListResp

func GetPublicSystemConfigurationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := publicapi.NewGetPublicSystemConfigurationListLogic(r.Context(), svcCtx)
		resp, err := l.GetPublicSystemConfigurationList()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
