package storageprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/storageprovider"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /storage_provider/create storageprovider CreateStorageProvider
//
// Create storage provider information | 创建服务提供商
//
// Create storage provider information | 创建服务提供商
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StorageProviderInfo
//
// Responses:
//  200: BaseMsgResp

func CreateStorageProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StorageProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storageprovider.NewCreateStorageProviderLogic(r.Context(), svcCtx)
		resp, err := l.CreateStorageProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
