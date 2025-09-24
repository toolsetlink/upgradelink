package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppUpgradeUpgradeReportLogModel = (*customUpgradeAppUpgradeUpgradeReportLogModel)(nil)

type (
	// UpgradeAppUpgradeUpgradeReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppUpgradeUpgradeReportLogModel.
	UpgradeAppUpgradeUpgradeReportLogModel interface {
		upgradeAppUpgradeUpgradeReportLogModel
		withSession(session sqlx.Session) UpgradeAppUpgradeUpgradeReportLogModel
	}

	customUpgradeAppUpgradeUpgradeReportLogModel struct {
		*defaultUpgradeAppUpgradeUpgradeReportLogModel
	}
)

// NewUpgradeAppUpgradeUpgradeReportLogModel returns a model for the database table.
func NewUpgradeAppUpgradeUpgradeReportLogModel(conn sqlx.SqlConn) UpgradeAppUpgradeUpgradeReportLogModel {
	return &customUpgradeAppUpgradeUpgradeReportLogModel{
		defaultUpgradeAppUpgradeUpgradeReportLogModel: newUpgradeAppUpgradeUpgradeReportLogModel(conn),
	}
}

func (m *customUpgradeAppUpgradeUpgradeReportLogModel) withSession(session sqlx.Session) UpgradeAppUpgradeUpgradeReportLogModel {
	return NewUpgradeAppUpgradeUpgradeReportLogModel(sqlx.NewSqlConnFromSession(session))
}
