package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkUpgradeStrategyModel = (*customUpgradeApkUpgradeStrategyModel)(nil)

type (
	// UpgradeApkUpgradeStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkUpgradeStrategyModel.
	UpgradeApkUpgradeStrategyModel interface {
		upgradeApkUpgradeStrategyModel
		withSession(session sqlx.Session) UpgradeApkUpgradeStrategyModel
	}

	customUpgradeApkUpgradeStrategyModel struct {
		*defaultUpgradeApkUpgradeStrategyModel
	}
)

// NewUpgradeApkUpgradeStrategyModel returns a model for the database table.
func NewUpgradeApkUpgradeStrategyModel(conn sqlx.SqlConn) UpgradeApkUpgradeStrategyModel {
	return &customUpgradeApkUpgradeStrategyModel{
		defaultUpgradeApkUpgradeStrategyModel: newUpgradeApkUpgradeStrategyModel(conn),
	}
}

func (m *customUpgradeApkUpgradeStrategyModel) withSession(session sqlx.Session) UpgradeApkUpgradeStrategyModel {
	return NewUpgradeApkUpgradeStrategyModel(sqlx.NewSqlConnFromSession(session))
}
