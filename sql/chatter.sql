/*
Navicat MySQL Data Transfer

Source Server         : root
Source Server Version : 50528
Source Host           : localhost:3306
Source Database       : chatter

Target Server Type    : MYSQL
Target Server Version : 50528
File Encoding         : 65001

Date: 2015-01-06 10:52:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for group_msg
-- ----------------------------
DROP TABLE IF EXISTS `group_msg`;
CREATE TABLE `group_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gid` int(10) unsigned NOT NULL,
  `mid` bigint(20) unsigned NOT NULL,
  `ttl` bigint(20) NOT NULL,
  `msg` blob NOT NULL,
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ux_group_msg_1` (`gid`,`mid`),
  KEY `ix_group_msg_1` (`ttl`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group_msg
-- ----------------------------

-- ----------------------------
-- Table structure for infraction
-- ----------------------------
DROP TABLE IF EXISTS `infraction`;
CREATE TABLE `infraction` (
  `id` varchar(32) NOT NULL,
  `user_id` varchar(32) DEFAULT NULL,
  `qun_id` varchar(32) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of infraction
-- ----------------------------

-- ----------------------------
-- Table structure for org
-- ----------------------------
DROP TABLE IF EXISTS `org`;
CREATE TABLE `org` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_name` varchar(255) NOT NULL,
  `parent_id` varchar(255) NOT NULL,
  `location` varchar(255) NOT NULL,
  `tenant_id` varchar(255) NOT NULL,
  `sort` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of org
-- ----------------------------

-- ----------------------------
-- Table structure for private_msg
-- ----------------------------
DROP TABLE IF EXISTS `private_msg`;
CREATE TABLE `private_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `skey` varchar(64) NOT NULL,
  `mid` bigint(20) unsigned NOT NULL,
  `ttl` bigint(20) NOT NULL,
  `msg` blob NOT NULL,
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ux_private_msg_1` (`skey`,`mid`),
  KEY `ix_private_msg_1` (`ttl`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of private_msg
-- ----------------------------

-- ----------------------------
-- Table structure for public_msg
-- ----------------------------
DROP TABLE IF EXISTS `public_msg`;
CREATE TABLE `public_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL,
  `ttl` bigint(20) NOT NULL,
  `msg` blob NOT NULL,
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `ix_public_msg_1` (`ttl`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of public_msg
-- ----------------------------

-- ----------------------------
-- Table structure for public_msg_log
-- ----------------------------
DROP TABLE IF EXISTS `public_msg_log`;
CREATE TABLE `public_msg_log` (
  `mid` bigint(20) unsigned NOT NULL,
  `stime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `ftime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  UNIQUE KEY `ux_public_msg_log_1` (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of public_msg_log
-- ----------------------------

-- ----------------------------
-- Table structure for qun
-- ----------------------------
DROP TABLE IF EXISTS `qun`;
CREATE TABLE `qun` (
  `id` varchar(64) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `type_id` varchar(64) DEFAULT NULL,
  `creator_id` varchar(64) DEFAULT NULL,
  `liveness` int(11) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `max_member` int(11) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of qun
-- ----------------------------
INSERT INTO `qun` VALUES ('1', '测试群1', '/static/images/psb.jpg', '1', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '1', '111', '100', '2014-12-21 15:39:54', '2014-12-23 15:39:57');

-- ----------------------------
-- Table structure for qun_type
-- ----------------------------
DROP TABLE IF EXISTS `qun_type`;
CREATE TABLE `qun_type` (
  `id` varchar(32) NOT NULL,
  `parent_id` varchar(32) DEFAULT NULL,
  `name` varchar(128) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of qun_type
-- ----------------------------

-- ----------------------------
-- Table structure for qun_user
-- ----------------------------
DROP TABLE IF EXISTS `qun_user`;
CREATE TABLE `qun_user` (
  `id` varchar(64) NOT NULL,
  `qun_id` varchar(64) DEFAULT NULL,
  `user_id` varchar(64) DEFAULT NULL,
  `role` int(1) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of qun_user
-- ----------------------------
INSERT INTO `qun_user` VALUES ('1', '1', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '1', '1', '2014-12-22 15:49:14', '2014-12-22 15:49:16');

-- ----------------------------
-- Table structure for tenant
-- ----------------------------
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant` (
  `id` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` int(11) NOT NULL,
  `customer_id` varchar(255) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tenant
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(64) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `nick_name` varchar(50) DEFAULT NULL,
  `signature` varchar(255) DEFAULT NULL,
  `avatar` varchar(45) DEFAULT NULL,
  `status` int(1) DEFAULT '0' COMMENT '验证/非验证',
  `password` varchar(32) DEFAULT NULL,
  `sex` int(1) DEFAULT NULL COMMENT '0表示男1表示女',
  `level` int(3) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `occupation` varchar(64) DEFAULT NULL,
  `url` varchar(128) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('03975bfc-7669-41e2-bc8f-80d03f469d63', 'a', '张三', '个性签名个性签名个性签名', '', '0', 'a', '0', '0', 'dizhi', '', '', '', '', '2014-12-19 15:55:32', '2014-12-19 15:55:32');
INSERT INTO `user` VALUES ('0aedbd65-5d21-436e-a8f6-2f58d393a46d', 'q', '李旭东', '好好学习，天天向上', '1', '1', 'q', '0', '0', '地址', '手机', '邮箱', '职业', 'url', '2014-12-19 14:09:50', '2014-12-19 14:09:50');
INSERT INTO `user` VALUES ('0e6b7c3f-ae1d-477b-9ab7-fcca2e286c49', 'z', '李四', '', '', '0', 'z', '0', '0', '', '', '', '', '', '2014-12-22 10:00:31', '2014-12-22 10:00:31');
INSERT INTO `user` VALUES ('179edfaf-1b3a-44e0-8b96-d5035cdcd145', 'w', '杨斌', '个性签名个性签名213123132签名', '', '0', 'w', '0', '0', '', '', '', '', '', '2014-12-22 03:00:24', '2014-12-22 03:00:24');
INSERT INTO `user` VALUES ('3a7dce46-e0fb-4d08-b305-3e7856182688', '1', '', '111111111111111签名', '', '0', '1', '0', '0', '', '', '', '', '', '2014-12-19 15:53:14', '2014-12-19 15:53:14');
INSERT INTO `user` VALUES ('56e2386e-02ab-404f-9180-70a2ae6e936d', '1111', '', '个性签名个性签名个性2222222222名', '', '0', '11111111', '0', '0', '', '', '', '', '', '2014-12-22 03:00:16', '2014-12-22 03:00:16');
INSERT INTO `user` VALUES ('5c6a2426-20f4-4a08-bd99-c81b77372d51', 'www', '', '个性签名个性签22222', '', '0', 'wwww', '0', '0', '', '', '', '', '', '2014-12-22 02:58:58', '2014-12-22 02:58:58');
INSERT INTO `user` VALUES ('6971780c-260c-402f-96df-936428a4284c', '2222', '', '个性签名个性签名个性签名', '', '0', '222222222', '0', '0', '', '', '', '', '', '2014-12-22 03:00:43', '2014-12-22 03:00:43');
INSERT INTO `user` VALUES ('85f44187-f6a3-4bbc-b619-dc000308f62e', '11', '', '个性签名个性签名', '', '0', '111', '0', '0', '', '', '', '', '', '2014-12-19 15:57:49', '2014-12-19 15:57:49');
INSERT INTO `user` VALUES ('dda873b4-64ad-4777-a75d-fe1fc6a46088', 'lixudong', '', '个性签名个性签名个性签名个性签名个性签名', '', '0', '111111', '0', '0', '', '', '', '', '', '2014-12-20 03:11:12', '2014-12-20 03:11:12');

-- ----------------------------
-- Table structure for user_user
-- ----------------------------
DROP TABLE IF EXISTS `user_user`;
CREATE TABLE `user_user` (
  `id` varchar(64) NOT NULL,
  `from_user_id` varchar(64) DEFAULT NULL,
  `to_user_id` varchar(64) DEFAULT NULL,
  `remark_name` varchar(45) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_user
-- ----------------------------
INSERT INTO `user_user` VALUES ('03975bfc-7669-41e2-bc8f-80d03f469d63', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '179edfaf-1b3a-44e0-8b96-d5035cdcd145', '', '1');
INSERT INTO `user_user` VALUES ('03975bfc-7669-41e2-bc8f-80d03f469d64', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '03975bfc-7669-41e2-bc8f-80d03f469d63', null, '1');
INSERT INTO `user_user` VALUES ('2', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '0e6b7c3f-ae1d-477b-9ab7-fcca2e286c49', null, '1');
INSERT INTO `user_user` VALUES ('3', '0e6b7c3f-ae1d-477b-9ab7-fcca2e286c49', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', null, '1');
INSERT INTO `user_user` VALUES ('4', '03975bfc-7669-41e2-bc8f-80d03f469d63', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', null, '1');
INSERT INTO `user_user` VALUES ('5', '179edfaf-1b3a-44e0-8b96-d5035cdcd145', '0aedbd65-5d21-436e-a8f6-2f58d393a46d', '', '1');
