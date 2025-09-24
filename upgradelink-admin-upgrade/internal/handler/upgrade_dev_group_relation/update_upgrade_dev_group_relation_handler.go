package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group_relation/update upgrade_dev_group_relation UpdateUpgradeDevGroupRelation
//
// Update upgrade dev group relation information | 更新UpgradeDevGroupRelation信息
//
// Update upgrade dev group relation information | 更新UpgradeDevGroupRelation信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevGroupRelationInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeDevGroupRelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevGroupRelationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewUpdateUpgradeDevGroupRelationLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeDevGroupRelation(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
