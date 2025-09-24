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
