package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeLnx struct {
	ent.Schema
}

func (UpgradeLnx) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.String("key").Comment("linux应用唯一标识"),
		field.String("name").Comment("linux应用名称"),
		field.String("package_name").Comment("linux应用包名"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeLnx) Edges() []ent.Edge {
	return nil
}

func (UpgradeLnx) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_lnx"},
	}
}
