package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeFileGithubModel = (*customUpgradeFileGithubModel)(nil)

type (
	// UpgradeFileGithubModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeFileGithubModel.
	UpgradeFileGithubModel interface {
		upgradeFileGithubModel
		withSession(session sqlx.Session) UpgradeFileGithubModel
	}

	customUpgradeFileGithubModel struct {
		*defaultUpgradeFileGithubModel
	}
)

// NewUpgradeFileGithubModel returns a model for the database table.
func NewUpgradeFileGithubModel(conn sqlx.SqlConn) UpgradeFileGithubModel {
	return &customUpgradeFileGithubModel{
		defaultUpgradeFileGithubModel: newUpgradeFileGithubModel(conn),
	}
}

func (m *customUpgradeFileGithubModel) withSession(session sqlx.Session) UpgradeFileGithubModel {
	return NewUpgradeFileGithubModel(sqlx.NewSqlConnFromSession(session))
}
