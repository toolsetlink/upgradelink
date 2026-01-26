package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeMacVersionModel = (*customUpgradeMacVersionModel)(nil)

type (
	// UpgradeMacVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeMacVersionModel.
	UpgradeMacVersionModel interface {
		upgradeMacVersionModel
		withSession(session sqlx.Session) UpgradeMacVersionModel
	}

	customUpgradeMacVersionModel struct {
		*defaultUpgradeMacVersionModel
	}
)

// NewUpgradeMacVersionModel returns a model for the database table.
func NewUpgradeMacVersionModel(conn sqlx.SqlConn) UpgradeMacVersionModel {
	return &customUpgradeMacVersionModel{
		defaultUpgradeMacVersionModel: newUpgradeMacVersionModel(conn),
	}
}

func (m *customUpgradeMacVersionModel) withSession(session sqlx.Session) UpgradeMacVersionModel {
	return NewUpgradeMacVersionModel(sqlx.NewSqlConnFromSession(session))
}
