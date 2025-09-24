package upgrade_url_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_url_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_url_version/list upgrade_url_version GetUpgradeUrlVersionList
//
// Get upgrade url version list | 获取UpgradeUrlVersion信息列表
//
// Get upgrade url version list | 获取UpgradeUrlVersion信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlVersionListReq
//
// Responses:
//  200: UpgradeUrlVersionListResp

func GetUpgradeUrlVersionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlVersionListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url_version.NewGetUpgradeUrlVersionListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeUrlVersionList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
