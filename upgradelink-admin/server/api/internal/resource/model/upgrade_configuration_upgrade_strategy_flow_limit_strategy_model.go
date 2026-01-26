package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel.
	UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeConfigurationUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel: newUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeConfigurationUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
