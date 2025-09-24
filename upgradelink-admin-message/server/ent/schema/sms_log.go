package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// SmsLog holds the schema definition for the SmsLog entity.
type SmsLog struct {
	ent.Schema
}

// Fields of the SmsLog.
func (SmsLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone_number").Comment("The target phone number | 目标电话"),
		field.String("content").Comment("The content | 发送的内容"),
		field.Uint8("send_status").Comment("The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败"),
		field.String("provider").Comment("The sms service provider | 短信服务提供商"),
	}
}

func (SmsLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

// Edges of the SmsLog.
func (SmsLog) Edges() []ent.Edge {
	return nil
}

func (SmsLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "mcms_sms_logs"},
	}
}
