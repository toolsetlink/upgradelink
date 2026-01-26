package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FmsCloudFilesModel = (*customFmsCloudFilesModel)(nil)

type (
	// FmsCloudFilesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFmsCloudFilesModel.
	FmsCloudFilesModel interface {
		fmsCloudFilesModel
		withSession(session sqlx.Session) FmsCloudFilesModel
	}

	customFmsCloudFilesModel struct {
		*defaultFmsCloudFilesModel
	}
)

// NewFmsCloudFilesModel returns a model for the database table.
func NewFmsCloudFilesModel(conn sqlx.SqlConn) FmsCloudFilesModel {
	return &customFmsCloudFilesModel{
		defaultFmsCloudFilesModel: newFmsCloudFilesModel(conn),
	}
}

func (m *customFmsCloudFilesModel) withSession(session sqlx.Session) FmsCloudFilesModel {
	return NewFmsCloudFilesModel(sqlx.NewSqlConnFromSession(session))
}
