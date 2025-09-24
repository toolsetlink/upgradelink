package upgrade_url

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_url"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_url/list upgrade_url GetUpgradeUrlList
//
// Get upgrade url list | 获取UpgradeUrl信息列表
//
// Get upgrade url list | 获取UpgradeUrl信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlListReq
//
// Responses:
//  200: UpgradeUrlListResp

func GetUpgradeUrlListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url.NewGetUpgradeUrlListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeUrlList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
