package upgrade_dev_model

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dev_model"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_dev_model/create upgrade_dev_model CreateUpgradeDevModel
//
// Create upgrade dev model information | 创建UpgradeDevModel信息
//
// Create upgrade dev model information | 创建UpgradeDevModel信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevModelInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeDevModelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevModelInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_model.NewCreateUpgradeDevModelLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeDevModel(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
