/*
Navicat MySQL Data Transfer

Source Server         : 本地数据
Source Server Version : 50725
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50725
File Encoding         : 65001

Date: 2020-12-22 17:07:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS `user_token`;
CREATE TABLE `user_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `rep_ip` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of user_token
-- ----------------------------

-- ----------------------------
-- Table structure for u_db_artical
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical`;
CREATE TABLE `u_db_artical` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `author` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `view` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `content` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT '0',
  `role_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=438 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_artical
-- ----------------------------
INSERT INTO `u_db_artical` VALUES ('1', 'fdasf', 'fadsf', '123', '123', '213132', '1');

-- ----------------------------
-- Table structure for u_db_artical_tags_copy
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical_tags_copy`;
CREATE TABLE `u_db_artical_tags_copy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `artical_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `tag_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_artical_tags_copy
-- ----------------------------
INSERT INTO `u_db_artical_tags_copy` VALUES ('5', '24', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES ('6', '25', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES ('7', '26', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('8', '26', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('9', '27', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('10', '27', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('11', '28', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('12', '28', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('13', '94', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('14', '94', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('15', '186', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('16', '186', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('17', '405', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('18', '405', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('19', '406', 'php');
INSERT INTO `u_db_artical_tags_copy` VALUES ('20', '407', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('21', '407', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('22', '408', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('23', '408', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('24', '409', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('25', '409', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('26', '410', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('27', '410', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('28', '411', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('29', '411', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('30', '412', 'gold');
INSERT INTO `u_db_artical_tags_copy` VALUES ('31', '412', 'cyan');
INSERT INTO `u_db_artical_tags_copy` VALUES ('32', '413', 'gold');
INSERT INTO `u_db_artical_tags_copy` VALUES ('33', '413', 'cyan');
INSERT INTO `u_db_artical_tags_copy` VALUES ('34', '414', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('35', '414', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('36', '415', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('37', '415', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('38', '416', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('39', '416', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('40', '417', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('41', '417', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('42', '418', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('43', '418', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('44', '419', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('45', '419', 'JAVA');
INSERT INTO `u_db_artical_tags_copy` VALUES ('46', '420', 'PHP');
INSERT INTO `u_db_artical_tags_copy` VALUES ('47', '420', 'JAVA');

-- ----------------------------
-- Table structure for u_db_artical_u_db_tags
-- ----------------------------
DROP TABLE IF EXISTS `u_db_artical_u_db_tags`;
CREATE TABLE `u_db_artical_u_db_tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `u_db_artical_id` int(11) NOT NULL,
  `u_db_tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_artical_u_db_tags
-- ----------------------------
INSERT INTO `u_db_artical_u_db_tags` VALUES ('1', '1', '1');

-- ----------------------------
-- Table structure for u_db_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `u_db_dictionary`;
CREATE TABLE `u_db_dictionary` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `type` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `value` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `describe` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `status` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_dictionary
-- ----------------------------
INSERT INTO `u_db_dictionary` VALUES ('1', 'PHP', 'tag', '1', '文章分类', '1');
INSERT INTO `u_db_dictionary` VALUES ('2', 'JAVA', 'tag', '2', '文章分类', '1');

-- ----------------------------
-- Table structure for u_db_menu
-- ----------------------------
DROP TABLE IF EXISTS `u_db_menu`;
CREATE TABLE `u_db_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `icon` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `path` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_menu
-- ----------------------------
INSERT INTO `u_db_menu` VALUES ('1', '首页', '1', '0', '', '0');
INSERT INTO `u_db_menu` VALUES ('2', '文章', '2', '1', '/admin/artical', '0');

-- ----------------------------
-- Table structure for u_db_role
-- ----------------------------
DROP TABLE IF EXISTS `u_db_role`;
CREATE TABLE `u_db_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_role
-- ----------------------------

-- ----------------------------
-- Table structure for u_db_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `u_db_role_menus`;
CREATE TABLE `u_db_role_menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT '0',
  `menu_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_role_menus
-- ----------------------------
INSERT INTO `u_db_role_menus` VALUES ('1', '1', '1');
INSERT INTO `u_db_role_menus` VALUES ('2', '1', '2');
INSERT INTO `u_db_role_menus` VALUES ('3', '1', '3');
INSERT INTO `u_db_role_menus` VALUES ('4', '1', '4');

-- ----------------------------
-- Table structure for u_db_tag
-- ----------------------------
DROP TABLE IF EXISTS `u_db_tag`;
CREATE TABLE `u_db_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_tag
-- ----------------------------
INSERT INTO `u_db_tag` VALUES ('1', 'php');

-- ----------------------------
-- Table structure for u_db_user
-- ----------------------------
DROP TABLE IF EXISTS `u_db_user`;
CREATE TABLE `u_db_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `gender` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `age` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `address` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `role` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of u_db_user
-- ----------------------------
INSERT INTO `u_db_user` VALUES ('1', 'admin', 'admin', '1', '11', '1', '1', '1');
