package url

import (
	"context"
	"errors"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource/model"

	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlVersionInfoLogic {
	return &GetUrlVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlVersionInfoLogic) GetUrlVersionInfo(req *types.GetUrlVersionInfoReq) (resp *types.GetUrlVersionInfoResp, err error) {
	// 请求参数效验
	if req.UrlKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "url.paramError"), l.svcCtx.Trans.Trans(l.ctx, "url.paramErrorDocs"))
	}
	if req.VersionCode == 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "url.paramError"), l.svcCtx.Trans.Trans(l.ctx, "url.paramErrorDocs"))
	}

	var res types.GetUrlVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	urlInfo, err := l.svcCtx.ResourceCtx.GetUrlInfoByKey(l.ctx, req.UrlKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.notFound"), l.svcCtx.Trans.Trans(l.ctx, "url.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	urlVersionInfo, err := l.svcCtx.ResourceCtx.GetUrlVersionInfoByUrlIdAndVersionCode(l.ctx, urlInfo.Id, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "url.versionNotFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetUrlVersionInfoRespData{
		UrlKey:      urlInfo.Key,
		VersionName: urlVersionInfo.VersionName,
		VersionCode: urlVersionInfo.VersionCode,
		Description: urlVersionInfo.Description,
	}

	return &res, nil
}
