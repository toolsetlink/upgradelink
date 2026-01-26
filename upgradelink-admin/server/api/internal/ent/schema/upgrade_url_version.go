package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeUrlVersion struct {
	ent.Schema
}

func (UpgradeUrlVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int("url_id").Comment("url应用ID"),
		field.String("url_path").Comment("url链接"),
		field.String("version_name").Comment("版本名"),
		field.Int("version_code").Comment("版本号"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeUrlVersion) Edges() []ent.Edge {
	return nil
}

func (UpgradeUrlVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_url_version"},
	}
}
