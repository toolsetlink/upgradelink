package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileModel = (*customUpgradeFileModel)(nil)

type (
	// UpgradeFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileModel.
	UpgradeFileModel interface {
		upgradeFileModel
		withSession(session sqlx.Session) UpgradeFileModel
	}

	customUpgradeFileModel struct {
		*defaultUpgradeFileModel
	}
)

// NewUpgradeFileModel returns a model for the database table.
func NewUpgradeFileModel(conn sqlx.SqlConn) UpgradeFileModel {
	return &customUpgradeFileModel{
		defaultUpgradeFileModel: newUpgradeFileModel(conn),
	}
}

func (m *customUpgradeFileModel) withSession(session sqlx.Session) UpgradeFileModel {
	return NewUpgradeFileModel(sqlx.NewSqlConnFromSession(session))
}
