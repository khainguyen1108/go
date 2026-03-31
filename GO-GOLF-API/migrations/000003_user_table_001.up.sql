/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `avatar` varchar(500) DEFAULT NULL,
  `nick_name` varchar(20) NOT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0: male, 1: female',
  `birthday` date DEFAULT NULL,
  `sns_type` tinyint(1) DEFAULT -1 COMMENT '-1: normal, 0: kakao, 1: google, 2: naver',
  `sns_id` varchar(100) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `skill_level` tinyint(1) NOT NULL DEFAULT -1 COMMENT '-1: unknown, 0: amateur, 1: professional',
  `average_par` float NOT NULL DEFAULT 0,
  `handicap` double NOT NULL,
  `handicap_issued_date_time` datetime DEFAULT NULL,
  `handicap_expired_date_time` datetime DEFAULT NULL,
  `score` int(11) NOT NULL DEFAULT 0,
  `score_cnt` int(11) NOT NULL DEFAULT 0,
  `ci` varchar(250) DEFAULT NULL,
  `introduction` varchar(150) DEFAULT NULL,
  `access_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1: 정상, 0: 접근차단',
  `public_profile` tinyint(1) NOT NULL DEFAULT 3 COMMENT '1: public, 2: partial, 3: private',
  `user_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0: disconfirm, 1: confirm',
  `use_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT ' 0: 탈퇴, 1: 일반, 2: 완전 삭제',
  `kcb_date_time` datetime(6) DEFAULT NULL,
  `join_date_time` datetime(6) NOT NULL,
  `left_date_time` datetime(6) DEFAULT NULL,
  `left_type` tinyint(1) NOT NULL DEFAULT -1 COMMENT '//0: not used, 1: not interested in, 2: hard to use, 3: private data leak, 4: has another account, 5: other',
  `left_reason` varchar(1500) DEFAULT NULL,
  `last_login` datetime(6) DEFAULT NULL,
  `login_cnt` int(11) NOT NULL DEFAULT 0,
  `push_token` varchar(500) DEFAULT NULL,
  `update_id` varchar(20) DEFAULT NULL,
  `create_id` varchar(20) DEFAULT NULL,
  `update_date_time` datetime(6) DEFAULT NULL,
  `create_date_time` datetime(6) DEFAULT NULL,
  `user_kind` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sns_type_sns_id` (`sns_type`,`sns_id`),
  UNIQUE KEY `user_id_sns_type` (`user_id`,`sns_type`),
  KEY `use_status` (`use_status`),
  KEY `phone_confirmed` (`user_type`) USING BTREE,
  KEY `access_type` (`access_type`),
  KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
