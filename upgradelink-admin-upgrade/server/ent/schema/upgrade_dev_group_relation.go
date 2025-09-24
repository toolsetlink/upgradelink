package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeDevGroupRelation struct {
	ent.Schema
}

func (UpgradeDevGroupRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("dev_id").Comment("设备id"),
		field.Int("dev_group_id").Comment("设备分组 id"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间"),
	}
}

func (UpgradeDevGroupRelation) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("upgrade_devs", UpgradeDev.Type).
		//	Unique().
		//	Required().
		//	Field("dev_id").
		//	Comment("设备ID"),
		//edge.To("upgrade_dev_groups", UpgradeDevGroup.Type).
		//	Unique().
		//	Required().
		//	Field("dev_group_id").
		//	Comment("设备分组ID"),
	}
}

func (UpgradeDevGroupRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_dev_group_relation"},
	}
}
