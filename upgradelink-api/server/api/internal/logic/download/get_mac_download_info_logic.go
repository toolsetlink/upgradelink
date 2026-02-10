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

type GetMacDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMacDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMacDownloadInfoLogic {
	return &GetMacDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMacDownloadInfoLogic) GetMacDownloadInfo(req *types.GetMacDownloadInfoReq) (resp string, err error) {
	// 请求参数效验
	if req.MacKey == "" {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "mac.paramError"), l.svcCtx.Trans.Trans(l.ctx, "mac.paramErrorDocs"))
	}
	if req.VersionCode < 0 {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "mac.paramError"), l.svcCtx.Trans.Trans(l.ctx, "mac.paramErrorDocs"))
	}

	// 通过唯一标识 获取到对应的应用信息
	macInfo, err := l.svcCtx.ResourceCtx.GetMacInfoByKey(l.ctx, req.MacKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "mac.notFound"), l.svcCtx.Trans.Trans(l.ctx, "mac.notFoundDocs"))
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var macVersionInfo *model.UpgradeMacVersion
	// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
	if req.VersionCode == 0 {
		macVersionInfo, err = l.svcCtx.ResourceCtx.GetMacVersionLastInfoByMacId(l.ctx, macInfo.Id)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "mac.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "mac.versionNotFoundDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}

	} else {
		macVersionInfo, err = l.svcCtx.ResourceCtx.GetMacVersionInfoByMacIdAndArchAndVersionCode(l.ctx, macInfo.Id, req.Arch, req.VersionCode)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "mac.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "mac.versionNotFoundDocs"))
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	}

	// 通过文件信息 获取云文件地址
	cloudFileInfo, err := l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, macVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExist"), l.svcCtx.Trans.Trans(l.ctx, "common.targetNotExistDocs"))
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           macInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              macInfo.Key,
		AppType:             "mac",
		AppVersionId:        macVersionInfo.Id,
		AppVersionCode:      macVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      macVersionInfo.Arch,
		DownloadCloudFileId: cloudFileInfo.Id,
	})
	if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"))
	}

	// 接口返回文件下载地址
	urlPath := ""
	urlPath = cloudFileInfo.Url

	return urlPath, nil
}
