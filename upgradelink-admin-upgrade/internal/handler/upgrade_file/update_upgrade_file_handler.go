package upgrade_file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_file"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_file/update upgrade_file UpdateUpgradeFile
//
// Update upgrade file information | 更新UpgradeFile信息
//
// Update upgrade file information | 更新UpgradeFile信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file.NewUpdateUpgradeFileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
