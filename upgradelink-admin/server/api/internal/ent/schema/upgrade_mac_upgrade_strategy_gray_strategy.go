package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeMacUpgradeStrategyGrayStrategy struct {
	ent.Schema
}

func (UpgradeMacUpgradeStrategyGrayStrategy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int32("enable").Comment("是否生效；可通过此控制策略是否生效0：失效；1：生效"),
		field.Time("begin_datetime").Comment("开始时间"),
		field.Time("end_datetime").Comment("结束时间"),
		field.Int("limit").Comment("数量限制"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeMacUpgradeStrategyGrayStrategy) Edges() []ent.Edge {
	return nil
}

func (UpgradeMacUpgradeStrategyGrayStrategy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_mac_upgrade_strategy_gray_strategy"},
	}
}
