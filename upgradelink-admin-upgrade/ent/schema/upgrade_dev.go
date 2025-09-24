package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeDev struct {
	ent.Schema
}

func (UpgradeDev) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.String("key").Comment("设备唯一标识"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}

func (UpgradeDev) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("upgrade_dev_groups", UpgradeDevGroup.Type).Through("upgrade_dev_group_relations", UpgradeDevGroupRelation.Type),
	}
}

func (UpgradeDev) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_devs"},
	}
}
