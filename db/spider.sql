/*
Navicat MariaDB Data Transfer

Source Server         : 3306
Source Server Version : 100209
Source Host           : localhost:3306
Source Database       : crawl

Target Server Type    : MariaDB
Target Server Version : 100209
File Encoding         : 65001

Date: 2020-04-11 23:13:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `Id` int(11) NOT NULL,
  `Name` char(255) NOT NULL DEFAULT '',
  `Type` int(11) NOT NULL,
  `Author` char(255) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name` (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of books
-- ----------------------------

-- ----------------------------
-- Table structure for chapters
-- ----------------------------
DROP TABLE IF EXISTS `chapters`;
CREATE TABLE `chapters` (
  `Id` int(11) NOT NULL,
  `Title` char(255) NOT NULL,
  `Content` longtext NOT NULL DEFAULT '',
  `BookId` int(11) NOT NULL,
  `Order` int(11) NOT NULL,
  `Url` char(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chapters
-- ----------------------------
