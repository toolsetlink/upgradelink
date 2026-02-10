.PHONY: gen-api
gen-api: # 生成代码
	goctl api go --api docs/tpl/all.api --style go_zero -dir server/api

.PHONY: gen-ent
gen-ent: # 生成数据库ent文件
	cd internal && go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./ent/template/*.tmpl" ./ent/schema --feature sql/execquery,intercept,sql/modifier

.PHONY: gen-model
gen-model: # 生成数据库文件
	goctl model mysql ddl --src docs/sql/tables.sql --dir server/api/internal/resource/model --style go_zero

.PHONY: help
help: # 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#' Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
