/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : apidoc

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-05-16 16:02:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for api_detail
-- ----------------------------
DROP TABLE IF EXISTS `api_detail`;
CREATE TABLE `api_detail` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `gid` int(20) NOT NULL COMMENT '组id',
  `name` varchar(255) NOT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `param` json DEFAULT NULL,
  `return_value` json DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `gid` (`gid`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of api_detail
-- ----------------------------

-- ----------------------------
-- Table structure for api_group
-- ----------------------------
DROP TABLE IF EXISTS `api_group`;
CREATE TABLE `api_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '项目id',
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of api_group
-- ----------------------------

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of project
-- ----------------------------

-- ----------------------------
-- Table structure for whitelist
-- ----------------------------
DROP TABLE IF EXISTS `whitelist`;
CREATE TABLE `whitelist` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `op` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of whitelist
-- ----------------------------
INSERT INTO `whitelist` VALUES ('1', '127.0.0.1', 'Localhost', '0');
DROP TRIGGER IF EXISTS `delete_api`;
DELIMITER ;;
CREATE TRIGGER `delete_api` AFTER DELETE ON `api_group` FOR EACH ROW begin
set @id=OLD.id;
delete from function where aid=@id;
end
;;
DELIMITER ;
DROP TRIGGER IF EXISTS `delete_project`;
DELIMITER ;;
CREATE TRIGGER `delete_project` AFTER DELETE ON `project` FOR EACH ROW begin
set @id=OLD.id;
delete from api where pid=@id;
delete from project_address where pid=@id;
end
;;
DELIMITER ;
