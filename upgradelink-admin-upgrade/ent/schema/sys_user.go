package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SysUser struct {
	ent.Schema
}

func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("UUID"),
		field.Time("created_at").Comment("Create Time | 创建日期"),
		field.Time("updated_at").Comment("Update Time | 修改日期"),
		field.Uint8("status").Optional().Comment("Status 1: normal 2: ban | 状态 1 正常 2 禁用"),
		field.Time("deleted_at").Optional().Comment("Delete Time | 删除日期"),
		field.String("username").Unique().Comment("User's login name | 登录名"),
		field.String("password").Comment("Password | 密码"),
		field.String("nickname").Unique().Comment("Nickname | 昵称"),
		field.String("description").Optional().Comment("The description of user | 用户的描述信息"),
		field.String("home_path").Comment("The home page that the user enters after logging in | 用户登陆后进入的首页"),
		field.String("mobile").Optional().Comment("Mobile number | 手机号"),
		field.String("email").Optional().Comment("Email | 邮箱号"),
		field.String("avatar").Optional().Comment("Avatar | 头像路径"),
		field.Uint64("department_id").Optional().Comment("Department ID | 部门ID"),
		field.Uint64("company_id").Optional().Comment("Company ID | 公司ID")}
}
func (SysUser) Edges() []ent.Edge {
	return nil
}

func (SysUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("nickname").Unique(),
		index.Fields("username").Unique(),
		index.Fields("email").Unique()}
}

func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_users"},
	}
}
