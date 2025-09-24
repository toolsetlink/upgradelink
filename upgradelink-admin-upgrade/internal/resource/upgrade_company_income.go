package resource

import (
	"context"
	"fmt"
	"upgradelink-admin-upgrade/internal/svc"
)

type GetIncomeAmountStruct struct {
	TotalPendingSettlement int `db:"total_pending_settlement"`
}

// GetIncomeAmount
// 获取为未结算金额
func GetIncomeAmount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int) (int, error) {

	query := fmt.Sprintf(`
		SELECT IFNULL (
		(
			SELECT
			  SUM(income_amount) AS total_pending_settlement  -- 待结算总金额（单位：分）
			FROM
			  upgrade_company_income
			WHERE
			  company_id = ?  -- 替换为实际的公司ID
			  AND status = 0    -- 0表示待结算
			  AND is_del = 0    -- 0表示未删除
			),
		0)`)

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetIncomeAmountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.TotalPendingSettlement, err

}

type GetSettleAmountStruct struct {
	TotalPendingSettlement int `db:"total_settle_settlement"`
}

// GetSettleAmount
// 获取为已结算金额
func GetSettleAmount(ctx context.Context, svcCtx *svc.ServiceContext, companyId int) (int, error) {

	query := fmt.Sprintf(`
		SELECT IFNULL (
		(
			SELECT
			  SUM(income_amount) AS total_pending_settlement  -- 待结算总金额（单位：分）
			FROM
			  upgrade_company_income
			WHERE
			  company_id = ?  -- 替换为实际的公司ID
			  AND status = 1    -- 0表示待结算  1已结算
			  AND is_del = 0    -- 0表示未删除
			),
		0)`)

	queryContext, err := svcCtx.DB.QueryContext(ctx, query, companyId)
	if err != nil {
		return 0, err
	}
	// 解析结构体
	var info GetSettleAmountStruct
	if err := queryContext.Scan(&info); err != nil {
		return 0, err
	}

	return info.TotalPendingSettlement, err

}
