package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// EmailProvider holds the schema definition for the EmailProvider entity.
type EmailProvider struct {
	ent.Schema
}

// Fields of the EmailProvider.
func (EmailProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("The email provider name | 电子邮件服务的提供商").
			Unique(),
		field.Uint8("auth_type").
			Comment("The auth type, supported plain, CRAMMD5 | 鉴权类型, 支持 plain, CRAMMD5"),
		field.String("email_addr").
			Comment("The email address | 邮箱地址"),
		field.String("password").
			Optional().
			Comment("The email's password | 电子邮件的密码"),
		field.String("host_name").
			Comment("The host name is the email service's host address | 电子邮箱服务的服务器地址"),
		field.String("identify").
			Optional().
			Comment("The identify info, for CRAMMD5 | 身份信息, 支持 CRAMMD5"),
		field.String("secret").
			Optional().
			Comment("The secret, for CRAMMD5 | 邮箱密钥, 用于 CRAMMD5"),
		field.Uint32("port").
			Optional().
			Comment("The port of the host | 服务器端口"),
		field.Bool("tls").Default(false).
			Comment("Whether to use TLS | 是否采用 tls 加密"),
		field.Bool("is_default").Default(false).
			Comment("Is it the default provider | 是否为默认提供商"),
	}
}

func (EmailProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the EmailProvider.
func (EmailProvider) Edges() []ent.Edge {
	return nil
}

func (EmailProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "mcms_email_providers"},
	}
}
