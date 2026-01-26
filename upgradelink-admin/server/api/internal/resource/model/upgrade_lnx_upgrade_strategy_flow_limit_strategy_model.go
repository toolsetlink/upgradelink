package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeLnxUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeLnxUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeLnxUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeLnxUpgradeStrategyFlowLimitStrategyModel.
	UpgradeLnxUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeLnxUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeLnxUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeLnxUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeLnxUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeLnxUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeLnxUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeLnxUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeLnxUpgradeStrategyFlowLimitStrategyModel: newUpgradeLnxUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeLnxUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeLnxUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
