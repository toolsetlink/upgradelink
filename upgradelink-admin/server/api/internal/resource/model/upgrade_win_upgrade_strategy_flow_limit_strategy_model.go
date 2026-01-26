package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeWinUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeWinUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeWinUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeWinUpgradeStrategyFlowLimitStrategyModel.
	UpgradeWinUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeWinUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeWinUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeWinUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeWinUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeWinUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeWinUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeWinUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeWinUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeWinUpgradeStrategyFlowLimitStrategyModel: newUpgradeWinUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeWinUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeWinUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeWinUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
