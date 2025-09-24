package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeDevsModel = (*customUpgradeDevsModel)(nil)

type (
	// UpgradeDevsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeDevsModel.
	UpgradeDevsModel interface {
		upgradeDevsModel
		withSession(session sqlx.Session) UpgradeDevsModel
	}

	customUpgradeDevsModel struct {
		*defaultUpgradeDevsModel
	}
)

// NewUpgradeDevsModel returns a model for the database table.
func NewUpgradeDevsModel(conn sqlx.SqlConn) UpgradeDevsModel {
	return &customUpgradeDevsModel{
		defaultUpgradeDevsModel: newUpgradeDevsModel(conn),
	}
}

func (m *customUpgradeDevsModel) withSession(session sqlx.Session) UpgradeDevsModel {
	return NewUpgradeDevsModel(sqlx.NewSqlConnFromSession(session))
}
