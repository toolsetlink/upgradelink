package upgrade_dashboard

import (
	"context"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeapk"
	"upgradelink-admin-upgrade/server/ent/upgradeelectron"
	"upgradelink-admin-upgrade/server/ent/upgradefile"
	"upgradelink-admin-upgrade/server/ent/upgradetauri"
	"upgradelink-admin-upgrade/server/ent/upgradeurl"
	"upgradelink-admin-upgrade/server/internal/resource"

	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/userctx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDashboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpgradeDashboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDashboardLogic {
	return &GetUpgradeDashboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetUpgradeDashboardLogic) GetUpgradeDashboard() (resp *types.UpgradeDashboardResp, err error) {

	newResp := &types.UpgradeDashboardResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.UpgradeDashboardInfo{},
	}
	// 获取公司 id
	userId, err := userctx.GetUserIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	userData, err := l.svcCtx.DB.SysUser.Get(l.ctx, userId)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "GetUpgradeDashboard")
	}
	companyID := int64(userData.CompanyID)
	if companyID == 0 {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "GetUpgradeDashboard")
	}

	// 获取昨日应用下载量
	yesterdayAppDownloadCount, err := resource.GetYesterdayDownloadCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.YesterdayDownloadCount = pointy.GetPointer(yesterdayAppDownloadCount)

	// 获取总应用下载量
	totalAppDownloadCount, err := resource.GetDownloadCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.TotalDownloadCount = pointy.GetPointer(totalAppDownloadCount)

	// 获取昨日应用获取升级次数
	yesterdayAppGetStrategyCount, err := resource.GetYesterdayAppGetStrategyCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.YesterdayAppGetStrategyCount = pointy.GetPointer(yesterdayAppGetStrategyCount)

	// 获取总应用获取升级次数
	totalAppGetStrategyCount, err := resource.GetAppGetStrategyCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.TotalAppGetStrategyCount = pointy.GetPointer(totalAppGetStrategyCount)

	// 获取昨日应用升级次数
	yesterdayAppUpgradeCount, err := resource.GetYesterdayAppUpgradeCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.YesterdayAppUpgradeCount = pointy.GetPointer(yesterdayAppUpgradeCount)

	// 获取总应用升级次数
	totalAppUpgradeCount, err := resource.GetAppUpgradeCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.TotalAppUpgradeCount = pointy.GetPointer(totalAppUpgradeCount)

	// 获取昨日应用启动次数
	yesterdayAppStartCount, err := resource.GetYesterdayAppStartCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.YesterdayAppStartCount = pointy.GetPointer(yesterdayAppStartCount)

	// 获取总应用启动次数
	totalAppStartCount, err := resource.GetAppStartCount(l.ctx, l.svcCtx, companyID)
	newResp.Data.TotalAppStartCount = pointy.GetPointer(totalAppStartCount)

	// 统计 近 7 天的数据

	var trafficUsageCount7Days types.TrafficUsageCount7Day
	trafficUsageCount7Days.Series = make([]types.TrafficUsageCount7DaySeriesItem, 0)
	trafficUsageCount7Days.TimeData = make([]string, 0)

	var downloadCount7Days types.DownloadCount7Day
	downloadCount7Days.Series = make([]types.DownloadCount7DaySeriesItem, 0)
	downloadCount7Days.TimeData = make([]string, 0)

	var appGetStrategyCount7Days types.AppGetStrategyCount7Day
	appGetStrategyCount7Days.Series = make([]types.AppGetStrategyCount7DaySeriesItem, 0)
	appGetStrategyCount7Days.TimeData = make([]string, 0)

	var appUpgradeCount7Days types.AppUpgradeCount7Day
	appUpgradeCount7Days.Series = make([]types.AppUpgradeCount7DaySeriesItem, 0)
	appUpgradeCount7Days.TimeData = make([]string, 0)

	var appStartCount7Days types.AppStartCount7Day
	appStartCount7Days.Series = make([]types.AppStartCount7DaySeriesItem, 0)
	appStartCount7Days.TimeData = make([]string, 0)

	// 获取公司下的所有url应用
	var urlPredicates []predicate.UpgradeUrl
	urlPredicates = append(urlPredicates, upgradeurl.CompanyIDEQ(int(companyID)))
	urlPredicates = append(urlPredicates, upgradeurl.IsDelEQ(0))
	urlList, err := l.svcCtx.DB.UpgradeUrl.Query().Where(urlPredicates...).Page(l.ctx, 1, 100)
	if err != nil {
		return nil, err
	}
	for infoKey, info := range urlList.List {
		// 获取应用近7天下载次数
		weeklyDownloadCountList, err := resource.GetWeeklyDownloadCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var download7DaySeriesItem types.DownloadCount7DaySeriesItem
		download7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyDownloadCountList); i++ {
			if infoKey == 0 && len(downloadCount7Days.TimeData) < len(weeklyDownloadCountList) {
				downloadCount7Days.TimeData = append(downloadCount7Days.TimeData, weeklyDownloadCountList[i].Date)
			}
			download7DaySeriesItem.Data = append(download7DaySeriesItem.Data, weeklyDownloadCountList[i].Count)
		}
		downloadCount7Days.Series = append(downloadCount7Days.Series, download7DaySeriesItem)

		// 获取应用近7天应用启动次数
		weeklyAppStartCountList, err := resource.GetWeeklyAppStartCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appStart7DaySeriesItem types.AppStartCount7DaySeriesItem
		appStart7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppStartCountList); i++ {
			if infoKey == 0 && len(appStartCount7Days.TimeData) < len(weeklyAppStartCountList) {
				appStartCount7Days.TimeData = append(appStartCount7Days.TimeData, weeklyAppStartCountList[i].Date)
			}
			appStart7DaySeriesItem.Data = append(appStart7DaySeriesItem.Data, weeklyAppStartCountList[i].Count)
		}
		appStartCount7Days.Series = append(appStartCount7Days.Series, appStart7DaySeriesItem)

		// 获取应用近7天获取应用升级次数
		weeklyAppGetStrategyCountList, err := resource.GetWeeklyAppGetStrategyCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appGetStrategy7DaySeriesItem types.AppGetStrategyCount7DaySeriesItem
		appGetStrategy7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppGetStrategyCountList); i++ {
			if infoKey == 0 && len(appGetStrategyCount7Days.TimeData) < len(weeklyAppGetStrategyCountList) {
				appGetStrategyCount7Days.TimeData = append(appGetStrategyCount7Days.TimeData, weeklyAppGetStrategyCountList[i].Date)
			}
			appGetStrategy7DaySeriesItem.Data = append(appGetStrategy7DaySeriesItem.Data, weeklyAppGetStrategyCountList[i].Count)
		}
		appGetStrategyCount7Days.Series = append(appGetStrategyCount7Days.Series, appGetStrategy7DaySeriesItem)

		// 获取应用近7天应用升级次数
		weeklyAppUpgradeCountList, err := resource.GetWeeklyAppUpgradeCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appUpgrade7DaySeriesItem types.AppUpgradeCount7DaySeriesItem
		appUpgrade7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppUpgradeCountList); i++ {
			if infoKey == 0 && len(appUpgradeCount7Days.TimeData) < len(weeklyAppUpgradeCountList) {
				appUpgradeCount7Days.TimeData = append(appUpgradeCount7Days.TimeData, weeklyAppUpgradeCountList[i].Date)
			}
			appUpgrade7DaySeriesItem.Data = append(appUpgrade7DaySeriesItem.Data, weeklyAppUpgradeCountList[i].Count)
		}
		appUpgradeCount7Days.Series = append(appUpgradeCount7Days.Series, appUpgrade7DaySeriesItem)

	}

	// 获取公司下的所有file应用
	var filePredicates []predicate.UpgradeFile
	filePredicates = append(filePredicates, upgradefile.CompanyIDEQ(int(companyID)))
	filePredicates = append(filePredicates, upgradefile.IsDelEQ(0))
	fileList, err := l.svcCtx.DB.UpgradeFile.Query().Where(filePredicates...).Page(l.ctx, 1, 100)
	if err != nil {
		return nil, err
	}
	for infoKey, info := range fileList.List {
		// 获取应用近7天流量占用
		weeklyTrafficUsageCountList, err := resource.GetFileWeeklyTrafficUsageCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var trafficUsage7DaySeriesItem types.TrafficUsageCount7DaySeriesItem
		trafficUsage7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyTrafficUsageCountList); i++ {
			if infoKey == 0 && len(trafficUsageCount7Days.TimeData) < len(weeklyTrafficUsageCountList) {
				trafficUsageCount7Days.TimeData = append(trafficUsageCount7Days.TimeData, weeklyTrafficUsageCountList[i].Date)
			}
			trafficUsage7DaySeriesItem.Data = append(trafficUsage7DaySeriesItem.Data, weeklyTrafficUsageCountList[i].Count)
		}
		trafficUsageCount7Days.Series = append(trafficUsageCount7Days.Series, trafficUsage7DaySeriesItem)

		// 获取应用近7天下载次数
		weeklyDownloadCountList, err := resource.GetWeeklyDownloadCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var download7DaySeriesItem types.DownloadCount7DaySeriesItem
		download7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyDownloadCountList); i++ {
			if infoKey == 0 && len(downloadCount7Days.TimeData) < len(weeklyDownloadCountList) {
				downloadCount7Days.TimeData = append(downloadCount7Days.TimeData, weeklyDownloadCountList[i].Date)
			}
			download7DaySeriesItem.Data = append(download7DaySeriesItem.Data, weeklyDownloadCountList[i].Count)
		}
		downloadCount7Days.Series = append(downloadCount7Days.Series, download7DaySeriesItem)

		// 获取应用近7天应用启动次数
		weeklyAppStartCountList, err := resource.GetWeeklyAppStartCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appStart7DaySeriesItem types.AppStartCount7DaySeriesItem
		appStart7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppStartCountList); i++ {
			if infoKey == 0 && len(appStartCount7Days.TimeData) < len(weeklyAppStartCountList) {
				appStartCount7Days.TimeData = append(appStartCount7Days.TimeData, weeklyAppStartCountList[i].Date)
			}
			appStart7DaySeriesItem.Data = append(appStart7DaySeriesItem.Data, weeklyAppStartCountList[i].Count)
		}
		appStartCount7Days.Series = append(appStartCount7Days.Series, appStart7DaySeriesItem)

		// 获取应用近7天获取应用升级次数
		weeklyAppGetStrategyCountList, err := resource.GetWeeklyAppGetStrategyCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appGetStrategy7DaySeriesItem types.AppGetStrategyCount7DaySeriesItem
		appGetStrategy7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppGetStrategyCountList); i++ {
			if infoKey == 0 && len(appGetStrategyCount7Days.TimeData) < len(weeklyAppGetStrategyCountList) {
				appGetStrategyCount7Days.TimeData = append(appGetStrategyCount7Days.TimeData, weeklyAppGetStrategyCountList[i].Date)
			}
			appGetStrategy7DaySeriesItem.Data = append(appGetStrategy7DaySeriesItem.Data, weeklyAppGetStrategyCountList[i].Count)
		}
		appGetStrategyCount7Days.Series = append(appGetStrategyCount7Days.Series, appGetStrategy7DaySeriesItem)

		// 获取应用近7天应用升级次数
		weeklyAppUpgradeCountList, err := resource.GetWeeklyAppUpgradeCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appUpgrade7DaySeriesItem types.AppUpgradeCount7DaySeriesItem
		appUpgrade7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppUpgradeCountList); i++ {
			if infoKey == 0 && len(appUpgradeCount7Days.TimeData) < len(weeklyAppUpgradeCountList) {
				appUpgradeCount7Days.TimeData = append(appUpgradeCount7Days.TimeData, weeklyAppUpgradeCountList[i].Date)
			}
			appUpgrade7DaySeriesItem.Data = append(appUpgrade7DaySeriesItem.Data, weeklyAppUpgradeCountList[i].Count)
		}
		appUpgradeCount7Days.Series = append(appUpgradeCount7Days.Series, appUpgrade7DaySeriesItem)

	}

	// 获取公司下的所有tauri应用
	var tauriPredicates []predicate.UpgradeTauri
	tauriPredicates = append(tauriPredicates, upgradetauri.CompanyIDEQ(int(companyID)))
	tauriPredicates = append(tauriPredicates, upgradetauri.IsDelEQ(0))
	tauriList, err := l.svcCtx.DB.UpgradeTauri.Query().Where(tauriPredicates...).Page(l.ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	for infoKey, info := range tauriList.List {

		// 获取应用近7天流量占用
		weeklyTrafficUsageCountList, err := resource.GetTauriWeeklyTrafficUsageCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var trafficUsage7DaySeriesItem types.TrafficUsageCount7DaySeriesItem
		trafficUsage7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyTrafficUsageCountList); i++ {
			if infoKey == 0 && len(trafficUsageCount7Days.TimeData) < len(weeklyTrafficUsageCountList) {
				trafficUsageCount7Days.TimeData = append(trafficUsageCount7Days.TimeData, weeklyTrafficUsageCountList[i].Date)
			}
			trafficUsage7DaySeriesItem.Data = append(trafficUsage7DaySeriesItem.Data, weeklyTrafficUsageCountList[i].Count)
		}
		trafficUsageCount7Days.Series = append(trafficUsageCount7Days.Series, trafficUsage7DaySeriesItem)

		// 获取应用近7天下载量
		weeklyDownloadCountList, err := resource.GetWeeklyDownloadCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}

		var download7DaySeriesItem types.DownloadCount7DaySeriesItem
		download7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyDownloadCountList); i++ {
			if infoKey == 0 && len(downloadCount7Days.TimeData) < len(weeklyDownloadCountList) {
				downloadCount7Days.TimeData = append(downloadCount7Days.TimeData, weeklyDownloadCountList[i].Date)
			}
			download7DaySeriesItem.Data = append(download7DaySeriesItem.Data, weeklyDownloadCountList[i].Count)
		}
		downloadCount7Days.Series = append(downloadCount7Days.Series, download7DaySeriesItem)

		// 获取应用近7天应用启动次数
		weeklyAppStartCountList, err := resource.GetWeeklyAppStartCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appStart7DaySeriesItem types.AppStartCount7DaySeriesItem
		appStart7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppStartCountList); i++ {
			if infoKey == 0 && len(appStartCount7Days.TimeData) < len(weeklyAppStartCountList) {
				appStartCount7Days.TimeData = append(appStartCount7Days.TimeData, weeklyAppStartCountList[i].Date)
			}
			appStart7DaySeriesItem.Data = append(appStart7DaySeriesItem.Data, weeklyAppStartCountList[i].Count)
		}
		appStartCount7Days.Series = append(appStartCount7Days.Series, appStart7DaySeriesItem)

		// 获取应用近7天获取应用升级次数
		weeklyAppGetStrategyCountList, err := resource.GetWeeklyAppGetStrategyCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appGetStrategy7DaySeriesItem types.AppGetStrategyCount7DaySeriesItem
		appGetStrategy7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppGetStrategyCountList); i++ {
			if infoKey == 0 && len(appGetStrategyCount7Days.TimeData) < len(weeklyAppGetStrategyCountList) {
				appGetStrategyCount7Days.TimeData = append(appGetStrategyCount7Days.TimeData, weeklyAppGetStrategyCountList[i].Date)
			}
			appGetStrategy7DaySeriesItem.Data = append(appGetStrategy7DaySeriesItem.Data, weeklyAppGetStrategyCountList[i].Count)
		}
		appGetStrategyCount7Days.Series = append(appGetStrategyCount7Days.Series, appGetStrategy7DaySeriesItem)

		// 获取应用近7天应用升级次数
		weeklyAppUpgradeCountList, err := resource.GetWeeklyAppUpgradeCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appUpgrade7DaySeriesItem types.AppUpgradeCount7DaySeriesItem
		appUpgrade7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppUpgradeCountList); i++ {
			if infoKey == 0 && len(appUpgradeCount7Days.TimeData) < len(weeklyAppUpgradeCountList) {
				appUpgradeCount7Days.TimeData = append(appUpgradeCount7Days.TimeData, weeklyAppUpgradeCountList[i].Date)
			}
			appUpgrade7DaySeriesItem.Data = append(appUpgrade7DaySeriesItem.Data, weeklyAppUpgradeCountList[i].Count)
		}
		appUpgradeCount7Days.Series = append(appUpgradeCount7Days.Series, appUpgrade7DaySeriesItem)

	}

	// 获取公司下的所有electron应用
	var electronPredicates []predicate.UpgradeElectron
	electronPredicates = append(electronPredicates, upgradeelectron.CompanyIDEQ(int(companyID)))
	electronPredicates = append(electronPredicates, upgradeelectron.IsDelEQ(0))
	electronList, err := l.svcCtx.DB.UpgradeElectron.Query().Where(electronPredicates...).Page(l.ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	for infoKey, info := range electronList.List {

		// 获取应用近7天流量占用
		weeklyTrafficUsageCountList, err := resource.GetElectronWeeklyTrafficUsageCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var trafficUsage7DaySeriesItem types.TrafficUsageCount7DaySeriesItem
		trafficUsage7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyTrafficUsageCountList); i++ {
			if infoKey == 0 && len(trafficUsageCount7Days.TimeData) < len(weeklyTrafficUsageCountList) {
				trafficUsageCount7Days.TimeData = append(trafficUsageCount7Days.TimeData, weeklyTrafficUsageCountList[i].Date)
			}
			trafficUsage7DaySeriesItem.Data = append(trafficUsage7DaySeriesItem.Data, weeklyTrafficUsageCountList[i].Count)
		}
		trafficUsageCount7Days.Series = append(trafficUsageCount7Days.Series, trafficUsage7DaySeriesItem)

		// 获取应用近7天下载量
		weeklyDownloadCountList, err := resource.GetWeeklyDownloadCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}

		var download7DaySeriesItem types.DownloadCount7DaySeriesItem
		download7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyDownloadCountList); i++ {
			if infoKey == 0 && len(downloadCount7Days.TimeData) < len(weeklyDownloadCountList) {
				downloadCount7Days.TimeData = append(downloadCount7Days.TimeData, weeklyDownloadCountList[i].Date)
			}
			download7DaySeriesItem.Data = append(download7DaySeriesItem.Data, weeklyDownloadCountList[i].Count)
		}
		downloadCount7Days.Series = append(downloadCount7Days.Series, download7DaySeriesItem)

		// 获取应用近7天应用启动次数
		weeklyAppStartCountList, err := resource.GetWeeklyAppStartCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appStart7DaySeriesItem types.AppStartCount7DaySeriesItem
		appStart7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppStartCountList); i++ {
			if infoKey == 0 && len(appStartCount7Days.TimeData) < len(weeklyAppStartCountList) {
				appStartCount7Days.TimeData = append(appStartCount7Days.TimeData, weeklyAppStartCountList[i].Date)
			}
			appStart7DaySeriesItem.Data = append(appStart7DaySeriesItem.Data, weeklyAppStartCountList[i].Count)
		}
		appStartCount7Days.Series = append(appStartCount7Days.Series, appStart7DaySeriesItem)

		// 获取应用近7天获取应用升级次数
		weeklyAppGetStrategyCountList, err := resource.GetWeeklyAppGetStrategyCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appGetStrategy7DaySeriesItem types.AppGetStrategyCount7DaySeriesItem
		appGetStrategy7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppGetStrategyCountList); i++ {
			if infoKey == 0 && len(appGetStrategyCount7Days.TimeData) < len(weeklyAppGetStrategyCountList) {
				appGetStrategyCount7Days.TimeData = append(appGetStrategyCount7Days.TimeData, weeklyAppGetStrategyCountList[i].Date)
			}
			appGetStrategy7DaySeriesItem.Data = append(appGetStrategy7DaySeriesItem.Data, weeklyAppGetStrategyCountList[i].Count)
		}
		appGetStrategyCount7Days.Series = append(appGetStrategyCount7Days.Series, appGetStrategy7DaySeriesItem)

		// 获取应用近7天应用升级次数
		weeklyAppUpgradeCountList, err := resource.GetWeeklyAppUpgradeCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appUpgrade7DaySeriesItem types.AppUpgradeCount7DaySeriesItem
		appUpgrade7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppUpgradeCountList); i++ {
			if infoKey == 0 && len(appUpgradeCount7Days.TimeData) < len(weeklyAppUpgradeCountList) {
				appUpgradeCount7Days.TimeData = append(appUpgradeCount7Days.TimeData, weeklyAppUpgradeCountList[i].Date)
			}
			appUpgrade7DaySeriesItem.Data = append(appUpgrade7DaySeriesItem.Data, weeklyAppUpgradeCountList[i].Count)
		}
		appUpgradeCount7Days.Series = append(appUpgradeCount7Days.Series, appUpgrade7DaySeriesItem)

	}

	// 获取公司下的所有安卓应用
	var apkPredicates []predicate.UpgradeApk
	apkPredicates = append(apkPredicates, upgradeapk.CompanyIDEQ(int(companyID)))
	apkPredicates = append(apkPredicates, upgradeapk.IsDelEQ(0))
	apkList, err := l.svcCtx.DB.UpgradeApk.Query().Where(apkPredicates...).Page(l.ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	for infoKey, info := range apkList.List {

		// 获取应用近7天流量占用
		weeklyTrafficUsageCountList, err := resource.GetApkWeeklyTrafficUsageCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var trafficUsage7DaySeriesItem types.TrafficUsageCount7DaySeriesItem
		trafficUsage7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyTrafficUsageCountList); i++ {
			if infoKey == 0 && len(trafficUsageCount7Days.TimeData) < len(weeklyTrafficUsageCountList) {
				trafficUsageCount7Days.TimeData = append(trafficUsageCount7Days.TimeData, weeklyTrafficUsageCountList[i].Date)
			}
			trafficUsage7DaySeriesItem.Data = append(trafficUsage7DaySeriesItem.Data, weeklyTrafficUsageCountList[i].Count)
		}
		trafficUsageCount7Days.Series = append(trafficUsageCount7Days.Series, trafficUsage7DaySeriesItem)

		// 获取应用近7天设备新增量
		weeklyDownloadCountList, err := resource.GetWeeklyDownloadCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}

		var download7DaySeriesItem types.DownloadCount7DaySeriesItem
		download7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyDownloadCountList); i++ {
			if infoKey == 0 && len(downloadCount7Days.TimeData) < len(weeklyDownloadCountList) {
				downloadCount7Days.TimeData = append(downloadCount7Days.TimeData, weeklyDownloadCountList[i].Date)
			}
			download7DaySeriesItem.Data = append(download7DaySeriesItem.Data, weeklyDownloadCountList[i].Count)
		}
		downloadCount7Days.Series = append(downloadCount7Days.Series, download7DaySeriesItem)

		// 获取应用近7天应用启动次数
		weeklyAppStartCountList, err := resource.GetWeeklyAppStartCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appStart7DaySeriesItem types.AppStartCount7DaySeriesItem
		appStart7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppStartCountList); i++ {
			if infoKey == 0 && len(appStartCount7Days.TimeData) < len(weeklyAppStartCountList) {
				appStartCount7Days.TimeData = append(appStartCount7Days.TimeData, weeklyAppStartCountList[i].Date)
			}
			appStart7DaySeriesItem.Data = append(appStart7DaySeriesItem.Data, weeklyAppStartCountList[i].Count)
		}
		appStartCount7Days.Series = append(appStartCount7Days.Series, appStart7DaySeriesItem)

		// 获取应用近7天获取应用升级次数
		weeklyAppGetStrategyCountList, err := resource.GetWeeklyAppGetStrategyCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appGetStrategy7DaySeriesItem types.AppGetStrategyCount7DaySeriesItem
		appGetStrategy7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppGetStrategyCountList); i++ {
			if infoKey == 0 && len(appGetStrategyCount7Days.TimeData) < len(weeklyAppGetStrategyCountList) {
				appGetStrategyCount7Days.TimeData = append(appGetStrategyCount7Days.TimeData, weeklyAppGetStrategyCountList[i].Date)
			}
			appGetStrategy7DaySeriesItem.Data = append(appGetStrategy7DaySeriesItem.Data, weeklyAppGetStrategyCountList[i].Count)
		}
		appGetStrategyCount7Days.Series = append(appGetStrategyCount7Days.Series, appGetStrategy7DaySeriesItem)

		// 获取应用近7天应用升级次数
		weeklyAppUpgradeCountList, err := resource.GetWeeklyAppUpgradeCount(l.ctx, l.svcCtx, info.Key)
		if err != nil {
			return nil, err
		}
		var appUpgrade7DaySeriesItem types.AppUpgradeCount7DaySeriesItem
		appUpgrade7DaySeriesItem.AppName = info.Name
		for i := 0; i < len(weeklyAppUpgradeCountList); i++ {
			if infoKey == 0 && len(appUpgradeCount7Days.TimeData) < len(weeklyAppUpgradeCountList) {
				appUpgradeCount7Days.TimeData = append(appUpgradeCount7Days.TimeData, weeklyAppUpgradeCountList[i].Date)
			}
			appUpgrade7DaySeriesItem.Data = append(appUpgrade7DaySeriesItem.Data, weeklyAppUpgradeCountList[i].Count)
		}
		appUpgradeCount7Days.Series = append(appUpgradeCount7Days.Series, appUpgrade7DaySeriesItem)

	}

	newResp.Data.TrafficUsageCount7Day = trafficUsageCount7Days
	newResp.Data.DownloadCount7Day = downloadCount7Days
	newResp.Data.AppGetStrategyCount7Day = appGetStrategyCount7Days
	newResp.Data.AppUpgradeCount7Day = appUpgradeCount7Days
	newResp.Data.AppStartCount7Day = appStartCount7Days

	return newResp, nil
}
