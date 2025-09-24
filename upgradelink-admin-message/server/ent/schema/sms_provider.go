package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// SmsProvider holds the schema definition for the SmsProvider entity.
type SmsProvider struct {
	ent.Schema
}

// Fields of the SmsProvider.
func (SmsProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("The SMS provider name | 短信服务的提供商").
			Unique(),
		field.String("secret_id").
			Comment("The secret ID | 密钥 ID"),
		field.String("secret_key").
			Comment("The secret key | 密钥 Key"),
		field.String("region").
			Comment("The service region | 服务器所在地区"),
		field.Bool("is_default").Default(false).
			Comment("Is it the default provider | 是否为默认提供商"),
	}
}

func (SmsProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the SmsProvider.
func (SmsProvider) Edges() []ent.Edge {
	return nil
}

func (SmsProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "mcms_sms_providers"},
	}
}
