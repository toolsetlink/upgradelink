package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkUpgradeStrategyFlowLimitStrategyModel = (*customUpgradeApkUpgradeStrategyFlowLimitStrategyModel)(nil)

type (
	// UpgradeApkUpgradeStrategyFlowLimitStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkUpgradeStrategyFlowLimitStrategyModel.
	UpgradeApkUpgradeStrategyFlowLimitStrategyModel interface {
		upgradeApkUpgradeStrategyFlowLimitStrategyModel
		withSession(session sqlx.Session) UpgradeApkUpgradeStrategyFlowLimitStrategyModel
	}

	customUpgradeApkUpgradeStrategyFlowLimitStrategyModel struct {
		*defaultUpgradeApkUpgradeStrategyFlowLimitStrategyModel
	}
)

// NewUpgradeApkUpgradeStrategyFlowLimitStrategyModel returns a model for the database table.
func NewUpgradeApkUpgradeStrategyFlowLimitStrategyModel(conn sqlx.SqlConn) UpgradeApkUpgradeStrategyFlowLimitStrategyModel {
	return &customUpgradeApkUpgradeStrategyFlowLimitStrategyModel{
		defaultUpgradeApkUpgradeStrategyFlowLimitStrategyModel: newUpgradeApkUpgradeStrategyFlowLimitStrategyModel(conn),
	}
}

func (m *customUpgradeApkUpgradeStrategyFlowLimitStrategyModel) withSession(session sqlx.Session) UpgradeApkUpgradeStrategyFlowLimitStrategyModel {
	return NewUpgradeApkUpgradeStrategyFlowLimitStrategyModel(sqlx.NewSqlConnFromSession(session))
}
