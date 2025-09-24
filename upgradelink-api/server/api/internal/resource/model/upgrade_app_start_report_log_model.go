package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppStartReportLogModel = (*customUpgradeAppStartReportLogModel)(nil)

type (
	// UpgradeAppStartReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppStartReportLogModel.
	UpgradeAppStartReportLogModel interface {
		upgradeAppStartReportLogModel
		withSession(session sqlx.Session) UpgradeAppStartReportLogModel
	}

	customUpgradeAppStartReportLogModel struct {
		*defaultUpgradeAppStartReportLogModel
	}
)

// NewUpgradeAppStartReportLogModel returns a model for the database table.
func NewUpgradeAppStartReportLogModel(conn sqlx.SqlConn) UpgradeAppStartReportLogModel {
	return &customUpgradeAppStartReportLogModel{
		defaultUpgradeAppStartReportLogModel: newUpgradeAppStartReportLogModel(conn),
	}
}

func (m *customUpgradeAppStartReportLogModel) withSession(session sqlx.Session) UpgradeAppStartReportLogModel {
	return NewUpgradeAppStartReportLogModel(sqlx.NewSqlConnFromSession(session))
}
