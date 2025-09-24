package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeTrafficPacket struct {
	ent.Schema
}

func (UpgradeTrafficPacket) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("流量包ID"),
		field.String("name").Comment("流量包名称"),
		field.String("key").Comment("流量包唯一标识"),
		field.Int("size").Comment("流量大小(单位:字节)"),
		field.Int32("price").Comment("价格"),
		field.Int32("valid_days").Comment("有效期(天)"),
		field.Int32("status").Comment("状态: 1=有效, 2=已兑换"),
		field.String("description").Optional().Comment("描述"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeTrafficPacket) Edges() []ent.Edge {
	return nil
}

func (UpgradeTrafficPacket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_traffic_packet"},
	}
}
