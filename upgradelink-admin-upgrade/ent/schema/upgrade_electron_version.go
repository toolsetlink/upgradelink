package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeElectronVersion struct {
	ent.Schema
}

func (UpgradeElectronVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int("electron_id").Comment("tauri应用ID"),
		field.String("cloud_file_id").Comment("云文件id"),
		field.String("sha512").Comment("生成的sha512"),
		field.String("install_cloud_file_id").Comment("云文件id 安装文件"),
		field.String("install_sha512").Comment("安装包生成的sha512"),
		field.String("version_name").Comment("版本名"),
		field.Int("version_code").Comment("版本号"),
		field.String("platform").Comment("操作平台:linux、darwin、windows"),
		field.String("arch").Comment("机器架构:x64、arm64"),
		field.String("description").Optional().Comment("描述信息"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeElectronVersion) Edges() []ent.Edge {
	return nil
}

func (UpgradeElectronVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_electron_version"},
	}
}
