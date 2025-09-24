package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkVersionModel = (*customUpgradeApkVersionModel)(nil)

type (
	// UpgradeApkVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkVersionModel.
	UpgradeApkVersionModel interface {
		upgradeApkVersionModel
		withSession(session sqlx.Session) UpgradeApkVersionModel
	}

	customUpgradeApkVersionModel struct {
		*defaultUpgradeApkVersionModel
	}
)

// NewUpgradeApkVersionModel returns a model for the database table.
func NewUpgradeApkVersionModel(conn sqlx.SqlConn) UpgradeApkVersionModel {
	return &customUpgradeApkVersionModel{
		defaultUpgradeApkVersionModel: newUpgradeApkVersionModel(conn),
	}
}

func (m *customUpgradeApkVersionModel) withSession(session sqlx.Session) UpgradeApkVersionModel {
	return NewUpgradeApkVersionModel(sqlx.NewSqlConnFromSession(session))
}
