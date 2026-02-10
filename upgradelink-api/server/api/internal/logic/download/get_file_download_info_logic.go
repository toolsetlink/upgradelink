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

type GetFileDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileDownloadInfoLogic {
	return &GetFileDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileDownloadInfoLogic) GetFileDownloadInfo(req *types.GetFileDownloadInfoReq) (resp *string, err error) {
	// 请求参数效验
	if req.FileKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "file.paramError"), l.svcCtx.Trans.Trans(l.ctx, "file.paramErrorDocs"))
	}
	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "file.paramError"), l.svcCtx.Trans.Trans(l.ctx, "file.paramErrorDocs"))
	}

	// 通过唯一标识 获取到对应的应用信息
	fileInfo, err := l.svcCtx.ResourceCtx.GetFileInfoByKey(l.ctx, req.FileKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "file.notFound"), l.svcCtx.Trans.Trans(l.ctx, "file.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var fileVersionInfo *model.UpgradeFileVersion

	// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
	if req.VersionCode == 0 {
		fileVersionInfo, err = l.svcCtx.ResourceCtx.GetFileVersionLastInfoByFileId(l.ctx, fileInfo.Id)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "file.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "file.versionNotFoundDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else {
		fileVersionInfo, err = l.svcCtx.ResourceCtx.GetFileVersionInfoByFileIdAndVersionCode(l.ctx, fileInfo.Id, req.VersionCode)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "file.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "file.versionNotFoundDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	}

	// 通过文件信息 获取云文件地址
	cloudFileInfo, err := l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, fileVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           fileInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              fileInfo.Key,
		AppType:             "file",
		AppVersionId:        fileVersionInfo.Id,
		AppVersionCode:      fileVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      "",
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
