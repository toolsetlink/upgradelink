goctls api new example -i -c -a -m github.com/suyuan32/simple-admin-example-api -p 8081 -e

go install github.com/suyuan32/goctls@latest

项目初始化生成代码
goctls api new upgrade --i18n=false --casbin=true  --module_name=simple-admin-upgrade-api  --port=8181 --ent=true --use_core_rpc=true

根据表 生成表 ent
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_url" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "sys_users" -p=false

goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_url_upgrade_strategy" -p=false

goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_file" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_file_upgrade_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_file_version" -p=false

goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_dev" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_dev_group" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_dev_group_relation" -p=false

goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_url_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_url_upgrade_strategy_flow_limit_strategy" -p=false

goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_file_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_file_upgrade_strategy_flow_limit_strategy" -p=false


goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_tauri" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_tauri_upgrade_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_tauri_version" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_tauri_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-test:8yBBSLypEDFBb24S@tcp(101.43.181.38:3306)/upgrade-test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_tauri_upgrade_strategy_flow_limit_strategy" -p=false

goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_app_download_report_log" -p=false



goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_configuration" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_configuration_version" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_configuration_upgrade_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_configuration_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_configuration_upgrade_strategy_flow_limit_strategy" -p=false


goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_apk" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_apk_version" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_apk_upgrade_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_apk_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_apk_upgrade_strategy_flow_limit_strategy" -p=false


goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_electron" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_electron_version" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_electron_upgrade_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_electron_upgrade_strategy_gray_strategy" -p=false
goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_electron_upgrade_strategy_flow_limit_strategy" -p=false


goctls extra ent import -d "mysql://upgrade-prod:bttiWWYjWZeYJcXm@tcp(101.43.181.38:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -t "upgrade_company_income" -p=false


goctls extra ent import -d 'mysql://songang:Sj13051570639!@tcp(rm-j6cx741ix46hwj4xa1o.mysql.cnhk.rds.aliyuncs.com:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai' -t "upgrade_traffic_packet" -p=false
goctls extra ent import -d 'mysql://songang:Sj13051570639!@tcp(rm-j6cx741ix46hwj4xa1o.mysql.cnhk.rds.aliyuncs.com:3306)/upgrade-prod?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai' -t "upgrade_company_traffic_packet" -p=false



生成 ent 代码
sudo go generate ./ent

生成
sudo make gen-ent

基于 ent 生成 api 代码
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeUrl --group=upgrade_url --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeDevModel --group=upgrade_dev_model --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeUrlUpgradeStrategy --group=upgrade_url_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeUrlVersion --group=upgrade_url_version --search_key_num=300 --overwrite=true

goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeFile --group=upgrade_file --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeFileUpgradeStrategy --group=upgrade_file_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeFileVersion --group=upgrade_file_version --search_key_num=300 --overwrite=true

goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeDev --group=upgrade_dev --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeDevGroup --group=upgrade_dev_group --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeDevGroupRelation --group=upgrade_dev_group_relation --search_key_num=300 --overwrite=true

goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeTauri --group=upgrade_tauri --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeTauriUpgradeStrategy --group=upgrade_tauri_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeTauriVersion --group=upgrade_tauri_version --search_key_num=300 --overwrite=true

goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=upgradeAppDownloadReportLog --group=upgrade_app_download_report_log --search_key_num=300 --overwrite=true


goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeConfiguration --group=upgrade_configuration --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeConfigurationUpgradeStrategy --group=upgrade_configuration_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeConfigurationVersion --group=upgrade_configuration_version --search_key_num=300 --overwrite=true


goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeApk --group=upgrade_apk --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeApkUpgradeStrategy --group=upgrade_apk_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeApkVersion --group=upgrade_apk_version --search_key_num=300 --overwrite=true


goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeElectron --group=upgrade_electron --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeElectronUpgradeStrategy --group=upgrade_electron_upgrade_strategy --search_key_num=300 --overwrite=true
goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeElectronVersion --group=upgrade_electron_version --search_key_num=300 --overwrite=true


goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeCompanyIncome --group=upgrade_company_income --search_key_num=300 --overwrite=true


goctls api ent --schema=./ent/schema --api_service_name=upgrade --output=./ --model=UpgradeCompanyTrafficPacket --group=upgrade_company_traffic_packet --search_key_num=300 --overwrite=true



# 需要修改 api 文件    id 的字段类型 莫名丢失

生成 api 代码   生成的代码没有逻辑 使用下面的   也需要执行   更新请求参数等
sudo make gen-api

// dashboard 数据数据
sudo make gen-api

