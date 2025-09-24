cor rpc 文件生成

更新表字段

make gen-ent   修改gen 文件

修改 schema 文件
make gen-rpc-ent-logic  model=User group user
make gen-rpc-ent-logic  model=Company group=company
make gen-rpc-ent-logic  model=CompanySecret  group=company_secret

生成文件
make gen-rpc




api 项目

// 生成 api 接入 rpc 代码  提供的脚本 生成不出来详细字段，  直接自己新 api文件
goctls extra proto2api -m company_secret  -j go_zero -p /Users/songang/LinkProjects/upgrade/upgradelink-admin-core/rpc/core.proto -a /Users/songang/LinkProjects/upgrade/upgradelink-admin-core/api/desc/core/api.api

// 生成 logic 相关代码 不包含具体逻辑的代码
make gen-api
