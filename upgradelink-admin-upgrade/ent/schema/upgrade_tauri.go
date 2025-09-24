package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeTauri struct {
	ent.Schema
}

func (UpgradeTauri) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.String("key").Comment("tauri应用唯一标识"),
		field.String("name").Comment("tauri应用名称"),
		field.String("description").Optional().Comment("描述信息"),
		field.String("github_url").Optional().Comment("github项目地址"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeTauri) Edges() []ent.Edge {
	return nil
}

func (UpgradeTauri) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_tauri"},
	}
}
