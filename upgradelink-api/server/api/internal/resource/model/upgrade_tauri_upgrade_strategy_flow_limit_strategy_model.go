package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeTauriUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeTauriUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeTauriUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeTauriUpgradeStrategyFlowLimitStrategyModel.
	UpgradeTauriUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeTauriUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeTauriUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeTauriUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeTauriUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeTauriUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeTauriUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeTauriUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeTauriUpgradeStrategyFlowLimitStrategyModel: newUpgradeTauriUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeTauriUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeTauriUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
