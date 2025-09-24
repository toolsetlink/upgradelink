package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeDevModel struct {
	ent.Schema
}

func (UpgradeDevModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID: 所属公司id"),
		field.String("key").Comment("设备机型唯一标识"),
		field.String("name").Comment("设备机型名称"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeDevModel) Edges() []ent.Edge {
	return nil
}

func (UpgradeDevModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_dev_model"},
	}
}
