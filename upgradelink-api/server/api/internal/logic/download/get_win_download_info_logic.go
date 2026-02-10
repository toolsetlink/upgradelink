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

type GetWinDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWinDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWinDownloadInfoLogic {
	return &GetWinDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWinDownloadInfoLogic) GetWinDownloadInfo(req *types.GetWinDownloadInfoReq) (resp string, err error) {
	// 请求参数效验
	if req.WinKey == "" {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "win.paramError"), l.svcCtx.Trans.Trans(l.ctx, "win.paramErrorDocs"))
	}
	if req.VersionCode < 0 {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "win.paramError"), l.svcCtx.Trans.Trans(l.ctx, "win.paramErrorDocs"))
	}

	// 通过唯一标识 获取到对应的应用信息
	winInfo, err := l.svcCtx.ResourceCtx.GetWinInfoByKey(l.ctx, req.WinKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "win.notFound"), l.svcCtx.Trans.Trans(l.ctx, "win.notFoundDocs"))
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var winVersionInfo *model.UpgradeWinVersion

	// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
	if req.VersionCode == 0 {
		winVersionInfo, err = l.svcCtx.ResourceCtx.GetWinVersionLastInfoByWinId(l.ctx, winInfo.Id)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFoundDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else {
		winVersionInfo, err = l.svcCtx.ResourceCtx.GetWinVersionInfoByWinIdAndArchAndVersionCode(l.ctx, winInfo.Id, req.Arch, req.VersionCode)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFoundDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	}

	// 通过文件信息 获取云文件地址
	cloudFileInfo, err := l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, winVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           winInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              winInfo.Key,
		AppType:             "win",
		AppVersionId:        winVersionInfo.Id,
		AppVersionCode:      winVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      winVersionInfo.Arch,
		DownloadCloudFileId: cloudFileInfo.Id,
	})
	if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 接口返回文件下载地址
	urlPath := ""
	urlPath = cloudFileInfo.Url

	return urlPath, nil
}
