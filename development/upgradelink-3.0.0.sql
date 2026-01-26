/*
 Navicat Premium Dump SQL

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80403 (8.4.3)
 Source Host           : localhost:3306
 Source Schema         : upgrade

 Target Server Type    : MySQL
 Target Server Version : 80403 (8.4.3)
 File Encoding         : 65001

 Date: 25/01/2026 21:22:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rules
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rules`;
CREATE TABLE `casbin_rules` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ptype` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v4` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=422495 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of casbin_rules
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420883, 'p', '003', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420884, 'p', '003', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420885, 'p', '003', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420886, 'p', '003', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420887, 'p', '003', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420888, 'p', '003', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420889, 'p', '003', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420890, 'p', '003', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420891, 'p', '003', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420892, 'p', '003', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420893, 'p', '003', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420894, 'p', '003', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420895, 'p', '003', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420896, 'p', '001', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420897, 'p', '001', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420898, 'p', '001', '/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420899, 'p', '001', '/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420900, 'p', '001', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420901, 'p', '001', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420902, 'p', '001', '/user/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420903, 'p', '001', '/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420904, 'p', '001', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420905, 'p', '001', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420906, 'p', '001', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420907, 'p', '001', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420908, 'p', '001', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420909, 'p', '001', '/user/refresh_token', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420910, 'p', '001', '/user/access_token', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420911, 'p', '001', '/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420912, 'p', '001', '/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420913, 'p', '001', '/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420914, 'p', '001', '/role/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420915, 'p', '001', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420916, 'p', '001', '/menu/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420917, 'p', '001', '/menu/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420918, 'p', '001', '/menu/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420919, 'p', '001', '/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420920, 'p', '001', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420921, 'p', '001', '/menu_param/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420922, 'p', '001', '/menu_param/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420923, 'p', '001', '/menu_param/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420924, 'p', '001', '/menu_param/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420925, 'p', '001', '/menu_param', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420926, 'p', '001', '/menu', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420927, 'p', '001', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420928, 'p', '001', '/authority/api/create_or_update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420929, 'p', '001', '/authority/api/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420930, 'p', '001', '/authority/menu/create_or_update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420931, 'p', '001', '/authority/menu/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420932, 'p', '001', '/api/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420933, 'p', '001', '/api/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420934, 'p', '001', '/api/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420935, 'p', '001', '/api/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420936, 'p', '001', '/api', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420937, 'p', '001', '/dictionary', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420938, 'p', '001', '/dictionary/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420939, 'p', '001', '/dictionary/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420940, 'p', '001', '/dictionary/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420941, 'p', '001', '/dictionary_detail/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420942, 'p', '001', '/dictionary_detail', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420943, 'p', '001', '/dictionary_detail/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420944, 'p', '001', '/dictionary_detail/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420945, 'p', '001', '/dictionary_detail/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420946, 'p', '001', '/dictionary/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420947, 'p', '001', '/dict/:name', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420948, 'p', '001', '/oauth_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420949, 'p', '001', '/oauth_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420950, 'p', '001', '/oauth_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420951, 'p', '001', '/oauth_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420952, 'p', '001', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420953, 'p', '001', '/oauth_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420954, 'p', '001', '/token/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420955, 'p', '001', '/token/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420956, 'p', '001', '/token/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420957, 'p', '001', '/token/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420958, 'p', '001', '/token/logout', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420959, 'p', '001', '/token', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420960, 'p', '001', '/department/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420961, 'p', '001', '/department/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420962, 'p', '001', '/department/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420963, 'p', '001', '/department/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420964, 'p', '001', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420965, 'p', '001', '/position/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420966, 'p', '001', '/position/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420967, 'p', '001', '/position/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420968, 'p', '001', '/position/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420969, 'p', '001', '/position', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420970, 'p', '001', '/task/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420971, 'p', '001', '/task/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420972, 'p', '001', '/task/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420973, 'p', '001', '/task/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420974, 'p', '001', '/task', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420975, 'p', '001', '/task_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420976, 'p', '001', '/task_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420977, 'p', '001', '/task_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420978, 'p', '001', '/task_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420979, 'p', '001', '/task_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420980, 'p', '001', '/configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420981, 'p', '001', '/configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420982, 'p', '001', '/configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420983, 'p', '001', '/configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420984, 'p', '001', '/configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420985, 'p', '001', '/email_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420986, 'p', '001', '/email_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420987, 'p', '001', '/email_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420988, 'p', '001', '/email_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420989, 'p', '001', '/email_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420990, 'p', '001', '/email_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420991, 'p', '001', '/email_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420992, 'p', '001', '/email_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420993, 'p', '001', '/email_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420994, 'p', '001', '/email_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420995, 'p', '001', '/sms_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420996, 'p', '001', '/sms_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420997, 'p', '001', '/sms_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420998, 'p', '001', '/sms_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (420999, 'p', '001', '/sms_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421000, 'p', '001', '/sms_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421001, 'p', '001', '/sms_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421002, 'p', '001', '/sms_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421003, 'p', '001', '/sms_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421004, 'p', '001', '/sms_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421005, 'p', '001', '/sms/send', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421006, 'p', '001', '/email/send', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421007, 'p', '001', '/member/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421008, 'p', '001', '/member/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421009, 'p', '001', '/member/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421010, 'p', '001', '/member/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421011, 'p', '001', '/member', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421012, 'p', '001', '/member_rank/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421013, 'p', '001', '/member_rank/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421014, 'p', '001', '/member_rank/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421015, 'p', '001', '/member_rank/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421016, 'p', '001', '/member_rank', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421017, 'p', '001', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421018, 'p', '001', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421019, 'p', '001', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421020, 'p', '001', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421021, 'p', '001', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421022, 'p', '001', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421023, 'p', '001', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421024, 'p', '001', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421025, 'p', '001', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421026, 'p', '001', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421027, 'p', '001', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421028, 'p', '001', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421029, 'p', '001', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421030, 'p', '001', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421031, 'p', '001', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421032, 'p', '001', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421033, 'p', '001', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421034, 'p', '001', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421035, 'p', '001', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421036, 'p', '001', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421037, 'p', '001', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421038, 'p', '001', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421039, 'p', '001', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421040, 'p', '001', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421041, 'p', '001', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421042, 'p', '001', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421043, 'p', '001', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421044, 'p', '001', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421045, 'p', '001', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421046, 'p', '001', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421047, 'p', '001', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421048, 'p', '001', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421049, 'p', '001', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421050, 'p', '001', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421051, 'p', '001', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421052, 'p', '001', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421053, 'p', '001', '/upload', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421054, 'p', '001', '/file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421055, 'p', '001', '/file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421056, 'p', '001', '/file/status', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421057, 'p', '001', '/file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421058, 'p', '001', '/file/download/:id', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421059, 'p', '001', '/file_tag/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421060, 'p', '001', '/file_tag/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421061, 'p', '001', '/file_tag/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421062, 'p', '001', '/file_tag/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421063, 'p', '001', '/file_tag', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421064, 'p', '001', '/storage_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421065, 'p', '001', '/storage_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421066, 'p', '001', '/storage_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421067, 'p', '001', '/storage_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421068, 'p', '001', '/storage_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421069, 'p', '001', '/cloud_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421070, 'p', '001', '/cloud_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421071, 'p', '001', '/cloud_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421072, 'p', '001', '/cloud_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421073, 'p', '001', '/cloud_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421074, 'p', '001', '/cloud_file/upload', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421075, 'p', '001', '/cloud_file_tag/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421076, 'p', '001', '/cloud_file_tag/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421077, 'p', '001', '/cloud_file_tag/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421078, 'p', '001', '/cloud_file_tag/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421079, 'p', '001', '/cloud_file_tag', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421080, 'p', '001', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421081, 'p', '001', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421082, 'p', '001', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421083, 'p', '001', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421084, 'p', '001', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421085, 'p', '001', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421086, 'p', '001', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421087, 'p', '001', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421088, 'p', '001', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421089, 'p', '001', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421090, 'p', '001', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421091, 'p', '001', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421092, 'p', '001', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421093, 'p', '001', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421094, 'p', '001', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421095, 'p', '001', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421096, 'p', '001', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421097, 'p', '001', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421098, 'p', '001', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421099, 'p', '001', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421100, 'p', '001', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421101, 'p', '001', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421102, 'p', '001', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421103, 'p', '001', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421104, 'p', '001', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421105, 'p', '001', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421106, 'p', '001', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421107, 'p', '001', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421108, 'p', '001', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421109, 'p', '001', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421110, 'p', '001', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421111, 'p', '001', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421112, 'p', '001', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421113, 'p', '001', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421114, 'p', '001', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421115, 'p', '001', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421116, 'p', '001', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421117, 'p', '001', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421118, 'p', '001', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421119, 'p', '001', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421120, 'p', '001', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421121, 'p', '001', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421122, 'p', '001', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421123, 'p', '001', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421124, 'p', '001', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421125, 'p', '001', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421126, 'p', '001', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421127, 'p', '001', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421128, 'p', '001', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421129, 'p', '001', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421130, 'p', '001', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421131, 'p', '001', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421132, 'p', '001', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421133, 'p', '001', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421134, 'p', '001', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421135, 'p', '001', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421136, 'p', '001', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421137, 'p', '001', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421138, 'p', '001', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421139, 'p', '001', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421140, 'p', '001', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421141, 'p', '001', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421142, 'p', '001', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421143, 'p', '001', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421144, 'p', '001', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421145, 'p', '001', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421146, 'p', '001', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421147, 'p', '001', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421148, 'p', '001', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421149, 'p', '001', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421150, 'p', '001', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421151, 'p', '001', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421152, 'p', '001', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421153, 'p', '001', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421154, 'p', '001', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421155, 'p', '001', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421156, 'p', '001', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421157, 'p', '001', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421158, 'p', '001', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421159, 'p', '001', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421160, 'p', '001', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421161, 'p', '001', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421162, 'p', '001', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421163, 'p', '001', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421164, 'p', '001', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421165, 'p', '001', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421166, 'p', '001', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421167, 'p', '001', '/upgrade_win', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421168, 'p', '001', '/upgrade_win/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421169, 'p', '001', '/upgrade_win/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421170, 'p', '001', '/upgrade_win/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421171, 'p', '001', '/upgrade_win/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421172, 'p', '001', '/upgrade_win_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421173, 'p', '001', '/upgrade_win_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421174, 'p', '001', '/upgrade_win_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421175, 'p', '001', '/upgrade_win_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421176, 'p', '001', '/upgrade_win_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421177, 'p', '001', '/upgrade_win_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421178, 'p', '001', '/upgrade_win_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421179, 'p', '001', '/upgrade_win_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421180, 'p', '001', '/upgrade_win_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421181, 'p', '001', '/upgrade_win_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421182, 'p', '001', '/upgrade_lnx', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421183, 'p', '001', '/upgrade_lnx/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421184, 'p', '001', '/upgrade_lnx/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421185, 'p', '001', '/upgrade_lnx/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421186, 'p', '001', '/upgrade_lnx/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421187, 'p', '001', '/upgrade_lnx_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421188, 'p', '001', '/upgrade_lnx_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421189, 'p', '001', '/upgrade_lnx_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421190, 'p', '001', '/upgrade_lnx_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421191, 'p', '001', '/upgrade_lnx_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421192, 'p', '001', '/upgrade_lnx_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421193, 'p', '001', '/upgrade_lnx_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421194, 'p', '001', '/upgrade_lnx_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421195, 'p', '001', '/upgrade_lnx_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421196, 'p', '001', '/upgrade_lnx_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421197, 'p', '001', '/upgrade_mac', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421198, 'p', '001', '/upgrade_mac/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421199, 'p', '001', '/upgrade_mac/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421200, 'p', '001', '/upgrade_mac/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421201, 'p', '001', '/upgrade_mac/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421202, 'p', '001', '/upgrade_mac_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421203, 'p', '001', '/upgrade_mac_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421204, 'p', '001', '/upgrade_mac_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421205, 'p', '001', '/upgrade_mac_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421206, 'p', '001', '/upgrade_mac_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421207, 'p', '001', '/upgrade_mac_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421208, 'p', '001', '/upgrade_mac_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421209, 'p', '001', '/upgrade_mac_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421210, 'p', '001', '/upgrade_mac_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421211, 'p', '001', '/upgrade_mac_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421958, 'p', '002', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421959, 'p', '002', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421960, 'p', '002', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421961, 'p', '002', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421962, 'p', '002', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421963, 'p', '002', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421964, 'p', '002', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421965, 'p', '002', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421966, 'p', '002', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421967, 'p', '002', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421968, 'p', '002', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421969, 'p', '002', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421970, 'p', '002', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421971, 'p', '002', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421972, 'p', '002', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421973, 'p', '002', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421974, 'p', '002', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421975, 'p', '002', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421976, 'p', '002', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421977, 'p', '002', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421978, 'p', '002', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421979, 'p', '002', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421980, 'p', '002', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421981, 'p', '002', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421982, 'p', '002', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421983, 'p', '002', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421984, 'p', '002', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421985, 'p', '002', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421986, 'p', '002', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421987, 'p', '002', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421988, 'p', '002', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421989, 'p', '002', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421990, 'p', '002', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421991, 'p', '002', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421992, 'p', '002', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421993, 'p', '002', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421994, 'p', '002', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421995, 'p', '002', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421996, 'p', '002', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421997, 'p', '002', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421998, 'p', '002', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (421999, 'p', '002', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422000, 'p', '002', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422001, 'p', '002', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422002, 'p', '002', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422003, 'p', '002', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422004, 'p', '002', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422005, 'p', '002', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422006, 'p', '002', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422007, 'p', '002', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422008, 'p', '002', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422009, 'p', '002', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422010, 'p', '002', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422011, 'p', '002', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422012, 'p', '002', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422013, 'p', '002', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422014, 'p', '002', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422015, 'p', '002', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422016, 'p', '002', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422017, 'p', '002', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422018, 'p', '002', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422019, 'p', '002', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422020, 'p', '002', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422021, 'p', '002', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422022, 'p', '002', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422023, 'p', '002', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422024, 'p', '002', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422025, 'p', '002', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422026, 'p', '002', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422027, 'p', '002', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422028, 'p', '002', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422029, 'p', '002', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422030, 'p', '002', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422031, 'p', '002', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422032, 'p', '002', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422033, 'p', '002', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422034, 'p', '002', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422035, 'p', '002', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422036, 'p', '002', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422037, 'p', '002', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422038, 'p', '002', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422039, 'p', '002', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422040, 'p', '002', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422041, 'p', '002', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422042, 'p', '002', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422043, 'p', '002', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422044, 'p', '002', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422045, 'p', '002', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422046, 'p', '002', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422047, 'p', '002', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422048, 'p', '002', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422049, 'p', '002', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422050, 'p', '002', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422051, 'p', '002', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422052, 'p', '002', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422053, 'p', '002', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422054, 'p', '002', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422055, 'p', '002', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422056, 'p', '002', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422057, 'p', '002', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422058, 'p', '002', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422059, 'p', '002', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422060, 'p', '002', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422061, 'p', '002', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422062, 'p', '002', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422063, 'p', '002', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422064, 'p', '002', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422065, 'p', '002', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422066, 'p', '002', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422067, 'p', '002', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422068, 'p', '002', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422069, 'p', '002', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422070, 'p', '002', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422071, 'p', '002', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422072, 'p', '002', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422073, 'p', '002', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422074, 'p', '002', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422075, 'p', '002', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422076, 'p', '002', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422077, 'p', '002', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422078, 'p', '002', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422079, 'p', '002', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422080, 'p', '002', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422081, 'p', '002', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422082, 'p', '002', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422083, 'p', '002', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422084, 'p', '002', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422085, 'p', '002', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422086, 'p', '002', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422087, 'p', '002', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422088, 'p', '002', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422089, 'p', '002', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422090, 'p', '002', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422091, 'p', '002', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422092, 'p', '002', '/upgrade_win', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422093, 'p', '002', '/upgrade_win/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422094, 'p', '002', '/upgrade_win/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422095, 'p', '002', '/upgrade_win/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422096, 'p', '002', '/upgrade_win/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422097, 'p', '002', '/upgrade_win_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422098, 'p', '002', '/upgrade_win_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422099, 'p', '002', '/upgrade_win_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422100, 'p', '002', '/upgrade_win_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422101, 'p', '002', '/upgrade_win_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422102, 'p', '002', '/upgrade_win_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422103, 'p', '002', '/upgrade_win_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422104, 'p', '002', '/upgrade_win_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422105, 'p', '002', '/upgrade_win_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422106, 'p', '002', '/upgrade_win_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422107, 'p', '002', '/upgrade_lnx', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422108, 'p', '002', '/upgrade_lnx/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422109, 'p', '002', '/upgrade_lnx/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422110, 'p', '002', '/upgrade_lnx/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422111, 'p', '002', '/upgrade_lnx/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422112, 'p', '002', '/upgrade_lnx_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422113, 'p', '002', '/upgrade_lnx_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422114, 'p', '002', '/upgrade_lnx_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422115, 'p', '002', '/upgrade_lnx_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422116, 'p', '002', '/upgrade_lnx_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422117, 'p', '002', '/upgrade_lnx_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422118, 'p', '002', '/upgrade_lnx_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422119, 'p', '002', '/upgrade_lnx_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422120, 'p', '002', '/upgrade_lnx_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422121, 'p', '002', '/upgrade_lnx_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422122, 'p', '002', '/upgrade_mac', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422123, 'p', '002', '/upgrade_mac/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422124, 'p', '002', '/upgrade_mac/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422125, 'p', '002', '/upgrade_mac/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422126, 'p', '002', '/upgrade_mac/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422127, 'p', '002', '/upgrade_mac_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422128, 'p', '002', '/upgrade_mac_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422129, 'p', '002', '/upgrade_mac_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422130, 'p', '002', '/upgrade_mac_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422131, 'p', '002', '/upgrade_mac_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422132, 'p', '002', '/upgrade_mac_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422133, 'p', '002', '/upgrade_mac_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422134, 'p', '002', '/upgrade_mac_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422135, 'p', '002', '/upgrade_mac_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422136, 'p', '002', '/upgrade_mac_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422137, 'p', '002', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422138, 'p', '002', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422139, 'p', '002', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422140, 'p', '002', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422141, 'p', '002', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422142, 'p', '002', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422143, 'p', '002', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422144, 'p', '002', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422145, 'p', '002', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422146, 'p', '002', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422147, 'p', '002', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422148, 'p', '002', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422149, 'p', '002', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422150, 'p', '002', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422151, 'p', '002', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422152, 'p', '002', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422153, 'p', '002', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422154, 'p', '002', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422155, 'p', '002', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422156, 'p', '002', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422157, 'p', '002', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422158, 'p', '002', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422159, 'p', '002', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422160, 'p', '002', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422161, 'p', '002', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422162, 'p', '002', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422163, 'p', '002', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422164, 'p', '002', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422165, 'p', '002', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422166, 'p', '002', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422167, 'p', '002', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422168, 'p', '002', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422169, 'p', '002', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422170, 'p', '002', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422171, 'p', '002', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422172, 'p', '002', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422173, 'p', '002', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422174, 'p', '002', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422175, 'p', '002', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422176, 'p', '002', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422177, 'p', '002', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422178, 'p', '002', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422179, 'p', '002', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422180, 'p', '002', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422181, 'p', '002', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422182, 'p', '002', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422183, 'p', '002', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422184, 'p', '002', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422185, 'p', '002', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422186, 'p', '002', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422187, 'p', '002', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422188, 'p', '002', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422189, 'p', '002', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422190, 'p', '002', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422191, 'p', '002', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422192, 'p', '002', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422193, 'p', '002', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422194, 'p', '002', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422195, 'p', '002', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422196, 'p', '002', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422197, 'p', '002', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422198, 'p', '002', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422199, 'p', '002', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422200, 'p', '002', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422201, 'p', '002', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422202, 'p', '002', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422203, 'p', '002', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422204, 'p', '002', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422205, 'p', '002', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422206, 'p', '002', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422207, 'p', '002', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422208, 'p', '002', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422209, 'p', '002', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422210, 'p', '002', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422211, 'p', '002', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422212, 'p', '002', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422213, 'p', '002', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422214, 'p', '002', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422215, 'p', '002', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422216, 'p', '002', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422217, 'p', '002', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422218, 'p', '002', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422219, 'p', '002', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422220, 'p', '002', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422221, 'p', '002', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422222, 'p', '002', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422223, 'p', '002', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422224, 'p', '002', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422225, 'p', '002', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422226, 'p', '002', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422227, 'p', '002', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422228, 'p', '002', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422229, 'p', '002', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422230, 'p', '002', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422231, 'p', '002', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422232, 'p', '002', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422233, 'p', '002', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422234, 'p', '002', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422235, 'p', '002', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422236, 'p', '002', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422237, 'p', '002', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422238, 'p', '002', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422239, 'p', '002', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422240, 'p', '002', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422241, 'p', '002', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422242, 'p', '002', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422243, 'p', '002', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422244, 'p', '002', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422245, 'p', '002', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422246, 'p', '002', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422247, 'p', '002', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422248, 'p', '002', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422249, 'p', '002', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422250, 'p', '002', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422251, 'p', '002', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422252, 'p', '002', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422253, 'p', '002', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422254, 'p', '002', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422255, 'p', '002', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422256, 'p', '002', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422257, 'p', '002', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422258, 'p', '002', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422259, 'p', '002', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422260, 'p', '002', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422261, 'p', '002', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422262, 'p', '002', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422263, 'p', '002', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422264, 'p', '002', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422265, 'p', '002', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422266, 'p', '002', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422267, 'p', '002', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422268, 'p', '002', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422269, 'p', '002', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422270, 'p', '002', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422271, 'p', '002', '/upgrade_win', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422272, 'p', '002', '/upgrade_win/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422273, 'p', '002', '/upgrade_win/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422274, 'p', '002', '/upgrade_win/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422275, 'p', '002', '/upgrade_win/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422276, 'p', '002', '/upgrade_win_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422277, 'p', '002', '/upgrade_win_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422278, 'p', '002', '/upgrade_win_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422279, 'p', '002', '/upgrade_win_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422280, 'p', '002', '/upgrade_win_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422281, 'p', '002', '/upgrade_win_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422282, 'p', '002', '/upgrade_win_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422283, 'p', '002', '/upgrade_win_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422284, 'p', '002', '/upgrade_win_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422285, 'p', '002', '/upgrade_win_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422286, 'p', '002', '/upgrade_lnx', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422287, 'p', '002', '/upgrade_lnx/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422288, 'p', '002', '/upgrade_lnx/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422289, 'p', '002', '/upgrade_lnx/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422290, 'p', '002', '/upgrade_lnx/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422291, 'p', '002', '/upgrade_lnx_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422292, 'p', '002', '/upgrade_lnx_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422293, 'p', '002', '/upgrade_lnx_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422294, 'p', '002', '/upgrade_lnx_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422295, 'p', '002', '/upgrade_lnx_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422296, 'p', '002', '/upgrade_lnx_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422297, 'p', '002', '/upgrade_lnx_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422298, 'p', '002', '/upgrade_lnx_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422299, 'p', '002', '/upgrade_lnx_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422300, 'p', '002', '/upgrade_lnx_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422301, 'p', '002', '/upgrade_mac', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422302, 'p', '002', '/upgrade_mac/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422303, 'p', '002', '/upgrade_mac/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422304, 'p', '002', '/upgrade_mac/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422305, 'p', '002', '/upgrade_mac/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422306, 'p', '002', '/upgrade_mac_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422307, 'p', '002', '/upgrade_mac_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422308, 'p', '002', '/upgrade_mac_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422309, 'p', '002', '/upgrade_mac_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422310, 'p', '002', '/upgrade_mac_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422311, 'p', '002', '/upgrade_mac_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422312, 'p', '002', '/upgrade_mac_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422313, 'p', '002', '/upgrade_mac_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422314, 'p', '002', '/upgrade_mac_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422315, 'p', '002', '/upgrade_mac_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422316, 'p', '002', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422317, 'p', '002', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422318, 'p', '002', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422319, 'p', '002', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422320, 'p', '002', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422321, 'p', '002', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422322, 'p', '002', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422323, 'p', '002', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422324, 'p', '002', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422325, 'p', '002', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422326, 'p', '002', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422327, 'p', '002', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422328, 'p', '002', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422329, 'p', '002', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422330, 'p', '002', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422331, 'p', '002', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422332, 'p', '002', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422333, 'p', '002', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422334, 'p', '002', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422335, 'p', '002', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422336, 'p', '002', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422337, 'p', '002', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422338, 'p', '002', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422339, 'p', '002', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422340, 'p', '002', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422341, 'p', '002', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422342, 'p', '002', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422343, 'p', '002', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422344, 'p', '002', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422345, 'p', '002', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422346, 'p', '002', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422347, 'p', '002', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422348, 'p', '002', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422349, 'p', '002', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422350, 'p', '002', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422351, 'p', '002', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422352, 'p', '002', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422353, 'p', '002', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422354, 'p', '002', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422355, 'p', '002', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422356, 'p', '002', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422357, 'p', '002', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422358, 'p', '002', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422359, 'p', '002', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422360, 'p', '002', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422361, 'p', '002', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422362, 'p', '002', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422363, 'p', '002', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422364, 'p', '002', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422365, 'p', '002', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422366, 'p', '002', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422367, 'p', '002', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422368, 'p', '002', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422369, 'p', '002', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422370, 'p', '002', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422371, 'p', '002', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422372, 'p', '002', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422373, 'p', '002', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422374, 'p', '002', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422375, 'p', '002', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422376, 'p', '002', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422377, 'p', '002', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422378, 'p', '002', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422379, 'p', '002', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422380, 'p', '002', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422381, 'p', '002', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422382, 'p', '002', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422383, 'p', '002', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422384, 'p', '002', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422385, 'p', '002', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422386, 'p', '002', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422387, 'p', '002', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422388, 'p', '002', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422389, 'p', '002', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422390, 'p', '002', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422391, 'p', '002', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422392, 'p', '002', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422393, 'p', '002', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422394, 'p', '002', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422395, 'p', '002', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422396, 'p', '002', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422397, 'p', '002', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422398, 'p', '002', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422399, 'p', '002', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422400, 'p', '002', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422401, 'p', '002', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422402, 'p', '002', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422403, 'p', '002', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422404, 'p', '002', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422405, 'p', '002', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422406, 'p', '002', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422407, 'p', '002', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422408, 'p', '002', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422409, 'p', '002', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422410, 'p', '002', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422411, 'p', '002', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422412, 'p', '002', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422413, 'p', '002', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422414, 'p', '002', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422415, 'p', '002', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422416, 'p', '002', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422417, 'p', '002', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422418, 'p', '002', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422419, 'p', '002', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422420, 'p', '002', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422421, 'p', '002', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422422, 'p', '002', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422423, 'p', '002', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422424, 'p', '002', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422425, 'p', '002', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422426, 'p', '002', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422427, 'p', '002', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422428, 'p', '002', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422429, 'p', '002', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422430, 'p', '002', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422431, 'p', '002', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422432, 'p', '002', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422433, 'p', '002', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422434, 'p', '002', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422435, 'p', '002', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422436, 'p', '002', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422437, 'p', '002', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422438, 'p', '002', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422439, 'p', '002', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422440, 'p', '002', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422441, 'p', '002', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422442, 'p', '002', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422443, 'p', '002', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422444, 'p', '002', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422445, 'p', '002', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422446, 'p', '002', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422447, 'p', '002', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422448, 'p', '002', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422449, 'p', '002', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422450, 'p', '002', '/upgrade_win', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422451, 'p', '002', '/upgrade_win/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422452, 'p', '002', '/upgrade_win/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422453, 'p', '002', '/upgrade_win/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422454, 'p', '002', '/upgrade_win/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422455, 'p', '002', '/upgrade_win_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422456, 'p', '002', '/upgrade_win_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422457, 'p', '002', '/upgrade_win_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422458, 'p', '002', '/upgrade_win_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422459, 'p', '002', '/upgrade_win_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422460, 'p', '002', '/upgrade_win_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422461, 'p', '002', '/upgrade_win_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422462, 'p', '002', '/upgrade_win_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422463, 'p', '002', '/upgrade_win_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422464, 'p', '002', '/upgrade_win_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422465, 'p', '002', '/upgrade_lnx', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422466, 'p', '002', '/upgrade_lnx/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422467, 'p', '002', '/upgrade_lnx/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422468, 'p', '002', '/upgrade_lnx/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422469, 'p', '002', '/upgrade_lnx/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422470, 'p', '002', '/upgrade_lnx_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422471, 'p', '002', '/upgrade_lnx_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422472, 'p', '002', '/upgrade_lnx_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422473, 'p', '002', '/upgrade_lnx_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422474, 'p', '002', '/upgrade_lnx_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422475, 'p', '002', '/upgrade_lnx_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422476, 'p', '002', '/upgrade_lnx_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422477, 'p', '002', '/upgrade_lnx_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422478, 'p', '002', '/upgrade_lnx_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422479, 'p', '002', '/upgrade_lnx_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422480, 'p', '002', '/upgrade_mac', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422481, 'p', '002', '/upgrade_mac/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422482, 'p', '002', '/upgrade_mac/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422483, 'p', '002', '/upgrade_mac/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422484, 'p', '002', '/upgrade_mac/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422485, 'p', '002', '/upgrade_mac_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422486, 'p', '002', '/upgrade_mac_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422487, 'p', '002', '/upgrade_mac_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422488, 'p', '002', '/upgrade_mac_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422489, 'p', '002', '/upgrade_mac_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422490, 'p', '002', '/upgrade_mac_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422491, 'p', '002', '/upgrade_mac_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422492, 'p', '002', '/upgrade_mac_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422493, 'p', '002', '/upgrade_mac_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (422494, 'p', '002', '/upgrade_mac_upgrade_strategy/update', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for fms_cloud_files
-- ----------------------------
DROP TABLE IF EXISTS `fms_cloud_files`;
CREATE TABLE `fms_cloud_files` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s name | 文件名',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s url | 文件地址',
  `size` bigint unsigned NOT NULL COMMENT 'The file''s size | 文件大小',
  `md5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'md5',
  `file_type` tinyint unsigned NOT NULL COMMENT 'The file''s type | 文件类型',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The user who upload the file | 上传用户的 ID',
  `cloud_file_storage_providers` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cloudfile_file_type` (`file_type`),
  KEY `cloudfile_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of fms_cloud_files
-- ----------------------------
BEGIN;
INSERT INTO `fms_cloud_files` (`id`, `created_at`, `updated_at`, `state`, `name`, `url`, `size`, `md5`, `file_type`, `user_id`, `cloud_file_storage_providers`) VALUES ('019bf528-5108-7a2e-a464-00ca5b9f2edd', '2026-01-25 12:36:59', '2026-01-25 12:36:59', 1, '019be9f6-139a-7b01-9653-b57c908a78fc', 'https://zcdn.upgrade.toolsetlink.com/upgrade/2026-01-25/image/019bf528-5108-7a2e-a464-00ca5b9f2edd.png', 7213, '044d01dc48cfd12bf6c4b0dfd7cbbfe7', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `role_id` bigint unsigned NOT NULL,
  `menu_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
BEGIN;
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 1);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 2);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 3);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 4);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 5);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 6);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 7);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 8);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 9);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 10);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 11);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 12);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 13);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 14);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 15);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 24);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 26);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 27);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 28);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 29);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 32);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 33);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 35);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 46);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 47);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 48);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 49);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 50);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 51);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 52);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 53);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 54);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 55);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 56);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 57);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 58);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 59);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 60);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 61);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 62);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 63);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 65);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 66);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 67);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 68);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 69);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 70);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 73);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 74);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 75);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 76);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 77);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 78);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 82);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 88);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 89);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 90);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 91);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 92);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 93);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 94);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 95);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 96);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 97);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 98);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 1);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 9);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 10);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 24);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 26);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 27);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 28);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 29);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 32);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 33);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 35);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 52);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 53);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 54);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 55);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 56);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 57);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 58);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 59);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 60);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 61);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 62);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 63);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 65);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 66);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 67);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 68);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 69);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 70);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 73);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 74);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 75);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 76);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 77);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 78);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 82);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 88);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 89);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 90);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 91);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 92);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 93);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 94);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 95);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 96);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 97);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (2, 98);
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'API path | API 路径',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'API description | API 描述',
  `api_group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'API group | API 分组',
  `service_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Other' COMMENT 'Service name | 服务名称',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'POST' COMMENT 'HTTP method | HTTP 请求类型',
  `is_required` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Whether is required | 是否必选',
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_path_method` (`path`,`method`)
) ENGINE=InnoDB AUTO_INCREMENT=359 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (1, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/login', 'apiDesc.userLogin', 'user', 'Core', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (2, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/register', 'apiDesc.userRegister', 'user', 'Core', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (3, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/create', 'apiDesc.createUser', 'user', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (4, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/update', 'apiDesc.updateUser', 'user', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (5, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/change_password', 'apiDesc.userChangePassword', 'user', 'Core', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (6, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/info', 'apiDesc.userInfo', 'user', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (7, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/list', 'apiDesc.userList', 'user', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (8, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/delete', 'apiDesc.deleteUser', 'user', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (9, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/perm', 'apiDesc.userPermissions', 'user', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (10, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/profile', 'apiDesc.userProfile', 'user', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (11, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/profile', 'apiDesc.updateProfile', 'user', 'Core', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (12, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/logout', 'apiDesc.logout', 'user', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (13, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user', 'apiDesc.getUserById', 'user', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (14, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/refresh_token', 'apiDesc.refreshToken', 'user', 'Core', 'GET', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (15, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/user/access_token', 'apiDesc.accessToken', 'user', 'Core', 'GET', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (16, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/role/create', 'apiDesc.createRole', 'role', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (17, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/role/update', 'apiDesc.updateRole', 'role', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (18, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/role/delete', 'apiDesc.deleteRole', 'role', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (19, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/role/list', 'apiDesc.roleList', 'role', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (20, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/role', 'apiDesc.getRoleById', 'role', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (21, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu/create', 'apiDesc.createMenu', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (22, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu/update', 'apiDesc.updateMenu', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (23, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu/delete', 'apiDesc.deleteMenu', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (24, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu/list', 'apiDesc.menuList', 'menu', 'Core', 'GET', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (25, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu/role/list', 'apiDesc.menuRoleList', 'authority', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (26, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu_param/create', 'apiDesc.createMenuParam', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (27, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu_param/update', 'apiDesc.updateMenuParam', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (28, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu_param/list', 'apiDesc.menuParamListByMenuId', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (29, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu_param/delete', 'apiDesc.deleteMenuParam', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (30, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu_param', 'apiDesc.getMenuParamById', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (31, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/menu', 'apiDesc.getMenuById', 'menu', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (32, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/captcha', 'apiDesc.captcha', 'captcha', 'Core', 'GET', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (33, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/authority/api/create_or_update', 'apiDesc.createOrUpdateApiAuthority', 'authority', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (34, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/authority/api/role', 'apiDesc.APIAuthorityOfRole', 'authority', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (35, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/authority/menu/create_or_update', 'apiDesc.createOrUpdateMenuAuthority', 'authority', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (36, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/authority/menu/role', 'apiDesc.menuAuthorityOfRole', 'authority', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (37, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/api/create', 'apiDesc.createApi', 'api', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (38, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/api/update', 'apiDesc.updateApi', 'api', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (39, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/api/delete', 'apiDesc.deleteAPI', 'api', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (40, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/api/list', 'apiDesc.APIList', 'api', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (41, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/api', 'apiDesc.getApiById', 'api', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (42, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary', 'apiDesc.getDictionaryById', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (43, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary/create', 'apiDesc.createDictionary', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (44, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary/update', 'apiDesc.updateDictionary', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (45, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary/delete', 'apiDesc.deleteDictionary', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (46, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary_detail/delete', 'apiDesc.deleteDictionaryDetail', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (47, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary_detail', 'apiDesc.getDictionaryDetailById', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (48, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary_detail/create', 'apiDesc.createDictionaryDetail', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (49, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary_detail/update', 'apiDesc.updateDictionaryDetail', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (50, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary_detail/list', 'apiDesc.getDictionaryListDetail', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (51, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dictionary/list', 'apiDesc.getDictionaryList', 'dictionary', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (52, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/dict/:name', 'apiDesc.getDictionaryDetailByDictionaryName', 'dictionary', 'Core', 'GET', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (53, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth_provider/create', 'apiDesc.createProvider', 'oauth', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (54, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth_provider/update', 'apiDesc.updateProvider', 'oauth', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (55, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth_provider/delete', 'apiDesc.deleteProvider', 'oauth', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (56, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth_provider/list', 'apiDesc.getProviderList', 'oauth', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (57, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth/login', 'apiDesc.oauthLogin', 'oauth', 'Core', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (58, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/oauth_provider', 'apiDesc.getProviderById', 'oauth', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (59, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token/create', 'apiDesc.createToken', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (60, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token/update', 'apiDesc.updateToken', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (61, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token/delete', 'apiDesc.deleteToken', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (62, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token/list', 'apiDesc.getTokenList', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (63, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token/logout', 'apiDesc.forceLoggingOut', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (64, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/token', 'apiDesc.getTokenById', 'token', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (65, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/department/create', 'apiDesc.createDepartment', 'department', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (66, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/department/update', 'apiDesc.updateDepartment', 'department', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (67, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/department/delete', 'apiDesc.deleteDepartment', 'department', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (68, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/department/list', 'apiDesc.getDepartmentList', 'department', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (69, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/department', 'apiDesc.getDepartmentById', 'department', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (70, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/position/create', 'apiDesc.createPosition', 'position', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (71, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/position/update', 'apiDesc.updatePosition', 'position', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (72, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/position/delete', 'apiDesc.deletePosition', 'position', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (73, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/position/list', 'apiDesc.getPositionList', 'position', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (74, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/position', 'apiDesc.getPositionById', 'position', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (75, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task/create', 'apiDesc.createTask', 'task', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (76, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task/update', 'apiDesc.updateTask', 'task', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (77, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task/delete', 'apiDesc.deleteTask', 'task', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (78, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task/list', 'apiDesc.getTaskList', 'task', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (79, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task', 'apiDesc.getTaskById', 'task', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (80, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task_log/create', 'apiDesc.createTaskLog', 'task_log', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (81, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task_log/update', 'apiDesc.updateTaskLog', 'task_log', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (82, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task_log/delete', 'apiDesc.deleteTaskLog', 'task_log', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (83, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task_log/list', 'apiDesc.getTaskLogList', 'task_log', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (84, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/task_log', 'apiDesc.getTaskLogById', 'task_log', 'Job', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (85, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/configuration/create', 'apiDesc.createConfiguration', 'configuration', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (86, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/configuration/update', 'apiDesc.updateConfiguration', 'configuration', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (87, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/configuration/delete', 'apiDesc.deleteConfiguration', 'configuration', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (88, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/configuration/list', 'apiDesc.getConfigurationList', 'configuration', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (89, '2025-02-10 09:33:13', '2025-02-10 09:33:13', '/configuration', 'apiDesc.getConfigurationById', 'configuration', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (90, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_log/create', 'apiDesc.createEmailLog', 'email_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (91, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_log/update', 'apiDesc.updateEmailLog', 'email_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (92, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_log/delete', 'apiDesc.deleteEmailLog', 'email_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (93, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_log/list', 'apiDesc.getEmailLogList', 'email_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (94, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_log', 'apiDesc.getEmailLogById', 'email_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (95, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_provider/create', 'apiDesc.createEmailProvider', 'email_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (96, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_provider/update', 'apiDesc.updateEmailProvider', 'email_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (97, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_provider/delete', 'apiDesc.deleteEmailProvider', 'email_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (98, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_provider/list', 'apiDesc.getEmailProviderList', 'email_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (99, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email_provider', 'apiDesc.getEmailProviderById', 'email_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (100, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_log/create', 'apiDesc.createSmsLog', 'sms_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (101, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_log/update', 'apiDesc.updateSmsLog', 'sms_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (102, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_log/delete', 'apiDesc.deleteSmsLog', 'sms_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (103, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_log/list', 'apiDesc.getSmsLogList', 'sms_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (104, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_log', 'apiDesc.getSmsLogById', 'sms_log', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (105, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_provider/create', 'apiDesc.createSmsProvider', 'sms_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (106, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_provider/update', 'apiDesc.updateSmsProvider', 'sms_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (107, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_provider/delete', 'apiDesc.deleteSmsProvider', 'sms_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (108, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_provider/list', 'apiDesc.getSmsProviderList', 'sms_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (109, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms_provider', 'apiDesc.getSmsProviderById', 'sms_provider', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (110, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/sms/send', 'apiDesc.sendSms', 'message_sender', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (111, '2025-02-10 09:50:08', '2025-02-10 09:50:08', '/email/send', 'apiDesc.sendEmail', 'message_sender', 'Mcms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (112, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member/create', 'apiDesc.createMember', 'member', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (113, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member/update', 'apiDesc.updateMember', 'member', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (114, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member/delete', 'apiDesc.deleteMember', 'member', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (115, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member/list', 'apiDesc.getMemberList', 'member', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (116, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member', 'apiDesc.getMemberById', 'member', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (117, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member_rank/create', 'apiDesc.createMemberRank', 'member_rank', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (118, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member_rank/update', 'apiDesc.updateMemberRank', 'member_rank', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (119, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member_rank/delete', 'apiDesc.deleteMemberRank', 'member_rank', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (120, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member_rank/list', 'apiDesc.getMemberRankList', 'member_rank', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (121, '2025-02-11 02:49:47', '2025-02-11 02:49:47', '/member_rank', 'apiDesc.getMemberRankById', 'member_rank', 'Mms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (122, '2025-02-20 09:25:42', '2025-10-30 15:59:12', '/upgrade_url/list', 'apiDesc.getUrlList', 'upgrade_url', 'Upgrade', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (123, '2025-02-20 09:25:42', '2025-10-30 15:59:13', '/upgrade_url/delete', 'apiDesc.deleteUrl', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (127, '2025-02-20 09:25:42', '2025-10-30 15:59:14', '/upgrade_url/create', 'apiDesc.createUrl', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (129, '2025-02-20 09:25:42', '2025-10-30 15:59:16', '/upgrade_url/update', 'apiDesc.updateUrl', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (130, '2025-02-20 09:25:42', '2025-10-30 15:59:17', '/upgrade_url', 'apiDesc.getUrlById', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (131, '2025-02-20 09:25:42', '2025-10-30 15:59:18', '/upgrade_dev_model', 'apiDesc.getDevModelById', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (132, '2025-02-20 09:25:42', '2025-10-30 15:59:19', '/upgrade_dev_model/list', 'apiDesc.getDevModelList', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (133, '2025-02-20 09:25:42', '2025-10-30 15:59:20', '/upgrade_dev_model/delete', 'apiDesc.deleteDevModel', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (134, '2025-02-20 09:25:42', '2025-10-30 15:59:22', '/upgrade_dev_model/create', 'apiDesc.createDevModel', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (135, '2025-02-20 09:25:42', '2025-10-30 15:59:23', '/upgrade_dev_model/update', 'apiDesc.updateDevModel', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (136, '2025-02-20 09:25:42', '2025-10-30 15:59:24', '/upgrade_url_upgrade_strategy', 'apiDesc.getUrlUpgradeStrategyById', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (137, '2025-02-20 09:25:42', '2025-10-30 15:59:25', '/upgrade_url_upgrade_strategy/list', 'apiDesc.getUrlUpgradeStrategyList', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (138, '2025-02-20 09:25:42', '2025-10-30 15:59:26', '/upgrade_url_upgrade_strategy/update', 'apiDesc.updateUrlUpgradeStrategy', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (139, '2025-02-20 09:25:42', '2025-10-30 15:59:27', '/upgrade_url_upgrade_strategy/create', 'apiDesc.createUrlUpgradeStrategy', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (140, '2025-02-20 09:25:42', '2025-10-30 15:59:28', '/upgrade_url_upgrade_strategy/delete', 'apiDesc.deleteUrlUpgradeStrategy', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (141, '2025-02-20 09:25:42', '2025-10-30 15:59:29', '/upgrade_url_version', 'apiDesc.getUrlVersionById', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (142, '2025-02-20 09:25:42', '2025-10-30 15:59:30', '/upgrade_url_version/list', 'apiDesc.getUrlVersionList', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (143, '2025-02-20 09:25:42', '2025-10-30 15:59:31', '/upgrade_url_version/delete', 'apiDesc.deleteUrlVersion', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (144, '2025-02-20 09:25:42', '2025-10-30 15:59:32', '/upgrade_url_version/create', 'apiDesc.createUrlVersion', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (145, '2025-02-20 09:25:42', '2025-10-30 15:59:33', '/upgrade_url_version/update', 'apiDesc.updateUrlVersion', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (146, '2025-02-20 09:25:42', '2025-10-30 15:59:33', '/company_secret', 'apiDesc.getSecretById', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (147, '2025-02-20 09:25:42', '2025-10-30 15:59:34', '/company_secret/list', 'apiDesc.getSecretList', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (148, '2025-02-20 09:25:42', '2025-10-30 15:59:35', '/company_secret/delete', 'apiDesc.deleteSecret', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (149, '2025-02-20 09:25:42', '2025-10-30 15:59:36', '/company_secret/create', 'apiDesc.createSecret', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (150, '2025-02-20 09:25:42', '2025-10-30 15:59:37', '/company_secret/update', 'apiDesc.updateSecret', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (151, '2025-02-20 09:25:42', '2025-10-30 15:59:38', '/upgrade_dev_swarm', 'apiDesc.getDevSwarmById', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (152, '2025-02-20 09:25:42', '2025-10-30 15:59:39', '/upgrade_dev_swarm/list', 'apiDesc.getDevSwarmList', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (153, '2025-02-20 09:25:42', '2025-10-30 15:59:40', '/upgrade_dev_swarm/delete', 'apiDesc.deleteDevSwarm', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (154, '2025-02-20 09:25:42', '2025-10-30 15:59:41', '/upgrade_dev_swarm/create', 'apiDesc.createDevSwarm', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (155, '2025-02-20 09:25:42', '2025-10-30 15:59:42', '/upgrade_dev_swarm/update', 'apiDesc.updateDevSwarm', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (156, '2025-02-20 09:25:42', '2025-10-30 15:59:44', '/upgrade_file', 'apiDesc.getFileAppById', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (157, '2025-02-20 09:25:42', '2025-10-30 15:59:45', '/upgrade_file/list', 'apiDesc.getFileAppList', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (158, '2025-02-20 09:25:42', '2025-10-30 15:59:46', '/upgrade_file/delete', 'apiDesc.deleteFileApp', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (159, '2025-02-20 09:25:42', '2025-10-30 15:59:47', '/upgrade_file/update', 'apiDesc.updateFileApp', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (160, '2025-02-20 09:25:42', '2025-10-30 15:59:48', '/upgrade_file/create', 'apiDesc.createFileApp', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (171, '2025-02-20 09:25:42', '2025-10-30 15:59:49', '/upgrade_dashboard', 'apiDesc.upgrade_dashboard', 'upgrade_dashboard', 'Upgrade', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (199, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/upload', 'apiDesc.uploadFile', 'file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (200, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file/list', 'apiDesc.fileList', 'file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (201, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file/update', 'apiDesc.updateFileInfo', 'file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (202, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file/status', 'apiDesc.setPublicStatus', 'file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (203, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file/delete', 'apiDesc.deleteFile', 'file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (204, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file/download/:id', 'apiDesc.downloadFile', 'file', 'Fms', 'GET', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (205, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file_tag/create', 'apiDesc.createFileTag', 'file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (206, '2025-03-10 02:45:02', '2025-03-10 02:45:02', '/file_tag/update', 'apiDesc.updateFileTag', 'file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (207, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/file_tag/delete', 'apiDesc.deleteFileTag', 'file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (208, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/file_tag/list', 'apiDesc.getFileTagList', 'file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (209, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/file_tag', 'apiDesc.getFileTagById', 'file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (210, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/storage_provider/create', 'apiDesc.createStorageProvider', 'storage_provider', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (211, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/storage_provider/update', 'apiDesc.updateStorageProvider', 'storage_provider', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (212, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/storage_provider/delete', 'apiDesc.deleteStorageProvider', 'storage_provider', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (213, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/storage_provider/list', 'apiDesc.getStorageProviderList', 'storage_provider', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (214, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/storage_provider', 'apiDesc.getStorageProviderById', 'storage_provider', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (215, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file/create', 'apiDesc.createCloudFile', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (216, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file/update', 'apiDesc.updateCloudFile', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (217, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file/delete', 'apiDesc.deleteCloudFile', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (218, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file/list', 'apiDesc.getCloudFileList', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (219, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file', 'apiDesc.getCloudFileById', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (220, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file/upload', 'apiDesc.uploadFileToCloud', 'cloud_file', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (221, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file_tag/create', 'apiDesc.createCloudFileTag', 'cloud_file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (222, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file_tag/update', 'apiDesc.updateCloudFileTag', 'cloud_file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (223, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file_tag/delete', 'apiDesc.deleteCloudFileTag', 'cloud_file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (224, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file_tag/list', 'apiDesc.getCloudFileTagList', 'cloud_file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (225, '2025-03-10 02:45:03', '2025-03-10 02:45:03', '/cloud_file_tag', 'apiDesc.getCloudFileTagById', 'cloud_file_tag', 'Fms', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (226, '2025-02-20 09:25:42', '2025-10-30 16:00:03', '/upgrade_file_version', 'apiDesc.getFileVersionById', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (227, '2025-02-20 09:25:42', '2025-10-30 16:00:04', '/upgrade_file_version/list', 'apiDesc.getFileVersionList', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (228, '2025-02-20 09:25:42', '2025-10-30 16:00:05', '/upgrade_file_version/delete', 'apiDesc.deleteFileVersion', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (229, '2025-02-20 09:25:42', '2025-10-30 16:00:06', '/upgrade_file_version/create', 'apiDesc.createFileVersion', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (230, '2025-02-20 09:25:42', '2025-10-30 16:00:07', '/upgrade_file_version/update', 'apiDesc.updateFileVersion', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (231, '2025-02-20 09:25:42', '2025-10-30 16:00:08', '/upgrade_file_upgrade_strategy', 'apiDesc.getFileUpgradeStrategyById', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (232, '2025-02-20 09:25:42', '2025-10-30 16:00:09', '/upgrade_file_upgrade_strategy/list', 'apiDesc.getFileUpgradeStrategyList', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (233, '2025-02-20 09:25:42', '2025-10-30 16:00:11', '/upgrade_file_upgrade_strategy/delete', 'apiDesc.deleteFileUpgradeStrategy', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (234, '2025-02-20 09:25:42', '2025-10-30 16:00:12', '/upgrade_file_upgrade_strategy/create', 'apiDesc.createFileUpgradeStrategy', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (235, '2025-02-20 09:25:42', '2025-10-30 16:00:13', '/upgrade_file_upgrade_strategy/update', 'apiDesc.updateFileUpgradeStrategy', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (236, '2025-02-20 09:25:42', '2025-10-30 16:00:14', '/upgrade_dev', 'apiDesc.getDevById', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (237, '2025-02-20 09:25:42', '2025-10-30 16:00:15', '/upgrade_dev/list', 'apiDesc.getDevList', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (238, '2025-02-20 09:25:42', '2025-10-30 16:00:17', '/upgrade_dev/delete', 'apiDesc.deleteDev', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (239, '2025-02-20 09:25:42', '2025-10-30 16:00:18', '/upgrade_dev/create', 'apiDesc.createDev', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (240, '2025-02-20 09:25:42', '2025-10-30 16:00:19', '/upgrade_dev/update', 'apiDesc.updateDev', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (242, '2025-02-20 09:25:42', '2025-10-30 16:00:20', '/upgrade_dev_group', 'apiDesc.getDevGroupById', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (243, '2025-02-20 09:25:42', '2025-10-30 16:00:21', '/upgrade_dev_group/list', 'apiDesc.getDevGroupList', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (244, '2025-02-20 09:25:42', '2025-10-30 16:00:22', '/upgrade_dev_group/delete', 'apiDesc.deleteDevGroup', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (245, '2025-02-20 09:25:42', '2025-10-30 16:00:23', '/upgrade_dev_group/create', 'apiDesc.createDevGroup', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (246, '2025-02-20 09:25:42', '2025-10-30 16:00:26', '/upgrade_dev_group/update', 'apiDesc.updateDevGroup', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (247, '2025-02-20 09:25:42', '2025-10-30 16:00:27', '/upgrade_tauri', 'apiDesc.getTauriAppById', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (248, '2025-02-20 09:25:42', '2025-10-30 16:00:28', '/upgrade_tauri/list', 'apiDesc.getTauriAppList', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (249, '2025-02-20 09:25:42', '2025-10-30 16:00:29', '/upgrade_tauri/delete', 'apiDesc.deleteTauriApp', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (250, '2025-02-20 09:25:42', '2025-10-30 16:00:30', '/upgrade_tauri/update', 'apiDesc.updateTauriApp', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (251, '2025-02-20 09:25:42', '2025-10-30 16:00:31', '/upgrade_tauri/create', 'apiDesc.createTauriApp', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (252, '2025-02-20 09:25:42', '2025-10-30 16:00:32', '/upgrade_tauri_version', 'apiDesc.getTauriVersionById', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (253, '2025-02-20 09:25:42', '2025-10-30 16:00:33', '/upgrade_tauri_version/list', 'apiDesc.getTauriVersionList', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (254, '2025-02-20 09:25:42', '2025-10-30 16:00:35', '/upgrade_tauri_version/delete', 'apiDesc.deleteTauriVersion', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (255, '2025-02-20 09:25:42', '2025-10-30 16:00:36', '/upgrade_tauri_version/create', 'apiDesc.createTauriVersion', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (256, '2025-02-20 09:25:42', '2025-10-30 16:00:37', '/upgrade_tauri_version/update', 'apiDesc.updateTauriVersion', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (257, '2025-02-20 09:25:42', '2025-10-30 16:00:38', '/upgrade_tauri_upgrade_strategy', 'apiDesc.getTauriUpgradeStrategyById', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (258, '2025-02-20 09:25:42', '2025-10-30 16:00:39', '/upgrade_tauri_upgrade_strategy/list', 'apiDesc.getTauriUpgradeStrategyList', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (259, '2025-02-20 09:25:42', '2025-10-30 16:00:40', '/upgrade_tauri_upgrade_strategy/delete', 'apiDesc.deleteTauriUpgradeStrategy', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (260, '2025-02-20 09:25:42', '2025-10-30 16:00:41', '/upgrade_tauri_upgrade_strategy/create', 'apiDesc.createTauriUpgradeStrategy', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (261, '2025-02-20 09:25:42', '2025-10-30 16:00:42', '/upgrade_tauri_upgrade_strategy/update', 'apiDesc.updateTauriUpgradeStrategy', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (262, '2025-02-20 09:25:42', '2025-10-30 16:00:43', '/upgrade_configuration', 'apiDesc.getConfigurationAppById', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (263, '2025-02-20 09:25:42', '2025-10-30 16:00:44', '/upgrade_configuration/list', 'apiDesc.getConfigurationAppList', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (264, '2025-02-20 09:25:42', '2025-10-30 16:00:45', '/upgrade_configuration/delete', 'apiDesc.deleteConfigurationApp', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (265, '2025-02-20 09:25:42', '2025-10-30 16:00:46', '/upgrade_configuration/create', 'apiDesc.createConfigurationApp', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (266, '2025-02-20 09:25:42', '2025-10-30 16:00:47', '/upgrade_configuration/update', 'apiDesc.updateConfigurationApp', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (267, '2025-02-20 09:25:42', '2025-10-30 16:00:48', '/upgrade_configuration_version', 'apiDesc.getConfigurationVersionById', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (268, '2025-02-20 09:25:42', '2025-10-30 16:00:49', '/upgrade_configuration_version/list', 'apiDesc.getConfigurationVersionList', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (269, '2025-02-20 09:25:42', '2025-10-30 16:00:50', '/upgrade_configuration_version/delete', 'apiDesc.deleteConfigurationVersion', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (270, '2025-02-20 09:25:42', '2025-10-30 16:00:52', '/upgrade_configuration_version/create', 'apiDesc.createConfigurationVersion', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (271, '2025-02-20 09:25:42', '2025-10-30 16:00:53', '/upgrade_configuration_version/update', 'apiDesc.updateConfigurationVersion', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (272, '2025-02-20 09:25:42', '2025-10-30 16:00:54', '/upgrade_configuration_upgrade_strategy', 'apiDesc.getConfigurationUpgradeStrategyById', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (273, '2025-02-20 09:25:42', '2025-10-30 16:00:55', '/upgrade_configuration_upgrade_strategy/list', 'apiDesc.getConfigurationUpgradeStrategyList', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (274, '2025-02-20 09:25:42', '2025-10-30 16:00:56', '/upgrade_configuration_upgrade_strategy/delete', 'apiDesc.deleteConfigurationUpgradeStrategy', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (275, '2025-02-20 09:25:42', '2025-10-30 16:00:57', '/upgrade_configuration_upgrade_strategy/create', 'apiDesc.createConfigurationUpgradeStrategy', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (276, '2025-02-20 09:25:42', '2025-10-30 16:00:58', '/upgrade_configuration_upgrade_strategy/update', 'apiDesc.updateConfigurationUpgradeStrategy', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (277, '2025-02-20 09:25:42', '2025-10-30 16:00:59', '/upgrade_apk', 'apiDesc.getApkAppById', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (278, '2025-02-20 09:25:42', '2025-10-30 16:01:00', '/upgrade_apk/list', 'apiDesc.getApkAppList', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (279, '2025-02-20 09:25:42', '2025-10-30 16:01:01', '/upgrade_apk/delete', 'apiDesc.deleteApkApp', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (280, '2025-02-20 09:25:42', '2025-10-30 16:01:02', '/upgrade_apk/create', 'apiDesc.createApkApp', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (281, '2025-02-20 09:25:42', '2025-10-30 16:01:04', '/upgrade_apk/update', 'apiDesc.updateApkApp', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (282, '2025-02-20 09:25:42', '2025-10-30 16:01:05', '/upgrade_apk_version', 'apiDesc.getApkVersionById', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (283, '2025-02-20 09:25:42', '2025-10-30 16:01:06', '/upgrade_apk_version/list', 'apiDesc.getApkVersionList', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (284, '2025-02-20 09:25:42', '2025-10-30 16:01:07', '/upgrade_apk_version/delete', 'apiDesc.deleteApkVersion', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (285, '2025-02-20 09:25:42', '2025-10-30 16:01:08', '/upgrade_apk_version/create', 'apiDesc.createApkVersion', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (286, '2025-02-20 09:25:42', '2025-10-30 16:01:09', '/upgrade_apk_version/update', 'apiDesc.updateApkVersion', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (287, '2025-02-20 09:25:42', '2025-10-30 16:01:10', '/upgrade_apk_upgrade_strategy', 'apiDesc.getApkUpgradeStrategyById', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (288, '2025-02-20 09:25:42', '2025-10-30 16:01:11', '/upgrade_apk_upgrade_strategy/list', 'apiDesc.getApkUpgradeStrategyList', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (289, '2025-02-20 09:25:42', '2025-10-30 16:01:12', '/upgrade_apk_upgrade_strategy/delete', 'apiDesc.deleteApkUpgradeStrategy', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (290, '2025-02-20 09:25:42', '2025-10-30 16:01:13', '/upgrade_apk_upgrade_strategy/create', 'apiDesc.createApkUpgradeStrategy', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (291, '2025-02-20 09:25:42', '2025-10-30 16:01:14', '/upgrade_apk_upgrade_strategy/update', 'apiDesc.updateApkUpgradeStrategy', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (292, '2025-02-20 09:25:42', '2025-10-30 16:01:15', '/upgrade_electron', 'apiDesc.getElectronAppById', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (293, '2025-02-20 09:25:42', '2025-10-30 16:01:16', '/upgrade_electron/list', 'apiDesc.getElectronAppList', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (294, '2025-02-20 09:25:42', '2025-10-30 16:01:18', '/upgrade_electron/delete', 'apiDesc.deleteElectronApp', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (295, '2025-02-20 09:25:42', '2025-10-30 16:01:19', '/upgrade_electron/create', 'apiDesc.createElectronApp', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (296, '2025-02-20 09:25:42', '2025-10-30 16:01:20', '/upgrade_electron/update', 'apiDesc.updateElectronApp', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (297, '2025-02-20 09:25:42', '2025-10-30 16:01:21', '/upgrade_electron_version', 'apiDesc.getElectronVersionById', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (298, '2025-02-20 09:25:42', '2025-10-30 16:01:23', '/upgrade_electron_version/list', 'apiDesc.getElectronVersionList', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (299, '2025-02-20 09:25:42', '2025-10-30 16:01:24', '/upgrade_electron_version/delete', 'apiDesc.deleteElectronVersion', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (300, '2025-02-20 09:25:42', '2025-10-30 16:01:25', '/upgrade_electron_version/create', 'apiDesc.createElectronVersion', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (301, '2025-02-20 09:25:42', '2025-10-30 16:01:26', '/upgrade_electron_version/update', 'apiDesc.updateElectronVersion', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (302, '2025-02-20 09:25:42', '2025-10-30 16:01:27', '/upgrade_electron_upgrade_strategy', 'apiDesc.getElectronUpgradeStrategyById', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (303, '2025-02-20 09:25:42', '2025-10-30 16:01:29', '/upgrade_electron_upgrade_strategy/list', 'apiDesc.getElectronUpgradeStrategyList', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (304, '2025-02-20 09:25:42', '2025-10-30 16:01:30', '/upgrade_electron_upgrade_strategy/delete', 'apiDesc.deleteElectronUpgradeStrategy', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (305, '2025-02-20 09:25:42', '2025-10-30 16:01:31', '/upgrade_electron_upgrade_strategy/create', 'apiDesc.createElectronUpgradeStrategy', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (306, '2025-02-20 09:25:42', '2025-10-30 16:01:33', '/upgrade_electron_upgrade_strategy/update', 'apiDesc.updateElectronUpgradeStrategy', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (307, '2025-02-20 09:25:42', '2025-10-30 16:01:34', '/upgrade_company_income', 'apiDesc.getIncomeById', 'upgrade_company_income', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (308, '2025-02-20 09:25:42', '2025-10-30 16:01:35', '/upgrade_company_income/list', 'apiDesc.getIncomeList', 'upgrade_company_income', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (309, '2025-09-09 09:16:34', '2025-10-30 16:01:36', '/upgrade_company_traffic_packet', 'apiDesc.getTrafficPacketById', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (310, '2025-09-09 09:16:26', '2025-10-30 16:01:37', '/upgrade_company_traffic_packet/list', 'apiDesc.getTrafficPacketList', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (311, '2025-09-09 09:16:28', '2025-10-30 16:01:38', '/upgrade_company_traffic_packet/delete', 'apiDesc.deleteTrafficPacket', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (312, '2025-09-09 09:16:29', '2025-10-30 16:01:39', '/upgrade_company_traffic_packet/create', 'apiDesc.createTrafficPacket', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (313, '2025-09-09 09:16:32', '2025-10-30 16:01:40', '/upgrade_company_traffic_packet/update', 'apiDesc.updateTrafficPacket', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (314, '2025-10-27 18:48:13', '2025-10-30 16:01:41', '/upgrade_win', 'apiDesc.getWinAppById', 'upgrade_win', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (315, '2025-10-27 18:48:19', '2025-10-30 16:01:42', '/upgrade_win/list', 'apiDesc.getWinAppList', 'upgrade_win', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (316, '2025-10-27 18:48:26', '2025-10-30 16:01:43', '/upgrade_win/delete', 'apiDesc.deleteWinApp', 'upgrade_win', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (317, '2025-10-27 18:48:31', '2025-10-30 16:01:44', '/upgrade_win/create', 'apiDesc.createWinApp', 'upgrade_win', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (318, '2025-10-27 18:48:37', '2025-10-30 16:01:45', '/upgrade_win/update', 'apiDesc.updateWinApp', 'upgrade_win', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (319, '2025-10-27 18:48:42', '2025-10-30 16:01:46', '/upgrade_win_version', 'apiDesc.getWinVersionById', 'upgrade_win_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (320, '2025-10-27 18:48:46', '2025-10-30 16:01:47', '/upgrade_win_version/list', 'apiDesc.getWinVersionList', 'upgrade_win_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (321, '2025-10-27 18:48:51', '2025-10-30 16:01:48', '/upgrade_win_version/delete', 'apiDesc.deleteWinVersion', 'upgrade_win_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (322, '2025-10-27 18:48:57', '2025-10-30 16:01:49', '/upgrade_win_version/create', 'apiDesc.createWinVersion', 'upgrade_win_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (323, '2025-10-27 18:49:06', '2025-10-30 16:01:50', '/upgrade_win_version/update', 'apiDesc.updateWinVersion', 'upgrade_win_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (324, '2025-10-27 18:49:11', '2025-10-30 16:01:51', '/upgrade_win_upgrade_strategy', 'apiDesc.getWinUpgradeStrategyById', 'upgrade_win_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (325, '2025-10-27 18:49:14', '2025-10-30 16:01:52', '/upgrade_win_upgrade_strategy/list', 'apiDesc.getWinUpgradeStrategyList', 'upgrade_win_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (326, '2025-10-27 18:49:20', '2025-10-30 16:01:53', '/upgrade_win_upgrade_strategy/delete', 'apiDesc.deleteWinUpgradeStrategy', 'upgrade_win_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (327, '2025-10-27 18:49:24', '2025-10-30 16:01:54', '/upgrade_win_upgrade_strategy/create', 'apiDesc.createWinUpgradeStrategy', 'upgrade_win_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (328, '2025-10-27 18:49:29', '2025-10-30 16:01:55', '/upgrade_win_upgrade_strategy/update', 'apiDesc.updateWinUpgradeStrategy', 'upgrade_win_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (329, '2025-10-27 18:49:46', '2025-10-30 16:01:56', '/upgrade_lnx', 'apiDesc.getLnxAppById', 'upgrade_lnx', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (330, '2025-10-27 18:49:51', '2025-10-30 16:01:56', '/upgrade_lnx/list', 'apiDesc.getLnxAppList', 'upgrade_lnx', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (331, '2025-10-27 18:49:56', '2025-10-30 16:01:57', '/upgrade_lnx/delete', 'apiDesc.deleteLnxApp', 'upgrade_lnx', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (332, '2025-10-27 18:50:00', '2025-10-30 16:01:58', '/upgrade_lnx/create', 'apiDesc.createLnxApp', 'upgrade_lnx', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (333, '2025-10-27 18:50:06', '2025-10-30 16:01:59', '/upgrade_lnx/update', 'apiDesc.updateLnxApp', 'upgrade_lnx', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (334, '2025-10-27 18:50:11', '2025-10-30 16:02:00', '/upgrade_lnx_version', 'apiDesc.getLnxVersionById', 'upgrade_lnx_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (335, '2025-10-27 18:50:15', '2025-10-30 16:02:01', '/upgrade_lnx_version/list', 'apiDesc.getLnxVersionList', 'upgrade_lnx_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (336, '2025-10-27 18:50:22', '2025-10-30 16:02:02', '/upgrade_lnx_version/delete', 'apiDesc.deleteLnxVersion', 'upgrade_lnx_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (337, '2025-10-27 18:50:29', '2025-10-30 16:02:03', '/upgrade_lnx_version/create', 'apiDesc.createLnxVersion', 'upgrade_lnx_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (338, '2025-10-27 18:50:36', '2025-10-30 16:02:04', '/upgrade_lnx_version/update', 'apiDesc.updateLnxVersion', 'upgrade_lnx_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (339, '2025-10-27 18:50:52', '2025-10-30 16:02:05', '/upgrade_lnx_upgrade_strategy', 'apiDesc.getLnxUpgradeStrategyById', 'upgrade_lnx_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (340, '2025-10-27 18:50:58', '2025-10-30 16:02:07', '/upgrade_lnx_upgrade_strategy/list', 'apiDesc.getLnxUpgradeStrategyList', 'upgrade_lnx_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (341, '2025-10-27 18:51:04', '2025-10-30 16:02:07', '/upgrade_lnx_upgrade_strategy/delete', 'apiDesc.deleteLnxUpgradeStrategy', 'upgrade_lnx_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (342, '2025-10-27 18:51:07', '2025-10-30 16:02:08', '/upgrade_lnx_upgrade_strategy/create', 'apiDesc.createLnxUpgradeStrategy', 'upgrade_lnx_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (343, '2025-10-27 18:51:12', '2025-10-30 16:02:09', '/upgrade_lnx_upgrade_strategy/update', 'apiDesc.updateLnxUpgradeStrategy', 'upgrade_lnx_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (344, '2025-10-27 18:51:17', '2025-10-30 16:02:10', '/upgrade_mac', 'apiDesc.getMacAppById', 'upgrade_mac', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (345, '2025-10-27 18:51:25', '2025-10-30 16:02:11', '/upgrade_mac/list', 'apiDesc.getMacAppList', 'upgrade_mac', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (346, '2025-10-27 18:51:30', '2025-10-30 16:02:12', '/upgrade_mac/delete', 'apiDesc.deleteMacApp', 'upgrade_mac', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (347, '2025-10-27 18:52:02', '2025-10-30 16:02:13', '/upgrade_mac/create', 'apiDesc.createMacApp', 'upgrade_mac', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (348, '2025-10-27 18:52:06', '2025-10-30 16:02:14', '/upgrade_mac/update', 'apiDesc.updateMacApp', 'upgrade_mac', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (349, '2025-10-27 18:52:10', '2025-10-30 16:02:15', '/upgrade_mac_version', 'apiDesc.getMacVersionById', 'upgrade_mac_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (350, '2025-10-27 18:52:16', '2025-10-30 16:02:16', '/upgrade_mac_version/list', 'apiDesc.getMacVersionList', 'upgrade_mac_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (351, '2025-10-27 18:52:22', '2025-10-30 16:02:17', '/upgrade_mac_version/delete', 'apiDesc.deleteMacVersion', 'upgrade_mac_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (352, '2025-10-27 18:52:29', '2025-10-30 16:02:18', '/upgrade_mac_version/create', 'apiDesc.createMacVersion', 'upgrade_mac_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (353, '2025-10-27 18:52:37', '2025-10-30 16:02:19', '/upgrade_mac_version/update', 'apiDesc.updateMacVersion', 'upgrade_mac_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (354, '2025-10-27 18:52:41', '2025-10-30 16:02:20', '/upgrade_mac_upgrade_strategy', 'apiDesc.getMacUpgradeStrategyById', 'upgrade_mac_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (355, '2025-10-27 18:52:46', '2025-10-30 16:02:21', '/upgrade_mac_upgrade_strategy/list', 'apiDesc.getMacUpgradeStrategyList', 'upgrade_mac_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (356, '2025-10-27 18:52:52', '2025-10-30 16:02:22', '/upgrade_mac_upgrade_strategy/delete', 'apiDesc.deleteMacUpgradeStrategy', 'upgrade_mac_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (357, '2025-10-27 18:52:58', '2025-10-30 16:02:23', '/upgrade_mac_upgrade_strategy/create', 'apiDesc.createMacUpgradeStrategy', 'upgrade_mac_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (358, '2025-10-27 18:53:03', '2025-10-30 16:02:25', '/upgrade_mac_upgrade_strategy/update', 'apiDesc.updateMacUpgradeStrategy', 'upgrade_mac_upgrade_strategy', 'Upgrade', 'POST', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_company_secret
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_secret`;
CREATE TABLE `sys_company_secret` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `access_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密钥id',
  `secret_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密钥key',
  `validity_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '有效期',
  `rule_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用权限',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='公司密钥表';

-- ----------------------------
-- Records of sys_company_secret
-- ----------------------------
BEGIN;
INSERT INTO `sys_company_secret` (`id`, `created_at`, `updated_at`, `company_id`, `access_key`, `secret_key`, `validity_datetime`, `rule_data`, `enable`, `is_del`, `description`) VALUES (1, '2025-02-24 07:58:27', '2026-01-25 09:48:14', 0, 'mui2W50H1j-OC4xD6PgQag', 'PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc', '2026-12-30 17:58:51', '', 1, 0, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_companys
-- ----------------------------
DROP TABLE IF EXISTS `sys_companys`;
CREATE TABLE `sys_companys` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公司名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='公司信息表';

-- ----------------------------
-- Records of sys_companys
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `menu_level` int unsigned NOT NULL COMMENT 'Menu level | 菜单层级',
  `menu_type` int unsigned NOT NULL COMMENT 'Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Index path | 菜单路由路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Index name | 菜单名称',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Redirect path | 跳转路径 （外链）',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT 'The path of vue file | 组件路径',
  `disabled` tinyint(1) DEFAULT '0' COMMENT 'Disable status | 是否停用',
  `service_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'Other' COMMENT 'Service Name | 服务名称',
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Permission symbol | 权限标识',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Menu name | 菜单显示标题',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Menu icon | 菜单图标',
  `hide_menu` tinyint(1) DEFAULT '0' COMMENT 'Hide menu | 是否隐藏菜单',
  `hide_breadcrumb` tinyint(1) DEFAULT '0' COMMENT 'Hide the breadcrumb | 隐藏面包屑',
  `ignore_keep_alive` tinyint(1) DEFAULT '0' COMMENT 'Do not keep alive the tab | 取消页面缓存',
  `hide_tab` tinyint(1) DEFAULT '0' COMMENT 'Hide the tab header | 隐藏页头',
  `frame_src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Show iframe | 内嵌 iframe',
  `carry_param` tinyint(1) DEFAULT '0' COMMENT 'The route carries parameters or not | 携带参数',
  `hide_children_in_menu` tinyint(1) DEFAULT '0' COMMENT 'Hide children menu or not | 隐藏所有子菜单',
  `affix` tinyint(1) DEFAULT '0' COMMENT 'Affix tab | Tab 固定',
  `dynamic_level` int unsigned DEFAULT '20' COMMENT 'The maximum number of pages the router can open | 能打开的子TAB数',
  `real_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT 'The real path of the route without dynamic part | 菜单路由不包含参数部分',
  `parent_id` bigint unsigned DEFAULT '100000' COMMENT 'Parent menu ID | 父菜单ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_name` (`name`),
  UNIQUE KEY `menu_path` (`path`)
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (1, '2025-02-10 09:33:13', '2025-09-22 14:33:07', 0, 1, 1, '/dashboard', 'Dashboard', '', '/dashboard/analytics/index', 0, 'Core', NULL, 'route.dashboard', 'ant-design:dashboard-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (2, '2025-02-10 09:33:13', '2025-09-22 14:56:10', 8, 1, 0, '/system', 'SystemManagement', '', 'LAYOUT', 0, 'Core', NULL, 'route.systemManagementTitle', 'ant-design:tool-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (3, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 1, 2, 1, '/menu', 'MenuManagement', '', '/sys/menu/index', 0, 'Core', NULL, 'route.menuManagementTitle', 'ant-design:bars-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (4, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 2, 2, 1, '/role', 'RoleManagement', '', '/sys/role/index', 0, 'Core', NULL, 'route.roleManagementTitle', 'ant-design:user-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (5, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 3, 2, 1, '/user', 'UserManagement', '', '/sys/user/index', 0, 'Core', NULL, 'route.userManagementTitle', 'ant-design:user-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (7, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 5, 2, 1, '/api', 'APIManagement', '', '/sys/api/index', 0, 'Core', NULL, 'route.apiManagementTitle', 'ant-design:api-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (9, '2025-02-10 09:33:13', '2025-09-22 14:56:17', 9, 1, 0, '/other', 'OtherPages', '', 'LAYOUT', 0, 'Core', NULL, 'route.otherPages', 'ant-design:question-circle-outlined', 1, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (10, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 3, 1, 1, '/profile', 'Profile', '', '/sys/profile/index', 0, 'Core', NULL, 'route.userProfileTitle', 'ant-design:profile-outlined', 1, 0, 0, 0, '', 0, 0, 0, 20, '', 9);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (12, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 7, 2, 1, '/token', 'TokenManagement', '', '/sys/token/index', 0, 'Core', NULL, 'route.tokenManagement', 'ant-design:lock-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (24, '2025-02-20 09:30:03', '2025-10-30 15:20:10', 400, 1, 1, '/upgrade_url_dir', 'UrlManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.urlManagement', 'ant-design:share-alt-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (26, '2025-02-20 09:52:10', '2025-10-30 15:21:04', 1, 2, 1, '/upgrade/url', 'UrlAppManagement', '', '/upgrade/url/index', 0, 'Upgrade', '', 'route.urlAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (27, '2025-09-11 11:28:22', '2025-10-30 15:21:09', 2, 2, 1, '/upgrade/url_version', 'UrlVersionManagement', '', '/upgrade/url_version/index', 0, 'Upgrade', NULL, 'route.urlVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (28, '2025-09-11 11:28:25', '2025-10-30 15:21:16', 3, 2, 1, '/upgrade/url_upgrade_strategy', 'UrlUpgradeStrategyManagement', '', '/upgrade/url_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.urlUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (29, '2025-02-22 16:00:09', '2025-10-30 15:21:21', 202, 1, 1, '/upgrade/dev_model', 'DevModelManagement', '', '/upgrade/dev_model/index', 0, 'Upgrade', '', 'route.devModelManagement', 'ant-design:laptop-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (32, '2025-02-25 03:11:59', '2025-10-30 15:21:26', 500, 1, 1, '/upgrade_file_dir', 'UpgradeFileManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.upgradeFileManagement', 'ant-design:file-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (35, '2025-02-25 03:17:26', '2025-10-30 15:21:37', 100, 1, 1, '/company_secret', 'CompanySecretManagement', '', '/sys/company_secret/index', 0, 'Core', '', 'route.companySecretManagement', 'ant-design:key-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (46, '2025-03-10 03:35:03', '2026-01-20 14:48:09', 3, 1, 1, '/file_dir', 'FileManagementDirectory', '', 'LAYOUT', 0, 'Fms', NULL, 'route.fileManagement', 'ant-design:folder-open-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (50, '2025-03-10 03:35:04', '2026-01-20 14:47:49', 4, 2, 1, '/cloud_file', 'CloudFileManagement', '', '/sys/cloudFile/index', 0, 'Fms', NULL, 'route.cloudFileManagement', 'ant-design:folder-open-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (52, '2025-02-21 06:19:59', '2025-10-30 15:22:22', 1, 2, 1, '/upgrade/file', 'FileAppManagement', '', '/upgrade/file/index', 0, 'Upgrade', NULL, 'route.fileAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (53, '2025-09-11 11:28:17', '2025-10-30 15:22:27', 2, 2, 1, '/upgrade/file_version', 'FileVersionManagement', '', '/upgrade/file_version/index', 0, 'Upgrade', NULL, 'route.fileVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (54, '2025-09-11 11:28:21', '2025-10-30 15:22:33', 3, 2, 1, '/upgrade/file_upgrade_strategy', 'FileUpgradeStrategyManagement', '', '/upgrade/file_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.fileUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (55, '2025-02-22 16:00:09', '2025-10-30 15:22:40', 200, 1, 1, '/upgrade/dev', 'DevManagement', '', '/upgrade/dev/index', 0, 'Upgrade', NULL, 'route.devManagement', 'ant-design:laptop-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (56, '2025-02-22 16:00:09', '2025-10-30 15:22:46', 201, 1, 1, '/upgrade/dev_group', 'DevGroupManagement', '', '/upgrade/dev_group/index', 0, 'Upgrade', NULL, 'route.devGroupManagement', 'ant-design:solution-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (57, '2025-02-21 06:19:59', '2025-10-30 15:22:50', 1, 2, 1, '/upgrade/tauri', 'TauriAppManagement', '', '/upgrade/tauri/index', 0, 'Upgrade', NULL, 'route.tauriAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (58, '2025-09-11 11:28:11', '2025-10-30 15:22:55', 2, 2, 1, '/upgrade/tauri_version', 'TauriVersionManagement', '', '/upgrade/tauri_version/index', 0, 'Upgrade', NULL, 'route.tauriVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (59, '2025-09-11 11:28:13', '2025-10-30 15:23:02', 3, 2, 1, '/upgrade/tauri_upgrade_strategy', 'TauriUpgradeStrategyManagement', '', '/upgrade/tauri_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.tauriUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (60, '2025-02-25 03:11:59', '2025-10-30 15:23:09', 700, 1, 1, '/upgrade_tauri_dir', 'UpgradeTauriManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.upgradeTauriManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (61, '2025-02-25 03:11:59', '2025-10-30 15:23:15', 600, 1, 1, '/upgrade_configuration_dir', 'UpgradeConfigurationManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.upgradeConfigurationManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (62, '2025-02-25 03:11:59', '2025-10-30 15:23:20', 1, 2, 1, '/upgrade/configuration', 'ConfigurationAppManagement', '', '/upgrade/configuration/index', 0, 'Upgrade', NULL, 'route.configurationAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (63, '2025-09-11 11:28:06', '2025-10-30 15:23:25', 2, 2, 1, '/upgrade/configuration_version', 'ConfigurationVersionManagement', '', '/upgrade/configuration_version/index', 0, 'Upgrade', NULL, 'route.configurationVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (65, '2025-09-11 11:28:08', '2025-10-30 15:23:31', 3, 2, 1, '/upgrade/configuration_upgrade_strategy', 'ConfigurationUpgradeStrategyManagement', '', '/upgrade/configuration_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.configurationUpgradeStrategyManagement', 'ant-design:home-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (66, '2025-02-25 03:11:59', '2025-10-30 15:23:37', 900, 1, 1, '/upgrade_apk_dir', 'UpgradeApkManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.upgradeApkManagement', 'ant-design:android-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (67, '2025-02-25 03:11:59', '2025-10-30 15:23:43', 1, 2, 1, '/upgrade/apk', 'ApkAppManagement', '', '/upgrade/apk/index', 0, 'Upgrade', NULL, 'route.apkAppManagement', 'ant-design:android-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (68, '2025-09-11 11:27:58', '2025-10-30 15:23:49', 2, 2, 1, '/upgrade/apk_version', 'ApkVersionManagement', '', '/upgrade/apk_version/index', 0, 'Upgrade', NULL, 'route.apkVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (69, '2025-09-11 11:28:00', '2025-10-30 15:23:56', 3, 2, 1, '/upgrade/apk_upgrade_strategy', 'ApkUpgradeStrategyManagement', '', '/upgrade/apk_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.apkUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (70, '2025-02-25 03:11:59', '2025-10-30 15:24:00', 800, 1, 1, '/upgrade_electron_dir', 'UpgradeElectronManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'route.upgradeElectronManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (73, '2025-02-25 03:11:59', '2025-10-30 15:24:05', 1, 2, 1, '/upgrade/electron', 'ElectronAppManagement', '', '/upgrade/electron/index', 0, 'Upgrade', NULL, 'route.electronAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (74, '2025-09-11 11:27:50', '2025-10-30 15:24:11', 2, 2, 1, '/upgrade/electron_version', 'ElectronVersionManagement', '', '/upgrade/electron_version/index', 0, 'Upgrade', NULL, 'route.electronVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (75, '2025-09-11 11:27:51', '2025-10-30 15:24:16', 3, 2, 1, '/upgrade/electron_upgrade_strategy', 'ElectronUpgradeStrategyManagement', '', '/upgrade/electron_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.electronUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (82, '2025-02-25 03:11:59', '2025-10-30 15:24:40', 910, 1, 1, '/upgrade_win_dir', 'UpgradeWinManagement', '', 'LAYOUT', 0, 'Upgrade', NULL, 'route.upgradeWinManagement', 'ant-design:windows-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (88, '2025-10-27 18:31:23', '2025-10-30 15:24:45', 1, 2, 1, '/upgrade/win', 'WinAppManagement', '', '/upgrade/win/index', 0, 'Upgrade', NULL, 'route.winAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 82);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (89, '2025-10-27 18:31:36', '2025-10-30 15:24:50', 2, 2, 1, '/upgrade/win_version', 'WinVersionManagement', '', '/upgrade/win_version/index', 0, 'Upgrade', NULL, 'route.winVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 82);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (90, '2025-10-27 18:31:50', '2025-10-30 15:24:56', 3, 2, 1, '/upgrade/win_upgrade_strategy', 'WinUpgradeStrategyManagement', '', '/upgrade/win_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.winUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 82);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (91, '2025-10-27 18:32:02', '2025-10-30 15:25:01', 920, 1, 1, '/upgrade_lnx_dir', 'UpgradeLnxManagement', '', 'LAYOUT', 0, 'Upgrade', NULL, 'route.upgradeLnxManagement', 'ant-design:linux-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (92, '2025-10-27 18:32:14', '2025-10-30 15:25:06', 1, 2, 1, '/upgrade/lnx', 'LnxAppManagement', '', '/upgrade/lnx/index', 0, 'Upgrade', '', 'route.lnxAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 91);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (93, '2025-10-27 18:32:26', '2025-10-30 15:25:11', 2, 2, 1, '/upgrade/lnx_version', 'LnxVersionManagement', '', '/upgrade/lnx_version/index', 0, 'Upgrade', NULL, 'route.lnxVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 91);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (94, '2025-10-27 18:32:36', '2025-10-30 15:25:17', 3, 2, 1, '/upgrade/lnx_upgrade_strategy', 'LnxUpgradeStrategyManagement', '', '/upgrade/lnx_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.lnxUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 91);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (95, '2025-10-27 18:32:50', '2025-10-30 15:25:20', 930, 1, 1, '/upgrade_mac_dir', 'UpgradeMacosManagement', '', 'LAYOUT', 0, 'Upgrade', NULL, 'route.upgradeMacosManagement', 'ant-design:apple-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (96, '2025-10-27 18:33:41', '2025-10-30 15:25:24', 1, 2, 1, '/upgrade/mac', 'MacAppManagement', '', '/upgrade/mac/index', 0, 'Upgrade', NULL, 'route.macAppManagement', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 95);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (97, '2025-10-27 18:33:54', '2025-10-30 15:25:30', 2, 2, 1, '/upgrade/mac_version', 'MacVersionManagement', '', '/upgrade/mac_version/index', 0, 'Upgrade', NULL, 'route.macVersionManagement', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 95);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (98, '2025-10-27 18:34:15', '2025-10-30 15:25:32', 3, 2, 1, '/upgrade/mac_upgrade_strategy', 'MacUpgradeStrategyManagement', '', '/upgrade/mac_upgrade_strategy/index', 0, 'Upgrade', NULL, 'route.macUpgradeStrategyManagement', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 95);
COMMIT;

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Role name | 角色名',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Role code for permission control in front end | 角色码，用于前端权限控制',
  `default_router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'dashboard' COMMENT 'Default menu : dashboard | 默认登录页面',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'Remark | 备注',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT 'Order number | 排序编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
BEGIN;
INSERT INTO `sys_roles` (`id`, `created_at`, `updated_at`, `status`, `name`, `code`, `default_router`, `remark`, `sort`) VALUES (1, '2025-09-09 10:23:13', '2025-10-27 17:51:04', 1, 'role.admin', '001', 'dashboard', '超级管理员', 1);
INSERT INTO `sys_roles` (`id`, `created_at`, `updated_at`, `status`, `name`, `code`, `default_router`, `remark`, `sort`) VALUES (2, '2025-09-09 17:39:56', '2026-01-23 15:22:58', 1, 'role.stuff', '002', 'dashboard', '普通用户', 2);
COMMIT;

-- ----------------------------
-- Table structure for sys_tokens
-- ----------------------------
DROP TABLE IF EXISTS `sys_tokens`;
CREATE TABLE `sys_tokens` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `uuid` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT ' User''s UUID | 用户的UUID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT 'Username | 用户名',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Token string | Token 字符串',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）',
  `expired_at` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT ' Expire time | 过期时间',
  PRIMARY KEY (`id`),
  KEY `token_expired_at` (`expired_at`),
  KEY `token_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_tokens
-- ----------------------------
BEGIN;
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf495-99a7-7625-88de-12b18c006777', '2026-01-25 17:56:44', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NTk0MjAzLCJpYXQiOjE3NjkzMzUwMDMsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.1SPVAC65Gp1rri6-pF8-RDaf0Bot0aX1mpUIvE4AQyo', 'core_user', '2026-01-28 17:56:44');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf497-6d00-7625-9eff-59cc351b8b93', '2026-01-25 17:58:43', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NTk0MzIzLCJpYXQiOjE3NjkzMzUxMjMsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.7M0P_pEIwopPtWnLueUGaWp0U054b1_EfLDsZvrY0EM', 'core_user', '2026-01-28 17:58:43');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf499-89d7-7625-a39f-e0c16690c8fc', '2026-01-25 18:01:02', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NTk0NDYxLCJpYXQiOjE3NjkzMzUyNjEsInJvbGVJZCI6IjAwMSwwMDIiLCJ1c2VySWQiOiIwMTk0ZWYzNC1hNmQxLTcyOGUtYTcwZC0zNjJmNTAzNWFhYjcifQ.XGllAJmHrJjE9vQz66pLoLXLLoYwLbUmSNMj-q3sfMA', 'core_user', '2026-01-28 18:01:02');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf49c-67ee-72b1-90d9-566ff0cf2eb6', '2026-01-25 18:04:10', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NTk0NjQ5LCJpYXQiOjE3NjkzMzU0NDksInJvbGVJZCI6IjAwMSwwMDIiLCJ1c2VySWQiOiIwMTk0ZWYzNC1hNmQxLTcyOGUtYTcwZC0zNjJmNTAzNWFhYjcifQ.dbRbKy6beU2XPPX1Om_8YX843jaF9gIc6yVU-QFlzK0', 'core_user', '2026-01-28 18:04:10');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf503-3373-7b63-840b-26973770d70a', '2026-01-25 19:56:26', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NjAxMzg2LCJpYXQiOjE3NjkzNDIxODYsInJvbGVJZCI6IjAwMSwwMDIiLCJ1c2VySWQiOiIwMTk0ZWYzNC1hNmQxLTcyOGUtYTcwZC0zNjJmNTAzNWFhYjcifQ.TZ_Pz9d4n3nGUEG0I5KIfte3e1HCCN5b6wvFutBKRNQ', 'core_user', '2026-01-28 19:56:26');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf51d-0deb-7b63-b4fd-88356a4bf051', '2026-01-25 20:24:41', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NjAzMDgwLCJpYXQiOjE3NjkzNDM4ODAsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.U_iDG8gdBalFDGJkoXfWyCeyrZAZ_7G1neDIzcC10ms', 'core_user', '2026-01-28 20:24:41');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf51d-abce-7b63-b03d-69640e432b57', '2026-01-25 20:25:21', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NjAzMTIxLCJpYXQiOjE3NjkzNDM5MjEsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.XlsOSk6wqFdhAT6_dblYYS_kylthD_aSZkEDWCvqbr0', 'core_user', '2026-01-28 20:25:21');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf51e-6d7b-7b63-bed0-a75144a97d60', '2026-01-25 20:26:11', '2026-01-25 20:33:07', 2, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NjAzMTcwLCJpYXQiOjE3NjkzNDM5NzAsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.P1lOk0BvQwj85JU9YLuhiqa5uO3qgvwR3qT4MQCIunM', 'core_user', '2026-01-28 20:26:11');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019bf525-b9e1-7afe-9aff-009893c2b85e', '2026-01-25 20:34:09', '2026-01-25 20:34:09', 1, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiIwIiwiZXhwIjoxNzY5NjAzNjQ4LCJpYXQiOjE3NjkzNDQ0NDgsInJvbGVJZCI6IjAwMSIsInVzZXJJZCI6IjAxOTRlZjM0LWE2ZDEtNzI4ZS1hNzBkLTM2MmY1MDM1YWFiNyJ9.U5cDeTDIIj_UxO_dXiyUZIX9JkE6oh9rJ2eQqnf-2xg', 'core_user', '2026-01-28 20:34:09');
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Delete Time | 删除日期',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'User''s login name | 登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Password | 密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Nickname | 昵称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The description of user | 用户的描述信息',
  `home_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '/dashboard' COMMENT 'The home page that the user enters after logging in | 用户登陆后进入的首页',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Mobile number | 手机号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Email | 邮箱号',
  `avatar` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Avatar | 头像路径',
  `company_id` bigint unsigned DEFAULT '0' COMMENT 'Company ID | 公司ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `user_username_email` (`username`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `status`, `deleted_at`, `username`, `password`, `nickname`, `description`, `home_path`, `mobile`, `email`, `avatar`, `company_id`) VALUES ('0194ef34-a6d1-728e-a70d-362f5035aab7', '2025-08-15 11:26:56', '2026-01-25 12:34:00', 1, NULL, 'admin', '$2a$10$1JMM5u4ch3ml9dNafqEZEu.vnOBy74lMZF22DaL10NmpN1V22TBc2', 'admin', '', '/dashboard', '', 'upgradelink@gmail.com', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='安卓应用';

-- ----------------------------
-- Records of upgrade_apk
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk_patch
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk_patch`;
CREATE TABLE `upgrade_apk_patch` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `apk_id` bigint NOT NULL DEFAULT '0' COMMENT '安卓应用ID',
  `high_apk_version_id` bigint NOT NULL DEFAULT '0' COMMENT '外键：apk_version.id',
  `low_apk_version_id` bigint NOT NULL DEFAULT '0' COMMENT '外键：apk_version.id',
  `patch_algo` int NOT NULL DEFAULT '0' COMMENT '差分算法 0:默认值无; 1 HDiffPatch;2 bsdiff;',
  `status` int NOT NULL DEFAULT '0' COMMENT '处理状态：0:尚未进行差分处理; 1:正在处理差分; 2:差分过程错误; 3:差分过程超时; 4:差分包有问题; 5:差分处理成功; 6:差分包大于新版本全量包; 7:上传文件中; 8:上传文件失败; 9:处理完成 ',
  `cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='差分信息表；记录APK的差分基本信息';

-- ----------------------------
-- Records of upgrade_apk_patch
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk_upgrade_strategy`;
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
  `gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '灰度策略id数据',
  `is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
  `flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '频控策略id数据',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='安卓应用 升级任务';

-- ----------------------------
-- Records of upgrade_apk_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='安卓应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_apk_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_apk_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='安卓应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_apk_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_apk_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_apk_version`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='安卓应用 版本库';

-- ----------------------------
-- Records of upgrade_apk_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_download_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_download_report_log`;
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
  PRIMARY KEY (`id`),
  KEY `idx_company_timestamp` (`company_id`,`timestamp`),
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='应用初次下载事件日志';

-- ----------------------------
-- Records of upgrade_app_download_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_start_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_start_report_log`;
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
  PRIMARY KEY (`id`),
  KEY `idx_company_timestamp` (`company_id`,`timestamp`),
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='应用启动事件日志';

-- ----------------------------
-- Records of upgrade_app_start_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_upgrade_download_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_upgrade_download_report_log`;
CREATE TABLE `upgrade_app_upgrade_download_report_log` (
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
  `download_version_code` bigint NOT NULL DEFAULT '0' COMMENT '下载版本号',
  `download_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `code` bigint NOT NULL DEFAULT '0' COMMENT '事件-状态码',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_company_timestamp` (`company_id`,`timestamp`),
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='应用升级下载事件日志';

-- ----------------------------
-- Records of upgrade_app_upgrade_download_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_upgrade_get_strategy_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_upgrade_get_strategy_report_log`;
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
  PRIMARY KEY (`id`),
  KEY `idx_company_timestamp` (`company_id`,`timestamp`),
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='应用升级获取升级策略事件日志';

-- ----------------------------
-- Records of upgrade_app_upgrade_get_strategy_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_upgrade_upgrade_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_upgrade_upgrade_report_log`;
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
  PRIMARY KEY (`id`),
  KEY `idx_company_timestamp` (`company_id`,`timestamp`),
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='应用升级升级事件日志';

-- ----------------------------
-- Records of upgrade_app_upgrade_upgrade_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_configuration
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_configuration`;
CREATE TABLE `upgrade_configuration` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '配置唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '配置名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='配置';

-- ----------------------------
-- Records of upgrade_configuration
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_configuration_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_configuration_upgrade_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='配置 升级任务';

-- ----------------------------
-- Records of upgrade_configuration_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_configuration_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_configuration_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='配置 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_configuration_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_configuration_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_configuration_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_configuration_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='配置 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_configuration_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_configuration_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_configuration_version`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='配置 版本库';

-- ----------------------------
-- Records of upgrade_configuration_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_dev_group
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_dev_group`;
CREATE TABLE `upgrade_dev_group` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备分组名称',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='设备分组';

-- ----------------------------
-- Records of upgrade_dev_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_dev_group_relation
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_dev_group_relation`;
CREATE TABLE `upgrade_dev_group_relation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `dev_id` bigint unsigned NOT NULL COMMENT '设备id',
  `dev_group_id` bigint unsigned NOT NULL COMMENT '设备分组 id',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='设备分组关系';

-- ----------------------------
-- Records of upgrade_dev_group_relation
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_dev_model
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_dev_model`;
CREATE TABLE `upgrade_dev_model` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL COMMENT '公司ID: 所属公司id',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型名称',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='设备机型表';

-- ----------------------------
-- Records of upgrade_dev_model
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_devs
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_devs`;
CREATE TABLE `upgrade_devs` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='设备管理';

-- ----------------------------
-- Records of upgrade_devs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_electron
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_electron`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='electron应用';

-- ----------------------------
-- Records of upgrade_electron
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_electron_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_electron_upgrade_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='electron应用 升级任务';

-- ----------------------------
-- Records of upgrade_electron_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_electron_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_electron_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='electron应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_electron_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_electron_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_electron_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_electron_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='electron应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_electron_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_electron_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_electron_version`;
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
  PRIMARY KEY (`id`),
  KEY `idx_electron_active_version` (`electron_id`,`is_del`,`version_code`,`create_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='electron应用 版本库';

-- ----------------------------
-- Records of upgrade_electron_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file`;
CREATE TABLE `upgrade_file` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件应用唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件应用名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='文件应用';

-- ----------------------------
-- Records of upgrade_file
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file_github
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file_github`;
CREATE TABLE `upgrade_file_github` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件id',
  `url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'github 文件地址',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_upgrade_file_github_file_id` (`file_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='github文件下载地址';

-- ----------------------------
-- Records of upgrade_file_github
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file_upgrade_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='文件应用 升级任务';

-- ----------------------------
-- Records of upgrade_file_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='文件应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_file_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_file_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='文件应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_file_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_file_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_file_version`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='文件应用 版本库';

-- ----------------------------
-- Records of upgrade_file_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_lnx
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_lnx`;
CREATE TABLE `upgrade_lnx` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'linux应用唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'linux应用名称',
  `package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'linux应用包名',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='linux应用';

-- ----------------------------
-- Records of upgrade_lnx
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_lnx_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_lnx_upgrade_strategy`;
CREATE TABLE `upgrade_lnx_upgrade_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
  `lnx_id` bigint NOT NULL DEFAULT '0' COMMENT 'linux应用ID',
  `lnx_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'lnx_version_id; 外键lnx_version.id',
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='linux应用 升级任务';

-- ----------------------------
-- Records of upgrade_lnx_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_lnx_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_lnx_upgrade_strategy_flow_limit_strategy`;
CREATE TABLE `upgrade_lnx_upgrade_strategy_flow_limit_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '9' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
  `end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
  `dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='linux应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_lnx_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_lnx_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_lnx_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_lnx_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='linux应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_lnx_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_lnx_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_lnx_version`;
CREATE TABLE `upgrade_lnx_version` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `lnx_id` bigint NOT NULL DEFAULT '0' COMMENT 'lnx应用ID',
  `cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
  `version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x64、arm64',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='linux应用 版本库';

-- ----------------------------
-- Records of upgrade_lnx_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_mac
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_mac`;
CREATE TABLE `upgrade_mac` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'mac应用唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'mac应用名称',
  `package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'mac应用包名',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='mac应用';

-- ----------------------------
-- Records of upgrade_mac
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_mac_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_mac_upgrade_strategy`;
CREATE TABLE `upgrade_mac_upgrade_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
  `mac_id` bigint NOT NULL DEFAULT '0' COMMENT 'mac应用ID',
  `mac_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'mac_version_id; 外键mac_version.id',
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='mac应用 升级任务';

-- ----------------------------
-- Records of upgrade_mac_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_mac_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_mac_upgrade_strategy_flow_limit_strategy`;
CREATE TABLE `upgrade_mac_upgrade_strategy_flow_limit_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '9' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
  `end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
  `dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='mac应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_mac_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_mac_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_mac_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_mac_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='mac应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_mac_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_mac_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_mac_version`;
CREATE TABLE `upgrade_mac_version` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `mac_id` bigint NOT NULL DEFAULT '0' COMMENT 'mac应用ID',
  `cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
  `version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x64、arm64',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='mac应用 版本库';

-- ----------------------------
-- Records of upgrade_mac_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_tauri
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_tauri`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='tauri应用';

-- ----------------------------
-- Records of upgrade_tauri
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_tauri_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_tauri_upgrade_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=779 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='tauri应用 升级任务';

-- ----------------------------
-- Records of upgrade_tauri_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_tauri_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_tauri_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='tauri应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_tauri_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_tauri_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_tauri_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_tauri_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='tauri应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_tauri_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_tauri_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_tauri_version`;
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
  PRIMARY KEY (`id`),
  KEY `idx_tauri_active_version` (`tauri_id`,`is_del`,`version_code`,`create_at`)
) ENGINE=InnoDB AUTO_INCREMENT=822 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='tauri应用 版本库';

-- ----------------------------
-- Records of upgrade_tauri_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_traffic_packet
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_traffic_packet`;
CREATE TABLE `upgrade_traffic_packet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '流量包ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '流量包名称',
  `key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '流量包唯一标识',
  `size` bigint NOT NULL DEFAULT '0' COMMENT '流量大小(单位:字节)',
  `price` int NOT NULL DEFAULT '1' COMMENT '价格',
  `valid_days` int NOT NULL DEFAULT '1' COMMENT '有效期(天)',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态: 1=有效, 2=已兑换',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='流量包表';

-- ----------------------------
-- Records of upgrade_traffic_packet
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_url
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_url`;
CREATE TABLE `upgrade_url` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'url唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'url应用名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='url应用';

-- ----------------------------
-- Records of upgrade_url
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_url_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_url_upgrade_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='url应用 升级任务';

-- ----------------------------
-- Records of upgrade_url_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_url_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_url_upgrade_strategy_flow_limit_strategy`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='url应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_url_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_url_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_url_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_url_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '升级任务结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='url应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_url_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_url_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_url_version`;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='url应用 版本库';

-- ----------------------------
-- Records of upgrade_url_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_win
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_win`;
CREATE TABLE `upgrade_win` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'win应用唯一标识',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'win应用名称',
  `package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'win应用包名',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='win应用';

-- ----------------------------
-- Records of upgrade_win
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_win_upgrade_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_win_upgrade_strategy`;
CREATE TABLE `upgrade_win_upgrade_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名称',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务描述信息',
  `win_id` bigint NOT NULL DEFAULT '0' COMMENT 'win应用ID',
  `win_version_id` bigint NOT NULL DEFAULT '0' COMMENT 'win_version_id; 外键win_version.id',
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='win应用 升级任务';

-- ----------------------------
-- Records of upgrade_win_upgrade_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_win_upgrade_strategy_flow_limit_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_win_upgrade_strategy_flow_limit_strategy`;
CREATE TABLE `upgrade_win_upgrade_strategy_flow_limit_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '9' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开始时间段: 时分秒',
  `end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '结束时间段: 时分秒',
  `dimension` int NOT NULL DEFAULT '1' COMMENT '流控维度；流控维度：1：秒；2：分；3：时；4：天',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '频控限制；在流控维度上的次数',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='win应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_win_upgrade_strategy_flow_limit_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_win_upgrade_strategy_gray_strategy
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_win_upgrade_strategy_gray_strategy`;
CREATE TABLE `upgrade_win_upgrade_strategy_gray_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `enable` int NOT NULL DEFAULT '0' COMMENT '是否生效；可通过此控制策略是否生效0：失效；1：生效',
  `begin_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_datetime` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `limit` bigint NOT NULL DEFAULT '10' COMMENT '数量限制',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='win应用 升级任务灰度策略表；';

-- ----------------------------
-- Records of upgrade_win_upgrade_strategy_gray_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_win_version
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_win_version`;
CREATE TABLE `upgrade_win_version` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `win_id` bigint NOT NULL DEFAULT '0' COMMENT 'win应用ID',
  `cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `version_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本名',
  `version_code` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `arch` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '机器架构:x64、arm64',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述信息',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='win应用 版本库';

-- ----------------------------
-- Records of upgrade_win_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0194ef34-a6d1-728e-a70d-362f5035aab7', 1);
COMMIT;

-- ----------------------------
-- View structure for daily_app_traffic_view
-- ----------------------------
DROP VIEW IF EXISTS `daily_app_traffic_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `daily_app_traffic_view` AS select `udrl`.`app_key` AS `app_key`,`udrl`.`company_id` AS `company_id`,cast(`udrl`.`timestamp` as date) AS `date`,sum(`fcf`.`size`) AS `daily_traffic_bytes`,round((sum(`fcf`.`size`) / (1024 * 1024)),2) AS `daily_traffic_mb` from (`upgrade_app_download_report_log` `udrl` join `fms_cloud_files` `fcf` on((`udrl`.`download_cloud_file_id` = `fcf`.`id`))) where ((`udrl`.`timestamp` >= (curdate() - interval 7 day)) and (`udrl`.`timestamp` < (curdate() + interval 1 day))) group by `udrl`.`company_id`,`udrl`.`app_key`,cast(`udrl`.`timestamp` as date) order by `date` desc,`daily_traffic_bytes` desc;

SET FOREIGN_KEY_CHECKS = 1;
