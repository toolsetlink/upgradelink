package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlModel = (*customUpgradeUrlModel)(nil)

type (
	// UpgradeUrlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlModel.
	UpgradeUrlModel interface {
		upgradeUrlModel
		withSession(session sqlx.Session) UpgradeUrlModel
	}

	customUpgradeUrlModel struct {
		*defaultUpgradeUrlModel
	}
)

// NewUpgradeUrlModel returns a model for the database table.
func NewUpgradeUrlModel(conn sqlx.SqlConn) UpgradeUrlModel {
	return &customUpgradeUrlModel{
		defaultUpgradeUrlModel: newUpgradeUrlModel(conn),
	}
}

func (m *customUpgradeUrlModel) withSession(session sqlx.Session) UpgradeUrlModel {
	return NewUpgradeUrlModel(sqlx.NewSqlConnFromSession(session))
}
