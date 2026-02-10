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

type GetApkDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApkDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApkDownloadInfoLogic {
	return &GetApkDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApkDownloadInfoLogic) GetApkDownloadInfo(req *types.GetApkDownloadInfoReq) (resp string, err error) {
	// 请求参数效验
	if req.ApkKey == "" {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "apk.paramError"), l.svcCtx.Trans.Trans(l.ctx, "apk.paramErrorDocs"))
	}
	if req.VersionCode < 0 {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "apk.paramError"), l.svcCtx.Trans.Trans(l.ctx, "apk.paramErrorDocs"))
	}

	// 通过唯一标识 获取到对应的应用信息
	apkInfo, err := l.svcCtx.ResourceCtx.GetApkInfoByKey(l.ctx, req.ApkKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "apk.notFound"), l.svcCtx.Trans.Trans(l.ctx, "apk.notFoundDocs"))
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var apkVersionInfo *model.UpgradeApkVersion
	var cloudFileInfo *model.FmsCloudFiles

	// 判断当前请求类型为下载全量包，还是补丁包
	if req.DownloadType == 1 {

		// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
		if req.VersionCode == 0 {
			apkVersionInfo, err = l.svcCtx.ResourceCtx.GetApkVersionLastInfoByApkId(l.ctx, apkInfo.Id)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "apk.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "apk.versionNotFoundDocs"))
			} else if err != nil {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
			}

		} else {
			apkVersionInfo, err = l.svcCtx.ResourceCtx.GetApkVersionInfoByApkIdAndVersionCode(l.ctx, apkInfo.Id, req.VersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "apk.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "apk.versionNotFoundDocs"))
			} else if err != nil {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
			}
		}

		// 通过文件信息 获取云文件地址
		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, apkVersionInfo.CloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else if req.DownloadType == 2 {
		// 通过 cloudFileId 获取到对应的补丁包信息
		if req.CloudFileId == "" {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "apk.paramError"), l.svcCtx.Trans.Trans(l.ctx, "apk.paramErrorDocs"))
		}

		// 通过传过来的 CloudFileId 获取到patch 包归属的 apk patch信息
		apkPatchInfo, err := l.svcCtx.ResourceCtx.GetPatchInfoByCloudFileId(l.ctx, apkInfo.Id, req.CloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

		// 通过apk patch信息 获取到对应 apk 版本信息
		apkVersionInfo, err = l.svcCtx.ResourceCtx.GetApkVersionInfoById(l.ctx, apkPatchInfo.HighApkVersionId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

		cloudFileInfo, err = l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, apkPatchInfo.CloudFileId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "apk.paramError"), l.svcCtx.Trans.Trans(l.ctx, "apk.paramErrorDocs"))
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           apkInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              apkInfo.Key,
		AppType:             "apk",
		AppVersionId:        apkVersionInfo.Id,
		AppVersionCode:      apkVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      "",
		DownloadCloudFileId: cloudFileInfo.Id,
	})
	if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 接口返回文件下载地址
	urlPath := cloudFileInfo.Url

	return urlPath, nil
}
