package tauri

import (
	"context"
	"errors"
	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource/model"

	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTauriVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTauriVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTauriVersionInfoLogic {
	return &GetTauriVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTauriVersionInfoLogic) GetTauriVersionInfo(req *types.GetTauriVersionInfoReq) (resp *types.GetTauriVersionInfoResp, err error) {
	// 请求参数效验
	if req.TauriKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "tauri.paramError"), l.svcCtx.Trans.Trans(l.ctx, "tauri.paramErrorDocs"))
	}

	if req.VersionName == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "tauri.paramError"), l.svcCtx.Trans.Trans(l.ctx, "tauri.paramErrorDocs"))
	}

	versionCode, err := common.SemVerToInt64(req.VersionName)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "common.versionFormatError"), l.svcCtx.Trans.Trans(l.ctx, "common.versionFormatErrorDocs"))
	}

	var res types.GetTauriVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	tauriInfo, err := l.svcCtx.ResourceCtx.GetTauriInfoByKey(l.ctx, req.TauriKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "tauri.notFound"), l.svcCtx.Trans.Trans(l.ctx, "tauri.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	// 获取版本信息
	tauriVersionInfo, err := l.svcCtx.ResourceCtx.GetTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch(l.ctx, tauriInfo.Id, versionCode, req.Target, req.Arch)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "tauri.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "tauri.versionNotFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetTauriVersionInfoRespData{
		TauriKey:    tauriInfo.Key,
		VersionName: tauriVersionInfo.VersionName,
		VersionCode: tauriVersionInfo.VersionCode,
		Description: tauriVersionInfo.Description,
		Target:      tauriVersionInfo.Target,
		Arch:        tauriVersionInfo.Arch,
	}

	return &res, nil
}
