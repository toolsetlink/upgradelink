package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeElectronVersionModel = (*customUpgradeElectronVersionModel)(nil)

type (
	// UpgradeElectronVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeElectronVersionModel.
	UpgradeElectronVersionModel interface {
		upgradeElectronVersionModel
		withSession(session sqlx.Session) UpgradeElectronVersionModel
	}

	customUpgradeElectronVersionModel struct {
		*defaultUpgradeElectronVersionModel
	}
)

// NewUpgradeElectronVersionModel returns a model for the database table.
func NewUpgradeElectronVersionModel(conn sqlx.SqlConn) UpgradeElectronVersionModel {
	return &customUpgradeElectronVersionModel{
		defaultUpgradeElectronVersionModel: newUpgradeElectronVersionModel(conn),
	}
}

func (m *customUpgradeElectronVersionModel) withSession(session sqlx.Session) UpgradeElectronVersionModel {
	return NewUpgradeElectronVersionModel(sqlx.NewSqlConnFromSession(session))
}
