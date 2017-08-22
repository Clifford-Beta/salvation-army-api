# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.1.21-MariaDB)
# Database: salvation_army_db
# Generation Time: 2017-08-18 13:14:36 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table i_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `i_type`;

CREATE TABLE `i_type` (
  `i_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `i_type_name` varchar(45) NOT NULL,
  `i_type_description` varchar(200) DEFAULT '""',
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `i_type_status` int(11) DEFAULT '1',
  PRIMARY KEY (`i_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `i_type` WRITE;
/*!40000 ALTER TABLE `i_type` DISABLE KEYS */;

INSERT INTO `i_type` (`i_type_id`, `i_type_name`, `i_type_description`, `timestamp`, `i_type_status`)
VALUES
	(1,'Electronics','Electical appliances','2017-08-17 07:42:28',1),
	(28,'Electronics','Electical appliances','2017-08-18 13:13:21',0);

/*!40000 ALTER TABLE `i_type` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
