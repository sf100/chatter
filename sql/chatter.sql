/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50624
Source Host           : localhost:3306
Source Database       : chatter

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2015-08-18 15:18:03
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
-- Table structure for history_msg
-- ----------------------------
DROP TABLE IF EXISTS `history_msg`;
CREATE TABLE `history_msg` (
  `id` int(11) NOT NULL,
  `source_id` varchar(64) DEFAULT NULL,
  `target_id` varchar(64) DEFAULT NULL,
  `content` varchar(500) DEFAULT NULL,
  `msg_type` int(11) DEFAULT NULL,
  `source_type` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `Created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
