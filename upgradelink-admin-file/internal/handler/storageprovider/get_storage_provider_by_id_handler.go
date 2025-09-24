package storageprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/storageprovider"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /storage_provider storageprovider GetStorageProviderById
//
// Get storage provider by ID | 通过ID获取服务提供商
//
// Get storage provider by ID | 通过ID获取服务提供商
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: StorageProviderInfoResp

func GetStorageProviderByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storageprovider.NewGetStorageProviderByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetStorageProviderById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
