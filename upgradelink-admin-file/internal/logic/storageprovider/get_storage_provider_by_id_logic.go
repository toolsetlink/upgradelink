package storageprovider

import (
	"context"

	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorageProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStorageProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorageProviderByIdLogic {
	return &GetStorageProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStorageProviderByIdLogic) GetStorageProviderById(req *types.IDReq) (*types.StorageProviderInfoResp, error) {
	data, err := l.svcCtx.DB.StorageProvider.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.StorageProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StorageProviderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			State:     &data.State,
			Name:      &data.Name,
			Bucket:    &data.Bucket,
			SecretId:  &data.SecretID,
			SecretKey: &data.SecretKey,
			Region:    &data.Region,
			IsDefault: &data.IsDefault,
			Folder:    &data.Folder,
			Endpoint:  &data.Endpoint,
			UseCdn:    &data.UseCdn,
			CdnUrl:    &data.CdnURL,
		},
	}, nil
}
