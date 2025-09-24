package cloudfile

import (
	"context"

	"upgradelink-admin-file/ent/cloudfile"

	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCloudFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileByIdLogic {
	return &GetCloudFileByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCloudFileByIdLogic) GetCloudFileById(req *types.UUIDReq) (*types.CloudFileInfoResp, error) {
	data, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.IDEQ(uuidx.ParseUUIDString(req.Id))).WithStorageProviders().
		First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			State:      &data.State,
			Name:       &data.Name,
			Url:        &data.URL,
			Size:       &data.Size,
			FileType:   &data.FileType,
			UserId:     &data.UserID,
			ProviderId: &data.Edges.StorageProviders.ID,
		},
	}, nil
}
