package download

import (
	"context"
	"errors"
	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/resource"

	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/config"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTauriDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTauriDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTauriDownloadInfoLogic {
	return &GetTauriDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTauriDownloadInfoLogic) GetTauriDownloadInfo(req *types.GetTauriDownloadInfoReq) (resp *string, err error) {
	// 请求参数效验
	if req.TauriKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
	}

	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err101Msg, config.Err101Docs)
	}

	// 通过唯一标识 获取到对应的应用信息
	tauriInfo, err := l.svcCtx.ResourceCtx.GetTauriInfoByKey(l.ctx, req.TauriKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err102Msg, config.Err102Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	var tauriVersionInfo *model.UpgradeTauriVersion
	// 判断是否传了 versionId， 如果传了，则直接选择数据
	if req.VersionId > 0 {
		tauriVersionInfo, err = l.svcCtx.ResourceCtx.GetTauriVersionInfoById(l.ctx, req.VersionId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
		}
	} else {
		// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
		if req.VersionCode == 0 {
			tauriVersionInfo, err = l.svcCtx.ResourceCtx.GetTauriVersionLastInfoByTauriIdAndTargetAndArch(l.ctx, tauriInfo.Id, req.Target, req.Arch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
			}

		} else {
			tauriVersionInfo, err = l.svcCtx.ResourceCtx.GetTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch(l.ctx, tauriInfo.Id, req.VersionCode, req.Target, req.Arch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err103Msg, config.Err103Docs)
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err100Msg, config.Err100Docs)
			}
		}
	}

	var cloudFileInfo *model.FmsCloudFiles
	// 通过文件信息 获取云文件地址，并生成预签名
	// 判断是否传了 DownloadType， 1: 下载安装版本，2: 下载升级文件
	if req.DownloadType == 2 {
		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, tauriVersionInfo.CloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
		}
	} else {
		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, tauriVersionInfo.InstallCloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
		}
	}

	// 通过文件信息 获取云文件地址
	cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, tauriVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 插入app 文件下载日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           tauriInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              tauriInfo.Key,
		AppType:             "tauri",
		AppVersionId:        tauriVersionInfo.Id,
		AppVersionCode:      tauriVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    tauriVersionInfo.Target,
		AppVersionArch:      tauriVersionInfo.Arch,
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
