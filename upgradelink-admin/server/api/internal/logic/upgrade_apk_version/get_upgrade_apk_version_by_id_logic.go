package upgrade_apk_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeApkVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeApkVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeApkVersionByIdLogic {
	return &GetUpgradeApkVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeApkVersionByIdLogic) GetUpgradeApkVersionById(req *types.IDReq) (*types.UpgradeApkVersionInfoResp, error) {
	data, err := l.svcCtx.DB.UpgradeApkVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	apkData, err := l.svcCtx.DB.UpgradeApk.Get(l.ctx, data.ApkID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeApkVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RespUpgradeApkVersionInfo{
			Id:            &data.ID,
			ApkName:       &apkData.Name,
			CloudFileId:   &data.CloudFileID,
			CloudFileName: &apkData.Name,
			VersionName:   &data.VersionName,
			VersionCode:   &data.VersionCode,
			Description:   &data.Description,
			IsDel:         &data.IsDel,
			CreateAt:      pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:      pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
