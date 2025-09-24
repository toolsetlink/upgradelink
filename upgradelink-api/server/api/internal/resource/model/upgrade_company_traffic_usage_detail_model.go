package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeCompanyTrafficUsageDetailModel = (*customUpgradeCompanyTrafficUsageDetailModel)(nil)

type (
	// UpgradeCompanyTrafficUsageDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeCompanyTrafficUsageDetailModel.
	UpgradeCompanyTrafficUsageDetailModel interface {
		upgradeCompanyTrafficUsageDetailModel
		withSession(session sqlx.Session) UpgradeCompanyTrafficUsageDetailModel
	}

	customUpgradeCompanyTrafficUsageDetailModel struct {
		*defaultUpgradeCompanyTrafficUsageDetailModel
	}
)

// NewUpgradeCompanyTrafficUsageDetailModel returns a model for the database table.
func NewUpgradeCompanyTrafficUsageDetailModel(conn sqlx.SqlConn) UpgradeCompanyTrafficUsageDetailModel {
	return &customUpgradeCompanyTrafficUsageDetailModel{
		defaultUpgradeCompanyTrafficUsageDetailModel: newUpgradeCompanyTrafficUsageDetailModel(conn),
	}
}

func (m *customUpgradeCompanyTrafficUsageDetailModel) withSession(session sqlx.Session) UpgradeCompanyTrafficUsageDetailModel {
	return NewUpgradeCompanyTrafficUsageDetailModel(sqlx.NewSqlConnFromSession(session))
}
