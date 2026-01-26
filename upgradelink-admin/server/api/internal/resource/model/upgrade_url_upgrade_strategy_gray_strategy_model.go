package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeUrlUpgradeStrategyGrayStrategyModel = (*customUpgradeUrlUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeUrlUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeUrlUpgradeStrategyGrayStrategyModel.
	UpgradeUrlUpgradeStrategyGrayStrategyModel interface {
		upgradeUrlUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyGrayStrategyModel
	}

	customUpgradeUrlUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeUrlUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeUrlUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeUrlUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeUrlUpgradeStrategyGrayStrategyModel {
	return &customUpgradeUrlUpgradeStrategyGrayStrategyModel{
		defaultUpgradeUrlUpgradeStrategyGrayStrategyModel: newUpgradeUrlUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeUrlUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeUrlUpgradeStrategyGrayStrategyModel {
	return NewUpgradeUrlUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
