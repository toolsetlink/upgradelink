package cloudfile

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/filex"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileByIdLogic {
	return &GetCloudFileByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCloudFileByIdLogic) GetCloudFileById(req *types.UUIDReq) (resp *types.CloudFileInfoResp, err error) {
	result, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, req.Id)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	fileType := filex.ConvertUint8ToFileType(result.FileType)

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(result.ID),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			},
			State:    pointy.GetPointer(result.State),
			Name:     &result.Name,
			Url:      &result.URL,
			Size:     pointy.GetPointer(result.Size),
			Md5:      &result.Md5,
			FileType: &fileType,
			UserId:   &result.UserID,
		},
	}, nil

	return
}
