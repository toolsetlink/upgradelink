package cloudfile

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/server/ent"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCloudFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCloudFileLogic {
	return &CreateCloudFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCloudFileLogic) CreateCloudFile(req *types.CloudFileInfo) (*types.BaseMsgResp, error) {
	// check storage provider exist
	_, err := l.svcCtx.DB.StorageProvider.Get(l.ctx, *req.ProviderId)
	switch {
	case ent.IsNotFound(err):
		return nil, errorx.NewCodeInvalidArgumentError("storage_provider.StorageProviderNotExist")
	case err != nil:
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	query := l.svcCtx.DB.CloudFile.Create().
		SetNotNilState(req.State).
		SetNotNilName(req.Name).
		SetNotNilURL(req.Url).
		SetNotNilSize(req.Size).
		SetNotNilFileType(req.FileType).
		SetNotNilUserID(req.UserId)

	if req.ProviderId != nil {
		query = query.SetStorageProvidersID(*req.ProviderId)
	}

	if req.TagIds != nil {
		query = query.AddTagIDs(req.TagIds...)
	}

	_, err = query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
