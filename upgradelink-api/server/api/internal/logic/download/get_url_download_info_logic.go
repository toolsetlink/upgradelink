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

type GetUrlDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlDownloadInfoLogic {
	return &GetUrlDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlDownloadInfoLogic) GetUrlDownloadInfo(req *types.GetUrlDownloadInfoReq) (resp *string, err error) {
	// 请求参数效验
	if req.UrlKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "url.paramError"), l.svcCtx.Trans.Trans(l.ctx, "url.paramErrorDocs"))
	}
	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "url.paramError"), l.svcCtx.Trans.Trans(l.ctx, "url.paramErrorDocs"))
	}

	// 通过唯一标识 获取到对应的应用信息
	urlInfo, err := l.svcCtx.ResourceCtx.GetUrlInfoByKey(l.ctx, req.UrlKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.notFound"), l.svcCtx.Trans.Trans(l.ctx, "url.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	var urlVersionInfo *model.UpgradeUrlVersion
	// 判断是否传了 versionId， 如果传了，则直接选择数据
	if req.VersionId > 0 {
		urlVersionInfo, err = l.svcCtx.ResourceCtx.GetUrlVersionInfoById(l.ctx, req.VersionId)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFoundDocs"))
		} else if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
		}
	} else {
		// 判断是否固定了版本号，如果没有固定 则获取详细的版本信息
		if req.VersionCode == 0 {
			urlVersionInfo, err = l.svcCtx.ResourceCtx.GetUrlVersionLastInfoByUrlId(l.ctx, urlInfo.Id)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFoundDocs"))
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
			}
		} else {
			urlVersionInfo, err = l.svcCtx.ResourceCtx.GetUrlVersionInfoByUrlIdAndVersionCode(l.ctx, urlInfo.Id, req.VersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFoundDocs"))
			} else if err != nil {
				return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
			}
		}
	}

	// 插入日志表
	_, err = l.svcCtx.ResourceCtx.AddAppDownloadReportLog(l.ctx, resource.AddAppDownloadReportLogReq{
		CompanyId:           urlInfo.CompanyId,
		Timestamp:           common.GetCurrentTime(),
		AppKey:              urlInfo.Key,
		AppType:             "url",
		AppVersionId:        urlVersionInfo.Id,
		AppVersionCode:      urlVersionInfo.VersionCode,
		AppVersionPlatform:  "",
		AppVersionTarget:    "",
		AppVersionArch:      "",
		DownloadCloudFileId: "",
	})
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	return &urlVersionInfo.UrlPath, nil
}
