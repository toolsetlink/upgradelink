CREATE TABLE `sys_companys` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公司名称',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公司信息表';

CREATE TABLE `sys_company_secret` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`access_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密钥id',
`secret_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密钥key',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公司密钥表';


CREATE TABLE `fms_cloud_files` (
`id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
`state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s name | 文件名',
`url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s url | 文件地址',
`size` bigint unsigned NOT NULL COMMENT 'The file''s size | 文件大小',
`file_type` tinyint unsigned NOT NULL COMMENT 'The file''s type | 文件类型',
`user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The user who upload the file | 上传用户的 ID',
`cloud_file_storage_providers` bigint unsigned DEFAULT NULL,
PRIMARY KEY (`id`) USING BTREE,
KEY `cloudfile_file_type` (`file_type`) USING BTREE,
KEY `cloudfile_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


CREATE TABLE `upgrade_devs` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备管理';


CREATE TABLE `upgrade_dev_model` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL COMMENT '公司ID: 所属公司id',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型名称',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备机型表';

CREATE TABLE `upgrade_dev_group` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备分组名称',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备分组';

CREATE TABLE `upgrade_dev_group_relation` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`dev_id` bigint unsigned NOT NULL COMMENT '设备id',
`dev_group_id` bigint unsigned NOT NULL COMMENT '设备分组 id',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备分组关系';

CREATE TABLE `upgrade_url` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'url唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'url应用名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用';


CREATE TABLE `upgrade_url_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`url_id` bigint NOT NULL DEFAULT '0' COMMENT 'url应用ID',
`url_path` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'url链接',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 版本库';


CREATE TABLE `upgrade_url_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`url_id` bigint NOT NULL DEFAULT '0' COMMENT 'url应用ID',
`url_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'url_version_id; 外键url_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务';

CREATE TABLE `upgrade_url_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务灰度策略表；';


CREATE TABLE `upgrade_url_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务灰度策略表；';



CREATE TABLE `upgrade_file` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件应用唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件应用名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用';

CREATE TABLE `upgrade_file_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`file_id` bigint NOT NULL DEFAULT '0' COMMENT '文件应用ID',
`cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 版本库';


CREATE TABLE `upgrade_file_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`file_id` bigint NOT NULL DEFAULT '0' COMMENT '文件应用ID',
`file_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'file_version_id; 外键file_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务';


CREATE TABLE `upgrade_file_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务灰度策略表；';

CREATE TABLE `upgrade_file_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '9' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务灰度策略表；';



CREATE TABLE `upgrade_tauri` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'tauri应用唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'tauri应用名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`github_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开源项目 github 地址',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用';


CREATE TABLE `upgrade_tauri_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`tauri_id` bigint NOT NULL DEFAULT '0' COMMENT 'tauri应用ID',
`cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
`install_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`signature` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '生成的 .sig 文件的内容',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_tauri_active_version` (`tauri_id`,`is_del`,`version_code`,`create_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 版本库';


CREATE TABLE `upgrade_tauri_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`tauri_id` bigint NOT NULL DEFAULT '0' COMMENT 'Tauri应用ID',
`tauri_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'tauri_version_id; 外键tauri_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务';


CREATE TABLE `upgrade_tauri_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务灰度策略表；';


CREATE TABLE `upgrade_tauri_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务灰度策略表；';


CREATE TABLE `fms_cloud_files` (
`id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
`state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s name | 文件名',
`url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s url | 文件地址',
`size` bigint unsigned NOT NULL COMMENT 'The file''s size | 文件大小',
`file_type` tinyint unsigned NOT NULL COMMENT 'The file''s type | 文件类型',
`user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The user who upload the file | 上传用户的 ID',
`cloud_file_storage_providers` bigint unsigned DEFAULT NULL,
PRIMARY KEY (`id`) USING BTREE,
KEY `cloudfile_file_type` (`file_type`) USING BTREE,
KEY `cloudfile_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


CREATE TABLE `upgrade_app_download_report_log` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
`app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
`app_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用类型 url,file,tauri,apk,configuration',
`app_version_id` bigint NOT NULL DEFAULT '0' COMMENT '应用版本ID',
`app_version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`app_version_platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '操作系统:linux、darwin、windows',
`app_version_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`app_version_arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`download_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '云文件id',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用初次下载事件日志';



CREATE TABLE `upgrade_configuration` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '配置唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '配置名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置';


CREATE TABLE `upgrade_configuration_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`configuration_id` bigint NOT NULL DEFAULT '0' COMMENT '配置ID',
`configuration_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'configuration_version_id; 外键configuration_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务';


CREATE TABLE `upgrade_configuration_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务灰度策略表；';


CREATE TABLE `upgrade_configuration_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务灰度策略表；';



CREATE TABLE `upgrade_configuration_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`configuration_id` bigint NOT NULL DEFAULT '0' COMMENT '配置ID',
`content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '内容',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 版本库';







CREATE TABLE `upgrade_apk` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '安卓应用唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '安卓应用名称',
`package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '安卓应用包名',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用';


CREATE TABLE `upgrade_apk_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`apk_id` bigint NOT NULL DEFAULT '0' COMMENT '安卓应用ID',
`cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 版本库';

CREATE TABLE `upgrade_apk_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`apk_id` bigint NOT NULL DEFAULT '0' COMMENT '安卓应用ID',
`apk_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'apk_version_id; 外键apk_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务';



CREATE TABLE `upgrade_apk_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '9' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务灰度策略表；';


CREATE TABLE `upgrade_apk_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务灰度策略表；';


CREATE TABLE `upgrade_app_start_report_log` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
`app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
`app_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用类型 url,file,tauri,apk,configuration',
`app_version_id` bigint NOT NULL DEFAULT '0' COMMENT '应用版本ID',
`app_version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`app_version_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`app_version_arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`dev_model_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
`dev_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
`launch_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '应用启动事件-应用启动时间',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用启动事件日志';



CREATE TABLE `upgrade_app_download_report_log` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
`app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
`app_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用类型 url,file,tauri,apk,configuration',
`app_version_id` bigint NOT NULL DEFAULT '0' COMMENT '应用版本ID',
`app_version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`app_version_platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '操作系统:linux、darwin、windows',
`app_version_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`app_version_arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`download_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '云文件id',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用初次下载事件日志';




CREATE TABLE `upgrade_app_upgrade_upgrade_report_log` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
`app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
`app_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用类型 url,file,tauri,apk,configuration',
`app_version_id` bigint NOT NULL DEFAULT '0' COMMENT '应用版本ID',
`app_version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`app_version_platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '操作系统:linux、darwin、windows',
`app_version_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`app_version_arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`dev_model_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
`dev_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
`upgrade_version_code` bigint NOT NULL DEFAULT '0' COMMENT '升级应用版本号',
`code` bigint NOT NULL DEFAULT '0' COMMENT '事件-状态码',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用升级升级事件日志';





CREATE TABLE `upgrade_app_upgrade_get_strategy_report_log` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
`app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
`app_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用类型 url,file,tauri,apk,configuration',
`app_version_id` bigint NOT NULL DEFAULT '0' COMMENT '应用版本ID',
`app_version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`app_version_platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '操作系统:linux、darwin、windows',
`app_version_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作系统:linux、darwin、windows',
`app_version_arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x86_64、i686、aarch64、armv7',
`dev_model_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
`dev_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
`strategy_version_id` bigint NOT NULL DEFAULT '0' COMMENT '升级策略应用版本id',
`strategy_version_code` bigint NOT NULL DEFAULT '0' COMMENT '升级策略应用版本号',
`strategy_id` bigint NOT NULL DEFAULT '0' COMMENT '升级策略应用版本号',
`code` bigint NOT NULL DEFAULT '0' COMMENT '事件-状态码',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用升级获取升级策略事件日志';



CREATE TABLE `upgrade_file_github` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件id',
`url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'github 文件地址',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_upgrade_file_github_file_id` (`file_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='github文件下载地址';



CREATE TABLE `upgrade_electron` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'electron应用唯一标识',
`name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'electron应用名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`github_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开源项目 github 地址',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用';



CREATE TABLE `upgrade_electron_upgrade_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
`electron_id` bigint NOT NULL DEFAULT '0' COMMENT 'electron应用ID',
`electron_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'electron_version_id; 外键electron_version.id',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
`upgrade_type` int NOT NULL DEFAULT '0' COMMENT '升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级',
`prompt_upgrade_content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提示升级描述内容',
`upgrade_dev_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型',
`upgrade_dev_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;',
`upgrade_version_type` int NOT NULL DEFAULT '0' COMMENT '指定升级的应用版本：0：全部版本；1：指定版本',
`upgrade_version_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;',
`is_gray` int NOT NULL DEFAULT '0' COMMENT '是否开启灰度 0：不开启；1：开启',
`gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '灰度策略id数据',
`is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
`flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '频控策略id数据',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务';




CREATE TABLE `upgrade_electron_upgrade_strategy_flow_limit_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
`end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
`dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务灰度策略表；';



CREATE TABLE `upgrade_electron_upgrade_strategy_gray_strategy` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
`begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
`end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
`limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务灰度策略表；';


CREATE TABLE `upgrade_electron_version` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
`company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
`electron_id` bigint NOT NULL DEFAULT '0' COMMENT 'tauri应用ID',
`cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
`sha512` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '生成的sha512',
`install_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '云文件id',
`install_sha512` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '安装包 生成的sha512',
`version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
`version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
`platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作平台:linux、darwin、windows',
`arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x64、arm64',
`description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
`is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
`create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE,
KEY `idx_electron_active_version` (`electron_id`,`is_del`,`version_code`,`create_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 版本库';
