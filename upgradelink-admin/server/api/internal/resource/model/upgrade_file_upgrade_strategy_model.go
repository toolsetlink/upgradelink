package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileUpgradeStrategyModel = (*customUpgradeFileUpgradeStrategyModel)(nil)

type (
	// UpgradeFileUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileUpgradeStrategyModel.
	UpgradeFileUpgradeStrategyModel interface {
		upgradeFileUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeFileUpgradeStrategyModel
	}

	customUpgradeFileUpgradeStrategyModel struct {
		*defaultUpgradeFileUpgradeStrategyModel
	}
)

// NewUpgradeFileUpgradeStrategyModel returns a model for the database table.
func NewUpgradeFileUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeFileUpgradeStrategyModel {
	return &customUpgradeFileUpgradeStrategyModel{
		defaultUpgradeFileUpgradeStrategyModel: newUpgradeFileUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeFileUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeFileUpgradeStrategyModel {
	return NewUpgradeFileUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
