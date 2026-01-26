package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeAppDownloadReportLogModel = (*customUpgradeAppDownloadReportLogModel)(nil)

type (
	// UpgradeAppDownloadReportLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeAppDownloadReportLogModel.
	UpgradeAppDownloadReportLogModel interface {
		upgradeAppDownloadReportLogModel
		withSession(session sqlx.Session) UpgradeAppDownloadReportLogModel
	}

	customUpgradeAppDownloadReportLogModel struct {
		*defaultUpgradeAppDownloadReportLogModel
	}
)

// NewUpgradeAppDownloadReportLogModel returns a model for the database table.
func NewUpgradeAppDownloadReportLogModel(conn sqlx.SqlConn) UpgradeAppDownloadReportLogModel {
	return &customUpgradeAppDownloadReportLogModel{
		defaultUpgradeAppDownloadReportLogModel: newUpgradeAppDownloadReportLogModel(conn),
	}
}

func (m *customUpgradeAppDownloadReportLogModel) withSession(session sqlx.Session) UpgradeAppDownloadReportLogModel {
	return NewUpgradeAppDownloadReportLogModel(sqlx.NewSqlConnFromSession(session))
}
