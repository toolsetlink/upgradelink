package resource

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin-upgrade/internal/svc"
)

type GetYesterdayDownloadCountStruct struct {
	Count int `db:"count"`
}

// GetYesterdayDownloadCount
// 获取昨日下载次数
func GetYesterdayDownloadCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("select count(*) AS count " +
		"FROM upgrade_app_download_report_log " +
		"WHERE company_id = ? " +
		"AND timestamp >= CURDATE() - INTERVAL 1 DAY " +
		"AND timestamp < CURDATE() ")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetYesterdayDownloadCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.Count, err
}

type GetDownloadCountStruct struct {
	Count int `db:"count"`
}

// GetDownloadCount
// 获取下载次数
func GetDownloadCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("select count(*) AS count " +
		"FROM upgrade_app_download_report_log " +
		"WHERE company_id = ? " +
		"AND timestamp < CURDATE() ")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetDownloadCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.Count, err

}

type WeeklyDownloadCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyDownloadCount 获取最近7天每日应用下载次数
func GetWeeklyDownloadCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyDownloadCount, error) {
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

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, appKey)
	if err != nil {
		return nil, err
	}

	var result []WeeklyDownloadCount

	// 解析结构体
	for queryContext.Next() {
		var info WeeklyDownloadCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}

type WeeklyTrafficUsageCount struct {
	Date  string  `db:"date"`
	Count float32 `db:"count"`
}

// GetFileWeeklyTrafficUsageCount 获取文件应用最近7天每日应用流量使用
func GetFileWeeklyTrafficUsageCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyTrafficUsageCount, error) {
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
               DATE(upgrade_app_download_report_log.timestamp) AS date,
               COALESCE(ROUND(SUM(fms_cloud_files.size) / POWER(1024, 3), 4), 0) AS count
           FROM upgrade_app_download_report_log
           left join fms_cloud_files on upgrade_app_download_report_log.download_cloud_file_id = fms_cloud_files.id
           WHERE upgrade_app_download_report_log.app_key = ?
           AND upgrade_app_download_report_log.timestamp >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND upgrade_app_download_report_log.timestamp < CURDATE()
           GROUP BY DATE(timestamp)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, appKey)
	if err != nil {
		return nil, err
	}

	var result []WeeklyTrafficUsageCount
	// 解析结构体
	for queryContext.Next() {
		var info WeeklyTrafficUsageCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}

// GetTauriWeeklyTrafficUsageCount 获取Tauri应用最近7天每日应用流量使用
func GetTauriWeeklyTrafficUsageCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyTrafficUsageCount, error) {
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
               DATE(upgrade_app_download_report_log.timestamp) AS date,
               COALESCE(ROUND(SUM(fms_cloud_files.size) / POWER(1024, 3), 4), 0) AS count
           FROM upgrade_app_download_report_log
           left join fms_cloud_files on upgrade_app_download_report_log.download_cloud_file_id = fms_cloud_files.id
           WHERE upgrade_app_download_report_log.app_key = ?
           AND upgrade_app_download_report_log.timestamp >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND upgrade_app_download_report_log.timestamp < CURDATE()
           GROUP BY DATE(timestamp)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, appKey)
	if err != nil {
		return nil, err
	}

	var result []WeeklyTrafficUsageCount
	// 解析结构体
	for queryContext.Next() {
		var info WeeklyTrafficUsageCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil

}

// GetElectronWeeklyTrafficUsageCount 获取Electron应用最近7天每日应用流量使用
func GetElectronWeeklyTrafficUsageCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyTrafficUsageCount, error) {
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
               DATE(upgrade_app_download_report_log.timestamp) AS date,
               COALESCE(ROUND(SUM(fms_cloud_files.size) / POWER(1024, 3), 4), 0) AS count
           FROM upgrade_app_download_report_log
           left join fms_cloud_files on upgrade_app_download_report_log.download_cloud_file_id = fms_cloud_files.id
           WHERE upgrade_app_download_report_log.app_key = ?
           AND upgrade_app_download_report_log.timestamp >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND upgrade_app_download_report_log.timestamp < CURDATE()
           GROUP BY DATE(timestamp)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, appKey)
	if err != nil {
		return nil, err
	}

	var result []WeeklyTrafficUsageCount
	// 解析结构体
	for queryContext.Next() {
		var info WeeklyTrafficUsageCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}

// GetApkWeeklyTrafficUsageCount 获取apk应用最近7天每日应用流量使用
func GetApkWeeklyTrafficUsageCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyTrafficUsageCount, error) {
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
               DATE(upgrade_app_download_report_log.timestamp) AS date,
               COALESCE(ROUND(SUM(fms_cloud_files.size) / POWER(1024, 3), 4), 0) AS count
           FROM upgrade_app_download_report_log
           left join fms_cloud_files on upgrade_app_download_report_log.download_cloud_file_id = fms_cloud_files.id
           WHERE upgrade_app_download_report_log.app_key = ?
           AND upgrade_app_download_report_log.timestamp >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND upgrade_app_download_report_log.timestamp < CURDATE()
           GROUP BY DATE(timestamp)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, appKey)
	if err != nil {
		return nil, err
	}
	var result []WeeklyTrafficUsageCount
	// 解析结构体
	for queryContext.Next() {
		var info WeeklyTrafficUsageCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}
