package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeTauriUpgradeStrategyGrayStrategyModel = (*customUpgradeTauriUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeTauriUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeTauriUpgradeStrategyGrayStrategyModel.
	UpgradeTauriUpgradeStrategyGrayStrategyModel interface {
		upgradeTauriUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyGrayStrategyModel
	}

	customUpgradeTauriUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeTauriUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeTauriUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeTauriUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeTauriUpgradeStrategyGrayStrategyModel {
	return &customUpgradeTauriUpgradeStrategyGrayStrategyModel{
		defaultUpgradeTauriUpgradeStrategyGrayStrategyModel: newUpgradeTauriUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeTauriUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeTauriUpgradeStrategyGrayStrategyModel {
	return NewUpgradeTauriUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
