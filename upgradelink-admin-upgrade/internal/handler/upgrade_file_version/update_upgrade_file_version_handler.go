package upgrade_file_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_file_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_file_version/update upgrade_file_version UpdateUpgradeFileVersion
//
// Update upgrade file version information | 更新UpgradeFileVersion信息
//
// Update upgrade file version information | 更新UpgradeFileVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileVersionInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeFileVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file_version.NewUpdateUpgradeFileVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeFileVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
