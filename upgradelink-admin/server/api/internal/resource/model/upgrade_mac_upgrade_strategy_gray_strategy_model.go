package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeMacUpgradeStrategyGrayStrategyModel = (*customUpgradeMacUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeMacUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeMacUpgradeStrategyGrayStrategyModel.
	UpgradeMacUpgradeStrategyGrayStrategyModel interface {
		upgradeMacUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeMacUpgradeStrategyGrayStrategyModel
	}

	customUpgradeMacUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeMacUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeMacUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeMacUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeMacUpgradeStrategyGrayStrategyModel {
	return &customUpgradeMacUpgradeStrategyGrayStrategyModel{
		defaultUpgradeMacUpgradeStrategyGrayStrategyModel: newUpgradeMacUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeMacUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeMacUpgradeStrategyGrayStrategyModel {
	return NewUpgradeMacUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
