# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.19)
# Database: godale
# Generation Time: 2018-03-13 19:16:56 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `post_id` varchar(150) NOT NULL DEFAULT '',
  `post_title` varchar(250) DEFAULT NULL,
  `post_content` blob,
  `post_added_by` varchar(150) DEFAULT NULL,
  `post_added_date` datetime DEFAULT NULL,
  `post_updated_by` varchar(150) DEFAULT NULL,
  `post_updated_date` datetime DEFAULT NULL,
  `post_archive` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




--
-- Dumping routines (PROCEDURE) for database 'godale'
--
DELIMITER ;;

# Dump of PROCEDURE spPostAdd
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `spPostAdd` */;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`localhost`*/ /*!50003 PROCEDURE `spPostAdd`(
	IN p_id varchar(150),
	IN p_title varchar(250),
	IN p_content blob,
	IN p_added_by varchar(150),
	IN p_added_date datetime,
	IN p_updated_by varchar(150),
	IN p_updated_date datetime,
	IN p_archive tinyint
)
BEGIN
   
   INSERT INTO Post(
		post_id,	
		post_title,
		post_content,
		post_added_by,
		post_added_date,
		post_updated_by,
		post_updated_date,
		post_archive
   )
    VALUES (
    	p_id,
    	P_title,
    	p_content,
    	p_added_by,
    	p_added_date,
    	p_updated_by,
    	p_updated_date,
    	p_archive
    );
  	
 	
   END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
# Dump of PROCEDURE spPostArchive
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `spPostArchive` */;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`localhost`*/ /*!50003 PROCEDURE `spPostArchive`(
	IN p_id varchar(150)
)
BEGIN
   
    UPDATE Post
    SET
    post_archive = 1
    
    WHERE
    post_id = p_id;
 	
 	
   END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
# Dump of PROCEDURE spPostUnArchive
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `spPostUnArchive` */;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`localhost`*/ /*!50003 PROCEDURE `spPostUnArchive`(
	IN p_id varchar(150)
)
BEGIN
   
    UPDATE Post
    SET
    post_archive = 0
    
    WHERE
    post_id = p_id;
 	
 	
   END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
# Dump of PROCEDURE spPostUpdate
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `spPostUpdate` */;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`localhost`*/ /*!50003 PROCEDURE `spPostUpdate`(
	IN p_id varchar(150),
	IN p_title varchar(250),
	IN p_content blob,
	IN p_updated_by varchar(150),
	IN p_updated_date datetime
)
BEGIN
   
    UPDATE Post
    SET
    post_title = p_title,
    post_content = p_content,
    post_updated_by = p_updated_by,
    post_updated_date = p_updated_date
    
    WHERE
    post_id = p_id;
 	
 	
   END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
DELIMITER ;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
