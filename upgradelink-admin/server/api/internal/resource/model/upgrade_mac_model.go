package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeMacModel = (*customUpgradeMacModel)(nil)

type (
	// UpgradeMacModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeMacModel.
	UpgradeMacModel interface {
		upgradeMacModel
		withSession(session sqlx.Session) UpgradeMacModel
	}

	customUpgradeMacModel struct {
		*defaultUpgradeMacModel
	}
)

// NewUpgradeMacModel returns a model for the database table.
func NewUpgradeMacModel(conn sqlx.SqlConn) UpgradeMacModel {
	return &customUpgradeMacModel{
		defaultUpgradeMacModel: newUpgradeMacModel(conn),
	}
}

func (m *customUpgradeMacModel) withSession(session sqlx.Session) UpgradeMacModel {
	return NewUpgradeMacModel(sqlx.NewSqlConnFromSession(session))
}
