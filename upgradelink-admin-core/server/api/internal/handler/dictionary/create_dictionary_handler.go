package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/dictionary"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /dictionary/create dictionary CreateDictionary
//
// Create dictionary information | 创建字典
//
// Create dictionary information | 创建字典
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryInfo
//
// Responses:
//  200: BaseMsgResp

func CreateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewCreateDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.CreateDictionary(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
