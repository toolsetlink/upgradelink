package storageprovider

import (
	"context"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/cloud"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStorageProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStorageProviderLogic {
	return &UpdateStorageProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStorageProviderLogic) UpdateStorageProvider(req *types.StorageProviderInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.StorageProvider.UpdateOneID(*req.Id).
		SetNotNilState(req.State).
		SetNotNilName(req.Name).
		SetNotNilBucket(req.Bucket).
		SetNotNilSecretID(req.SecretId).
		SetNotNilSecretKey(req.SecretKey).
		SetNotNilRegion(req.Region).
		SetNotNilIsDefault(req.IsDefault).
		SetNotNilFolder(req.Folder).
		SetNotNilEndpoint(req.Endpoint).
		SetNotNilUseCdn(req.UseCdn).
		SetNotNilCdnURL(req.CdnUrl).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudStorage = cloud.NewCloudServiceGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
