package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeUrlUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeUrlUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlUpgradeStrategyFlowLimitStrategyModel.
	UpgradeUrlUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeUrlUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeUrlUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeUrlUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeUrlUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeUrlUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeUrlUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeUrlUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeUrlUpgradeStrategyFlowLimitStrategyModel: newUpgradeUrlUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeUrlUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeUrlUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
