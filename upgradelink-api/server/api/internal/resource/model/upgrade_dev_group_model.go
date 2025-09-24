package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeDevGroupModel = (*customUpgradeDevGroupModel)(nil)

type (
	// UpgradeDevGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeDevGroupModel.
	UpgradeDevGroupModel interface {
		upgradeDevGroupModel
		withSession(session sqlx.Session) UpgradeDevGroupModel
	}

	customUpgradeDevGroupModel struct {
		*defaultUpgradeDevGroupModel
	}
)

// NewUpgradeDevGroupModel returns a model for the database table.
func NewUpgradeDevGroupModel(conn sqlx.SqlConn) UpgradeDevGroupModel {
	return &customUpgradeDevGroupModel{
		defaultUpgradeDevGroupModel: newUpgradeDevGroupModel(conn),
	}
}

func (m *customUpgradeDevGroupModel) withSession(session sqlx.Session) UpgradeDevGroupModel {
	return NewUpgradeDevGroupModel(sqlx.NewSqlConnFromSession(session))
}
