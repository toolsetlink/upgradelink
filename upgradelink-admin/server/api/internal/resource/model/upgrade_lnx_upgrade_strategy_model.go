package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeLnxUpgradeStrategyModel = (*customUpgradeLnxUpgradeStrategyModel)(nil)

type (
	// UpgradeLnxUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeLnxUpgradeStrategyModel.
	UpgradeLnxUpgradeStrategyModel interface {
		upgradeLnxUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyModel
	}

	customUpgradeLnxUpgradeStrategyModel struct {
		*defaultUpgradeLnxUpgradeStrategyModel
	}
)

// NewUpgradeLnxUpgradeStrategyModel returns a model for the database table.
func NewUpgradeLnxUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeLnxUpgradeStrategyModel {
	return &customUpgradeLnxUpgradeStrategyModel{
		defaultUpgradeLnxUpgradeStrategyModel: newUpgradeLnxUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeLnxUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyModel {
	return NewUpgradeLnxUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
