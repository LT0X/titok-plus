/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : titok_plus_user

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 13/05/2023 16:55:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_favor_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_favor_videos`;
CREATE TABLE `user_favor_videos`  (
  `user_info_id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  PRIMARY KEY (`user_info_id`, `video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_follows
-- ----------------------------
DROP TABLE IF EXISTS `user_follows`;
CREATE TABLE `user_follows`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `to_user_id` bigint(20) NOT NULL COMMENT '关注的用户',
  `action_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '默认关注为1，取消关注为2',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '关注表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_infos
-- ----------------------------
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(70) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `follow_count` bigint(20) NULL DEFAULT NULL,
  `follower_count` bigint(20) NULL DEFAULT NULL,
  `avatar` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
  `background_image` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
  `signature` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `total_favorited` bigint(20) NULL DEFAULT NULL,
  `work_count` bigint(20) NULL DEFAULT NULL,
  `favorite_count` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 70 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_logins
-- ----------------------------
DROP TABLE IF EXISTS `user_logins`;
CREATE TABLE `user_logins`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_info_id` bigint(20) NULL DEFAULT NULL,
  `username` varchar(70) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `password` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`, `username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 70 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
