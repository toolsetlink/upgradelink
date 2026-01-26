package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeLnxVersionModel = (*customUpgradeLnxVersionModel)(nil)

type (
	// UpgradeLnxVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeLnxVersionModel.
	UpgradeLnxVersionModel interface {
		upgradeLnxVersionModel
		withSession(session sqlx.Session) UpgradeLnxVersionModel
	}

	customUpgradeLnxVersionModel struct {
		*defaultUpgradeLnxVersionModel
	}
)

// NewUpgradeLnxVersionModel returns a model for the database table.
func NewUpgradeLnxVersionModel(conn sqlx.SqlConn) UpgradeLnxVersionModel {
	return &customUpgradeLnxVersionModel{
		defaultUpgradeLnxVersionModel: newUpgradeLnxVersionModel(conn),
	}
}

func (m *customUpgradeLnxVersionModel) withSession(session sqlx.Session) UpgradeLnxVersionModel {
	return NewUpgradeLnxVersionModel(sqlx.NewSqlConnFromSession(session))
}
