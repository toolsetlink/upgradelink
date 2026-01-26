# 生成代码
goctl api go --api docs/tpl/all.api --style go_zero -dir server/api

# 生成数据库ent文件  在internal 目录下执行
go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./ent/template/*.tmpl" ./ent/schema --feature sql/execquery,intercept,sql/modifier

# 生成数据库文件
goctl model mysql ddl --src docs/sql/tables.sql --dir server/api/internal/resource/model --style go_zero