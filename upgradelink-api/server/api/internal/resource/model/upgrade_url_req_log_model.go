package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlReqLogModel = (*customUpgradeUrlReqLogModel)(nil)

type (
	// UpgradeUrlReqLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlReqLogModel.
	UpgradeUrlReqLogModel interface {
		upgradeUrlReqLogModel
		withSession(session sqlx.Session) UpgradeUrlReqLogModel
	}

	customUpgradeUrlReqLogModel struct {
		*defaultUpgradeUrlReqLogModel
	}
)

// NewUpgradeUrlReqLogModel returns a model for the database table.
func NewUpgradeUrlReqLogModel(conn sqlx.SqlConn) UpgradeUrlReqLogModel {
	return &customUpgradeUrlReqLogModel{
		defaultUpgradeUrlReqLogModel: newUpgradeUrlReqLogModel(conn),
	}
}

func (m *customUpgradeUrlReqLogModel) withSession(session sqlx.Session) UpgradeUrlReqLogModel {
	return NewUpgradeUrlReqLogModel(sqlx.NewSqlConnFromSession(session))
}
