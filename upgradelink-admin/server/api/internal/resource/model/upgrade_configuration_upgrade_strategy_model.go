package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeConfigurationUpgradeStrategyModel = (*customUpgradeConfigurationUpgradeStrategyModel)(nil)

type (
	// UpgradeConfigurationUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeConfigurationUpgradeStrategyModel.
	UpgradeConfigurationUpgradeStrategyModel interface {
		upgradeConfigurationUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyModel
	}

	customUpgradeConfigurationUpgradeStrategyModel struct {
		*defaultUpgradeConfigurationUpgradeStrategyModel
	}
)

// NewUpgradeConfigurationUpgradeStrategyModel returns a model for the database table.
func NewUpgradeConfigurationUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeConfigurationUpgradeStrategyModel {
	return &customUpgradeConfigurationUpgradeStrategyModel{
		defaultUpgradeConfigurationUpgradeStrategyModel: newUpgradeConfigurationUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeConfigurationUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyModel {
	return NewUpgradeConfigurationUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
