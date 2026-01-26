package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeWinModel = (*customUpgradeWinModel)(nil)

type (
	// UpgradeWinModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeWinModel.
	UpgradeWinModel interface {
		upgradeWinModel
		withSession(session sqlx.Session) UpgradeWinModel
	}

	customUpgradeWinModel struct {
		*defaultUpgradeWinModel
	}
)

// NewUpgradeWinModel returns a model for the database table.
func NewUpgradeWinModel(conn sqlx.SqlConn) UpgradeWinModel {
	return &customUpgradeWinModel{
		defaultUpgradeWinModel: newUpgradeWinModel(conn),
	}
}

func (m *customUpgradeWinModel) withSession(session sqlx.Session) UpgradeWinModel {
	return NewUpgradeWinModel(sqlx.NewSqlConnFromSession(session))
}
