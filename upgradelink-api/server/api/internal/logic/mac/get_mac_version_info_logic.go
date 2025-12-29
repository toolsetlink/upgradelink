package mac

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

type GetMacVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMacVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMacVersionInfoLogic {
	return &GetMacVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMacVersionInfoLogic) GetMacVersionInfo(req *types.GetMacVersionInfoReq) (resp *types.GetMacVersionInfoResp, err error) {

	// 请求参数效验
	if req.MacKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrMac4Msg, common.ErrMac4Docs)
	}
	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrMac4Msg, common.ErrMac4Docs)
	}

	var res types.GetMacVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	macInfo, err := l.svcCtx.ResourceCtx.GetMacInfoByKey(l.ctx, req.MacKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrMac2Msg, common.ErrMac2Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	macVersionInfo, err := l.svcCtx.ResourceCtx.GetMacVersionInfoByMacIdAndArchAndVersionCode(l.ctx, macInfo.Id, req.Arch, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrMac3Msg, common.ErrMac3Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetMacVersionInfoRespData{
		MacKey:      macInfo.Key,
		PackageName: macInfo.PackageName,
		VersionName: macVersionInfo.VersionName,
		VersionCode: macVersionInfo.VersionCode,
		Description: macVersionInfo.Description,
	}
	return &res, nil
}
