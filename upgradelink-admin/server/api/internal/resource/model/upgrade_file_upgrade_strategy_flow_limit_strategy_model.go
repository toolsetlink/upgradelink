package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeFileUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeFileUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileUpgradeStrategyFlowLimitStrategyModel.
	UpgradeFileUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeFileUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeFileUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeFileUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeFileUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeFileUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeFileUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeFileUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeFileUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeFileUpgradeStrategyFlowLimitStrategyModel: newUpgradeFileUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeFileUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeFileUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeFileUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
