package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeTauriVersionModel = (*customUpgradeTauriVersionModel)(nil)

type (
	// UpgradeTauriVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeTauriVersionModel.
	UpgradeTauriVersionModel interface {
		upgradeTauriVersionModel
		withSession(session sqlx.Session) UpgradeTauriVersionModel
	}

	customUpgradeTauriVersionModel struct {
		*defaultUpgradeTauriVersionModel
	}
)

// NewUpgradeTauriVersionModel returns a model for the database table.
func NewUpgradeTauriVersionModel(conn sqlx.SqlConn) UpgradeTauriVersionModel {
	return &customUpgradeTauriVersionModel{
		defaultUpgradeTauriVersionModel: newUpgradeTauriVersionModel(conn),
	}
}

func (m *customUpgradeTauriVersionModel) withSession(session sqlx.Session) UpgradeTauriVersionModel {
	return NewUpgradeTauriVersionModel(sqlx.NewSqlConnFromSession(session))
}
