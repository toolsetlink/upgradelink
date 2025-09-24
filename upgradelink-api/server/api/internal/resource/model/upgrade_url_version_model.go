package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlVersionModel = (*customUpgradeUrlVersionModel)(nil)

type (
	// UpgradeUrlVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlVersionModel.
	UpgradeUrlVersionModel interface {
		upgradeUrlVersionModel
		withSession(session sqlx.Session) UpgradeUrlVersionModel
	}

	customUpgradeUrlVersionModel struct {
		*defaultUpgradeUrlVersionModel
	}
)

// NewUpgradeUrlVersionModel returns a model for the database table.
func NewUpgradeUrlVersionModel(conn sqlx.SqlConn) UpgradeUrlVersionModel {
	return &customUpgradeUrlVersionModel{
		defaultUpgradeUrlVersionModel: newUpgradeUrlVersionModel(conn),
	}
}

func (m *customUpgradeUrlVersionModel) withSession(session sqlx.Session) UpgradeUrlVersionModel {
	return NewUpgradeUrlVersionModel(sqlx.NewSqlConnFromSession(session))
}
