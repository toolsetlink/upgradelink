package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppReportLogModel = (*customUpgradeAppReportLogModel)(nil)

type (
	// UpgradeAppReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppReportLogModel.
	UpgradeAppReportLogModel interface {
		upgradeAppReportLogModel
		withSession(session sqlx.Session) UpgradeAppReportLogModel
	}

	customUpgradeAppReportLogModel struct {
		*defaultUpgradeAppReportLogModel
	}
)

// NewUpgradeAppReportLogModel returns a model for the database table.
func NewUpgradeAppReportLogModel(conn sqlx.SqlConn) UpgradeAppReportLogModel {
	return &customUpgradeAppReportLogModel{
		defaultUpgradeAppReportLogModel: newUpgradeAppReportLogModel(conn),
	}
}

func (m *customUpgradeAppReportLogModel) withSession(session sqlx.Session) UpgradeAppReportLogModel {
	return NewUpgradeAppReportLogModel(sqlx.NewSqlConnFromSession(session))
}
