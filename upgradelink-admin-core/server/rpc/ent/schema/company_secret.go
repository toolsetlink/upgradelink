package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type CompanySecret struct {
	ent.Schema
}

func (CompanySecret) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("company_id").Optional().Default(1).
			Comment("Company ID | 公司ID"),
		field.String("access_key").Unique().
			Comment("access_key | 密钥id"),
		field.String("secret_key").Unique().
			Comment("secret_key | 密钥key"),
		field.Time("validity_datetime").
			Optional().
			Comment("Validity Time | 有效期"),
		field.String("rule_data").Optional().Comment("规则数据"),
		field.Uint32("enable").Default(0).
			Comment("是否生效；可通过此控制策略是否生效0：失效；1：生效"),
		field.String("description").Optional().Comment("描述信息"),
		field.Uint32("is_del").Default(0).
			Comment("是否删除 0：正常；1：已删除"),
	}
}

func (CompanySecret) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (CompanySecret) Edges() []ent.Edge {
	return nil
}

func (CompanySecret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("company_id"),
	}
}

func (CompanySecret) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_company_secret"},
	}
}
