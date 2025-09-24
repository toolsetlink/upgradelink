package upgrade_file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_file"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_file/create upgrade_file CreateUpgradeFile
//
// Create upgrade file information | 创建UpgradeFile信息
//
// Create upgrade file information | 创建UpgradeFile信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file.NewCreateUpgradeFileLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
