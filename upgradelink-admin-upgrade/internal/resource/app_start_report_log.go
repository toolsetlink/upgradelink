package resource

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin-upgrade/internal/svc"
)

type GetYesterdayAppStartCountStruct struct {
	Count int `db:"count"`
}

// GetYesterdayAppStartCount
// 获取昨日应用启动次数
func GetYesterdayAppStartCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("select count(*) AS count " +
		"FROM upgrade_app_start_report_log " +
		"WHERE company_id = ? " +
		"AND timestamp >= CURDATE() - INTERVAL 1 DAY " +
		"AND timestamp < CURDATE() ")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetYesterdayAppStartCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.Count, err

}

type GetAppStartCountStruct struct {
	Count int `db:"count"`
}

// GetAppStartCount
// 获取应用启动总次数
func GetAppStartCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("select count(*) AS count " +
		"FROM upgrade_app_start_report_log " +
		"WHERE company_id = ? " +
		"AND timestamp < CURDATE() ")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetYesterdayAppStartCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.Count, err
}

type WeeklyAppStartCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyAppStartCount 获取最近7天每日应用启动次数
func GetWeeklyAppStartCount(ctx context.Context, svcCtx *svc.ServiceContext, appKey string) ([]WeeklyAppStartCount, error) {
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
           FROM upgrade_app_start_report_log
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
	// 解析结构体
	var result []WeeklyAppStartCount
	// 解析结构体
	for queryContext.Next() {
		var info WeeklyAppStartCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}

type WeeklyAppUpgradeCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}
