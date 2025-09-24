
# 生成代码
goctl api go --api docs/tpl/all.api --style go_zero -dir server/api

# 生成数据库文件
goctl model mysql ddl --src docs/sql/tables.sql --dir server/api/internal/resource/model --style go_zero