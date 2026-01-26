package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeDevModelModel = (*customUpgradeDevModelModel)(nil)

type (
	// UpgradeDevModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeDevModelModel.
	UpgradeDevModelModel interface {
		upgradeDevModelModel
		withSession(session sqlx.Session) UpgradeDevModelModel
	}

	customUpgradeDevModelModel struct {
		*defaultUpgradeDevModelModel
	}
)

// NewUpgradeDevModelModel returns a model for the database table.
func NewUpgradeDevModelModel(conn sqlx.SqlConn) UpgradeDevModelModel {
	return &customUpgradeDevModelModel{
		defaultUpgradeDevModelModel: newUpgradeDevModelModel(conn),
	}
}

func (m *customUpgradeDevModelModel) withSession(session sqlx.Session) UpgradeDevModelModel {
	return NewUpgradeDevModelModel(sqlx.NewSqlConnFromSession(session))
}
