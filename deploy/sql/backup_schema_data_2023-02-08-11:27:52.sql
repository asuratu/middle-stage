-- MySQL dump 10.13  Distrib 8.0.31, for Linux (aarch64)
--
-- Host: localhost    Database: ms_user
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category_updated_at` (`updated_at`),
  KEY `idx_category_name` (`name`),
  KEY `idx_category_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'yHSvBNy','Voluptatem aut perferendis sit accusantium consequatur.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(2,'oxTspdC','Perferendis accusantium consequatur voluptatem sit aut.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(3,'ErNVJEt','Consequatur voluptatem perferendis sit aut accusantium.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(4,'mHscDCx','Voluptatem aut consequatur perferendis sit accusantium.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(5,'JpKTbyM','Accusantium perferendis sit voluptatem aut consequatur.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(6,'IrvuFpt','Sit consequatur perferendis voluptatem aut accusantium.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(7,'gNBlkvc','Accusantium perferendis voluptatem sit aut consequatur.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(8,'jkUrRZC','Perferendis accusantium voluptatem aut consequatur sit.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(9,'ErwVfWA','Perferendis aut consequatur voluptatem accusantium sit.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228'),(10,'jXbxVst','Sit aut consequatur perferendis accusantium voluptatem.','2023-01-31 15:38:32.228','2023-01-31 15:38:32.228');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `link`
--

DROP TABLE IF EXISTS `link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `link` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_link_updated_at` (`updated_at`),
  KEY `idx_link_name` (`name`),
  KEY `idx_link_email` (`email`),
  KEY `idx_link_phone` (`phone`),
  KEY `idx_link_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `link`
--

LOCK TABLES `link` WRITE;
/*!40000 ALTER TABLE `link` DISABLE KEYS */;
/*!40000 ALTER TABLE `link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `link`
--

DROP TABLE IF EXISTS `link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `link` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_link_created_at` (`created_at`),
  KEY `idx_link_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `link`
--

LOCK TABLES `link` WRITE;
/*!40000 ALTER TABLE `link` DISABLE KEYS */;
INSERT INTO `link` VALUES (1,'EMFepmh','http://ssVRcBL.org/kFykIhM.php','2023-01-31 15:38:32.234','2023-01-31 15:38:32.234'),(2,'gouaOip','http://www.UpelNyq.ru/','2023-01-31 15:38:32.234','2023-01-31 15:38:32.234'),(3,'uMKHrfp','http://CcLsSGk.info/fDMWEMD.html','2023-01-31 15:38:32.234','2023-01-31 15:38:32.234'),(4,'SayZTsH','http://cceImxG.biz/SRkomoW','2023-01-31 15:38:32.234','2023-01-31 15:38:32.234'),(5,'QrCDZpx','http://UkutcFK.info/mkZJuGI','2023-01-31 15:38:32.234','2023-01-31 15:38:32.234');
/*!40000 ALTER TABLE `link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `topic`
--

DROP TABLE IF EXISTS `topic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `topic` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `body` longtext NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `category_id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_topic_title` (`title`),
  KEY `idx_topic_user_id` (`user_id`),
  KEY `idx_topic_category_id` (`category_id`),
  KEY `idx_topic_created_at` (`created_at`),
  KEY `idx_topic_updated_at` (`updated_at`),
  CONSTRAINT `fk_topic_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `fk_topic_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `topic`
--

LOCK TABLES `topic` WRITE;
/*!40000 ALTER TABLE `topic` DISABLE KEYS */;
INSERT INTO `topic` VALUES (1,'Consequatur voluptatem sit accusantium perferendis aut.','Aut consequatur accusantium sit perferendis voluptatem. Consequatur sit accusantium voluptatem aut perferendis. Voluptatem aut sit perferendis accusantium consequatur.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(2,'Sit consequatur voluptatem accusantium aut perferendis.','Perferendis accusantium aut sit consequatur voluptatem. Sit perferendis voluptatem aut consequatur accusantium. Consequatur sit perferendis voluptatem aut accusantium. Voluptatem perferendis sit accusantium aut consequatur. Consequatur aut perferendis accusantium voluptatem sit.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(3,'Aut voluptatem sit consequatur perferendis accusantium.','Voluptatem sit aut perferendis consequatur accusantium. Accusantium voluptatem perferendis consequatur sit aut. Perferendis consequatur aut sit accusantium voluptatem. Sit perferendis aut voluptatem consequatur accusantium. Aut consequatur accusantium sit perferendis voluptatem. Perferendis consequatur accusantium voluptatem sit aut. Sit voluptatem aut consequatur perferendis accusantium.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(4,'Consequatur aut perferendis sit voluptatem accusantium.','Sit voluptatem accusantium aut consequatur perferendis. Perferendis voluptatem consequatur sit accusantium aut. Consequatur voluptatem sit accusantium aut perferendis. Accusantium consequatur aut perferendis sit voluptatem.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(5,'Voluptatem perferendis consequatur accusantium aut sit.','Perferendis consequatur accusantium sit voluptatem aut. Aut accusantium sit perferendis consequatur voluptatem. Consequatur perferendis aut voluptatem accusantium sit. Consequatur perferendis voluptatem sit accusantium aut. Aut sit perferendis consequatur voluptatem accusantium. Aut accusantium consequatur sit perferendis voluptatem. Sit voluptatem consequatur perferendis aut accusantium. Accusantium sit voluptatem aut perferendis consequatur. Voluptatem perferendis accusantium consequatur aut sit. Voluptatem aut perferendis consequatur accusantium sit.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(6,'Accusantium sit aut consequatur perferendis voluptatem.','Aut accusantium consequatur perferendis voluptatem sit. Consequatur perferendis voluptatem aut sit accusantium. Consequatur sit accusantium voluptatem perferendis aut. Perferendis voluptatem consequatur accusantium aut sit. Voluptatem consequatur perferendis aut sit accusantium.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(7,'Perferendis voluptatem accusantium sit aut consequatur.','Sit perferendis voluptatem accusantium aut consequatur. Voluptatem sit aut perferendis consequatur accusantium. Consequatur accusantium aut sit voluptatem perferendis.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(8,'Accusantium voluptatem aut consequatur perferendis sit.','Consequatur voluptatem aut perferendis sit accusantium. Consequatur accusantium perferendis aut sit voluptatem. Accusantium voluptatem consequatur aut sit perferendis. Sit accusantium voluptatem perferendis consequatur aut. Accusantium voluptatem aut perferendis sit consequatur. Perferendis aut sit accusantium consequatur voluptatem. Consequatur voluptatem accusantium perferendis sit aut. Perferendis aut consequatur sit voluptatem accusantium. Accusantium perferendis aut sit voluptatem consequatur. Sit consequatur perferendis accusantium aut voluptatem.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(9,'Accusantium perferendis consequatur sit voluptatem aut.','Aut perferendis accusantium voluptatem sit consequatur. Accusantium sit perferendis voluptatem consequatur aut. Perferendis aut consequatur sit accusantium voluptatem. Accusantium aut voluptatem sit perferendis consequatur. Consequatur accusantium perferendis aut sit voluptatem. Accusantium sit consequatur voluptatem perferendis aut. Voluptatem accusantium perferendis sit consequatur aut.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238'),(10,'Aut consequatur accusantium perferendis voluptatem sit.','Perferendis voluptatem consequatur aut accusantium sit.',1,3,'2023-01-31 15:38:32.238','2023-01-31 15:38:32.238');
/*!40000 ALTER TABLE `topic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `city` varchar(10) DEFAULT NULL,
  `introduction` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_created_at` (`created_at`),
  KEY `idx_user_updated_at` (`updated_at`),
  KEY `idx_user_name` (`name`),
  KEY `idx_user_email` (`email`),
  KEY `idx_user_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'ruGhtcL','tVtXeHe@JyKEIpO.biz','10954703106','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(2,'efPjXXd','toGhrlI@Ewfvlmj.com','19094914140','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(3,'NLesflM','UiarfrP@bZEmmON.com','11813659129','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(4,'PHVVgrA','TZpdpJK@ptcyjIo.info','18116945864','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(5,'qDdOrTL','WLpDBaI@yBaGiPI.biz','12854400368','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(6,'CsJvewW','TUNPtYG@LUQetcV.info','17980737666','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(7,'hSGxYRr','LGWQajB@ywBFTDo.net','16934477920','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(8,'UnMjsCs','uucETJB@FKwRIsI.org','16876237308','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(9,'AcPANNT','PWYYWTv@haqEyIR.org','12798925465','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL),(10,'sEpeETi','qhTUlGu@CFtuVpu.info','16936705661','$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe','2023-01-31 15:38:32.211','2023-01-31 15:38:32.211',NULL,NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'ms_user'
--

--
-- Dumping routines for database 'ms_user'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-02-08 11:27:52
