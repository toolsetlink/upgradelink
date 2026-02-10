package download

import (
	"context"
	"errors"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoLogic {
	return &GetElectronDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoLogic) GetElectronDownloadInfo(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// 请求参数效验
	if req.ElectronKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "electron.paramError"), l.svcCtx.Trans.Trans(l.ctx, "electron.paramErrorDocs"))
	}

	// 客户端windows 系统字段传的值为  win32
	if req.Platform == "win32" {
		req.Platform = "windows"
	}

	versionCode := int64(0)
	if req.VersionName != "" {
		versionCode, err = common.SemVerToInt64(req.VersionName)
		if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "electron.paramError"), l.svcCtx.Trans.Trans(l.ctx, "electron.paramErrorDocs"))
		}
	}

	// 通过唯一标识 获取到对应的应用信息
	electronInfo, err := l.svcCtx.ResourceCtx.GetElectronInfoByKey(l.ctx, req.ElectronKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "electron.notFound"), l.svcCtx.Trans.Trans(l.ctx, "electron.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var electronVersionInfo *model.UpgradeElectronVersion

	// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
	if versionCode == 0 {
		electronVersionInfo, err = l.svcCtx.ResourceCtx.GetElectronVersionLastInfoByElectronIdAndPlatformAndArch(l.ctx, electronInfo.Id, req.Platform, req.Arch)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFoundDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else {
		electronVersionInfo, err = l.svcCtx.ResourceCtx.GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch(l.ctx, electronInfo.Id, versionCode, req.Platform, req.Arch)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFoundDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	}

	var cloudFileInfo *model.FmsCloudFiles
	// 通过文件信息 获取云文件地址，并生成预签名
	// 判断是否传了 DownloadType， 1: 下载安装版本，2: 下载升级文件
	if req.DownloadType == 2 {
		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, electronVersionInfo.CloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	} else {
		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, electronVersionInfo.InstallCloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           electronInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              electronInfo.Key,
		AppType:             "electron",
		AppVersionId:        electronVersionInfo.Id,
		AppVersionCode:      electronVersionInfo.VersionCode,
		AppVersionPlatform:  electronVersionInfo.Platform,
		AppVersionTarget:    "",
		AppVersionArch:      electronVersionInfo.Arch,
		DownloadCloudFileId: cloudFileInfo.Id,
	})
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 接口返回文件下载地址
	urlPath := ""
	urlPath = cloudFileInfo.Url

	return &urlPath, nil
}
