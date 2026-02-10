package electron

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

type GetElectronVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronVersionInfoLogic {
	return &GetElectronVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronVersionInfoLogic) GetElectronVersionInfo(req *types.GetElectronVersionInfoReq) (resp *types.GetElectronVersionInfoResp, err error) {
	// 请求参数效验
	if req.ElectronKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "electron.paramError"), l.svcCtx.Trans.Trans(l.ctx, "electron.paramErrorDocs"))
	}

	if req.VersionName == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "electron.paramError"), l.svcCtx.Trans.Trans(l.ctx, "electron.paramErrorDocs"))
	}
	versionCode, err := common.SemVerToInt64(req.VersionName)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "electron.paramError"), l.svcCtx.Trans.Trans(l.ctx, "electron.paramErrorDocs"))
	}

	var res types.GetElectronVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	electronInfo, err := l.svcCtx.ResourceCtx.GetElectronInfoByKey(l.ctx, req.ElectronKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "electron.notFound"), l.svcCtx.Trans.Trans(l.ctx, "electron.notFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	// 获取版本信息
	electronVersionInfo, err := l.svcCtx.ResourceCtx.GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch(l.ctx, electronInfo.Id, versionCode, req.Platform, req.Arch)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFound"), l.svcCtx.Trans.Trans(l.ctx, "electron.versionNotFoundDocs"))
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetElectronVersionInfoRespData{
		ElectronKey: electronInfo.Key,
		Platform:    electronVersionInfo.Platform,
		Arch:        electronVersionInfo.Arch,
		VersionName: electronVersionInfo.VersionName,
		VersionCode: electronVersionInfo.VersionCode,
		Description: electronVersionInfo.Description,
	}

	return &res, nil
}
