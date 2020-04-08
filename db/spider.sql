/*
Navicat MySQL Data Transfer

Source Server         : localhost mysql 3306
Source Server Version : 80012
Source Host           : 127.0.0.1:3306
Source Database       : spider

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2020-04-08 16:15:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for spider_books
-- ----------------------------
DROP TABLE IF EXISTS `spider_books`;
CREATE TABLE `spider_books` (
  `Id` int(11) NOT NULL,
  `Name` char(255) DEFAULT NULL,
  `Type` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of spider_books
-- ----------------------------

-- ----------------------------
-- Table structure for spider_chapters
-- ----------------------------
DROP TABLE IF EXISTS `spider_chapters`;
CREATE TABLE `spider_chapters` (
  `Id` int(11) NOT NULL,
  `Content` int(11) DEFAULT NULL,
  `BookId` varchar(10240) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of spider_chapters
-- ----------------------------
