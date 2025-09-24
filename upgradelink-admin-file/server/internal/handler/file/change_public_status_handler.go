package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/file"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /file/status file ChangePublicStatus
//
// Change file public status | 改变文件公开状态
//
// Change file public status | 改变文件公开状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeReq
//
// Responses:
//  200: BaseMsgResp

func ChangePublicStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewChangePublicStatusLogic(r.Context(), svcCtx)
		resp, err := l.ChangePublicStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
