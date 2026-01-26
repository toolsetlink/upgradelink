package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeWinUpgradeStrategyGrayStrategyModel = (*customUpgradeWinUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeWinUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeWinUpgradeStrategyGrayStrategyModel.
	UpgradeWinUpgradeStrategyGrayStrategyModel interface {
		upgradeWinUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeWinUpgradeStrategyGrayStrategyModel
	}

	customUpgradeWinUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeWinUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeWinUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeWinUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeWinUpgradeStrategyGrayStrategyModel {
	return &customUpgradeWinUpgradeStrategyGrayStrategyModel{
		defaultUpgradeWinUpgradeStrategyGrayStrategyModel: newUpgradeWinUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeWinUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeWinUpgradeStrategyGrayStrategyModel {
	return NewUpgradeWinUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
