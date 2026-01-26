package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeTauriUpgradeStrategyModel = (*customUpgradeTauriUpgradeStrategyModel)(nil)

type (
	// UpgradeTauriUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeTauriUpgradeStrategyModel.
	UpgradeTauriUpgradeStrategyModel interface {
		upgradeTauriUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyModel
	}

	customUpgradeTauriUpgradeStrategyModel struct {
		*defaultUpgradeTauriUpgradeStrategyModel
	}
)

// NewUpgradeTauriUpgradeStrategyModel returns a model for the database table.
func NewUpgradeTauriUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeTauriUpgradeStrategyModel {
	return &customUpgradeTauriUpgradeStrategyModel{
		defaultUpgradeTauriUpgradeStrategyModel: newUpgradeTauriUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeTauriUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyModel {
	return NewUpgradeTauriUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
