/*
Navicat MariaDB Data Transfer

Source Server         : 3306
Source Server Version : 100209
Source Host           : localhost:3306
Source Database       : spider

Target Server Type    : MariaDB
Target Server Version : 100209
File Encoding         : 65001

Date: 2020-04-12 00:38:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(255) NOT NULL DEFAULT '',
  `type` int(11) NOT NULL,
  `author` char(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for chapters
-- ----------------------------
DROP TABLE IF EXISTS `chapters`;
CREATE TABLE `chapters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` char(255) NOT NULL,
  `content` longtext NOT NULL,
  `book_id` int(11) NOT NULL,
  `order` int(11) NOT NULL,
  `url` char(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
