package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// StorageProvider holds the schema definition for the StorageProvider entity.
type StorageProvider struct {
	ent.Schema
}

// Fields of the StorageProvider.
func (StorageProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Comment("The cloud storage service name | 服务名称"),
		field.String("bucket").
			Comment("The cloud storage bucket name | 云存储服务的存储桶"),
		field.String("secret_id").
			Comment("The secret ID | 密钥 ID"),
		field.String("secret_key").
			Comment("The secret key | 密钥 Key"),
		field.String("endpoint").
			Comment("The service URL | 服务器地址"),
		field.String("folder").
			Optional().
			Comment("The folder in cloud | 云服务目标文件夹"),
		field.String("region").
			Comment("The service region | 服务器所在地区"),
		field.Bool("is_default").Default(false).
			Comment("Is it the default provider | 是否为默认提供商"),
		field.Bool("use_cdn").Default(false).
			Comment("Does it use CDN | 是否使用 CDN"),
		field.String("cdn_url").Optional().
			Comment("CDN URL | CDN 地址"),
	}
}

func (StorageProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StateMixin{},
	}
}

// Edges of the StorageProvider.
func (StorageProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cloudfiles", CloudFile.Type).Ref("storage_providers"),
	}
}

func (StorageProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "fms_storage_providers"},
	}
}
