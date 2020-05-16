/*
 Navicat Premium Data Transfer

 Source Server         : 虚拟机216
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 192.168.0.216:3306
 Source Schema         : apidoc

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 11/05/2020 18:28:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_detail
-- ----------------------------
DROP TABLE IF EXISTS `api_detail`;
CREATE TABLE `api_detail`  (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `gid` int(20) NOT NULL COMMENT '组id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `uri` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `param` json NULL,
  `return_value` json NULL,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gid`(`gid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_group
-- ----------------------------
DROP TABLE IF EXISTS `api_group`;
CREATE TABLE `api_group`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '项目id',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Triggers structure for table api_group
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_api`;
delimiter ;;
CREATE TRIGGER `delete_api` AFTER DELETE ON `api_group` FOR EACH ROW begin
set @id=OLD.id;
delete from function where aid=@id;
end
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table project
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_project`;
delimiter ;;
CREATE TRIGGER `delete_project` AFTER DELETE ON `project` FOR EACH ROW begin
set @id=OLD.id;
delete from api where pid=@id;
delete from project_address where pid=@id;
end
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
