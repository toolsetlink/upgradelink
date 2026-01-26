package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileReqLogModel = (*customUpgradeFileReqLogModel)(nil)

type (
	// UpgradeFileReqLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileReqLogModel.
	UpgradeFileReqLogModel interface {
		upgradeFileReqLogModel
		withSession(session sqlx.Session) UpgradeFileReqLogModel
	}

	customUpgradeFileReqLogModel struct {
		*defaultUpgradeFileReqLogModel
	}
)

// NewUpgradeFileReqLogModel returns a model for the database table.
func NewUpgradeFileReqLogModel(conn sqlx.SqlConn) UpgradeFileReqLogModel {
	return &customUpgradeFileReqLogModel{
		defaultUpgradeFileReqLogModel: newUpgradeFileReqLogModel(conn),
	}
}

func (m *customUpgradeFileReqLogModel) withSession(session sqlx.Session) UpgradeFileReqLogModel {
	return NewUpgradeFileReqLogModel(sqlx.NewSqlConnFromSession(session))
}
