package win

import (
	"context"
	"errors"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource/model"

	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWinVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWinVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWinVersionInfoLogic {
	return &GetWinVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWinVersionInfoLogic) GetWinVersionInfo(req *types.GetWinVersionInfoReq) (resp *types.GetWinVersionInfoResp, err error) {

	// 请求参数效验
	if req.WinKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "win.paramError"), l.svcCtx.Trans.Trans(l.ctx, "win.paramErrorDocs"))
	}
	if req.VersionCode == 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "win.paramError"), l.svcCtx.Trans.Trans(l.ctx, "win.paramErrorDocs"))
	}

	var res types.GetWinVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	winInfo, err := l.svcCtx.ResourceCtx.GetWinInfoByKey(l.ctx, req.WinKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "win.notFound"), l.svcCtx.Trans.Trans(l.ctx, "win.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	winVersionInfo, err := l.svcCtx.ResourceCtx.GetWinVersionInfoByWinIdAndArchAndVersionCode(l.ctx, winInfo.Id, req.Arch, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "win.versionNotFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetWinVersionInfoRespData{
		WinKey:      winInfo.Key,
		PackageName: winInfo.PackageName,
		VersionName: winVersionInfo.VersionName,
		VersionCode: winVersionInfo.VersionCode,
		Description: winVersionInfo.Description,
	}
	return &res, nil
}
