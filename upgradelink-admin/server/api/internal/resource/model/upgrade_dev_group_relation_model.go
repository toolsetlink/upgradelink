package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeDevGroupRelationModel = (*customUpgradeDevGroupRelationModel)(nil)

type (
	// UpgradeDevGroupRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeDevGroupRelationModel.
	UpgradeDevGroupRelationModel interface {
		upgradeDevGroupRelationModel
		withSession(session sqlx.Session) UpgradeDevGroupRelationModel
	}

	customUpgradeDevGroupRelationModel struct {
		*defaultUpgradeDevGroupRelationModel
	}
)

// NewUpgradeDevGroupRelationModel returns a model for the database table.
func NewUpgradeDevGroupRelationModel(conn sqlx.SqlConn) UpgradeDevGroupRelationModel {
	return &customUpgradeDevGroupRelationModel{
		defaultUpgradeDevGroupRelationModel: newUpgradeDevGroupRelationModel(conn),
	}
}

func (m *customUpgradeDevGroupRelationModel) withSession(session sqlx.Session) UpgradeDevGroupRelationModel {
	return NewUpgradeDevGroupRelationModel(sqlx.NewSqlConnFromSession(session))
}
