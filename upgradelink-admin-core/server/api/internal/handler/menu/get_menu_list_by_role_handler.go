package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/menu"
	"upgradelink-admin-core/server/api/internal/svc"
)

// swagger:route get /menu/role/list menu GetMenuListByRole
//
// Get menu list by role | 获取菜单列表
//
// Get menu list by role | 获取菜单列表
//
// Responses:
//  200: MenuListResp

func GetMenuListByRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuListByRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuListByRole()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
