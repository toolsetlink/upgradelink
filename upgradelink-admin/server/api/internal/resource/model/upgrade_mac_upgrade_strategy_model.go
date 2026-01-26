package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeMacUpgradeStrategyModel = (*customUpgradeMacUpgradeStrategyModel)(nil)

type (
	// UpgradeMacUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeMacUpgradeStrategyModel.
	UpgradeMacUpgradeStrategyModel interface {
		upgradeMacUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeMacUpgradeStrategyModel
	}

	customUpgradeMacUpgradeStrategyModel struct {
		*defaultUpgradeMacUpgradeStrategyModel
	}
)

// NewUpgradeMacUpgradeStrategyModel returns a model for the database table.
func NewUpgradeMacUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeMacUpgradeStrategyModel {
	return &customUpgradeMacUpgradeStrategyModel{
		defaultUpgradeMacUpgradeStrategyModel: newUpgradeMacUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeMacUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeMacUpgradeStrategyModel {
	return NewUpgradeMacUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
