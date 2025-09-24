package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkModel = (*customUpgradeApkModel)(nil)

type (
	// UpgradeApkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkModel.
	UpgradeApkModel interface {
		upgradeApkModel
		withSession(session sqlx.Session) UpgradeApkModel
	}

	customUpgradeApkModel struct {
		*defaultUpgradeApkModel
	}
)

// NewUpgradeApkModel returns a model for the database table.
func NewUpgradeApkModel(conn sqlx.SqlConn) UpgradeApkModel {
	return &customUpgradeApkModel{
		defaultUpgradeApkModel: newUpgradeApkModel(conn),
	}
}

func (m *customUpgradeApkModel) withSession(session sqlx.Session) UpgradeApkModel {
	return NewUpgradeApkModel(sqlx.NewSqlConnFromSession(session))
}
