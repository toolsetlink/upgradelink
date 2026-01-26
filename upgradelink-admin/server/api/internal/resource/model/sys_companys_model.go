package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysCompanysModel = (*customSysCompanysModel)(nil)

type (
	// SysCompanysModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysCompanysModel.
	SysCompanysModel interface {
		sysCompanysModel
		withSession(session sqlx.Session) SysCompanysModel
	}

	customSysCompanysModel struct {
		*defaultSysCompanysModel
	}
)

// NewSysCompanysModel returns a model for the database table.
func NewSysCompanysModel(conn sqlx.SqlConn) SysCompanysModel {
	return &customSysCompanysModel{
		defaultSysCompanysModel: newSysCompanysModel(conn),
	}
}

func (m *customSysCompanysModel) withSession(session sqlx.Session) SysCompanysModel {
	return NewSysCompanysModel(sqlx.NewSqlConnFromSession(session))
}
