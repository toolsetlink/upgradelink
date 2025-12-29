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

type GetLnxDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLnxDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLnxDownloadInfoLogic {
	return &GetLnxDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLnxDownloadInfoLogic) GetLnxDownloadInfo(req *types.GetLnxDownloadInfoReq) (resp string, err error) {
	// 请求参数效验
	if req.LnxKey == "" {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx5Msg, common.ErrLnx5Docs)
	}
	if req.VersionCode < 0 {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx5Msg, common.ErrLnx5Docs)
	}

	// 通过唯一标识 获取到对应的应用信息
	lnxInfo, err := l.svcCtx.ResourceCtx.GetLnxInfoByKey(l.ctx, req.LnxKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx2Msg, common.ErrLnx2Docs)
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	var lnxVersionInfo *model.UpgradeLnxVersion
	// 判断是否传了 versionId， 如果传了，则直接选择数据
	if req.VersionId > 0 {
		lnxVersionInfo, err = l.svcCtx.ResourceCtx.GetLnxVersionInfoById(l.ctx, req.VersionId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx3Msg, common.ErrLnx3Docs)
		} else if err != nil {
			return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
		}
	} else {
		// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
		if req.VersionCode == 0 {
			lnxVersionInfo, err = l.svcCtx.ResourceCtx.GetLnxVersionLastInfoByLnxId(l.ctx, lnxInfo.Id)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx3Msg, common.ErrLnx3Docs)
			} else if err != nil {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
			}

		} else {
			lnxVersionInfo, err = l.svcCtx.ResourceCtx.GetLnxVersionInfoByLnxIdAndArchAndVersionCode(l.ctx, lnxInfo.Id, req.Arch, req.VersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx3Msg, common.ErrLnx3Docs)
			} else if err != nil {
				return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
			}
		}
	}

	// 通过文件信息 获取云文件地址
	cloudFileInfo, err := l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, lnxVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.ErrLnx4Msg, common.ErrLnx4Docs)
	} else if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.ErrLnx4Msg, common.ErrLnx4Docs)
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           lnxInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              lnxInfo.Key,
		AppType:             "lnx",
		AppVersionId:        lnxVersionInfo.Id,
		AppVersionCode:      lnxVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      lnxVersionInfo.Arch,
		DownloadCloudFileId: cloudFileInfo.Id,
	})
	if err != nil {
		return "", http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	// 接口返回文件下载地址
	urlPath := ""
	urlPath = cloudFileInfo.Url

	return urlPath, nil
}
