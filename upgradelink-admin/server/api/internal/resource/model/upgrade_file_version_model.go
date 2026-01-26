package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileVersionModel = (*customUpgradeFileVersionModel)(nil)

type (
	// UpgradeFileVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileVersionModel.
	UpgradeFileVersionModel interface {
		upgradeFileVersionModel
		withSession(session sqlx.Session) UpgradeFileVersionModel
	}

	customUpgradeFileVersionModel struct {
		*defaultUpgradeFileVersionModel
	}
)

// NewUpgradeFileVersionModel returns a model for the database table.
func NewUpgradeFileVersionModel(conn sqlx.SqlConn) UpgradeFileVersionModel {
	return &customUpgradeFileVersionModel{
		defaultUpgradeFileVersionModel: newUpgradeFileVersionModel(conn),
	}
}

func (m *customUpgradeFileVersionModel) withSession(session sqlx.Session) UpgradeFileVersionModel {
	return NewUpgradeFileVersionModel(sqlx.NewSqlConnFromSession(session))
}
