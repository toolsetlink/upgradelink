package app

import (
	"context"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/config"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportLogic) Report(req *types.ReportReq) (resp *types.ReportResp, err error) {

	// 获取公司 id
	companyId, err := l.svcCtx.ResourceCtx.GetCompanyIdByAppKey(l.ctx, req.AppKey)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	timestamp, err := common.ParseRFC3339ToTime(req.Timestamp)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 获取应用版本 id
	appVersionId, err := l.svcCtx.ResourceCtx.GetAppVersionIdByReport(l.ctx, resource.GetAppVersionIdByReportReq{
		AppKey:           req.AppKey,
		AppVersionCode:   req.EventData.VersionCode,
		DevModelKey:      req.EventData.DevModelKey,
		DevKey:           req.EventData.DevKey,
		AppVersionTarget: req.EventData.Target,
		AppVersionArch:   req.EventData.Arch,
	})

	// 获取应用版本类型
	appType, err := l.svcCtx.ResourceCtx.GetAppTypeByAppKey(l.ctx, req.AppKey)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	var res types.ReportResp

	// app 启动事件 开启事件
	if req.EventType == "app_start" {

		launchTime, err := common.ParseRFC3339ToTime(req.EventData.LaunchTime)
		if err != nil {
			return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
		}

		_, err = l.svcCtx.ResourceCtx.AddAppStartReportLog(l.ctx, resource.AddAppStartReportLogReq{
			CompanyId:        companyId,
			AppKey:           req.AppKey,
			AppType:          appType,
			Timestamp:        *timestamp,
			AppVersionId:     appVersionId,
			AppVersionCode:   req.EventData.VersionCode,
			DevModelKey:      req.EventData.DevModelKey,
			DevKey:           req.EventData.DevKey,
			AppVersionTarget: req.EventData.Target,
			AppVersionArch:   req.EventData.Arch,
			LaunchTime:       *launchTime,
		})
	}

	//app_upgrade_download 应用升级-下载事件
	if req.EventType == "app_upgrade_download" {

		_, err = l.svcCtx.ResourceCtx.AddAppUpgradeDownloadReportLog(l.ctx, resource.AddAppUpgradeDownloadReportLogReq{
			CompanyId:           companyId,
			AppKey:              req.AppKey,
			AppType:             appType,
			Timestamp:           *timestamp,
			AppVersionId:        appVersionId,
			AppVersionCode:      req.EventData.VersionCode,
			DevModelKey:         req.EventData.DevModelKey,
			DevKey:              req.EventData.DevKey,
			AppVersionTarget:    req.EventData.Target,
			AppVersionArch:      req.EventData.Arch,
			DownloadVersionCode: req.EventData.DownloadVersionCode,
			Code:                req.EventData.Code,
		})
	}

	//app_upgrade_upgrade 应用升级-升级事件
	if req.EventType == "app_upgrade_upgrade" {
		_, err = l.svcCtx.ResourceCtx.AddAppUpgradeUpgradeReportLog(l.ctx, resource.AddAppUpgradeUpgradeReportLogReq{
			CompanyId:          companyId,
			AppKey:             req.AppKey,
			AppType:            appType,
			Timestamp:          *timestamp,
			AppVersionId:       appVersionId,
			AppVersionCode:     req.EventData.VersionCode,
			DevModelKey:        req.EventData.DevModelKey,
			DevKey:             req.EventData.DevKey,
			AppVersionTarget:   req.EventData.Target,
			AppVersionArch:     req.EventData.Arch,
			UpgradeVersionCode: req.EventData.UpgradeVersionCode,
			Code:               req.EventData.Code,
		})
	}

	res.Code = 0
	res.Msg = "上报成功"
	return &res, nil
}
