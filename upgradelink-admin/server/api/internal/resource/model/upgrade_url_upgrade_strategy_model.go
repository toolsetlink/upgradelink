package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlUpgradeStrategyModel = (*customUpgradeUrlUpgradeStrategyModel)(nil)

type (
	// UpgradeUrlUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlUpgradeStrategyModel.
	UpgradeUrlUpgradeStrategyModel interface {
		upgradeUrlUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyModel
	}

	customUpgradeUrlUpgradeStrategyModel struct {
		*defaultUpgradeUrlUpgradeStrategyModel
	}
)

// NewUpgradeUrlUpgradeStrategyModel returns a model for the database table.
func NewUpgradeUrlUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeUrlUpgradeStrategyModel {
	return &customUpgradeUrlUpgradeStrategyModel{
		defaultUpgradeUrlUpgradeStrategyModel: newUpgradeUrlUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeUrlUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyModel {
	return NewUpgradeUrlUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
