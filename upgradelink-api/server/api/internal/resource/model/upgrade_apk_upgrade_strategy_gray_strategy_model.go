package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkUpgradeStrategyGrayStrategyModel = (*customUpgradeApkUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeApkUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkUpgradeStrategyGrayStrategyModel.
	UpgradeApkUpgradeStrategyGrayStrategyModel interface {
		upgradeApkUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeApkUpgradeStrategyGrayStrategyModel
	}

	customUpgradeApkUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeApkUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeApkUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeApkUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeApkUpgradeStrategyGrayStrategyModel {
	return &customUpgradeApkUpgradeStrategyGrayStrategyModel{
		defaultUpgradeApkUpgradeStrategyGrayStrategyModel: newUpgradeApkUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeApkUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeApkUpgradeStrategyGrayStrategyModel {
	return NewUpgradeApkUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
