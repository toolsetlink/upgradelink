package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyStatisticsYesterdayAppUpgradeCountKey = PREFIX + "StatisticsYesterdayAppUpgradeCount:APP_KEY:%v"
	CacheKeyStatisticsAppUpgradeCountKey          = PREFIX + "StatisticsAppUpgradeCount:APP_KEY:%v"
	CacheKeyStatisticsWeeklyAppUpgradeCountKey    = PREFIX + "StatisticsWeeklyAppUpgradeCount:APP_KEY:%v"
)

type AddAppUpgradeUpgradeReportLogReq struct {
	Id                 int64     `db:"id"`                   // ID
	CompanyId          int64     `db:"company_id"`           // 公司ID
	Timestamp          time.Time `db:"timestamp"`            // 事件发生时间
	AppKey             string    `db:"app_key"`              // 应用Key
	AppType            string    `db:"app_type"`             // 应用类型
	AppVersionId       int64     `db:"app_version_id"`       // 应用版本ID
	AppVersionCode     int64     `db:"app_version_code"`     // 应用版本号
	DevModelKey        string    `db:"dev_model_key"`        // 设备机型唯一标识
	DevKey             string    `db:"dev_key"`              // 设备唯一标识
	AppVersionTarget   string    `db:"app_version_target"`   // 操作系统
	AppVersionArch     string    `db:"app_version_arch"`     // 机器架构
	UpgradeVersionCode int64     `db:"upgrade_version_code"` // 应用升级版本号
	Code               int64     `db:"code"`                 // 事件-状态码
	CreateAt           time.Time `db:"create_at"`            // 创建时间
}

func (c *Ctx) AddAppUpgradeUpgradeReportLog(ctx context.Context, req AddAppUpgradeUpgradeReportLogReq) (*model.UpgradeAppUpgradeUpgradeReportLog, error) {

	query := fmt.Sprintf("insert into %s (company_id, timestamp, app_key, app_type, app_version_id, app_version_code, dev_model_key, dev_key, app_version_target, app_version_arch, upgrade_version_code, code) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_upgrade_upgrade_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Timestamp, req.AppKey, req.AppType, req.AppVersionId, req.AppVersionCode, req.DevModelKey, req.DevKey, req.AppVersionTarget, req.AppVersionArch, req.UpgradeVersionCode, req.Code)

	fmt.Println(ret.RowsAffected())

	return nil, err
}

type GetYesterdayAppUpgradeCountStruct struct {
	Count int `db:"count"`
}

// GetYesterdayAppUpgradeCount
// 获取昨日应用升级次数
func (c *Ctx) GetYesterdayAppUpgradeCount(ctx context.Context, appKey string) (int, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsYesterdayAppUpgradeCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info GetYesterdayAppUpgradeCountStruct
		query := fmt.Sprintf("select count(*) AS count " +
			"FROM upgrade_app_upgrade_upgrade_report_log " +
			"WHERE app_key = ? " +
			"AND timestamp >= CURDATE() - INTERVAL 1 DAY " +
			"AND timestamp < CURDATE() ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, appKey)
		})

		if err != nil {
			return nil, err
		}
		return info.Count, nil
	})

	if v != nil {
		return v.(int), err
	}

	return 0, err
}

type GetAppUpgradeCountStruct struct {
	Count int `db:"count"`
}

// GetAppUpgradeCount
// 获取应用升级总次数
func (c *Ctx) GetAppUpgradeCount(ctx context.Context, appKey string) (int, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsAppUpgradeCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info GetAppUpgradeCountStruct
		query := fmt.Sprintf("select count(*) AS count " +
			"FROM upgrade_app_upgrade_upgrade_report_log " +
			"WHERE app_key = ? " +
			"AND timestamp < CURDATE() ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, appKey)
		})

		if err != nil {
			return nil, err
		}
		return info.Count, nil
	})

	if v != nil {
		return v.(int), err
	}

	return 0, err
}

type WeeklyAppUpgradeCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyAppUpgradeCount 获取最近7天每日应用升级次数
func (c *Ctx) GetWeeklyAppUpgradeCount(ctx context.Context, appKey string) ([]WeeklyAppUpgradeCount, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsWeeklyAppUpgradeCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ApkVersionList []WeeklyAppUpgradeCount
		const query = `
			  SELECT
				  DATE_FORMAT(dates.date, '%y%m%d') AS date,
				  IFNULL(tmp.count, 0) AS count
			  FROM (
				  SELECT DATE_SUB(CURDATE(), INTERVAL 7 DAY) + INTERVAL t.n DAY AS date
				  FROM (
					  SELECT 0 AS n UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL
					  SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL SELECT 6
				  ) t
			  ) dates
			  LEFT JOIN (
				  SELECT
					  DATE(timestamp) AS date,
					  COUNT(1) AS count
				  FROM upgrade_app_upgrade_upgrade_report_log
				  WHERE app_key = ?
				  AND timestamp >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
				  AND timestamp < CURDATE()
				  GROUP BY DATE(timestamp)
			  ) tmp ON dates.date = tmp.date
			  ORDER BY dates.date ASC
		  `
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ApkVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ApkVersionList, query, appKey)
		})

		if err != nil {
			return nil, err
		}
		return ApkVersionList, nil
	})

	if v != nil {
		list := v.([]WeeklyAppUpgradeCount)
		return list, err
	}

	return nil, err
}
