package app

import (
	"context"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"upgradelink-api/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLogic {
	return &StatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsLogic) Statistics(req *types.StatisticsReq) (resp *types.StatisticsResp, err error) {
	var res types.StatisticsResp

	// 设置响应状态码和消息
	res.Code = 200
	res.Msg = l.svcCtx.Trans.Trans(l.ctx, "common.success")

	// 获取昨日应用下载量
	yesterdayAppDownloadCount, err := l.svcCtx.ResourceCtx.GetYesterdayDownloadCount(l.ctx, req.AppKey)
	res.Data.YesterdayDownloadCount = pointy.GetPointer(yesterdayAppDownloadCount)

	// 获取总应用下载量
	totalAppDownloadCount, err := l.svcCtx.ResourceCtx.GetDownloadCount(l.ctx, req.AppKey)
	res.Data.TotalDownloadCount = pointy.GetPointer(totalAppDownloadCount)

	// 获取昨日应用获取升级次数
	yesterdayAppGetStrategyCount, err := l.svcCtx.ResourceCtx.GetYesterdayAppGetStrategyCount(l.ctx, req.AppKey)
	res.Data.YesterdayAppGetStrategyCount = pointy.GetPointer(yesterdayAppGetStrategyCount)

	// 获取总应用获取升级次数
	totalAppGetStrategyCount, err := l.svcCtx.ResourceCtx.GetAppGetStrategyCount(l.ctx, req.AppKey)
	res.Data.TotalAppGetStrategyCount = pointy.GetPointer(totalAppGetStrategyCount)

	// 获取昨日应用升级次数
	yesterdayAppUpgradeCount, err := l.svcCtx.ResourceCtx.GetYesterdayAppUpgradeCount(l.ctx, req.AppKey)
	res.Data.YesterdayAppUpgradeCount = pointy.GetPointer(yesterdayAppUpgradeCount)

	// 获取总应用升级次数
	totalAppUpgradeCount, err := l.svcCtx.ResourceCtx.GetAppUpgradeCount(l.ctx, req.AppKey)
	res.Data.TotalAppUpgradeCount = pointy.GetPointer(totalAppUpgradeCount)

	// 获取昨日应用启动次数
	yesterdayAppStartCount, err := l.svcCtx.ResourceCtx.GetYesterdayAppStartCount(l.ctx, req.AppKey)
	res.Data.YesterdayAppStartCount = pointy.GetPointer(yesterdayAppStartCount)

	// 获取总应用启动次数
	totalAppStartCount, err := l.svcCtx.ResourceCtx.GetAppStartCount(l.ctx, req.AppKey)
	res.Data.TotalAppStartCount = pointy.GetPointer(totalAppStartCount)

	// 近7天应用下载次数
	weeklyDownloadCount, err := l.svcCtx.ResourceCtx.GetWeeklyDownloadCount(l.ctx, req.AppKey)
	for i := 0; i < len(weeklyDownloadCount); i++ {
		res.Data.DownloadCount7Day = append(res.Data.DownloadCount7Day, types.DownloadCount7DaySeriesItem{
			TimeData: weeklyDownloadCount[i].Date,
			Data:     weeklyDownloadCount[i].Count,
		})
	}

	// 近7天应用获取升级次数
	weeklyAppGetStrategyCount, err := l.svcCtx.ResourceCtx.GetWeeklyAppGetStrategyCount(l.ctx, req.AppKey)
	for i := 0; i < len(weeklyAppGetStrategyCount); i++ {
		res.Data.AppGetStrategyCount7Day = append(res.Data.AppGetStrategyCount7Day, types.AppGetStrategyCount7DaySeriesItem{
			TimeData: weeklyAppGetStrategyCount[i].Date,
			Data:     weeklyAppGetStrategyCount[i].Count,
		})
	}

	// 近 7 天应用升级次数
	weeklyAppUpgradeCount, err := l.svcCtx.ResourceCtx.GetWeeklyAppUpgradeCount(l.ctx, req.AppKey)
	for i := 0; i < len(weeklyAppUpgradeCount); i++ {
		res.Data.AppUpgradeCount7Day = append(res.Data.AppUpgradeCount7Day, types.AppUpgradeCount7DaySeriesItem{
			TimeData: weeklyAppUpgradeCount[i].Date,
			Data:     weeklyAppUpgradeCount[i].Count,
		})
	}

	// 近 7 天应用启动次数
	weeklyAppStartCount, err := l.svcCtx.ResourceCtx.GetWeeklyAppStartCount(l.ctx, req.AppKey)
	for i := 0; i < len(weeklyAppStartCount); i++ {
		res.Data.AppStartCount7Day = append(res.Data.AppStartCount7Day, types.AppStartCount7DaySeriesItem{
			TimeData: weeklyAppStartCount[i].Date,
			Data:     weeklyAppStartCount[i].Count,
		})
	}

	return &res, nil
}
