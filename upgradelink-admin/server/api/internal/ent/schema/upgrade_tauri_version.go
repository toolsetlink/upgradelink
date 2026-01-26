package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeTauriVersion struct {
	ent.Schema
}

func (UpgradeTauriVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int("tauri_id").Comment("tauri应用ID"),
		field.String("cloud_file_id").Comment("云文件id"),
		field.String("install_cloud_file_id").Comment("云文件id 安装文件"),
		field.String("version_name").Comment("版本名"),
		field.Int("version_code").Comment("版本号"),
		field.String("target").Comment("操作系统:linux、darwin、windows"),
		field.String("arch").Comment("机器架构:x86_64、i686、aarch64、armv7"),
		field.String("signature").Comment("生成的 .sig 文件的内容"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}

}
func (UpgradeTauriVersion) Edges() []ent.Edge {
	return nil
}

func (UpgradeTauriVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_tauri_version"},
	}
}
