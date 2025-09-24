package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group_relation upgrade_dev_group_relation GetUpgradeDevGroupRelationById
//
// Get upgrade dev group relation by ID | 通过ID获取UpgradeDevGroupRelation信息
//
// Get upgrade dev group relation by ID | 通过ID获取UpgradeDevGroupRelation信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeDevGroupRelationInfoResp

func GetUpgradeDevGroupRelationByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewGetUpgradeDevGroupRelationByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDevGroupRelationById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
