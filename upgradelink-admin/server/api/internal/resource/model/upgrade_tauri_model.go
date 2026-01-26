package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeTauriModel = (*customUpgradeTauriModel)(nil)

type (
	// UpgradeTauriModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeTauriModel.
	UpgradeTauriModel interface {
		upgradeTauriModel
		withSession(session sqlx.Session) UpgradeTauriModel
	}

	customUpgradeTauriModel struct {
		*defaultUpgradeTauriModel
	}
)

// NewUpgradeTauriModel returns a model for the database table.
func NewUpgradeTauriModel(conn sqlx.SqlConn) UpgradeTauriModel {
	return &customUpgradeTauriModel{
		defaultUpgradeTauriModel: newUpgradeTauriModel(conn),
	}
}

func (m *customUpgradeTauriModel) withSession(session sqlx.Session) UpgradeTauriModel {
	return NewUpgradeTauriModel(sqlx.NewSqlConnFromSession(session))
}
