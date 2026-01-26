package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeElectronUpgradeStrategyGrayStrategyModel = (*customUpgradeElectronUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeElectronUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeElectronUpgradeStrategyGrayStrategyModel.
	UpgradeElectronUpgradeStrategyGrayStrategyModel interface {
		upgradeElectronUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyGrayStrategyModel
	}

	customUpgradeElectronUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeElectronUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeElectronUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeElectronUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeElectronUpgradeStrategyGrayStrategyModel {
	return &customUpgradeElectronUpgradeStrategyGrayStrategyModel{
		defaultUpgradeElectronUpgradeStrategyGrayStrategyModel: newUpgradeElectronUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeElectronUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeElectronUpgradeStrategyGrayStrategyModel {
	return NewUpgradeElectronUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
