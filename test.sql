/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 05/07/2021 06:43:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for u_db_artical
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical`;
CREATE TABLE `u_db_artical`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `view` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `role_id` bigint(20) NOT NULL,
  `image_id` int(11) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for u_db_artical_tags_copy
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical_tags_copy`;
CREATE TABLE `u_db_artical_tags_copy`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `artical_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `tag_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_artical_tags_copy
-- ----------------------------
INSERT INTO `u_db_artical_tags_copy` VALUES (5, '24', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES (6, '25', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES (7, '26', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (8, '26', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (9, '27', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (10, '27', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (11, '28', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (12, '28', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (13, '94', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (14, '94', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (15, '186', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (16, '186', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (17, '405', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (18, '405', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (19, '406', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES (20, '407', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (21, '407', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (22, '408', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (23, '408', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (24, '409', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (25, '409', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (26, '410', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (27, '410', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (28, '411', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (29, '411', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (30, '412', 'gold');
INSERT INTO `u_db_artical_tags_copy` VALUES (31, '412', 'cyan');
INSERT INTO `u_db_artical_tags_copy` VALUES (32, '413', 'gold');
INSERT INTO `u_db_artical_tags_copy` VALUES (33, '413', 'cyan');
INSERT INTO `u_db_artical_tags_copy` VALUES (34, '414', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (35, '414', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (36, '415', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (37, '415', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (38, '416', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (39, '416', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (40, '417', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (41, '417', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (42, '418', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (43, '418', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (44, '419', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (45, '419', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES (46, '420', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES (47, '420', 'JAVA');

-- ----------------------------
-- Table structure for u_db_artical_u_db_tags
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical_u_db_tags`;
CREATE TABLE `u_db_artical_u_db_tags`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `u_db_artical_id` int(11) NOT NULL,
  `u_db_tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_artical_u_db_tags
-- ----------------------------
INSERT INTO `u_db_artical_u_db_tags` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for u_db_attendance
-- ----------------------------
DROP TABLE IF EXISTS `u_db_attendance`;
CREATE TABLE `u_db_attendance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `mouth` tinyint(3) NULL DEFAULT NULL COMMENT '月份',
  `attendance` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '每月工时',
  `create_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_attendance
-- ----------------------------
INSERT INTO `u_db_attendance` VALUES (2, 1, 7, '{\"25\":8,\"26\":8,\"27\":8,\"28\":8,\"29\":8,\"30\":8}', '2021-07-04 22:10:50');

-- ----------------------------
-- Table structure for u_db_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `u_db_dictionary`;
CREATE TABLE `u_db_dictionary`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '',
  `describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '',
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_dictionary
-- ----------------------------
INSERT INTO `u_db_dictionary` VALUES (2, 'JAVA', 'tag', '2', '文章分类', '1');
INSERT INTO `u_db_dictionary` VALUES (3, '人力资源', 'department', '100001', '部门名称', '1');
INSERT INTO `u_db_dictionary` VALUES (4, '挨罚部', 'department', '100002', '部门名称', '1');

-- ----------------------------
-- Table structure for u_db_image
-- ----------------------------
DROP TABLE IF EXISTS `u_db_image`;
CREATE TABLE `u_db_image`  (
  `id` int(255) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_image
-- ----------------------------
INSERT INTO `u_db_image` VALUES (NULL, 'e1164a4edec9781df446d2cf7577326d.ico', 'artical', 'http://localhost:8080/static/upload/2021/07/04/e1164a4edec9781df446d2cf7577326d.ico', '2021-07-03 17:03:55', '2021-07-03 17:03:55');
INSERT INTO `u_db_image` VALUES (NULL, 'be2f7991a1f1fbf887ad352edb914560.ico', 'artical', 'http://localhost:8080/static/upload/2021/07/04/be2f7991a1f1fbf887ad352edb914560.ico', '2021-07-03 17:06:53', '2021-07-03 17:06:53');
INSERT INTO `u_db_image` VALUES (NULL, '5ff7a52f4e18590733a6922f672ee2e4.ico', 'artical', 'http://localhost:8080/static/upload/2021/07/04/5ff7a52f4e18590733a6922f672ee2e4.ico', '2021-07-03 17:07:09', '2021-07-03 17:07:09');
INSERT INTO `u_db_image` VALUES (NULL, '78c06f4cbd769b41bc99f101fc5fef61.xls', 'artical', 'http://localhost:8080/static/upload/2021/07/05/78c06f4cbd769b41bc99f101fc5fef61.xls', '2021-07-04 20:55:02', '2021-07-04 20:55:02');
INSERT INTO `u_db_image` VALUES (NULL, '553511e18fc94db83c7778f94477ccc5.xls', 'artical', 'http://localhost:8080/static/upload/2021/07/05/553511e18fc94db83c7778f94477ccc5.xls', '2021-07-04 21:05:54', '2021-07-04 21:05:54');

-- ----------------------------
-- Table structure for u_db_leave
-- ----------------------------
DROP TABLE IF EXISTS `u_db_leave`;
CREATE TABLE `u_db_leave`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `type` int(11) NOT NULL DEFAULT 0,
  `reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  `result` int(11) NOT NULL DEFAULT 0,
  `start_time` datetime(0) NOT NULL,
  `end_time` datetime(0) NOT NULL,
  `create_time` datetime(0) NOT NULL,
  `deal_user_id` int(11) NOT NULL,
  `status` tinyint(3) NULL DEFAULT NULL,
  `deal_reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_leave
-- ----------------------------
INSERT INTO `u_db_leave` VALUES (1, 1, 1, '5', 1, '2021-07-01 23:48:34', '2021-07-16 23:48:38', '2021-07-01 23:48:26', 1, 1, '');
INSERT INTO `u_db_leave` VALUES (2, 1, 1, '213123', 0, '2021-07-04 10:47:54', '2021-07-21 10:47:56', '2021-07-04 18:47:58', 1, 0, '');
INSERT INTO `u_db_leave` VALUES (3, 1, 1, '213123', 0, '2021-07-03 16:00:00', '2021-07-20 16:00:00', '2021-07-04 20:10:09', 1, 2, '');
INSERT INTO `u_db_leave` VALUES (4, 1, 1, '213123', 0, '2021-06-29 16:00:00', '2021-07-18 16:00:00', '2021-07-04 20:13:16', 1, 1, '213123');

-- ----------------------------
-- Table structure for u_db_menu
-- ----------------------------
DROP TABLE IF EXISTS `u_db_menu`;
CREATE TABLE `u_db_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `parent_id` int(11) NOT NULL DEFAULT 0,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_menu
-- ----------------------------
INSERT INTO `u_db_menu` VALUES (1, '首页', '1', 0, '', 0);
INSERT INTO `u_db_menu` VALUES (2, '文章', '2', 1, '/admin/artical', 0);
INSERT INTO `u_db_menu` VALUES (3, '部门管理', '3', 1, '/admin/department/list', 0);
INSERT INTO `u_db_menu` VALUES (4, '员工管理', '4', 1, '/admin/user/list', 0);
INSERT INTO `u_db_menu` VALUES (5, '考勤管理', '5', 1, '', 0);
INSERT INTO `u_db_menu` VALUES (6, '工时管理', '6', 5, '/admin/attendance/list', 0);
INSERT INTO `u_db_menu` VALUES (7, '请假审核', '7', 5, '/admin/attendance/review', 0);
INSERT INTO `u_db_menu` VALUES (8, '表单控件', '8', 1, '/admin/datatable', 0);
INSERT INTO `u_db_menu` VALUES (9, '字段控制', '9', 1, '/admin/column', 0);
INSERT INTO `u_db_menu` VALUES (10, '请假申请', '10', 5, '/admin/attendance/apply', 0);
INSERT INTO `u_db_menu` VALUES (11, '表单管理', '11', 1, '/admin/form', 0);

-- ----------------------------
-- Table structure for u_db_role
-- ----------------------------
DROP TABLE IF EXISTS `u_db_role`;
CREATE TABLE `u_db_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_role
-- ----------------------------
INSERT INTO `u_db_role` VALUES (1, '超管');
INSERT INTO `u_db_role` VALUES (2, '部门管理员');
INSERT INTO `u_db_role` VALUES (3, '员工');

-- ----------------------------
-- Table structure for u_db_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `u_db_role_menus`;
CREATE TABLE `u_db_role_menus`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT 0,
  `menu_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_role_menus
-- ----------------------------
INSERT INTO `u_db_role_menus` VALUES (1, 1, 1);
INSERT INTO `u_db_role_menus` VALUES (2, 1, 2);
INSERT INTO `u_db_role_menus` VALUES (3, 1, 3);
INSERT INTO `u_db_role_menus` VALUES (4, 1, 4);
INSERT INTO `u_db_role_menus` VALUES (5, 1, 5);
INSERT INTO `u_db_role_menus` VALUES (6, 1, 6);
INSERT INTO `u_db_role_menus` VALUES (7, 1, 7);
INSERT INTO `u_db_role_menus` VALUES (8, 1, 8);
INSERT INTO `u_db_role_menus` VALUES (9, 1, 9);
INSERT INTO `u_db_role_menus` VALUES (10, 1, 10);
INSERT INTO `u_db_role_menus` VALUES (11, 1, 11);

-- ----------------------------
-- Table structure for u_db_tag
-- ----------------------------
DROP TABLE IF EXISTS `u_db_tag`;
CREATE TABLE `u_db_tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `type` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  `artical_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_tag
-- ----------------------------
INSERT INTO `u_db_tag` VALUES (1, 'php', 'artical', 0);
INSERT INTO `u_db_tag` VALUES (2, '人力资源部', 'organization', 0);

-- ----------------------------
-- Table structure for u_db_user
-- ----------------------------
DROP TABLE IF EXISTS `u_db_user`;
CREATE TABLE `u_db_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `gender` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `age` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '',
  `role` int(11) NULL DEFAULT 0,
  `salf` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  `role_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of u_db_user
-- ----------------------------
INSERT INTO `u_db_user` VALUES (1, 'admin', '$2a$10$a66cbFqCJR5fEymRjnci2ecfAa4NTF2Rk6GcHmmoITDdZCr9zmn.2', '1', '11', '1', '1', 1, 'efcc', 1, 3);
INSERT INTO `u_db_user` VALUES (3, 'test', 'test', '1', '12', '12315', '12', 1, 'asdf', 2, 2);

SET FOREIGN_KEY_CHECKS = 1;
