package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeElectronUpgradeStrategyModel = (*customUpgradeElectronUpgradeStrategyModel)(nil)

type (
	// UpgradeElectronUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeElectronUpgradeStrategyModel.
	UpgradeElectronUpgradeStrategyModel interface {
		upgradeElectronUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyModel
	}

	customUpgradeElectronUpgradeStrategyModel struct {
		*defaultUpgradeElectronUpgradeStrategyModel
	}
)

// NewUpgradeElectronUpgradeStrategyModel returns a model for the database table.
func NewUpgradeElectronUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeElectronUpgradeStrategyModel {
	return &customUpgradeElectronUpgradeStrategyModel{
		defaultUpgradeElectronUpgradeStrategyModel: newUpgradeElectronUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeElectronUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyModel {
	return NewUpgradeElectronUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
