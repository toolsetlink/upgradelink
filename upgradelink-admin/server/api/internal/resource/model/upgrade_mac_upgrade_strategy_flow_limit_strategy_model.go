package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeMacUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeMacUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeMacUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeMacUpgradeStrategyFlowLimitStrategyModel.
	UpgradeMacUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeMacUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeMacUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeMacUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeMacUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeMacUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeMacUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeMacUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeMacUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeMacUpgradeStrategyFlowLimitStrategyModel: newUpgradeMacUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeMacUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeMacUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeMacUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
