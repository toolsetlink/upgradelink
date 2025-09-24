package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group_relation/delete upgrade_dev_group_relation DeleteUpgradeDevGroupRelation
//
// Delete upgrade dev group relation information | 删除UpgradeDevGroupRelation信息
//
// Delete upgrade dev group relation information | 删除UpgradeDevGroupRelation信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteUpgradeDevGroupRelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewDeleteUpgradeDevGroupRelationLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUpgradeDevGroupRelation(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
