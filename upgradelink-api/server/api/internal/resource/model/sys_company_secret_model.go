package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysCompanySecretModel = (*customSysCompanySecretModel)(nil)

type (
	// SysCompanySecretModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysCompanySecretModel.
	SysCompanySecretModel interface {
		sysCompanySecretModel
		withSession(session sqlx.Session) SysCompanySecretModel
	}

	customSysCompanySecretModel struct {
		*defaultSysCompanySecretModel
	}
)

// NewSysCompanySecretModel returns a model for the database table.
func NewSysCompanySecretModel(conn sqlx.SqlConn) SysCompanySecretModel {
	return &customSysCompanySecretModel{
		defaultSysCompanySecretModel: newSysCompanySecretModel(conn),
	}
}

func (m *customSysCompanySecretModel) withSession(session sqlx.Session) SysCompanySecretModel {
	return NewSysCompanySecretModel(sqlx.NewSqlConnFromSession(session))
}
