package upgrade_dev_group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group/list upgrade_dev_group GetUpgradeDevGroupList
//
// Get upgrade dev group list | 获取UpgradeDevGroup信息列表
//
// Get upgrade dev group list | 获取UpgradeDevGroup信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevGroupListReq
//
// Responses:
//  200: UpgradeDevGroupListResp

func GetUpgradeDevGroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevGroupListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group.NewGetUpgradeDevGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDevGroupList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
