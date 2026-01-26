package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeConfigurationModel = (*customUpgradeConfigurationModel)(nil)

type (
	// UpgradeConfigurationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeConfigurationModel.
	UpgradeConfigurationModel interface {
		upgradeConfigurationModel
		withSession(session sqlx.Session) UpgradeConfigurationModel
	}

	customUpgradeConfigurationModel struct {
		*defaultUpgradeConfigurationModel
	}
)

// NewUpgradeConfigurationModel returns a model for the database table.
func NewUpgradeConfigurationModel(conn sqlx.SqlConn) UpgradeConfigurationModel {
	return &customUpgradeConfigurationModel{
		defaultUpgradeConfigurationModel: newUpgradeConfigurationModel(conn),
	}
}

func (m *customUpgradeConfigurationModel) withSession(session sqlx.Session) UpgradeConfigurationModel {
	return NewUpgradeConfigurationModel(sqlx.NewSqlConnFromSession(session))
}
