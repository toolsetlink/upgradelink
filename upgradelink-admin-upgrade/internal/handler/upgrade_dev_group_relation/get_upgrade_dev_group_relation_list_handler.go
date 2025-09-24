package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group_relation/list upgrade_dev_group_relation GetUpgradeDevGroupRelationList
//
// Get upgrade dev group relation list | 获取UpgradeDevGroupRelation信息列表
//
// Get upgrade dev group relation list | 获取UpgradeDevGroupRelation信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevGroupRelationListReq
//
// Responses:
//  200: UpgradeDevGroupRelationListResp

func GetUpgradeDevGroupRelationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevGroupRelationListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewGetUpgradeDevGroupRelationListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDevGroupRelationList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
