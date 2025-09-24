package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeCompanyIncome struct {
	ent.Schema
}

func (UpgradeCompanyIncome) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int8("income_type").Comment("收入类型：0 - 广告收入；"),
		field.Int("income_amount").Comment("收入金额（单位分）"),
		field.Time("income_time").Optional().Comment("收入产生时间"),
		field.String("remark").Optional().Comment("备注"),
		field.Int32("status").Comment("收入状态：0 - 待结算；1 - 已结算；2 - 失效（如订单取消导致收入作废）"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}

}
func (UpgradeCompanyIncome) Edges() []ent.Edge {
	return nil
}

func (UpgradeCompanyIncome) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_company_income"},
	}
}
