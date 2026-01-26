package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeMacVersion struct {
	ent.Schema
}

func (UpgradeMacVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int("mac_id").Comment("mac应用ID"),
		field.String("cloud_file_id").Comment("云文件id"),
		field.String("version_name").Comment("版本名"),
		field.Int("version_code").Comment("版本号"),
		field.String("arch").Comment("机器架构:x64、arm64"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeMacVersion) Edges() []ent.Edge {
	return nil
}

func (UpgradeMacVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_mac_version"},
	}
}
