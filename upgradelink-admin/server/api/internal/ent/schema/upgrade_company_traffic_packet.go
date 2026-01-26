package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeCompanyTrafficPacket struct {
	ent.Schema
}

func (UpgradeCompanyTrafficPacket) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("记录ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int("packet_id").Comment("流量包ID"),
		field.Time("start_time").Comment("开始时间"),
		field.Time("end_time").Comment("结束时间"),
		field.Int("initial_size").Comment("初始流量大小(字节)"),
		field.Int("remaining_size").Comment("剩余流量大小(字节)"),
		field.Int32("status").Comment("状态: 1=有效, 0=已过期, 2=已用完"),
		field.Time("exchange_time").Comment("兑换流量包时间"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeCompanyTrafficPacket) Edges() []ent.Edge {
	return nil
}

func (UpgradeCompanyTrafficPacket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_company_traffic_packet"},
	}
}
