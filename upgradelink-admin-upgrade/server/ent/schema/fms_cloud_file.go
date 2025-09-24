package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FmsCloudFile struct {
	ent.Schema
}

func (FmsCloudFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("UUID"),
		field.Time("created_at").Optional().Comment("Create Time | 创建日期"),
		field.Time("updated_at").Optional().Comment("Update Time | 修改日期"),
		field.Bool("state").Optional().Comment("State true: normal false: ban | 状态 true 正常 false 禁用"),
		field.String("name").Comment("The file's name | 文件名"),
		field.String("url").Comment("The file's url | 文件地址"),
		field.Uint64("size").Comment("The file's size | 文件大小"),
		field.Uint8("file_type").Comment("The file's type | 文件类型"),
		field.String("user_id").Comment("The user who upload the file | 上传用户的 ID"),
		field.Uint64("cloud_file_storage_providers").Optional()}
}
func (FmsCloudFile) Edges() []ent.Edge {
	return nil
}

func (FmsCloudFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("file_type"),
		index.Fields("name")}
}

func (FmsCloudFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "fms_cloud_files"},
	}
}
