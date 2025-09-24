package upgrade_dev_group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_dev_group"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_dev_group/update upgrade_dev_group UpdateUpgradeDevGroup
//
// Update upgrade dev group information | 更新UpgradeDevGroup信息
//
// Update upgrade dev group information | 更新UpgradeDevGroup信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevGroupInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeDevGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevGroupInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group.NewUpdateUpgradeDevGroupLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeDevGroup(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
