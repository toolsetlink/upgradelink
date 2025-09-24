package user

import (
	"net/http"

	"upgradelink-admin-core/server/api/internal/logic/user"
	"upgradelink-admin-core/server/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /user/info user GetUserInfo
//
// Get user basic information | 获取用户基本信息
//
// Get user basic information | 获取用户基本信息
//
// Responses:
//  200: UserBaseIDInfoResp

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
