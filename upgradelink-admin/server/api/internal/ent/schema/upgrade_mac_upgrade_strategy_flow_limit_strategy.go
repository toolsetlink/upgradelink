package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeMacUpgradeStrategyFlowLimitStrategy struct {
	ent.Schema
}

func (UpgradeMacUpgradeStrategyFlowLimitStrategy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int32("enable").Comment("是否生效；可通过此控制策略是否生效0：失效；1：生效"),
		field.String("begin_time").Comment("开始时间段: 时分秒"),
		field.String("end_time").Comment("结束时间段: 时分秒"),
		field.Int32("dimension").Comment("流控维度；流控维度：1：秒；2：分；3：时；4：天"),
		field.Int("limit").Comment("频控限制；在流控维度上的次数"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeMacUpgradeStrategyFlowLimitStrategy) Edges() []ent.Edge {
	return nil
}

func (UpgradeMacUpgradeStrategyFlowLimitStrategy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_mac_upgrade_strategy_flow_limit_strategy"},
	}
}
