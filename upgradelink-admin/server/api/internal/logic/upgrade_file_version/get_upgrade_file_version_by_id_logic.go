package upgrade_file_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeFileVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeFileVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeFileVersionByIdLogic {
	return &GetUpgradeFileVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeFileVersionByIdLogic) GetUpgradeFileVersionById(req *types.IDReq) (*types.UpgradeFileVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeFileVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	fileData, err := l.svcCtx.DB.UpgradeFile.Get(l.ctx, data.FileID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeFileVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RespUpgradeFileVersionInfo{
			Id:            &data.ID,
			FileName:      &fileData.Name,
			CloudFileId:   &data.CloudFileID,
			CloudFileName: &fileData.Name,
			VersionName:   &data.VersionName,
			VersionCode:   &data.VersionCode,
			Description:   &data.Description,
			IsDel:         &data.IsDel,
			CreateAt:      pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:      pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
