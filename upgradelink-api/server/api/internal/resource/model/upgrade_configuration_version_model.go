package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeConfigurationVersionModel = (*customUpgradeConfigurationVersionModel)(nil)

type (
	// UpgradeConfigurationVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeConfigurationVersionModel.
	UpgradeConfigurationVersionModel interface {
		upgradeConfigurationVersionModel
		withSession(session sqlx.Session) UpgradeConfigurationVersionModel
	}

	customUpgradeConfigurationVersionModel struct {
		*defaultUpgradeConfigurationVersionModel
	}
)

// NewUpgradeConfigurationVersionModel returns a model for the database table.
func NewUpgradeConfigurationVersionModel(conn sqlx.SqlConn) UpgradeConfigurationVersionModel {
	return &customUpgradeConfigurationVersionModel{
		defaultUpgradeConfigurationVersionModel: newUpgradeConfigurationVersionModel(conn),
	}
}

func (m *customUpgradeConfigurationVersionModel) withSession(session sqlx.Session) UpgradeConfigurationVersionModel {
	return NewUpgradeConfigurationVersionModel(sqlx.NewSqlConnFromSession(session))
}
