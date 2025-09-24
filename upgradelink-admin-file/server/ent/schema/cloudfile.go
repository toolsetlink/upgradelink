package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// CloudFile holds the schema definition for the CloudFile entity.
type CloudFile struct {
	ent.Schema
}

// Fields of the CloudFile.
func (CloudFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("The file's name | 文件名"),
		field.String("url").
			Comment("The file's url | 文件地址"),
		field.Uint64("size").
			Comment("The file's size | 文件大小"),
		field.Uint8("file_type").
			Comment("The file's type | 文件类型"),
		field.String("user_id").
			Comment("The user who upload the file | 上传用户的 ID"),
	}
}

// Edges of the CloudFile.
func (CloudFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("storage_providers", StorageProvider.Type).Unique(),
		edge.From("tags", CloudFileTag.Type).Ref("cloud_files"),
	}
}

// Mixin of the CloudFile.
func (CloudFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StateMixin{},
	}
}

// Indexes of the CloudFile.
func (CloudFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("file_type"),
	}
}

// Annotations of the CloudFile
func (CloudFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "fms_cloud_files"},
	}
}
