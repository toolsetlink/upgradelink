package storageprovider

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/ent/cloudfile"
	"upgradelink-admin-file/ent/storageprovider"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/cloud"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStorageProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStorageProviderLogic {
	return &DeleteStorageProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStorageProviderLogic) DeleteStorageProvider(req *types.IDsReq) (*types.BaseMsgResp, error) {
	check, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.HasStorageProvidersWith(storageprovider.IDIn(req.Ids...))).
		Count(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if check != 0 {
		return nil, errorx.NewCodeInvalidArgumentError("storage_provider.hasFileError")
	}

	_, err = l.svcCtx.DB.StorageProvider.Delete().Where(storageprovider.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudStorage = cloud.NewCloudServiceGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
