package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeApkPatchModel = (*customUpgradeApkPatchModel)(nil)

type (
	// UpgradeApkPatchModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeApkPatchModel.
	UpgradeApkPatchModel interface {
		upgradeApkPatchModel
		withSession(session sqlx.Session) UpgradeApkPatchModel
	}

	customUpgradeApkPatchModel struct {
		*defaultUpgradeApkPatchModel
	}
)

// NewUpgradeApkPatchModel returns a model for the database table.
func NewUpgradeApkPatchModel(conn sqlx.SqlConn) UpgradeApkPatchModel {
	return &customUpgradeApkPatchModel{
		defaultUpgradeApkPatchModel: newUpgradeApkPatchModel(conn),
	}
}

func (m *customUpgradeApkPatchModel) withSession(session sqlx.Session) UpgradeApkPatchModel {
	return NewUpgradeApkPatchModel(sqlx.NewSqlConnFromSession(session))
}
