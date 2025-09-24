package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeConfigurationUpgradeStrategyGrayStrategyModel = (*customUpgradeConfigurationUpgradeStrategyGrayStrategyModel)(nil)

type (
	// UpgradeConfigurationUpgradeStrategyGrayStrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeConfigurationUpgradeStrategyGrayStrategyModel.
	UpgradeConfigurationUpgradeStrategyGrayStrategyModel interface {
		upgradeConfigurationUpgradeStrategyGrayStrategyModel
		withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyGrayStrategyModel
	}

	customUpgradeConfigurationUpgradeStrategyGrayStrategyModel struct {
		*defaultUpgradeConfigurationUpgradeStrategyGrayStrategyModel
	}
)

// NewUpgradeConfigurationUpgradeStrategyGrayStrategyModel returns a model for the database table.
func NewUpgradeConfigurationUpgradeStrategyGrayStrategyModel(conn sqlx.SqlConn) UpgradeConfigurationUpgradeStrategyGrayStrategyModel {
	return &customUpgradeConfigurationUpgradeStrategyGrayStrategyModel{
		defaultUpgradeConfigurationUpgradeStrategyGrayStrategyModel: newUpgradeConfigurationUpgradeStrategyGrayStrategyModel(conn),
	}
}

func (m *customUpgradeConfigurationUpgradeStrategyGrayStrategyModel) withSession(session sqlx.Session) UpgradeConfigurationUpgradeStrategyGrayStrategyModel {
	return NewUpgradeConfigurationUpgradeStrategyGrayStrategyModel(sqlx.NewSqlConnFromSession(session))
}
