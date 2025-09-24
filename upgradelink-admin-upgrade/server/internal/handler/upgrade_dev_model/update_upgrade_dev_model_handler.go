package upgrade_dev_model

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dev_model"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_dev_model/update upgrade_dev_model UpdateUpgradeDevModel
//
// Update upgrade dev model information | 更新UpgradeDevModel信息
//
// Update upgrade dev model information | 更新UpgradeDevModel信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevModelInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeDevModelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevModelInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_model.NewUpdateUpgradeDevModelLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeDevModel(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
