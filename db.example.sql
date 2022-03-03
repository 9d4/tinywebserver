-- Adminer 4.8.1 MySQL 8.0.27 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(2048) COLLATE utf8_bin DEFAULT NULL,
  `done` tinyint DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

INSERT INTO `todos` (`id`, `content`, `done`, `created_at`, `updated_at`) VALUES
(29,	'09 11',	0,	'2021-12-15 02:11:57',	'2021-12-15 02:32:54'),
(30,	'09 12',	0,	'2021-12-15 02:12:02',	'2021-12-15 02:34:50');

-- 2021-12-15 03:53:17