package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppUpgradeGetStrategyReportLogModel = (*customUpgradeAppUpgradeGetStrategyReportLogModel)(nil)

type (
	// UpgradeAppUpgradeGetStrategyReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppUpgradeGetStrategyReportLogModel.
	UpgradeAppUpgradeGetStrategyReportLogModel interface {
		upgradeAppUpgradeGetStrategyReportLogModel
		withSession(session sqlx.Session) UpgradeAppUpgradeGetStrategyReportLogModel
	}

	customUpgradeAppUpgradeGetStrategyReportLogModel struct {
		*defaultUpgradeAppUpgradeGetStrategyReportLogModel
	}
)

// NewUpgradeAppUpgradeGetStrategyReportLogModel returns a model for the database table.
func NewUpgradeAppUpgradeGetStrategyReportLogModel(conn sqlx.SqlConn) UpgradeAppUpgradeGetStrategyReportLogModel {
	return &customUpgradeAppUpgradeGetStrategyReportLogModel{
		defaultUpgradeAppUpgradeGetStrategyReportLogModel: newUpgradeAppUpgradeGetStrategyReportLogModel(conn),
	}
}

func (m *customUpgradeAppUpgradeGetStrategyReportLogModel) withSession(session sqlx.Session) UpgradeAppUpgradeGetStrategyReportLogModel {
	return NewUpgradeAppUpgradeGetStrategyReportLogModel(sqlx.NewSqlConnFromSession(session))
}
