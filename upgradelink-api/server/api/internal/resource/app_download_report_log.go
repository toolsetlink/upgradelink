package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyStatisticsAppYesterdayDownloadCountKey = PREFIX + "StatisticsAppYesterdayDownloadCount:APP_KEY:%v"
	CacheKeyStatisticsAppDownloadCountKey          = PREFIX + "StatisticsAppDownloadCount:APP_KEY:%v"
	CacheKeyStatisticsWeeklyDownloadCountKey       = PREFIX + "StatisticsWeeklyDownloadCount:APP_KEY:%v"
)

type AddAppDownloadReportLogReq struct {
	Id                  int64     `db:"id"`                     // ID
	CompanyId           int64     `db:"company_id"`             // 公司ID
	Timestamp           time.Time `db:"timestamp"`              // 事件发生时间
	AppKey              string    `db:"app_key"`                // 应用Key
	AppType             string    `db:"app_type"`               // 应用类型
	AppVersionId        int64     `db:"app_version_id"`         // 应用版本ID
	AppVersionCode      int64     `db:"app_version_code"`       // 应用版本号
	AppVersionPlatform  string    `db:"app_version_platform"`   // 应用版本操作系统:linux、darwin、windows
	AppVersionTarget    string    `db:"app_version_target"`     // 应用版本操作系统:linux、darwin、windows
	AppVersionArch      string    `db:"app_version_arch"`       // 应用版本机器架构:x86_64、i686、aarch64、armv7
	DownloadCloudFileId string    `db:"download_cloud_file_id"` // 下载云文件ID
	CreateAt            time.Time `db:"create_at"`              // 创建时间
}

func (c *Ctx) AddAppDownloadReportLog(ctx context.Context, req AddAppDownloadReportLogReq) (int64, error) {

	query := fmt.Sprintf("insert into %s (company_id, timestamp, app_key, app_type, app_version_id, app_version_code, app_version_platform, app_version_target, app_version_arch, download_cloud_file_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_download_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Timestamp, req.AppKey, req.AppType, req.AppVersionId, req.AppVersionCode, req.AppVersionPlatform, req.AppVersionTarget, req.AppVersionArch, req.DownloadCloudFileId)
	fmt.Println(ret.RowsAffected())
	if err != nil {
		return 0, err
	}

	lastId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

type GetYesterdayDownloadCountStruct struct {
	Count int `db:"count"`
}

// GetYesterdayDownloadCount
// 获取昨日下载次数
func (c *Ctx) GetYesterdayDownloadCount(ctx context.Context, appKey string) (int, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsAppYesterdayDownloadCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info GetYesterdayDownloadCountStruct
		query := fmt.Sprintf("select count(*) AS count " +
			"FROM upgrade_app_download_report_log " +
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

type GetDownloadCountStruct struct {
	Count int `db:"count"`
}

// GetDownloadCount
// 获取下载次数
func (c *Ctx) GetDownloadCount(ctx context.Context, appKey string) (int, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsAppDownloadCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info GetDownloadCountStruct
		query := fmt.Sprintf("select count(*) AS count " +
			"FROM upgrade_app_download_report_log " +
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

type WeeklyDownloadCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyDownloadCount 获取最近7天每日应用下载次数
func (c *Ctx) GetWeeklyDownloadCount(ctx context.Context, appKey string) ([]WeeklyDownloadCount, error) {

	cacheKey := fmt.Sprintf(CacheKeyStatisticsWeeklyDownloadCountKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ApkVersionList []WeeklyDownloadCount
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
				   FROM upgrade_app_download_report_log
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
		list := v.([]WeeklyDownloadCount)
		return list, err
	}

	return nil, err
}
