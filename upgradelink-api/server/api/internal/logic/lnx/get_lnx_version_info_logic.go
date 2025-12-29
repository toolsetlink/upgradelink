package lnx

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

type GetLnxVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLnxVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLnxVersionInfoLogic {
	return &GetLnxVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLnxVersionInfoLogic) GetLnxVersionInfo(req *types.GetLnxVersionInfoReq) (resp *types.GetLnxVersionInfoResp, err error) {

	// 请求参数效验
	if req.LnxKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx4Msg, common.ErrLnx4Docs)
	}
	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx4Msg, common.ErrLnx4Docs)
	}

	var res types.GetLnxVersionInfoResp

	// 通过唯一标识 获取到对应的应用信息
	lnxInfo, err := l.svcCtx.ResourceCtx.GetLnxInfoByKey(l.ctx, req.LnxKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx2Msg, common.ErrLnx2Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	lnxVersionInfo, err := l.svcCtx.ResourceCtx.GetLnxVersionInfoByLnxIdAndArchAndVersionCode(l.ctx, lnxInfo.Id, req.Arch, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx3Msg, common.ErrLnx3Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	res.Code = 200
	res.Msg = ""
	res.Data = types.GetLnxVersionInfoRespData{
		LnxKey:      lnxInfo.Key,
		PackageName: lnxInfo.PackageName,
		VersionName: lnxVersionInfo.VersionName,
		VersionCode: lnxVersionInfo.VersionCode,
		Description: lnxVersionInfo.Description,
	}
	return &res, nil
}
