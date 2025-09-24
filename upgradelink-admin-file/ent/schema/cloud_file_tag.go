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

// CloudFileTag holds the schema definition for the CloudFileTag entity.
type CloudFileTag struct {
	ent.Schema
}

// Fields of the CloudFileTag.
func (CloudFileTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("CloudFileTag's name | 标签名称"),
		field.String("remark").Comment("The remark of tag | 标签的备注").
			Optional(),
	}
}

func (CloudFileTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the CloudFileTag.
func (CloudFileTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cloud_files", CloudFile.Type),
	}
}

func (CloudFileTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

func (CloudFileTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "fms_cloud_file_tags"}, // fms means CloudFileTag management service
	}
}
