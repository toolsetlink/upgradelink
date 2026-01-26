package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileUpgradeStrategyGrayStrategyModel = (*customUpgradeFileUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeFileUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileUpgradeStrategyGrayStrategyModel.
	UpgradeFileUpgradeStrategyGrayStrategyModel interface {
		upgradeFileUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeFileUpgradeStrategyGrayStrategyModel
	}

	customUpgradeFileUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeFileUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeFileUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeFileUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeFileUpgradeStrategyGrayStrategyModel {
	return &customUpgradeFileUpgradeStrategyGrayStrategyModel{
		defaultUpgradeFileUpgradeStrategyGrayStrategyModel: newUpgradeFileUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeFileUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeFileUpgradeStrategyGrayStrategyModel {
	return NewUpgradeFileUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
