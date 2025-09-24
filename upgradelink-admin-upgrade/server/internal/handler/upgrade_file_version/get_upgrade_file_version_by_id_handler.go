package upgrade_file_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_file_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_file_version upgrade_file_version GetUpgradeFileVersionById
//
// Get upgrade file version by ID | 通过ID获取UpgradeFileVersion信息
//
// Get upgrade file version by ID | 通过ID获取UpgradeFileVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeFileVersionInfoResp

func GetUpgradeFileVersionByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file_version.NewGetUpgradeFileVersionByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeFileVersionById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
