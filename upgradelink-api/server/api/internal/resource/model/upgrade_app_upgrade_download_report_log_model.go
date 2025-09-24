package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppUpgradeDownloadReportLogModel = (*customUpgradeAppUpgradeDownloadReportLogModel)(nil)

type (
	// UpgradeAppUpgradeDownloadReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppUpgradeDownloadReportLogModel.
	UpgradeAppUpgradeDownloadReportLogModel interface {
		upgradeAppUpgradeDownloadReportLogModel
		withSession(session sqlx.Session) UpgradeAppUpgradeDownloadReportLogModel
	}

	customUpgradeAppUpgradeDownloadReportLogModel struct {
		*defaultUpgradeAppUpgradeDownloadReportLogModel
	}
)

// NewUpgradeAppUpgradeDownloadReportLogModel returns a model for the database table.
func NewUpgradeAppUpgradeDownloadReportLogModel(conn sqlx.SqlConn) UpgradeAppUpgradeDownloadReportLogModel {
	return &customUpgradeAppUpgradeDownloadReportLogModel{
		defaultUpgradeAppUpgradeDownloadReportLogModel: newUpgradeAppUpgradeDownloadReportLogModel(conn),
	}
}

func (m *customUpgradeAppUpgradeDownloadReportLogModel) withSession(session sqlx.Session) UpgradeAppUpgradeDownloadReportLogModel {
	return NewUpgradeAppUpgradeDownloadReportLogModel(sqlx.NewSqlConnFromSession(session))
}
