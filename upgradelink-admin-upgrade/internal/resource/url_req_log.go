package resource

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin-upgrade/internal/svc"
)

type GetYesterdayNewDevCountStruct struct {
	NewDeviceCount int `db:"new_device_count"`
}

// GetYesterdayNewDevCount
// 获取设备数量
func GetYesterdayNewDevCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("SELECT COUNT(DISTINCT dev_key) AS new_device_count " +
		"FROM upgrade_url_req_log AS current " +
		"WHERE company_id = ? " +
		"AND create_at >= DATE_SUB(CURDATE(), INTERVAL 1 DAY) " +
		"AND create_at < CURDATE() " +
		"AND NOT EXISTS ( " +
		"SELECT 1 FROM upgrade_url_req_log AS previous " +
		"WHERE previous.company_id = current.company_id " +
		"AND previous.dev_key = current.dev_key " +
		"AND previous.create_at < CURDATE() - INTERVAL 1 DAY " +
		")")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetYesterdayNewDevCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.NewDeviceCount, err
}

type GetDevCountByCompanyIdStruct struct {
	DeviceCount int `db:"device_count"`
}

// GetDevCountByCompanyId
// 获取公司下设备总数量
func GetDevCountByCompanyId(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("SELECT COUNT(DISTINCT dev_key) AS device_count " +
		"FROM upgrade_url_req_log " +
		"WHERE company_id = ?")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetDevCountByCompanyIdStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.DeviceCount, err
}

type GetYesterdayReqCountStruct struct {
	NewDeviceCount int `db:"new_req_count"`
}

// GetYesterdayReqCount
// 获取昨日请求次数
func GetYesterdayReqCount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("SELECT COUNT(1) AS new_req_count " +
		"FROM upgrade_url_req_log AS current " +
		"WHERE company_id = ? " +
		"AND create_at >= DATE_SUB(CURDATE(), INTERVAL 1 DAY) " +
		"AND create_at < CURDATE() ")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetYesterdayNewDevCountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.NewDeviceCount, err
}

type GetReqCountByCompanyIdStruct struct {
	ReqCount int `db:"req_count"`
}

// GetReqCountByCompanyId
// 获取公司下总请求次数
func GetReqCountByCompanyId(ctx context.Context, svcCtx *svc.ServiceContext, companyId int64) (int, error) {

	query := fmt.Sprintf("SELECT COUNT(1) AS req_count " +
		"FROM upgrade_url_req_log " +
		"WHERE company_id = ?")

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetReqCountByCompanyIdStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.ReqCount, err
}

// 新增结构体用于接收查询结果
type DailyDevCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyNewDevCount 获取最近7天每日新增设备数
func GetWeeklyNewDevCount(ctx context.Context, svcCtx *svc.ServiceContext, urlId int64) ([]DailyDevCount, error) {
	const query = `
       SELECT
           DATE_FORMAT(dates.date, '%Y-%m-%d') AS date,
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
               DATE(create_at) AS date,
               COUNT(DISTINCT dev_key) AS count
           FROM upgrade_url_req_log
           WHERE url_id = ?
           AND create_at >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND create_at < CURDATE()
           AND NOT EXISTS (
               SELECT 1
               FROM upgrade_url_req_log AS previous
               WHERE previous.url_id = upgrade_url_req_log.url_id
               AND previous.dev_key = upgrade_url_req_log.dev_key
               AND previous.create_at < DATE(upgrade_url_req_log.create_at)
           )
           GROUP BY DATE(create_at)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, urlId)
	if err != nil {
		return nil, err
	}
	// 解析结构体
	var result []DailyDevCount
	// 解析结构体
	for queryContext.Next() {
		var info DailyDevCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil

}

type DailyReqCount struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// GetWeeklyReqCount 获取近7天请求次数
func GetWeeklyReqCount(ctx context.Context, svcCtx *svc.ServiceContext, urlId int64) ([]DailyReqCount, error) {
	const query = `
       SELECT
           DATE_FORMAT(dates.date, '%Y-%m-%d') AS date,
           IFNULL(tmp.count, 0) AS count
       FROM (
           SELECT DATE_SUB(CURDATE(), INTERVAL 7 DAY) + INTERVAL t.n DAY AS date
           FROM (
               SELECT 0 AS n UNION SELECT 1 UNION SELECT 2 UNION
               SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6
           ) t
       ) dates
       LEFT JOIN (
           SELECT
               DATE(create_at) AS date,
               COUNT(1) AS count
           FROM upgrade_url_req_log
           WHERE url_id = ?
           AND create_at >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
           AND create_at < CURDATE() + INTERVAL 1 DAY
           GROUP BY DATE(create_at)
       ) tmp ON dates.date = tmp.date
       ORDER BY dates.date ASC
   `

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, urlId)
	if err != nil {
		return nil, err
	}
	// 解析结构体
	var result []DailyReqCount
	// 解析结构体
	for queryContext.Next() {
		var info DailyReqCount
		if err := queryContext.Scan(&info.Date, &info.Count); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}
