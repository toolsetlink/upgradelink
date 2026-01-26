package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeLnxModel = (*customUpgradeLnxModel)(nil)

type (
	// UpgradeLnxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeLnxModel.
	UpgradeLnxModel interface {
		upgradeLnxModel
		withSession(session sqlx.Session) UpgradeLnxModel
	}

	customUpgradeLnxModel struct {
		*defaultUpgradeLnxModel
	}
)

// NewUpgradeLnxModel returns a model for the database table.
func NewUpgradeLnxModel(conn sqlx.SqlConn) UpgradeLnxModel {
	return &customUpgradeLnxModel{
		defaultUpgradeLnxModel: newUpgradeLnxModel(conn),
	}
}

func (m *customUpgradeLnxModel) withSession(session sqlx.Session) UpgradeLnxModel {
	return NewUpgradeLnxModel(sqlx.NewSqlConnFromSession(session))
}
