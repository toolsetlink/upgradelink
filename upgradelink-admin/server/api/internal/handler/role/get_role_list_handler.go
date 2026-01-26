package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/role"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"
)

func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetRoleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
