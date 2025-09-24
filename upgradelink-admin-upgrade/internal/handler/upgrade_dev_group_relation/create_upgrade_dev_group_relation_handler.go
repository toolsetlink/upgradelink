package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group_relation/create upgrade_dev_group_relation CreateUpgradeDevGroupRelation
//
// Create upgrade dev group relation information | 创建UpgradeDevGroupRelation信息
//
// Create upgrade dev group relation information | 创建UpgradeDevGroupRelation信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevGroupRelationInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeDevGroupRelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevGroupRelationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewCreateUpgradeDevGroupRelationLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeDevGroupRelation(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
