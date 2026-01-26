package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeDevGroup struct {
	ent.Schema
}

func (UpgradeDevGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.String("name").Comment("设备分组名称"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}

func (UpgradeDevGroup) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("upgrade_devs", UpgradeDev.Type).Ref("upgrade_dev_groups").Through("upgrade_dev_group_relations", UpgradeDevGroupRelation.Type),
	}
}

func (UpgradeDevGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_dev_group"},
	}
}
