package storageprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/storageprovider"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /storage_provider/update storageprovider UpdateStorageProvider
//
// Update storage provider information | 更新服务提供商
//
// Update storage provider information | 更新服务提供商
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StorageProviderInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateStorageProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StorageProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storageprovider.NewUpdateStorageProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStorageProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
