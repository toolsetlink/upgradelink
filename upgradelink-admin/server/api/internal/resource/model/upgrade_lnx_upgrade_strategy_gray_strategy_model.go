package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeLnxUpgradeStrategyGrayStrategyModel = (*customUpgradeLnxUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeLnxUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeLnxUpgradeStrategyGrayStrategyModel.
	UpgradeLnxUpgradeStrategyGrayStrategyModel interface {
		upgradeLnxUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyGrayStrategyModel
	}

	customUpgradeLnxUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeLnxUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeLnxUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeLnxUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeLnxUpgradeStrategyGrayStrategyModel {
	return &customUpgradeLnxUpgradeStrategyGrayStrategyModel{
		defaultUpgradeLnxUpgradeStrategyGrayStrategyModel: newUpgradeLnxUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeLnxUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeLnxUpgradeStrategyGrayStrategyModel {
	return NewUpgradeLnxUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
