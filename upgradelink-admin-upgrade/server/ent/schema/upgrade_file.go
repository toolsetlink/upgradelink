package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeFile struct {
	ent.Schema
}

func (UpgradeFile) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.String("key").Comment("文件应用唯一标识"),
		field.String("name").Comment("文件应用名称"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeFile) Edges() []ent.Edge {
	return nil
}

func (UpgradeFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_file"},
	}
}
