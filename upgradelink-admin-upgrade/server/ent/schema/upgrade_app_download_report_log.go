package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeAppDownloadReportLog struct {
	ent.Schema
}

func (UpgradeAppDownloadReportLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Time("timestamp").Optional().Comment("事件发生时间"),
		field.String("app_key").Comment("应用Key"),
		field.Int("app_version_id").Comment("应用版本ID"),
		field.Int("app_version_code").Comment("版本号"),
		field.String("app_version_target").Optional().Comment("操作系统:linux、darwin、windows"),
		field.String("app_version_arch").Optional().Comment("机器架构:x86_64、i686、aarch64、armv7"),
		field.Time("create_at").Optional().Comment("创建时间")}
}
func (UpgradeAppDownloadReportLog) Edges() []ent.Edge {
	return nil
}

func (UpgradeAppDownloadReportLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_app_download_report_log"},
	}
}
