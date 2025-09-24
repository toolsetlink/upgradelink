package download

import (
	"context"
	"errors"
	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/config"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoDmgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoDmgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoDmgLogic {
	return &GetElectronDownloadInfoDmgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoDmgLogic) GetElectronDownloadInfoDmg(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// 请求参数效验
	if req.ElectronKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
	}

	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err101Msg, config.Err101Docs)
	}

	// 通过唯一标识 获取到对应的应用信息
	electronInfo, err := l.svcCtx.ResourceCtx.GetElectronInfoByKey(l.ctx, req.ElectronKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err102Msg, config.Err102Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	var electronVersionInfo *model.UpgradeElectronVersion
	// 判断是否传了 versionId， 如果传了，则直接选择数据
	if req.VersionId > 0 {
		electronVersionInfo, err = l.svcCtx.ResourceCtx.GetElectronVersionInfoById(l.ctx, req.VersionId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
		}
	} else {
		// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
		if req.VersionCode == 0 {
			electronVersionInfo, err = l.svcCtx.ResourceCtx.GetElectronVersionLastInfoByElectronIdAndPlatformAndArch(l.ctx, electronInfo.Id, req.Platform, req.Arch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
			}

		} else {
			electronVersionInfo, err = l.svcCtx.ResourceCtx.GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch(l.ctx, electronInfo.Id, req.VersionCode, req.Platform, req.Arch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
			}
		}
	}

	var cloudFileInfo *model.FmsCloudFiles
	// 通过文件信息 获取云文件地址
	cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, electronVersionInfo.InstallCloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
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
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 接口返回文件下载地址
	urlPath := ""
	urlPath = cloudFileInfo.Url

	return &urlPath, nil
}
