package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpgradeCompanyTrafficPacketModel = (*customUpgradeCompanyTrafficPacketModel)(nil)

type (
	// UpgradeCompanyTrafficPacketModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpgradeCompanyTrafficPacketModel.
	UpgradeCompanyTrafficPacketModel interface {
		upgradeCompanyTrafficPacketModel
		withSession(session sqlx.Session) UpgradeCompanyTrafficPacketModel
	}

	customUpgradeCompanyTrafficPacketModel struct {
		*defaultUpgradeCompanyTrafficPacketModel
	}
)

// NewUpgradeCompanyTrafficPacketModel returns a model for the database table.
func NewUpgradeCompanyTrafficPacketModel(conn sqlx.SqlConn) UpgradeCompanyTrafficPacketModel {
	return &customUpgradeCompanyTrafficPacketModel{
		defaultUpgradeCompanyTrafficPacketModel: newUpgradeCompanyTrafficPacketModel(conn),
	}
}

func (m *customUpgradeCompanyTrafficPacketModel) withSession(session sqlx.Session) UpgradeCompanyTrafficPacketModel {
	return NewUpgradeCompanyTrafficPacketModel(sqlx.NewSqlConnFromSession(session))
}
