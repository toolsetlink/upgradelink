package upgrade_dev_model

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_dev_model"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_dev_model/list upgrade_dev_model GetUpgradeDevModelList
//
// Get upgrade dev model list | 获取UpgradeDevModel信息列表
//
// Get upgrade dev model list | 获取UpgradeDevModel信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeDevModelListReq
//
// Responses:
//  200: UpgradeDevModelListResp

func GetUpgradeDevModelListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeDevModelListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_model.NewGetUpgradeDevModelListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeDevModelList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
