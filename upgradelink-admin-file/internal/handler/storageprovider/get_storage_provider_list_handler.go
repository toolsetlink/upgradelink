package storageprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/storageprovider"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /storage_provider/list storageprovider GetStorageProviderList
//
// Get storage provider list | 获取服务提供商列表
//
// Get storage provider list | 获取服务提供商列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StorageProviderListReq
//
// Responses:
//  200: StorageProviderListResp

func GetStorageProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StorageProviderListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storageprovider.NewGetStorageProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetStorageProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
