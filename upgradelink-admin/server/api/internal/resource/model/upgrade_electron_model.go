package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeElectronModel = (*customUpgradeElectronModel)(nil)

type (
	// UpgradeElectronModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeElectronModel.
	UpgradeElectronModel interface {
		upgradeElectronModel
		withSession(session sqlx.Session) UpgradeElectronModel
	}

	customUpgradeElectronModel struct {
		*defaultUpgradeElectronModel
	}
)

// NewUpgradeElectronModel returns a model for the database table.
func NewUpgradeElectronModel(conn sqlx.SqlConn) UpgradeElectronModel {
	return &customUpgradeElectronModel{
		defaultUpgradeElectronModel: newUpgradeElectronModel(conn),
	}
}

func (m *customUpgradeElectronModel) withSession(session sqlx.Session) UpgradeElectronModel {
	return NewUpgradeElectronModel(sqlx.NewSqlConnFromSession(session))
}
