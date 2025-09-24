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

// FileTag holds the schema definition for the FileTag entity.
type FileTag struct {
	ent.Schema
}

// Fields of the FileTag.
func (FileTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("FileTag's name | 标签名称"),
		field.String("remark").Comment("The remark of tag | 标签的备注").
			Optional(),
	}
}

func (FileTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the FileTag.
func (FileTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type),
	}
}

func (FileTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

func (FileTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "fms_file_tags"}, // fms means FileTag management service
	}
}
