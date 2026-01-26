package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeElectronUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeElectronUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeElectronUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeElectronUpgradeStrategyFlowLimitStrategyModel.
	UpgradeElectronUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeElectronUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeElectronUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeElectronUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeElectronUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeElectronUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeElectronUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeElectronUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeElectronUpgradeStrategyFlowLimitStrategyModel: newUpgradeElectronUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeElectronUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeElectronUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
