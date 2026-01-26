package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeWinUpgradeStrategyModel = (*customUpgradeWinUpgradeStrategyModel)(nil)

type (
	// UpgradeWinUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeWinUpgradeStrategyModel.
	UpgradeWinUpgradeStrategyModel interface {
		upgradeWinUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeWinUpgradeStrategyModel
	}

	customUpgradeWinUpgradeStrategyModel struct {
		*defaultUpgradeWinUpgradeStrategyModel
	}
)

// NewUpgradeWinUpgradeStrategyModel returns a model for the database table.
func NewUpgradeWinUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeWinUpgradeStrategyModel {
	return &customUpgradeWinUpgradeStrategyModel{
		defaultUpgradeWinUpgradeStrategyModel: newUpgradeWinUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeWinUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeWinUpgradeStrategyModel {
	return NewUpgradeWinUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
