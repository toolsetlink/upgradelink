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

 Date: 26/09/2025 15:49:10
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
) ENGINE=InnoDB AUTO_INCREMENT=336688 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of casbin_rules
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336260, 'p', '003', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336261, 'p', '003', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336262, 'p', '003', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336263, 'p', '003', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336264, 'p', '003', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336265, 'p', '003', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336266, 'p', '003', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336267, 'p', '003', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336268, 'p', '003', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336269, 'p', '003', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336270, 'p', '003', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336271, 'p', '003', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336272, 'p', '003', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336273, 'p', '001', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336274, 'p', '001', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336275, 'p', '001', '/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336276, 'p', '001', '/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336277, 'p', '001', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336278, 'p', '001', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336279, 'p', '001', '/user/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336280, 'p', '001', '/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336281, 'p', '001', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336282, 'p', '001', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336283, 'p', '001', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336284, 'p', '001', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336285, 'p', '001', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336286, 'p', '001', '/user/refresh_token', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336287, 'p', '001', '/user/access_token', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336288, 'p', '001', '/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336289, 'p', '001', '/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336290, 'p', '001', '/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336291, 'p', '001', '/role/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336292, 'p', '001', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336293, 'p', '001', '/menu/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336294, 'p', '001', '/menu/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336295, 'p', '001', '/menu/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336296, 'p', '001', '/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336297, 'p', '001', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336298, 'p', '001', '/menu_param/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336299, 'p', '001', '/menu_param/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336300, 'p', '001', '/menu_param/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336301, 'p', '001', '/menu_param/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336302, 'p', '001', '/menu_param', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336303, 'p', '001', '/menu', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336304, 'p', '001', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336305, 'p', '001', '/authority/api/create_or_update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336306, 'p', '001', '/authority/api/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336307, 'p', '001', '/authority/menu/create_or_update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336308, 'p', '001', '/authority/menu/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336309, 'p', '001', '/api/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336310, 'p', '001', '/api/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336311, 'p', '001', '/api/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336312, 'p', '001', '/api/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336313, 'p', '001', '/api', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336314, 'p', '001', '/dictionary', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336315, 'p', '001', '/dictionary/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336316, 'p', '001', '/dictionary/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336317, 'p', '001', '/dictionary/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336318, 'p', '001', '/dictionary_detail/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336319, 'p', '001', '/dictionary_detail', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336320, 'p', '001', '/dictionary_detail/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336321, 'p', '001', '/dictionary_detail/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336322, 'p', '001', '/dictionary_detail/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336323, 'p', '001', '/dictionary/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336324, 'p', '001', '/dict/:name', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336325, 'p', '001', '/oauth_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336326, 'p', '001', '/oauth_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336327, 'p', '001', '/oauth_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336328, 'p', '001', '/oauth_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336329, 'p', '001', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336330, 'p', '001', '/oauth_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336331, 'p', '001', '/token/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336332, 'p', '001', '/token/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336333, 'p', '001', '/token/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336334, 'p', '001', '/token/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336335, 'p', '001', '/token/logout', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336336, 'p', '001', '/token', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336337, 'p', '001', '/department/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336338, 'p', '001', '/department/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336339, 'p', '001', '/department/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336340, 'p', '001', '/department/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336341, 'p', '001', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336342, 'p', '001', '/position/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336343, 'p', '001', '/position/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336344, 'p', '001', '/position/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336345, 'p', '001', '/position/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336346, 'p', '001', '/position', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336347, 'p', '001', '/task/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336348, 'p', '001', '/task/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336349, 'p', '001', '/task/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336350, 'p', '001', '/task/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336351, 'p', '001', '/task', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336352, 'p', '001', '/task_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336353, 'p', '001', '/task_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336354, 'p', '001', '/task_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336355, 'p', '001', '/task_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336356, 'p', '001', '/task_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336357, 'p', '001', '/configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336358, 'p', '001', '/configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336359, 'p', '001', '/configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336360, 'p', '001', '/configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336361, 'p', '001', '/configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336362, 'p', '001', '/email_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336363, 'p', '001', '/email_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336364, 'p', '001', '/email_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336365, 'p', '001', '/email_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336366, 'p', '001', '/email_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336367, 'p', '001', '/email_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336368, 'p', '001', '/email_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336369, 'p', '001', '/email_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336370, 'p', '001', '/email_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336371, 'p', '001', '/email_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336372, 'p', '001', '/sms_log/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336373, 'p', '001', '/sms_log/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336374, 'p', '001', '/sms_log/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336375, 'p', '001', '/sms_log/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336376, 'p', '001', '/sms_log', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336377, 'p', '001', '/sms_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336378, 'p', '001', '/sms_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336379, 'p', '001', '/sms_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336380, 'p', '001', '/sms_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336381, 'p', '001', '/sms_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336382, 'p', '001', '/sms/send', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336383, 'p', '001', '/email/send', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336384, 'p', '001', '/member/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336385, 'p', '001', '/member/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336386, 'p', '001', '/member/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336387, 'p', '001', '/member/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336388, 'p', '001', '/member', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336389, 'p', '001', '/member_rank/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336390, 'p', '001', '/member_rank/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336391, 'p', '001', '/member_rank/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336392, 'p', '001', '/member_rank/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336393, 'p', '001', '/member_rank', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336394, 'p', '001', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336395, 'p', '001', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336396, 'p', '001', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336397, 'p', '001', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336398, 'p', '001', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336399, 'p', '001', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336400, 'p', '001', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336401, 'p', '001', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336402, 'p', '001', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336403, 'p', '001', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336404, 'p', '001', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336405, 'p', '001', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336406, 'p', '001', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336407, 'p', '001', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336408, 'p', '001', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336409, 'p', '001', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336410, 'p', '001', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336411, 'p', '001', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336412, 'p', '001', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336413, 'p', '001', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336414, 'p', '001', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336415, 'p', '001', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336416, 'p', '001', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336417, 'p', '001', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336418, 'p', '001', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336419, 'p', '001', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336420, 'p', '001', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336421, 'p', '001', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336422, 'p', '001', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336423, 'p', '001', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336424, 'p', '001', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336425, 'p', '001', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336426, 'p', '001', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336427, 'p', '001', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336428, 'p', '001', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336429, 'p', '001', '/upgrade_apple', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336430, 'p', '001', '/upgrade_apple/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336431, 'p', '001', '/upgrade_apple/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336432, 'p', '001', '/upgrade_apple/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336433, 'p', '001', '/upgrade_apple/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336434, 'p', '001', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336435, 'p', '001', '/upload', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336436, 'p', '001', '/file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336437, 'p', '001', '/file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336438, 'p', '001', '/file/status', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336439, 'p', '001', '/file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336440, 'p', '001', '/file/download/:id', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336441, 'p', '001', '/file_tag/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336442, 'p', '001', '/file_tag/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336443, 'p', '001', '/file_tag/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336444, 'p', '001', '/file_tag/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336445, 'p', '001', '/file_tag', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336446, 'p', '001', '/storage_provider/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336447, 'p', '001', '/storage_provider/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336448, 'p', '001', '/storage_provider/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336449, 'p', '001', '/storage_provider/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336450, 'p', '001', '/storage_provider', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336451, 'p', '001', '/cloud_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336452, 'p', '001', '/cloud_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336453, 'p', '001', '/cloud_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336454, 'p', '001', '/cloud_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336455, 'p', '001', '/cloud_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336456, 'p', '001', '/cloud_file/upload', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336457, 'p', '001', '/cloud_file_tag/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336458, 'p', '001', '/cloud_file_tag/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336459, 'p', '001', '/cloud_file_tag/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336460, 'p', '001', '/cloud_file_tag/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336461, 'p', '001', '/cloud_file_tag', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336462, 'p', '001', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336463, 'p', '001', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336464, 'p', '001', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336465, 'p', '001', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336466, 'p', '001', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336467, 'p', '001', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336468, 'p', '001', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336469, 'p', '001', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336470, 'p', '001', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336471, 'p', '001', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336472, 'p', '001', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336473, 'p', '001', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336474, 'p', '001', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336475, 'p', '001', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336476, 'p', '001', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336477, 'p', '001', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336478, 'p', '001', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336479, 'p', '001', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336480, 'p', '001', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336481, 'p', '001', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336482, 'p', '001', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336483, 'p', '001', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336484, 'p', '001', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336485, 'p', '001', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336486, 'p', '001', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336487, 'p', '001', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336488, 'p', '001', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336489, 'p', '001', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336490, 'p', '001', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336491, 'p', '001', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336492, 'p', '001', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336493, 'p', '001', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336494, 'p', '001', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336495, 'p', '001', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336496, 'p', '001', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336497, 'p', '001', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336498, 'p', '001', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336499, 'p', '001', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336500, 'p', '001', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336501, 'p', '001', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336502, 'p', '001', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336503, 'p', '001', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336504, 'p', '001', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336505, 'p', '001', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336506, 'p', '001', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336507, 'p', '001', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336508, 'p', '001', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336509, 'p', '001', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336510, 'p', '001', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336511, 'p', '001', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336512, 'p', '001', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336513, 'p', '001', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336514, 'p', '001', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336515, 'p', '001', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336516, 'p', '001', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336517, 'p', '001', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336518, 'p', '001', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336519, 'p', '001', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336520, 'p', '001', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336521, 'p', '001', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336522, 'p', '001', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336523, 'p', '001', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336524, 'p', '001', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336525, 'p', '001', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336526, 'p', '001', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336527, 'p', '001', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336528, 'p', '001', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336529, 'p', '001', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336530, 'p', '001', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336531, 'p', '001', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336532, 'p', '001', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336533, 'p', '001', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336534, 'p', '001', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336535, 'p', '001', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336536, 'p', '001', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336537, 'p', '001', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336538, 'p', '001', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336539, 'p', '001', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336540, 'p', '001', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336541, 'p', '001', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336542, 'p', '001', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336543, 'p', '001', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336544, 'p', '001', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336545, 'p', '001', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336546, 'p', '001', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336547, 'p', '001', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336548, 'p', '001', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336549, 'p', '002', '/user/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336550, 'p', '002', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336551, 'p', '002', '/user/change_password', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336552, 'p', '002', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336553, 'p', '002', '/user/perm', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336554, 'p', '002', '/user/profile', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336555, 'p', '002', '/user/profile', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336556, 'p', '002', '/user/logout', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336557, 'p', '002', '/menu/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336558, 'p', '002', '/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336559, 'p', '002', '/oauth/login', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336560, 'p', '002', '/upgrade_url/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336561, 'p', '002', '/upgrade_url/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336562, 'p', '002', '/upgrade_url/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336563, 'p', '002', '/upgrade_url/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336564, 'p', '002', '/upgrade_url', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336565, 'p', '002', '/upgrade_dev_model', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336566, 'p', '002', '/upgrade_dev_model/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336567, 'p', '002', '/upgrade_dev_model/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336568, 'p', '002', '/upgrade_dev_model/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336569, 'p', '002', '/upgrade_dev_model/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336570, 'p', '002', '/upgrade_url_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336571, 'p', '002', '/upgrade_url_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336572, 'p', '002', '/upgrade_url_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336573, 'p', '002', '/upgrade_url_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336574, 'p', '002', '/upgrade_url_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336575, 'p', '002', '/upgrade_url_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336576, 'p', '002', '/upgrade_url_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336577, 'p', '002', '/upgrade_url_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336578, 'p', '002', '/upgrade_url_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336579, 'p', '002', '/upgrade_url_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336580, 'p', '002', '/company_secret', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336581, 'p', '002', '/company_secret/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336582, 'p', '002', '/company_secret/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336583, 'p', '002', '/company_secret/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336584, 'p', '002', '/company_secret/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336585, 'p', '002', '/upgrade_dev_swarm', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336586, 'p', '002', '/upgrade_dev_swarm/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336587, 'p', '002', '/upgrade_dev_swarm/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336588, 'p', '002', '/upgrade_dev_swarm/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336589, 'p', '002', '/upgrade_dev_swarm/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336590, 'p', '002', '/upgrade_file', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336591, 'p', '002', '/upgrade_file/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336592, 'p', '002', '/upgrade_file/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336593, 'p', '002', '/upgrade_file/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336594, 'p', '002', '/upgrade_file/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336595, 'p', '002', '/upgrade_apple', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336596, 'p', '002', '/upgrade_apple/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336597, 'p', '002', '/upgrade_apple/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336598, 'p', '002', '/upgrade_apple/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336599, 'p', '002', '/upgrade_apple/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336600, 'p', '002', '/upgrade_dashboard', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336601, 'p', '002', '/upgrade_file_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336602, 'p', '002', '/upgrade_file_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336603, 'p', '002', '/upgrade_file_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336604, 'p', '002', '/upgrade_file_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336605, 'p', '002', '/upgrade_file_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336606, 'p', '002', '/upgrade_file_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336607, 'p', '002', '/upgrade_file_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336608, 'p', '002', '/upgrade_file_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336609, 'p', '002', '/upgrade_file_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336610, 'p', '002', '/upgrade_file_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336611, 'p', '002', '/upgrade_dev', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336612, 'p', '002', '/upgrade_dev/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336613, 'p', '002', '/upgrade_dev/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336614, 'p', '002', '/upgrade_dev/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336615, 'p', '002', '/upgrade_dev/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336616, 'p', '002', '/upgrade_dev_group', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336617, 'p', '002', '/upgrade_dev_group/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336618, 'p', '002', '/upgrade_dev_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336619, 'p', '002', '/upgrade_dev_group/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336620, 'p', '002', '/upgrade_dev_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336621, 'p', '002', '/upgrade_tauri', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336622, 'p', '002', '/upgrade_tauri/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336623, 'p', '002', '/upgrade_tauri/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336624, 'p', '002', '/upgrade_tauri/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336625, 'p', '002', '/upgrade_tauri/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336626, 'p', '002', '/upgrade_tauri_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336627, 'p', '002', '/upgrade_tauri_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336628, 'p', '002', '/upgrade_tauri_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336629, 'p', '002', '/upgrade_tauri_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336630, 'p', '002', '/upgrade_tauri_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336631, 'p', '002', '/upgrade_tauri_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336632, 'p', '002', '/upgrade_tauri_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336633, 'p', '002', '/upgrade_tauri_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336634, 'p', '002', '/upgrade_tauri_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336635, 'p', '002', '/upgrade_tauri_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336636, 'p', '002', '/upgrade_configuration', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336637, 'p', '002', '/upgrade_configuration/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336638, 'p', '002', '/upgrade_configuration/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336639, 'p', '002', '/upgrade_configuration/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336640, 'p', '002', '/upgrade_configuration/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336641, 'p', '002', '/upgrade_configuration_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336642, 'p', '002', '/upgrade_configuration_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336643, 'p', '002', '/upgrade_configuration_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336644, 'p', '002', '/upgrade_configuration_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336645, 'p', '002', '/upgrade_configuration_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336646, 'p', '002', '/upgrade_configuration_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336647, 'p', '002', '/upgrade_configuration_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336648, 'p', '002', '/upgrade_configuration_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336649, 'p', '002', '/upgrade_configuration_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336650, 'p', '002', '/upgrade_configuration_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336651, 'p', '002', '/upgrade_apk', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336652, 'p', '002', '/upgrade_apk/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336653, 'p', '002', '/upgrade_apk/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336654, 'p', '002', '/upgrade_apk/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336655, 'p', '002', '/upgrade_apk/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336656, 'p', '002', '/upgrade_apk_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336657, 'p', '002', '/upgrade_apk_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336658, 'p', '002', '/upgrade_apk_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336659, 'p', '002', '/upgrade_apk_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336660, 'p', '002', '/upgrade_apk_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336661, 'p', '002', '/upgrade_apk_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336662, 'p', '002', '/upgrade_apk_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336663, 'p', '002', '/upgrade_apk_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336664, 'p', '002', '/upgrade_apk_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336665, 'p', '002', '/upgrade_apk_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336666, 'p', '002', '/upgrade_electron', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336667, 'p', '002', '/upgrade_electron/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336668, 'p', '002', '/upgrade_electron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336669, 'p', '002', '/upgrade_electron/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336670, 'p', '002', '/upgrade_electron/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336671, 'p', '002', '/upgrade_electron_version', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336672, 'p', '002', '/upgrade_electron_version/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336673, 'p', '002', '/upgrade_electron_version/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336674, 'p', '002', '/upgrade_electron_version/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336675, 'p', '002', '/upgrade_electron_version/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336676, 'p', '002', '/upgrade_electron_upgrade_strategy', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336677, 'p', '002', '/upgrade_electron_upgrade_strategy/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336678, 'p', '002', '/upgrade_electron_upgrade_strategy/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336679, 'p', '002', '/upgrade_electron_upgrade_strategy/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336680, 'p', '002', '/upgrade_electron_upgrade_strategy/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336681, 'p', '002', '/upgrade_company_income', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336682, 'p', '002', '/upgrade_company_income/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336683, 'p', '002', '/upgrade_company_traffic_packet', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336684, 'p', '002', '/upgrade_company_traffic_packet/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336685, 'p', '002', '/upgrade_company_traffic_packet/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336686, 'p', '002', '/upgrade_company_traffic_packet/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (336687, 'p', '002', '/upgrade_company_traffic_packet/update', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for cloud_file_tag_cloud_files
-- ----------------------------
DROP TABLE IF EXISTS `cloud_file_tag_cloud_files`;
CREATE TABLE `cloud_file_tag_cloud_files` (
  `cloud_file_tag_id` bigint unsigned NOT NULL,
  `cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`cloud_file_tag_id`,`cloud_file_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of cloud_file_tag_cloud_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for file_tag_files
-- ----------------------------
DROP TABLE IF EXISTS `file_tag_files`;
CREATE TABLE `file_tag_files` (
  `file_tag_id` bigint unsigned NOT NULL,
  `file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`file_tag_id`,`file_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of file_tag_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for fms_cloud_file_tags
-- ----------------------------
DROP TABLE IF EXISTS `fms_cloud_file_tags`;
CREATE TABLE `fms_cloud_file_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'CloudFileTag''s name | 标签名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The remark of tag | 标签的备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `cloudfiletag_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of fms_cloud_file_tags
-- ----------------------------
BEGIN;
INSERT INTO `fms_cloud_file_tags` (`id`, `created_at`, `updated_at`, `status`, `name`, `remark`) VALUES (1, '2025-04-25 06:52:26', '2025-04-25 06:52:26', 1, '测试 1', '测试 1');
INSERT INTO `fms_cloud_file_tags` (`id`, `created_at`, `updated_at`, `status`, `name`, `remark`) VALUES (2, '2025-04-25 06:52:30', '2025-04-25 06:52:30', 1, '测试 2', '测试 2');
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
  `file_type` tinyint unsigned NOT NULL COMMENT 'The file''s type | 文件类型',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The user who upload the file | 上传用户的 ID',
  `cloud_file_storage_providers` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `cloudfile_file_type` (`file_type`) USING BTREE,
  KEY `cloudfile_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of fms_cloud_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for fms_file_tags
-- ----------------------------
DROP TABLE IF EXISTS `fms_file_tags`;
CREATE TABLE `fms_file_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'FileTag''s name | 标签名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The remark of tag | 标签的备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `filetag_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of fms_file_tags
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for fms_files
-- ----------------------------
DROP TABLE IF EXISTS `fms_files`;
CREATE TABLE `fms_files` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'File''s name | 文件名称',
  `file_type` tinyint unsigned NOT NULL COMMENT 'File''s type | 文件类型',
  `size` bigint unsigned NOT NULL COMMENT 'File''s size | 文件大小',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'File''s path | 文件路径',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'User''s UUID | 用户的 UUID',
  `md5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The md5 of the file | 文件的 md5',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `file_user_id` (`user_id`) USING BTREE,
  KEY `file_file_type` (`file_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of fms_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for fms_storage_providers
-- ----------------------------
DROP TABLE IF EXISTS `fms_storage_providers`;
CREATE TABLE `fms_storage_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The cloud storage service name | 服务名称',
  `bucket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The cloud storage bucket name | 云存储服务的存储桶',
  `secret_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret ID | 密钥 ID',
  `secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret key | 密钥 Key',
  `endpoint` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The service URL | 服务器地址',
  `folder` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The folder in cloud | 云服务目标文件夹',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The service region | 服务器所在地区',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  `use_cdn` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Does it use CDN | 是否使用 CDN',
  `cdn_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'CDN URL | CDN 地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of fms_storage_providers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mcms_email_logs
-- ----------------------------
DROP TABLE IF EXISTS `mcms_email_logs`;
CREATE TABLE `mcms_email_logs` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The target email address | 目标邮箱地址',
  `subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The subject | 发送的标题',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The content | 发送的内容',
  `send_status` tinyint unsigned NOT NULL COMMENT 'The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败',
  `provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The sms service provider | 短信服务提供商',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of mcms_email_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mcms_email_providers
-- ----------------------------
DROP TABLE IF EXISTS `mcms_email_providers`;
CREATE TABLE `mcms_email_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The email provider name | 电子邮件服务的提供商',
  `auth_type` tinyint unsigned NOT NULL COMMENT 'The auth type, supported plain, CRAMMD5 | 鉴权类型, 支持 plain, CRAMMD5',
  `email_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The email address | 邮箱地址',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The email''s password | 电子邮件的密码',
  `host_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The host name is the email service''s host address | 电子邮箱服务的服务器地址',
  `identify` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The identify info, for CRAMMD5 | 身份信息, 支持 CRAMMD5',
  `secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The secret, for CRAMMD5 | 邮箱密钥, 用于 CRAMMD5',
  `port` int unsigned DEFAULT NULL COMMENT 'The port of the host | 服务器端口',
  `tls` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Whether to use TLS | 是否采用 tls 加密',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of mcms_email_providers
-- ----------------------------
BEGIN;
INSERT INTO `mcms_email_providers` (`id`, `created_at`, `updated_at`, `name`, `auth_type`, `email_addr`, `password`, `host_name`, `identify`, `secret`, `port`, `tls`, `is_default`) VALUES (1, '2025-02-10 09:50:08', '2025-02-10 09:50:08', 'tencent', 1, 'input your email address', 'input your password', 'smtp.qq.com', NULL, NULL, 465, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for mcms_sms_logs
-- ----------------------------
DROP TABLE IF EXISTS `mcms_sms_logs`;
CREATE TABLE `mcms_sms_logs` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `phone_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The target phone number | 目标电话',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The content | 发送的内容',
  `send_status` tinyint unsigned NOT NULL COMMENT 'The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败',
  `provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The sms service provider | 短信服务提供商',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of mcms_sms_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mcms_sms_providers
-- ----------------------------
DROP TABLE IF EXISTS `mcms_sms_providers`;
CREATE TABLE `mcms_sms_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The SMS provider name | 短信服务的提供商',
  `secret_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret ID | 密钥 ID',
  `secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret key | 密钥 Key',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The service region | 服务器所在地区',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of mcms_sms_providers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `role_id` bigint unsigned NOT NULL,
  `menu_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

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
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 19);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 20);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 21);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 22);
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES (1, 23);
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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `api_path_method` (`path`,`method`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=314 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

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
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (122, '2025-02-20 09:25:42', '2025-09-22 14:45:50', '/upgrade_url/list', 'list', 'upgrade_url', 'Upgrade', 'POST', 1);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (123, '2025-02-20 09:25:42', '2025-09-22 14:45:44', '/upgrade_url/delete', 'delete', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (127, '2025-02-20 09:25:42', '2025-09-22 14:45:41', '/upgrade_url/create', 'create', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (129, '2025-02-20 09:25:42', '2025-09-22 14:45:39', '/upgrade_url/update', 'update', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (130, '2025-02-20 09:25:42', '2025-09-22 14:46:53', '/upgrade_url', 'detail', 'upgrade_url', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (131, '2025-02-20 09:25:42', '2025-09-22 14:46:52', '/upgrade_dev_model', 'detail', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (132, '2025-02-20 09:25:42', '2025-09-22 14:45:56', '/upgrade_dev_model/list', 'list', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (133, '2025-02-20 09:25:42', '2025-09-22 14:45:58', '/upgrade_dev_model/delete', 'delete', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (134, '2025-02-20 09:25:42', '2025-09-22 14:46:00', '/upgrade_dev_model/create', 'create', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (135, '2025-02-20 09:25:42', '2025-09-22 14:46:04', '/upgrade_dev_model/update', 'update', 'upgrade_dev_model', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (136, '2025-02-20 09:25:42', '2025-09-22 14:51:42', '/upgrade_url_upgrade_strategy', 'detail', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (137, '2025-02-20 09:25:42', '2025-09-22 14:47:09', '/upgrade_url_upgrade_strategy/list', 'list', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (138, '2025-02-20 09:25:42', '2025-09-22 14:49:34', '/upgrade_url_upgrade_strategy/update', 'update', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (139, '2025-02-20 09:25:42', '2025-09-22 14:51:23', '/upgrade_url_upgrade_strategy/create', 'create', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (140, '2025-02-20 09:25:42', '2025-09-22 14:48:43', '/upgrade_url_upgrade_strategy/delete', 'delete', 'upgrade_url_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (141, '2025-02-20 09:25:42', '2025-09-22 14:51:19', '/upgrade_url_version', 'detail', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (142, '2025-02-20 09:25:42', '2025-09-22 14:47:11', '/upgrade_url_version/list', 'list', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (143, '2025-02-20 09:25:42', '2025-09-22 14:48:46', '/upgrade_url_version/delete', 'delete', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (144, '2025-02-20 09:25:42', '2025-09-22 14:51:25', '/upgrade_url_version/create', 'create', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (145, '2025-02-20 09:25:42', '2025-09-22 14:49:37', '/upgrade_url_version/update', 'update', 'upgrade_url_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (146, '2025-02-20 09:25:42', '2025-09-22 14:51:43', '/company_secret', 'detail', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (147, '2025-02-20 09:25:42', '2025-09-22 14:47:14', '/company_secret/list', 'list', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (148, '2025-02-20 09:25:42', '2025-09-22 14:48:48', '/company_secret/delete', 'delete', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (149, '2025-02-20 09:25:42', '2025-09-22 14:51:28', '/company_secret/create', 'create', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (150, '2025-02-20 09:25:42', '2025-09-22 14:49:40', '/company_secret/update', 'update', 'sys_company_secret', 'Core', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (151, '2025-02-20 09:25:42', '2025-09-22 14:51:45', '/upgrade_dev_swarm', 'detail', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (152, '2025-02-20 09:25:42', '2025-09-22 14:47:16', '/upgrade_dev_swarm/list', 'list', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (153, '2025-02-20 09:25:42', '2025-09-22 14:48:50', '/upgrade_dev_swarm/delete', 'delete', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (154, '2025-02-20 09:25:42', '2025-09-22 14:51:30', '/upgrade_dev_swarm/create', 'create', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (155, '2025-02-20 09:25:42', '2025-09-22 14:49:42', '/upgrade_dev_swarm/update', 'update', 'upgrade_dev_swarm', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (156, '2025-02-20 09:25:42', '2025-09-22 14:51:47', '/upgrade_file', 'detail', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (157, '2025-02-20 09:25:42', '2025-09-22 14:47:19', '/upgrade_file/list', 'list', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (158, '2025-02-20 09:25:42', '2025-09-22 14:48:55', '/upgrade_file/delete', 'delete', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (159, '2025-02-20 09:25:42', '2025-09-22 14:49:46', '/upgrade_file/update', 'update', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (160, '2025-02-20 09:25:42', '2025-09-22 14:51:31', '/upgrade_file/create', 'create', 'upgrade_file', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (171, '2025-02-20 09:25:42', '2025-09-22 14:47:22', '/upgrade_dashboard', 'upgrade_dashboard', 'upgrade_dashboard', 'Upgrade', 'POST', 1);
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
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (226, '2025-02-20 09:25:42', '2025-09-22 14:51:53', '/upgrade_file_version', 'detail', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (227, '2025-02-20 09:25:42', '2025-09-22 14:48:05', '/upgrade_file_version/list', 'list', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (228, '2025-02-20 09:25:42', '2025-09-22 14:48:56', '/upgrade_file_version/delete', 'delete', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (229, '2025-02-20 09:25:42', '2025-09-22 14:50:53', '/upgrade_file_version/create', 'create', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (230, '2025-02-20 09:25:42', '2025-09-22 14:49:47', '/upgrade_file_version/update', 'update', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (231, '2025-02-20 09:25:42', '2025-09-22 14:51:54', '/upgrade_file_upgrade_strategy', 'detail', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (232, '2025-02-20 09:25:42', '2025-09-22 14:51:10', '/upgrade_file_upgrade_strategy/list', 'list', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (233, '2025-02-20 09:25:42', '2025-09-22 14:48:58', '/upgrade_file_upgrade_strategy/delete', 'delete', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (234, '2025-02-20 09:25:42', '2025-09-22 14:50:51', '/upgrade_file_upgrade_strategy/create', 'create', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (235, '2025-02-20 09:25:42', '2025-09-22 14:49:50', '/upgrade_file_upgrade_strategy/update', 'update', 'upgrade_file_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (236, '2025-02-20 09:25:42', '2025-09-22 14:51:56', '/upgrade_dev', 'detail', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (237, '2025-02-20 09:25:42', '2025-09-22 14:48:07', '/upgrade_dev/list', 'list', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (238, '2025-02-20 09:25:42', '2025-09-22 14:49:00', '/upgrade_dev/delete', 'delete', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (239, '2025-02-20 09:25:42', '2025-09-22 14:50:48', '/upgrade_dev/create', 'create', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (240, '2025-02-20 09:25:42', '2025-09-22 14:49:51', '/upgrade_dev/update', 'update', 'upgrade_dev', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (242, '2025-02-20 09:25:42', '2025-09-22 14:51:58', '/upgrade_dev_group', 'detail', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (243, '2025-02-20 09:25:42', '2025-09-22 14:48:11', '/upgrade_dev_group/list', 'list', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (244, '2025-02-20 09:25:42', '2025-09-22 14:49:02', '/upgrade_dev_group/delete', 'delete', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (245, '2025-02-20 09:25:42', '2025-09-22 14:50:40', '/upgrade_dev_group/create', 'create', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (246, '2025-02-20 09:25:42', '2025-09-22 14:49:53', '/upgrade_dev_group/update', 'update', 'upgrade_dev_group', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (247, '2025-02-20 09:25:42', '2025-09-22 14:52:00', '/upgrade_tauri', 'detail', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (248, '2025-02-20 09:25:42', '2025-09-22 14:48:09', '/upgrade_tauri/list', 'list', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (249, '2025-02-20 09:25:42', '2025-09-22 14:49:03', '/upgrade_tauri/delete', 'delete', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (250, '2025-02-20 09:25:42', '2025-09-22 14:49:56', '/upgrade_tauri/update', 'update', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (251, '2025-02-20 09:25:42', '2025-09-22 14:50:42', '/upgrade_tauri/create', 'create', 'upgrade_tauri', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (252, '2025-02-20 09:25:42', '2025-09-22 14:52:01', '/upgrade_tauri_version', 'detail', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (253, '2025-02-20 09:25:42', '2025-09-22 14:48:12', '/upgrade_tauri_version/list', 'list', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (254, '2025-02-20 09:25:42', '2025-09-22 14:49:05', '/upgrade_tauri_version/delete', 'delete', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (255, '2025-02-20 09:25:42', '2025-09-22 14:50:38', '/upgrade_tauri_version/create', 'create', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (256, '2025-02-20 09:25:42', '2025-09-22 14:49:57', '/upgrade_tauri_version/update', 'update', 'upgrade_tauri_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (257, '2025-02-20 09:25:42', '2025-09-22 14:52:03', '/upgrade_tauri_upgrade_strategy', 'detail', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (258, '2025-02-20 09:25:42', '2025-09-22 14:48:14', '/upgrade_tauri_upgrade_strategy/list', 'list', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (259, '2025-02-20 09:25:42', '2025-09-22 14:49:06', '/upgrade_tauri_upgrade_strategy/delete', 'delete', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (260, '2025-02-20 09:25:42', '2025-09-22 14:50:36', '/upgrade_tauri_upgrade_strategy/create', 'create', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (261, '2025-02-20 09:25:42', '2025-09-22 14:49:59', '/upgrade_tauri_upgrade_strategy/update', 'update', 'upgrade_tauri_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (262, '2025-02-20 09:25:42', '2025-09-22 14:52:05', '/upgrade_configuration', 'detail', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (263, '2025-02-20 09:25:42', '2025-09-22 14:48:15', '/upgrade_configuration/list', 'list', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (264, '2025-02-20 09:25:42', '2025-09-22 14:49:08', '/upgrade_configuration/delete', 'delete', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (265, '2025-02-20 09:25:42', '2025-09-22 14:52:08', '/upgrade_configuration/create', 'create', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (266, '2025-02-20 09:25:42', '2025-09-22 14:50:01', '/upgrade_configuration/update', 'update', 'upgrade_configuration', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (267, '2025-02-20 09:25:42', '2025-09-22 14:52:06', '/upgrade_configuration_version', 'detail', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (268, '2025-02-20 09:25:42', '2025-09-22 14:48:17', '/upgrade_configuration_version/list', 'list', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (269, '2025-02-20 09:25:42', '2025-09-22 14:49:11', '/upgrade_configuration_version/delete', 'delete', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (270, '2025-02-20 09:25:42', '2025-09-22 14:50:34', '/upgrade_configuration_version/create', 'create', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (271, '2025-02-20 09:25:42', '2025-09-22 14:50:02', '/upgrade_configuration_version/update', 'update', 'upgrade_configuration_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (272, '2025-02-20 09:25:42', '2025-09-22 14:52:13', '/upgrade_configuration_upgrade_strategy', 'detail', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (273, '2025-02-20 09:25:42', '2025-09-22 14:48:18', '/upgrade_configuration_upgrade_strategy/list', 'list', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (274, '2025-02-20 09:25:42', '2025-09-22 14:49:14', '/upgrade_configuration_upgrade_strategy/delete', 'delete', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (275, '2025-02-20 09:25:42', '2025-09-22 14:50:32', '/upgrade_configuration_upgrade_strategy/create', 'create', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (276, '2025-02-20 09:25:42', '2025-09-22 14:50:04', '/upgrade_configuration_upgrade_strategy/update', 'update', 'upgrade_configuration_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (277, '2025-02-20 09:25:42', '2025-09-22 14:52:15', '/upgrade_apk', 'detail', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (278, '2025-02-20 09:25:42', '2025-09-22 14:48:19', '/upgrade_apk/list', 'list', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (279, '2025-02-20 09:25:42', '2025-09-22 14:49:17', '/upgrade_apk/delete', 'delete', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (280, '2025-02-20 09:25:42', '2025-09-22 14:50:30', '/upgrade_apk/create', 'create', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (281, '2025-02-20 09:25:42', '2025-09-22 14:50:07', '/upgrade_apk/update', 'update', 'upgrade_apk', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (282, '2025-02-20 09:25:42', '2025-09-22 14:52:16', '/upgrade_apk_version', 'detail', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (283, '2025-02-20 09:25:42', '2025-09-22 14:48:21', '/upgrade_apk_version/list', 'list', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (284, '2025-02-20 09:25:42', '2025-09-22 14:49:18', '/upgrade_apk_version/delete', 'delete', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (285, '2025-02-20 09:25:42', '2025-09-22 14:50:28', '/upgrade_apk_version/create', 'create', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (286, '2025-02-20 09:25:42', '2025-09-22 14:50:09', '/upgrade_apk_version/update', 'update', 'upgrade_apk_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (287, '2025-02-20 09:25:42', '2025-09-22 14:52:18', '/upgrade_apk_upgrade_strategy', 'detail', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (288, '2025-02-20 09:25:42', '2025-09-22 14:48:23', '/upgrade_apk_upgrade_strategy/list', 'list', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (289, '2025-02-20 09:25:42', '2025-09-22 14:49:20', '/upgrade_apk_upgrade_strategy/delete', 'delete', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (290, '2025-02-20 09:25:42', '2025-09-22 14:50:27', '/upgrade_apk_upgrade_strategy/create', 'create', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (291, '2025-02-20 09:25:42', '2025-09-22 14:52:20', '/upgrade_apk_upgrade_strategy/update', 'update', 'upgrade_apk_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (292, '2025-02-20 09:25:42', '2025-09-22 14:52:23', '/upgrade_electron', 'detail', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (293, '2025-02-20 09:25:42', '2025-09-22 14:52:24', '/upgrade_electron/list', 'list', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (294, '2025-02-20 09:25:42', '2025-09-22 14:49:22', '/upgrade_electron/delete', 'delete', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (295, '2025-02-20 09:25:42', '2025-09-22 14:50:25', '/upgrade_electron/create', 'create', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (296, '2025-02-20 09:25:42', '2025-09-22 14:50:12', '/upgrade_electron/update', 'update', 'upgrade_electron', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (297, '2025-02-20 09:25:42', '2025-09-22 14:52:31', '/upgrade_electron_version', 'detail', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (298, '2025-02-20 09:25:42', '2025-09-22 14:48:26', '/upgrade_electron_version/list', 'list', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (299, '2025-02-20 09:25:42', '2025-09-22 14:49:24', '/upgrade_electron_version/delete', 'delete', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (300, '2025-02-20 09:25:42', '2025-09-22 14:50:24', '/upgrade_electron_version/create', 'create', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (301, '2025-02-20 09:25:42', '2025-09-22 14:50:15', '/upgrade_electron_version/update', 'update', 'upgrade_electron_version', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (302, '2025-02-20 09:25:42', '2025-09-22 14:52:33', '/upgrade_electron_upgrade_strategy', 'detail', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (303, '2025-02-20 09:25:42', '2025-09-22 14:48:27', '/upgrade_electron_upgrade_strategy/list', 'list', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (304, '2025-02-20 09:25:42', '2025-09-22 14:49:26', '/upgrade_electron_upgrade_strategy/delete', 'delete', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (305, '2025-02-20 09:25:42', '2025-09-22 14:50:22', '/upgrade_electron_upgrade_strategy/create', 'create', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (306, '2025-02-20 09:25:42', '2025-09-22 14:52:37', '/upgrade_electron_upgrade_strategy/update', 'update', 'upgrade_electron_upgrade_strategy', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (307, '2025-02-20 09:25:42', '2025-09-22 14:52:40', '/upgrade_company_income', 'detail', 'upgrade_company_income', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (308, '2025-02-20 09:25:42', '2025-09-22 14:48:30', '/upgrade_company_income/list', 'list', 'upgrade_company_income', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (309, '2025-09-09 09:16:34', '2025-09-22 14:52:46', '/upgrade_company_traffic_packet', 'detail', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (310, '2025-09-09 09:16:26', '2025-09-22 14:48:31', '/upgrade_company_traffic_packet/list', 'list', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (311, '2025-09-09 09:16:28', '2025-09-22 14:49:30', '/upgrade_company_traffic_packet/delete', 'delete', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (312, '2025-09-09 09:16:29', '2025-09-22 14:50:19', '/upgrade_company_traffic_packet/create', 'create', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `service_name`, `method`, `is_required`) VALUES (313, '2025-09-09 09:16:32', '2025-09-22 14:50:16', '/upgrade_company_traffic_packet/update', 'update', 'upgrade_company_traffic_packet', 'Upgrade', 'POST', 0);
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
  `access_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密钥id',
  `secret_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密钥key',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公司密钥表';

-- ----------------------------
-- Records of sys_company_secret
-- ----------------------------
BEGIN;
INSERT INTO `sys_company_secret` (`id`, `created_at`, `updated_at`, `company_id`, `access_key`, `secret_key`) VALUES (1, '2025-02-24 07:58:27', '2025-09-23 09:25:48', 1, 'mui2W50H1j-OC4xD6PgQag', 'PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc');
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公司信息表';

-- ----------------------------
-- Records of sys_companys
-- ----------------------------
BEGIN;
INSERT INTO `sys_companys` (`id`, `created_at`, `updated_at`, `name`) VALUES (1, '2025-02-24 07:55:58', '2025-09-23 09:51:51', '默认');
COMMIT;

-- ----------------------------
-- Table structure for sys_configuration
-- ----------------------------
DROP TABLE IF EXISTS `sys_configuration`;
CREATE TABLE `sys_configuration` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Configurarion name | 配置名称',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuration key | 配置的键名',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuraion value | 配置的值',
  `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuration category | 配置的分类',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `configuration_key` (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_configuration
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_departments
-- ----------------------------
DROP TABLE IF EXISTS `sys_departments`;
CREATE TABLE `sys_departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Department name | 部门名称',
  `ancestors` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Parents'' IDs | 父级列表',
  `leader` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Department leader | 部门负责人',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Leader''s phone number | 负责人电话',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Leader''s email | 部门负责人电子邮箱',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT 'Parent department ID | 父级部门ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_departments
-- ----------------------------
BEGIN;
INSERT INTO `sys_departments` (`id`, `created_at`, `updated_at`, `status`, `sort`, `name`, `ancestors`, `leader`, `phone`, `email`, `remark`, `parent_id`) VALUES (1, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 1, 1, 'department.managementDepartment', '', 'admin', '18888888888', 'simpleadmin@gmail.com', 'Super Administrator', 1000000);
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The title shown in the ui | 展示名称 （建议配合i18n）',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of dictionary for search | 字典搜索名称',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The description of dictionary | 字典的描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The title shown in the ui | 展示名称 （建议配合i18n）',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'key | 键',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'value | 值',
  `dictionary_id` bigint unsigned DEFAULT NULL COMMENT 'Dictionary ID | 字典ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_dictionary_details
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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `menu_name` (`name`) USING BTREE,
  UNIQUE KEY `menu_path` (`path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (1, '2025-02-10 09:33:13', '2025-09-22 14:33:07', 0, 1, 1, '/dashboard', 'Dashboard', '', '/dashboard/analytics/index', 0, 'Core', NULL, 'route.dashboard', 'ant-design:dashboard-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (2, '2025-02-10 09:33:13', '2025-09-22 14:56:10', 8, 1, 0, '/system', 'SystemManagement', '', 'LAYOUT', 0, 'Core', NULL, 'route.systemManagementTitle', 'ant-design:tool-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (3, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 1, 2, 1, '/menu', 'MenuManagement', '', '/sys/menu/index', 0, 'Core', NULL, 'route.menuManagementTitle', 'ant-design:bars-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (4, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 2, 2, 1, '/role', 'RoleManagement', '', '/sys/role/index', 0, 'Core', NULL, 'route.roleManagementTitle', 'ant-design:user-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (5, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 3, 2, 1, '/user', 'UserManagement', '', '/sys/user/index', 0, 'Core', NULL, 'route.userManagementTitle', 'ant-design:user-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (6, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 4, 2, 1, '/department', 'DepartmentManagement', '', '/sys/department/index', 0, 'Core', NULL, 'route.departmentManagement', 'ic:outline-people-alt', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (7, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 5, 2, 1, '/api', 'APIManagement', '', '/sys/api/index', 0, 'Core', NULL, 'route.apiManagementTitle', 'ant-design:api-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (8, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 6, 2, 1, '/dictionary', 'DictionaryManagement', '', '/sys/dictionary/index', 0, 'Core', NULL, 'route.dictionaryManagementTitle', 'ant-design:book-outlined', 0, 0, 0, 0, '', 0, 1, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (9, '2025-02-10 09:33:13', '2025-09-22 14:56:17', 9, 1, 0, '/other', 'OtherPages', '', 'LAYOUT', 0, 'Core', NULL, 'route.otherPages', 'ant-design:question-circle-outlined', 1, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (10, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 3, 1, 1, '/profile', 'Profile', '', '/sys/profile/index', 0, 'Core', NULL, 'route.userProfileTitle', 'ant-design:profile-outlined', 1, 0, 0, 0, '', 0, 0, 0, 20, '', 9);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (11, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 6, 2, 1, '/oauth', 'OauthManagement', '', '/sys/oauth/index', 0, 'Core', NULL, 'route.oauthManagement', 'ant-design:unlock-filled', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (12, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 7, 2, 1, '/token', 'TokenManagement', '', '/sys/token/index', 0, 'Core', NULL, 'route.tokenManagement', 'ant-design:lock-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (13, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 8, 2, 1, '/position', 'PositionManagement', '', '/sys/position/index', 0, 'Core', NULL, 'route.positionManagement', 'ic:twotone-work-outline', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (14, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 8, 2, 1, '/task', 'TaskManagement', '', '/sys/task/index', 0, 'Job', NULL, 'route.taskManagement', 'ic:baseline-access-alarm', 1, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (15, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 9, 2, 1, '/configuration', 'ConfigurationManagement', '', '/sys/configuration/index', 0, 'Core', NULL, 'route.configurationManagement', 'carbon:data-2', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 2);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (16, '2025-02-10 09:50:08', '2025-02-10 09:50:08', 4, 1, 1, '/mcms_dir', 'MessageCenterManagement', '', 'LAYOUT', 0, 'Mcms', NULL, 'route.messageCenterManagement', 'clarity:email-line', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (17, '2025-02-10 09:50:08', '2025-02-10 09:50:08', 1, 2, 2, '/mcms_email_provider', 'EmailProviderManagement', '', '/mcms/emailProvider/index', 0, 'Mcms', NULL, 'route.emailProviderManagement', 'clarity:email-line', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 16);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (18, '2025-02-10 09:50:08', '2025-02-20 09:48:06', 2, 2, 1, '/mcms_sms_provider', 'SmsProviderManagement', '', '/mcms/smsProvider/index', 0, 'Mcms', '', 'route.smsProviderManagement', 'clarity:mobile-line', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 16);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (24, '2025-02-20 09:30:03', '2025-09-22 14:35:25', 400, 1, 1, '/upgrade_url_dir', 'UrlManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'URL升级', 'ant-design:share-alt-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (26, '2025-02-20 09:52:10', '2025-09-22 14:20:54', 1, 2, 1, '/upgrade/url', 'UrlAppManagement', '', '/upgrade/url/index', 0, 'Upgrade', '', 'URL应用', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (27, '2025-09-11 11:28:22', '2025-09-22 14:35:22', 2, 2, 1, '/upgrade/url_version', 'UrlVersionManagement', '', '/upgrade/url_version/index', 0, 'Upgrade', NULL, 'URL应用版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (28, '2025-09-11 11:28:25', '2025-09-22 14:35:20', 3, 2, 1, '/upgrade/url_upgrade_strategy', 'UrlUpgradeStrategyManagement', '', '/upgrade/url_upgrade_strategy/index', 0, 'Upgrade', NULL, 'URL升级任务', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 24);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (29, '2025-02-22 16:00:09', '2025-09-22 14:41:25', 202, 1, 1, '/upgrade/dev_model', 'DevModelManagement', '', '/upgrade/dev_model/index', 0, 'Upgrade', '', '设备机型管理', 'ant-design:laptop-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (32, '2025-02-25 03:11:59', '2025-09-22 14:18:19', 500, 1, 1, '/upgrade_file_dir', 'UpgradeFileManagement', '', 'LAYOUT', 0, 'Upgrade', '', '文件升级', 'ant-design:file-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (35, '2025-02-25 03:17:26', '2025-09-22 14:17:18', 100, 1, 1, '/company_secret', 'CompanySecretManagement', '', '/sys/company_secret/index', 0, 'Core', '', '密钥', 'ant-design:key-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (46, '2025-03-10 03:35:03', '2025-03-10 03:35:03', 3, 1, 1, '/fms_dir', 'FileManagementDirectory', '', 'LAYOUT', 0, 'Fms', NULL, 'route.fileManagement', 'ant-design:folder-open-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (47, '2025-03-10 03:35:03', '2025-03-10 03:35:03', 1, 2, 1, '/fms/file', 'FileManagement', '', '/fms/file/index', 0, 'Fms', NULL, 'route.fileManagement', 'ant-design:folder-open-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (48, '2025-03-10 03:35:03', '2025-03-10 03:35:03', 2, 2, 1, '/fms/file_tag', 'FileTagManagement', '', '/fms/fileTag/index', 0, 'Fms', NULL, 'route.fileTagManagement', 'ant-design:book-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (49, '2025-03-10 03:35:03', '2025-03-10 03:35:03', 3, 2, 1, '/fms/storage_provider', 'StorageProviderManagement', '', '/fms/storageProvider/index', 0, 'Fms', NULL, 'route.storageProviderManagement', 'mdi:cloud-outline', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (50, '2025-03-10 03:35:04', '2025-03-10 03:35:04', 4, 2, 1, '/fms/cloud_file', 'CloudFileManagement', '', '/fms/cloudFile/index', 0, 'Fms', NULL, 'route.cloudFileManagement', 'ant-design:folder-open-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (51, '2025-03-10 03:35:04', '2025-03-10 03:35:04', 5, 2, 1, '/fms/cloud_file_tag', 'CloudFileTagManagement', '', '/fms/cloudFileTag/index', 0, 'Fms', NULL, 'route.cloudFileTagManagement', 'ant-design:book-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 46);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (52, '2025-02-21 06:19:59', '2025-09-22 14:21:04', 1, 2, 1, '/upgrade/file', 'FileAppManagement', '', '/upgrade/file/index', 0, 'Upgrade', NULL, '文件', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (53, '2025-09-11 11:28:17', '2025-09-22 14:35:18', 2, 2, 1, '/upgrade/file_version', 'FileVersionManagement', '', '/upgrade/file_version/index', 0, 'Upgrade', NULL, '文件应用版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (54, '2025-09-11 11:28:21', '2025-09-22 14:35:19', 3, 2, 1, '/upgrade/file_upgrade_strategy', 'FileUpgradeStrategyManagement', '', '/upgrade/file_upgrade_strategy/index', 0, 'Upgrade', NULL, '文件升级任务', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 32);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (55, '2025-02-22 16:00:09', '2025-09-22 14:28:03', 200, 1, 1, '/upgrade/dev', 'DevManagement', '', '/upgrade/dev/index', 0, 'Upgrade', NULL, '设备管理', 'ant-design:laptop-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (56, '2025-02-22 16:00:09', '2025-09-22 14:27:15', 201, 1, 1, '/upgrade/dev_group', 'DevGroupManagement', '', '/upgrade/dev_group/index', 0, 'Upgrade', NULL, '设备分组管理', 'ant-design:solution-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (57, '2025-02-21 06:19:59', '2025-09-22 14:21:07', 1, 2, 1, '/upgrade/tauri', 'TauriAppManagement', '', '/upgrade/tauri/index', 0, 'Upgrade', NULL, 'Tauri应用', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (58, '2025-09-11 11:28:11', '2025-09-22 14:29:24', 2, 2, 1, '/upgrade/tauri_version', 'TauriVersionManagement', '', '/upgrade/tauri_version/index', 0, 'Upgrade', NULL, 'Tauri应用版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (59, '2025-09-11 11:28:13', '2025-09-22 14:28:38', 3, 2, 1, '/upgrade/tauri_upgrade_strategy', 'TauriUpgradeStrategyManagement', '', '/upgrade/tauri_upgrade_strategy/index', 0, 'Upgrade', NULL, 'Tauri升级任务', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 60);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (60, '2025-02-25 03:11:59', '2025-09-22 14:59:14', 600, 1, 1, '/upgrade_tauri_dir', 'UpgradeTauriManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'Tauri升级', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (61, '2025-02-25 03:11:59', '2025-09-22 14:59:31', 700, 1, 1, '/upgrade_configuration_dir', 'UpgradeConfigurationManagement', '', 'LAYOUT', 0, 'Upgrade', '', '配置升级', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (62, '2025-02-25 03:11:59', '2025-09-22 14:39:38', 1, 2, 1, '/upgrade/configuration', 'ConfigurationAppManagement', '', '/upgrade/configuration/index', 0, 'Upgrade', NULL, '配置', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (63, '2025-09-11 11:28:06', '2025-09-22 14:29:28', 2, 2, 1, '/upgrade/configuration_version', 'ConfigurationVersionManagement', '', '/upgrade/configuration_version/index', 0, 'Upgrade', NULL, '配置版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (65, '2025-09-11 11:28:08', '2025-02-21 06:19:59', 3, 2, 1, '/upgrade/configuration_upgrade_strategy', 'ConfigurationUpgradeStrategyManagement', '', '/upgrade/configuration_upgrade_strategy/index', 0, 'Upgrade', NULL, '配置升级任务', 'ant-design:home-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 61);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (66, '2025-02-25 03:11:59', '2025-09-22 14:39:11', 800, 1, 1, '/upgrade_apk_dir', 'UpgradeApkManagement', '', 'LAYOUT', 0, 'Upgrade', '', '安卓应用升级', 'ant-design:android-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (67, '2025-02-25 03:11:59', '2025-09-22 14:21:35', 1, 2, 1, '/upgrade/apk', 'ApkAppManagement', '', '/upgrade/apk/index', 0, 'Upgrade', NULL, '安卓应用', 'ant-design:android-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (68, '2025-09-11 11:27:58', '2025-09-22 14:29:30', 2, 2, 1, '/upgrade/apk_version', 'ApkVersionManagement', '', '/upgrade/apk_version/index', 0, 'Upgrade', NULL, '安卓应用版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (69, '2025-09-11 11:28:00', '2025-09-22 14:28:35', 3, 2, 1, '/upgrade/apk_upgrade_strategy', 'ApkUpgradeStrategyManagement', '', '/upgrade/apk_upgrade_strategy/index', 0, 'Upgrade', NULL, '安卓应用升级任务', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 66);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (70, '2025-02-25 03:11:59', '2025-09-22 06:38:38', 900, 1, 1, '/upgrade_electron_dir', 'UpgradeElectronManagement', '', 'LAYOUT', 0, 'Upgrade', '', 'Electron应用升级', 'ant-design:home-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 1000000);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (73, '2025-02-25 03:11:59', '2025-09-22 14:21:17', 1, 2, 1, '/upgrade/electron', 'ElectronAppManagement', '', '/upgrade/electron/index', 0, 'Upgrade', NULL, 'Electron应用', 'ant-design:appstore-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (74, '2025-09-11 11:27:50', '2025-09-22 14:59:39', 2, 2, 1, '/upgrade/electron_version', 'ElectronVersionManagement', '', '/upgrade/electron_version/index', 0, 'Upgrade', NULL, 'Electron应用版本', 'ant-design:appstore-add-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
INSERT INTO `sys_menus` (`id`, `created_at`, `updated_at`, `sort`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `disabled`, `service_name`, `permission`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) VALUES (75, '2025-09-11 11:27:51', '2025-09-22 14:59:42', 3, 2, 1, '/upgrade/electron_upgrade_strategy', 'ElectronUpgradeStrategyManagement', '', '/upgrade/electron_upgrade_strategy/index', 0, 'Upgrade', NULL, 'Electron应用升级任务', 'ant-design:paper-clip-outlined', 0, 0, 0, 0, '', 0, 0, 0, 20, '', 70);
COMMIT;

-- ----------------------------
-- Table structure for sys_oauth_providers
-- ----------------------------
DROP TABLE IF EXISTS `sys_oauth_providers`;
CREATE TABLE `sys_oauth_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The provider''s name | 提供商名称',
  `client_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The client id | 客户端 id',
  `client_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The client secret | 客户端密钥',
  `redirect_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The redirect url | 跳转地址',
  `scopes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The scopes | 权限范围',
  `auth_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The auth url of the provider | 认证地址',
  `token_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The token url of the provider | 获取 token地址',
  `auth_style` bigint unsigned NOT NULL COMMENT 'The auth style, 0: auto detect 1: third party log in 2: log in with username and password | 鉴权方式 0 自动 1 第三方登录 2 使用用户名密码',
  `info_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The URL to request user information by token | 用户信息请求地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_oauth_providers
-- ----------------------------
BEGIN;
INSERT INTO `sys_oauth_providers` (`id`, `created_at`, `updated_at`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) VALUES (1, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 'google', 'your client id', 'your client secret', 'http://localhost:3100/oauth/login/callback', 'email openid', 'https://accounts.google.com/o/oauth2/auth', 'https://oauth2.googleapis.com/token', 1, 'https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN');
INSERT INTO `sys_oauth_providers` (`id`, `created_at`, `updated_at`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) VALUES (2, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 'github', 'your client id', 'your client secret', 'http://localhost:3100/oauth/login/callback', 'email openid', 'https://github.com/login/oauth/authorize', 'https://github.com/login/oauth/access_token', 2, 'https://api.github.com/user');
COMMIT;

-- ----------------------------
-- Table structure for sys_positions
-- ----------------------------
DROP TABLE IF EXISTS `sys_positions`;
CREATE TABLE `sys_positions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Position Name | 职位名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'The code of position | 职位编码',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `position_code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_positions
-- ----------------------------
BEGIN;
INSERT INTO `sys_positions` (`id`, `created_at`, `updated_at`, `status`, `sort`, `name`, `code`, `remark`) VALUES (1, '2025-02-10 09:33:13', '2025-02-10 09:33:13', 1, 1, 'position.ceo', '001', 'CEO');
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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `role_code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
BEGIN;
INSERT INTO `sys_roles` (`id`, `created_at`, `updated_at`, `status`, `name`, `code`, `default_router`, `remark`, `sort`) VALUES (1, '2025-09-09 10:23:13', '2025-09-22 06:54:49', 1, 'role.admin', '001', 'dashboard', '超级管理员', 1);
INSERT INTO `sys_roles` (`id`, `created_at`, `updated_at`, `status`, `name`, `code`, `default_router`, `remark`, `sort`) VALUES (2, '2025-09-09 17:39:56', '2025-09-09 09:39:56', 1, 'role.stuff', '002', 'dashboard', '普通用户', 2);
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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `token_expired_at` (`expired_at`) USING BTREE,
  KEY `token_uuid` (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_tokens
-- ----------------------------
BEGIN;
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019975e7-7d5d-7df3-908a-6112b21a1423', '2025-09-23 09:28:49', '2025-09-23 09:28:49', 1, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXB0SWQiOjEsImV4cCI6MTc1ODg3ODkyOSwiaWF0IjoxNzU4NjE5NzI5LCJyb2xlSWQiOiIwMDEiLCJ1c2VySWQiOiIwMTk0ZWYzNC1hNmQxLTcyOGUtYTcwZC0zNjJmNTAzNWFhYjcifQ.bUi4ac3i5r-AeDNi_UDBjiz2tBL2X4SysFoIjNo8Ys4', 'core_user', '2025-09-26 09:28:49');
INSERT INTO `sys_tokens` (`id`, `created_at`, `updated_at`, `status`, `uuid`, `username`, `token`, `source`, `expired_at`) VALUES ('019984a2-8cb4-7df3-b7ce-afab08c34749', '2025-09-26 06:07:49', '2025-09-26 06:07:49', 1, '0194ef34-a6d1-728e-a70d-362f5035aab7', 'admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXB0SWQiOjEsImV4cCI6MTc1OTEyNjA2OSwiaWF0IjoxNzU4ODY2ODY5LCJyb2xlSWQiOiIwMDEiLCJ1c2VySWQiOiIwMTk0ZWYzNC1hNmQxLTcyOGUtYTcwZC0zNjJmNTAzNWFhYjcifQ.C9oWS2QBw6vUNk7azJbn9jPAk2Gta9rVowgBkfqCWOs', 'core_user', '2025-09-29 06:07:49');
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
  `department_id` bigint unsigned DEFAULT '1' COMMENT 'Department ID | 部门ID',
  `company_id` bigint unsigned DEFAULT '1' COMMENT 'Company ID | 公司ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `nickname` (`nickname`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  UNIQUE KEY `user_username_email` (`username`,`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `status`, `deleted_at`, `username`, `password`, `nickname`, `description`, `home_path`, `mobile`, `email`, `avatar`, `department_id`, `company_id`) VALUES ('0194ef34-a6d1-728e-a70d-362f5035aab7', '2025-08-15 11:26:56', '2025-09-23 09:27:45', 1, NULL, 'admin', '$2a$10$1JMM5u4ch3ml9dNafqEZEu.vnOBy74lMZF22DaL10NmpN1V22TBc2', 'admin', NULL, '/dashboard', '13113131111', 'simple_admin@gmail.com', '', 1, 1);
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用';

-- ----------------------------
-- Records of upgrade_apk
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
  `gray_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '灰度策略id数据',
  `is_flow_limit` int NOT NULL DEFAULT '0' COMMENT '是否开启频控 0：不开启；1：开启',
  `flow_limit_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '频控策略id数据',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='安卓应用 版本库';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用初次下载事件日志';

-- ----------------------------
-- Records of upgrade_app_download_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_app_report_log
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_app_report_log`;
CREATE TABLE `upgrade_app_report_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `event_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '事件类型',
  `app_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用Key',
  `version_code` bigint NOT NULL DEFAULT '0' COMMENT '当前应用版本号',
  `dev_model_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备机型唯一标识',
  `dev_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备唯一标识',
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '事件发生时间',
  `launch_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '应用启动事件-应用启动时间',
  `code` bigint NOT NULL DEFAULT '0' COMMENT '事件-状态码',
  `download_version_code` bigint NOT NULL DEFAULT '0' COMMENT '应用升级-下载事件-应用版本号',
  `upgrade_version_code` bigint NOT NULL DEFAULT '0' COMMENT '应用升级-升级事件-应用版本号',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_event_time` (`company_id`,`event_type`,`timestamp`) USING BTREE,
  KEY `idx_event_type` (`event_type`) USING BTREE,
  KEY `idx_appkey_event_time` (`app_key`,`event_type`,`timestamp`) USING BTREE,
  KEY `idx_company_id` (`company_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用事件日志';

-- ----------------------------
-- Records of upgrade_app_report_log
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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用启动事件日志';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用升级下载事件日志';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用升级获取升级策略事件日志';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_timestamp` (`company_id`,`timestamp`) USING BTREE,
  KEY `idx_appkey_timestamp` (`app_key`,`timestamp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用升级升级事件日志';

-- ----------------------------
-- Records of upgrade_app_upgrade_upgrade_report_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_company_income
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_company_income`;
CREATE TABLE `upgrade_company_income` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `income_type` tinyint NOT NULL DEFAULT '0' COMMENT '收入类型：0 - 广告收入；',
  `income_amount` bigint NOT NULL DEFAULT '0' COMMENT '收入金额（单位分）',
  `income_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收入产生时间',
  `remark` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
  `status` int NOT NULL DEFAULT '0' COMMENT '收入状态：0 - 待结算；1 - 已结算；2 - 失效（如订单取消导致收入作废）',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '是否删除 0：正常；1：已删除',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='收入表';

-- ----------------------------
-- Records of upgrade_company_income
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_company_traffic_packet
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_company_traffic_packet`;
CREATE TABLE `upgrade_company_traffic_packet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `packet_id` bigint NOT NULL DEFAULT '0' COMMENT '流量包ID',
  `start_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始时间',
  `end_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '结束时间',
  `initial_size` bigint NOT NULL DEFAULT '0' COMMENT '初始流量大小(字节)',
  `remaining_size` bigint NOT NULL DEFAULT '0' COMMENT '剩余流量大小(字节)',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态: 1=有效, 0=已过期, 2=已用完',
  `exchange_time` timestamp NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '兑换流量包时间',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_company_status_time` (`company_id`,`status`,`start_time`,`end_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户购买的流量包记录';

-- ----------------------------
-- Records of upgrade_company_traffic_packet
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upgrade_company_traffic_usage_detail
-- ----------------------------
DROP TABLE IF EXISTS `upgrade_company_traffic_usage_detail`;
CREATE TABLE `upgrade_company_traffic_usage_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `company_id` bigint NOT NULL DEFAULT '0' COMMENT '公司ID',
  `app_download_report_id` bigint NOT NULL DEFAULT '0' COMMENT '下载记录ID',
  `company_traffic_packet_id` bigint NOT NULL DEFAULT '0' COMMENT '流量包记录ID',
  `used_size` bigint NOT NULL DEFAULT '0' COMMENT '使用流量(字节)',
  `used_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '使用时间',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='流量使用明细表';

-- ----------------------------
-- Records of upgrade_company_traffic_usage_detail
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置 版本库';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备分组';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备分组关系';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备机型表';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备管理';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='electron应用 升级任务灰度策略表；';

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
  `sha512` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '生成的sha512',
  `install_cloud_file_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '云文件id',
  `install_sha512` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '安装包 生成的sha512',
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_upgrade_file_github_file_id` (`file_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='github文件下载地址';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文件应用 版本库';

-- ----------------------------
-- Records of upgrade_file_version
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tauri_active_version` (`tauri_id`,`is_del`,`version_code`,`create_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tauri应用 版本库';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='流量包表';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用';

-- ----------------------------
-- Records of upgrade_url
-- ----------------------------
BEGIN;
INSERT INTO `upgrade_url` (`id`, `company_id`, `key`, `name`, `description`, `is_del`, `create_at`, `update_at`) VALUES (1, 1, 'WT2DZ4vJy9H30VPlM6AqjA', '测试url应用1', '', 0, '2025-09-26 06:08:17', '2025-09-26 06:08:17');
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 升级任务灰度策略表；';

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='url应用 版本库';

-- ----------------------------
-- Records of upgrade_url_version
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_positions
-- ----------------------------
DROP TABLE IF EXISTS `user_positions`;
CREATE TABLE `user_positions` (
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `position_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`position_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user_positions
-- ----------------------------
BEGIN;
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0194ef34-a6d1-728e-a70d-362f5035aab7', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0194f2c8-9ea7-76d2-9e3e-985c51e94b0d', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0194f3b7-5198-7e59-8f64-6fb42ee129cd', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019536d9-fbbe-7930-a7f3-287b00b754c9', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019536df-5075-7930-9fcb-f9aa1bf959e8', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019536f6-ec6c-7f53-aa23-b42fc71c149f', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019559ea-16de-7bb8-a652-c7af6e98a4ec', 0);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01955ac4-6ffa-7bb8-a4b7-a87bd2d3a561', 0);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01955ace-2bd8-7bb8-b373-d9f227fe8d98', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01962b48-b5fc-7bb8-84c8-f4b4d1db97b4', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01967682-ad93-7bb8-9499-e80886194fe1', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196b320-ab43-7bb8-9ec9-9a3af4949aec', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196d1ab-da36-7bb8-9290-0a73f609a790', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196d7e0-3752-7bb8-ade6-25c8bffcbd5d', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196d7e0-a0ed-7bb8-96b5-476631433f93', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196e7ba-0ddf-7bb8-a386-2665a2033060', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0196f1f0-7a1a-7bb8-8eda-65e7ce55afd9', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019769f9-185a-714f-82b9-bfe75880f729', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01979065-a792-714f-833e-f90315058e04', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01979bd1-1186-714f-9778-a388a32771d7', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01979ded-b714-714f-ab6d-c74eddb0e0d4', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01979f71-fb7f-714f-936d-862967640581', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197a139-e5cc-714f-8484-a6c3729cc6d8', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197ab9c-3061-714f-8110-2da85081b734', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197aff0-9e78-714f-b3a0-9af2461ba20f', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197b960-5e9d-714f-8510-6e35fd521c7c', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197becf-efaf-714f-aae0-ed9def6030bd', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197d885-a35d-714f-9aa1-1d51e33e561b', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197e84e-0ed4-714f-8c6c-4c05a7a5c803', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197e897-fc4a-714f-8af7-6965e9233752', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197ee15-1993-714f-8f61-021d78268403', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197eebf-6263-714f-869b-60942a59d214', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0197f810-c84b-714f-b40f-6e866f954752', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019808d2-9cca-714f-b02e-4af00e1b20d5', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01980bf9-f497-70a1-94ca-061912abb7db', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01980c1d-9472-70a1-a4c0-f41785a83451', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01980d23-558d-70a1-bff6-a870b58f4cef', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01981c72-08fa-70a1-aae4-86b5f2784541', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198220d-a0df-70a1-baab-649bb23d5e75', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019825e7-4d9a-70a1-8453-52447af754c4', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198263e-3fe1-70a1-b566-3f8563fd189f', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01983ff6-00c3-70a1-a79f-fb19c8204d0c', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019849ff-56e0-70a1-a25a-5fe2466f9c80', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01984eb2-c3c0-70a1-bf0f-d31399078b9d', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01984eca-5c08-70a1-afd8-d2f9638c27cb', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01985020-7a41-70a1-9311-46b8b5dc553c', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198505b-8e31-70a1-9f80-ab483d54a50d', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01985947-c27e-70a1-b08d-692a65b5ff1a', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019882a3-b194-70a1-87bb-c8b090166c04', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01988906-1ba7-70a1-b394-e7a95d8a374e', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01988e3f-7aec-70a1-bdbd-c3aa220ede86', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01989293-d7b8-7dce-8711-e8539f3b6a06', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198ac2a-c4c1-7dce-a844-af0a65994280', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198acf7-0d1f-7dce-b24d-c623f14e6e35', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198b0a4-d747-7dce-96f9-4f4c1325ba1b', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198b0e1-41d1-7dce-9803-5503eccc5de8', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198c6c2-f33b-7dce-8297-d4f5f8251133', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198eb66-fac9-7dce-85c4-1f798e6f9a51', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198ee4b-9fb0-7dce-a3aa-4826de862abd', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198f01c-23f5-7dce-91f9-ca308a89eeae', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198f906-6605-7dce-bbfa-f7448d84eb1a', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198f9dd-41fd-7dce-9922-b5c6225f32fa', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0198fbbe-a850-7dce-b169-495ed025e397', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0199100e-8050-7dce-9fb3-74de4ce74231', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('0199135d-1497-7dce-8933-906ba7a6c3a8', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('019919d2-3ef5-7dce-a506-7aefdd850245', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01993645-37ee-7dce-b25b-a6d830507946', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01993e19-730a-7dce-9f06-366a8f96252c', 1);
INSERT INTO `user_positions` (`user_id`, `position_id`) VALUES ('01997157-69f0-7850-9e67-049d89fcac79', 1);
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0194ef34-a6d1-728e-a70d-362f5035aab7', 1);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0194f2c8-9ea7-76d2-9e3e-985c51e94b0d', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0194f3b7-5198-7e59-8f64-6fb42ee129cd', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019536d9-fbbe-7930-a7f3-287b00b754c9', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019536df-5075-7930-9fcb-f9aa1bf959e8', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019536f6-ec6c-7f53-aa23-b42fc71c149f', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019559ea-16de-7bb8-a652-c7af6e98a4ec', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01955ac4-6ffa-7bb8-a4b7-a87bd2d3a561', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01955ace-2bd8-7bb8-b373-d9f227fe8d98', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01962b48-b5fc-7bb8-84c8-f4b4d1db97b4', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01967682-ad93-7bb8-9499-e80886194fe1', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196b320-ab43-7bb8-9ec9-9a3af4949aec', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196d1ab-da36-7bb8-9290-0a73f609a790', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196d7e0-3752-7bb8-ade6-25c8bffcbd5d', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196d7e0-a0ed-7bb8-96b5-476631433f93', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196e7ba-0ddf-7bb8-a386-2665a2033060', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0196f1f0-7a1a-7bb8-8eda-65e7ce55afd9', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019769f9-185a-714f-82b9-bfe75880f729', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01979065-a792-714f-833e-f90315058e04', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01979bd1-1186-714f-9778-a388a32771d7', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01979ded-b714-714f-ab6d-c74eddb0e0d4', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01979f71-fb7f-714f-936d-862967640581', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197a139-e5cc-714f-8484-a6c3729cc6d8', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197ab9c-3061-714f-8110-2da85081b734', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197aff0-9e78-714f-b3a0-9af2461ba20f', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197b960-5e9d-714f-8510-6e35fd521c7c', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197becf-efaf-714f-aae0-ed9def6030bd', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197d885-a35d-714f-9aa1-1d51e33e561b', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197e84e-0ed4-714f-8c6c-4c05a7a5c803', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197e897-fc4a-714f-8af7-6965e9233752', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197ee15-1993-714f-8f61-021d78268403', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197eebf-6263-714f-869b-60942a59d214', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0197f810-c84b-714f-b40f-6e866f954752', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019808d2-9cca-714f-b02e-4af00e1b20d5', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01980bf9-f497-70a1-94ca-061912abb7db', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01980c1d-9472-70a1-a4c0-f41785a83451', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01980d23-558d-70a1-bff6-a870b58f4cef', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01981c72-08fa-70a1-aae4-86b5f2784541', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198220d-a0df-70a1-baab-649bb23d5e75', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019825e7-4d9a-70a1-8453-52447af754c4', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198263e-3fe1-70a1-b566-3f8563fd189f', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01983ff6-00c3-70a1-a79f-fb19c8204d0c', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019849ff-56e0-70a1-a25a-5fe2466f9c80', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01984eb2-c3c0-70a1-bf0f-d31399078b9d', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01984eca-5c08-70a1-afd8-d2f9638c27cb', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01985020-7a41-70a1-9311-46b8b5dc553c', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198505b-8e31-70a1-9f80-ab483d54a50d', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01985947-c27e-70a1-b08d-692a65b5ff1a', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019882a3-b194-70a1-87bb-c8b090166c04', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01988906-1ba7-70a1-b394-e7a95d8a374e', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01988e3f-7aec-70a1-bdbd-c3aa220ede86', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01989293-d7b8-7dce-8711-e8539f3b6a06', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198ac2a-c4c1-7dce-a844-af0a65994280', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198acf7-0d1f-7dce-b24d-c623f14e6e35', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198b0a4-d747-7dce-96f9-4f4c1325ba1b', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198b0e1-41d1-7dce-9803-5503eccc5de8', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198c6c2-f33b-7dce-8297-d4f5f8251133', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198eb66-fac9-7dce-85c4-1f798e6f9a51', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198ee4b-9fb0-7dce-a3aa-4826de862abd', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198f01c-23f5-7dce-91f9-ca308a89eeae', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198f906-6605-7dce-bbfa-f7448d84eb1a', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198f9dd-41fd-7dce-9922-b5c6225f32fa', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0198fbbe-a850-7dce-b169-495ed025e397', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0199100e-8050-7dce-9fb3-74de4ce74231', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('0199135d-1497-7dce-8933-906ba7a6c3a8', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('019919d2-3ef5-7dce-a506-7aefdd850245', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01993645-37ee-7dce-b25b-a6d830507946', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01993e19-730a-7dce-9f06-366a8f96252c', 2);
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES ('01997157-69f0-7850-9e67-049d89fcac79', 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
