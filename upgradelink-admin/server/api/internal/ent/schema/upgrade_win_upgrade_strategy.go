package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UpgradeWinUpgradeStrategy struct {
	ent.Schema
}

func (UpgradeWinUpgradeStrategy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("company_id").Comment("公司ID"),
		field.Int32("enable").Comment("是否生效；可通过此控制策略是否生效0：失效；1：生效"),
		field.String("name").Comment("任务名称"),
		field.String("description").Comment("任务描述信息"),
		field.Int("win_id").Comment("win应用ID"),
		field.Int("win_version_id").Comment("win_version_id; 外键win_version.id"),
		field.Time("begin_datetime").Comment("升级任务开始时间"),
		field.Time("end_datetime").Comment("升级任务结束时间"),
		field.Int32("upgrade_type").Comment("升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级"),
		field.String("prompt_upgrade_content").Comment("提示升级描述内容"),
		field.Int32("upgrade_dev_type").Comment("指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型"),
		field.String("upgrade_dev_data").Comment("升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;"),
		field.Int32("upgrade_version_type").Comment("指定升级的应用版本：0：全部版本；1：指定版本"),
		field.String("upgrade_version_data").Comment("升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;"),
		field.Int32("is_gray").Comment("是否开启灰度 0：不开启；1：开启"),
		field.String("gray_data").Comment("灰度策略id数据"),
		field.Int32("is_flow_limit").Comment("是否开启频控 0：不开启；1：开启"),
		field.String("flow_limit_data").Comment("频控策略id数据"),
		field.Int32("is_del").Comment("是否删除 0：正常；1：已删除"),
		field.Time("create_at").Optional().Comment("创建时间"),
		field.Time("update_at").Optional().Comment("修改时间")}
}
func (UpgradeWinUpgradeStrategy) Edges() []ent.Edge {
	return nil
}

func (UpgradeWinUpgradeStrategy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "upgrade_win_upgrade_strategy"},
	}
}
