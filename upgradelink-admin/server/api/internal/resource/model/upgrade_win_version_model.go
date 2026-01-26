package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeWinVersionModel = (*customUpgradeWinVersionModel)(nil)

type (
	// UpgradeWinVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeWinVersionModel.
	UpgradeWinVersionModel interface {
		upgradeWinVersionModel
		withSession(session sqlx.Session) UpgradeWinVersionModel
	}

	customUpgradeWinVersionModel struct {
		*defaultUpgradeWinVersionModel
	}
)

// NewUpgradeWinVersionModel returns a model for the database table.
func NewUpgradeWinVersionModel(conn sqlx.SqlConn) UpgradeWinVersionModel {
	return &customUpgradeWinVersionModel{
		defaultUpgradeWinVersionModel: newUpgradeWinVersionModel(conn),
	}
}

func (m *customUpgradeWinVersionModel) withSession(session sqlx.Session) UpgradeWinVersionModel {
	return NewUpgradeWinVersionModel(sqlx.NewSqlConnFromSession(session))
}
