/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE IF NOT EXISTS `account` (
  `id` varchar(20) NOT NULL,
  `name` varchar(30) NOT NULL,
  `email` varchar(64) NOT NULL,
  `password` varchar(200) NOT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `type` tinyint(1) NOT NULL DEFAULT 2,
  `privileges` varchar(250) DEFAULT NULL,
  `use_status` tinyint(1) NOT NULL DEFAULT 1,
  `last_login` datetime(6) DEFAULT NULL,
  `login_failed_attempts` smallint(6) NOT NULL DEFAULT 0,
  `create_date_time` datetime(6) NOT NULL,
  `update_date_time` datetime(6) NOT NULL,
  `update_id` varchar(20) NOT NULL,
  `create_id` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `use_status` (`use_status`),
  KEY `type` (`type`),
  KEY `FK_account_account` (`create_id`),
  KEY `FK_account_account_2` (`update_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
