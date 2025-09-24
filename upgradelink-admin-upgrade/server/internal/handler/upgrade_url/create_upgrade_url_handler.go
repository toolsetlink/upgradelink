package upgrade_url

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_url"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_url/create upgrade_url CreateUpgradeUrl
//
// Create upgrade url information | 创建UpgradeUrl信息
//
// Create upgrade url information | 创建UpgradeUrl信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url.NewCreateUpgradeUrlLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeUrl(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
