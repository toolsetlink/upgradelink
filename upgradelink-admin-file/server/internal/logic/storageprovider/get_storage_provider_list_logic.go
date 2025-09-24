package storageprovider

import (
	"context"

	"upgradelink-admin-file/server/ent/predicate"
	"upgradelink-admin-file/server/ent/storageprovider"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorageProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStorageProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorageProviderListLogic {
	return &GetStorageProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStorageProviderListLogic) GetStorageProviderList(req *types.StorageProviderListReq) (*types.StorageProviderListResp, error) {
	var predicates []predicate.StorageProvider
	if req.Name != nil {
		predicates = append(predicates, storageprovider.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.StorageProvider.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.StorageProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.StorageProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				State:     &v.State,
				Name:      &v.Name,
				Bucket:    &v.Bucket,
				SecretId:  &v.SecretID,
				SecretKey: &v.SecretKey,
				Region:    &v.Region,
				IsDefault: &v.IsDefault,
				Folder:    &v.Folder,
				Endpoint:  &v.Endpoint,
				UseCdn:    &v.UseCdn,
				CdnUrl:    &v.CdnURL,
			})
	}

	return resp, nil
}
