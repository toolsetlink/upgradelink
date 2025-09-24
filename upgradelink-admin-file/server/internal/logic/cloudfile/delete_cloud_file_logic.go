package cloudfile

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/server/ent/cloudfile"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCloudFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCloudFileLogic {
	return &DeleteCloudFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCloudFileLogic) DeleteCloudFile(req *types.UUIDsReq) (*types.BaseMsgResp, error) {
	if l.svcCtx.Config.UploadConf.DeleteFileWithCloud {
		data, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).
			WithStorageProviders().All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for _, v := range data {
			if client, ok := l.svcCtx.CloudStorage.CloudStorage[v.Edges.StorageProviders.Name]; ok {
				if !l.svcCtx.CloudStorage.ProviderData[v.Edges.StorageProviders.Name].UseCdn {
					keyData := strings.Split(v.URL, *l.svcCtx.CloudStorage.CloudStorage[v.Edges.StorageProviders.Name].Config.Endpoint)
					if len(keyData) != 2 {
						logx.Errorw("failed to find the key of the cloud file", logx.Field("data", req))
						return nil, errorx.NewCodeInternalError(i18n.Failed)
					}
					_, err = client.DeleteObject(&s3.DeleteObjectInput{
						Bucket: aws.String(l.svcCtx.CloudStorage.ProviderData[v.Edges.StorageProviders.Name].Bucket),
						Key:    aws.String(keyData[1]),
					})
				} else {
					keyData := strings.TrimPrefix(v.URL, l.svcCtx.CloudStorage.ProviderData[v.Edges.StorageProviders.Name].CdnUrl)
					if len(keyData) < 2 {
						logx.Errorw("failed to find the key of the cloud file", logx.Field("data", req))
						return nil, errorx.NewCodeInternalError(i18n.Failed)
					}
					_, err = client.DeleteObject(&s3.DeleteObjectInput{
						Bucket: aws.String(l.svcCtx.CloudStorage.ProviderData[v.Edges.StorageProviders.Name].Bucket),
						Key:    aws.String(keyData),
					})
				}
				if err != nil {
					logx.Errorw("failed to delete the cloud file", logx.Field("detail", err), logx.Field("data", req))
					return nil, errorx.NewCodeInternalError(i18n.Failed)
				}
			}
		}
	}

	_, err := l.svcCtx.DB.CloudFile.Delete().Where(cloudfile.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
