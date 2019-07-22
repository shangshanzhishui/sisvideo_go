-- MySQL dump 10.13  Distrib 5.7.26, for Linux (x86_64)
--
-- Host: localhost    Database: sis_video
-- ------------------------------------------------------
-- Server version	5.7.26-0ubuntu0.18.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL,
  `video_id` varchar(64) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `content` text,
  `time` datetime DEFAULT NULL,
  `display_time` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES ('049b3e3e-09fb-4fea-82cf-869a4aa3c499','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:33'),('146f2c78-95f0-4841-bcba-ec0d0f1a203f','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:27'),('235476bb-39db-4797-a0bc-7f54822a225c','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:50:39'),('2758564f-bb27-41bb-ae75-8378381fc571','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,20:08:01'),('28a87f93-882f-426e-bcc1-18cf64e95e58','28456704-5a5c-4f42-b5c2-f35127511cdd',30,'啊哈',NULL,'jan 22 2019,00:23:04'),('34be33c3-60d9-4bb5-9fff-d76d06c1e643','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('36ab0026-c305-49f4-ac2b-2f2b896f9c7d','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:29:35'),('389cf97f-3235-4605-97f7-c043ae4d3205','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:35'),('395b507a-63c3-4a76-a746-3aad1c4bf5ac','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:44:38'),('3b41c021-ec6d-49b0-a6d3-34e4af7c453d','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,20:08:09'),('3d3ab17b-5dd2-4652-a215-9616486d03ef','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:39'),('3e206e3a-40e6-4fd0-8961-00fac23a3121','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:40'),('4c63b02f-ba63-4b19-82e6-8369e2146306','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:44:44'),('5253074f-9aec-4eae-978d-6cb2ae8a03df','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,20:10:36'),('58b4e0c8-76fb-4dd1-9a4e-b0786f413c42','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh 哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈   \n',NULL,'jan 21 2019,19:52:34'),('5be15258-aeb7-4083-aa3d-9f04951607d1','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('5dc41636-9633-4d0a-b0c3-b2d7d1044bcd','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:45:36'),('66d4b830-22d6-488d-9849-f97b261dadc5','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:58:12'),('6df46fad-5161-435a-a483-7c7e811b3264','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh 哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈   \n',NULL,'jan 21 2019,19:54:24'),('72bd2792-6c4e-4135-b9fa-78f66a5fc999','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:57:09'),('7bd50e86-1671-407d-a1fd-e35ef125c2f6','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:35'),('7cf51f48-7aa6-4b17-9195-c99a9708fabe','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('82c61ee2-c4f8-47a9-a7db-f9d8524e2652','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:45:07'),('86372bd9-e638-4177-87a0-ac475da19524','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,20:02:55'),('9448b025-382a-480a-8ca8-a508c774258e','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:39'),('94c3dac8-7419-49ca-a07d-c3985801b58b','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee\n',NULL,'jan 21 2019,19:50:42'),('9baa6d4b-cfe9-40ec-8fdc-467d37b33f23','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('a82cacdf-632e-4ae1-8643-5b90c2310d3a','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('ab56668b-231d-423b-87de-4b3c664dd0d2','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:37'),('ae26d67f-02b6-4f7d-a122-238322923a61','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:36'),('b2fc06a5-6bf5-4c38-8511-9b9ffa48aea6','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:30'),('c1193f69-6b57-4594-b627-8911c6c36009','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('c12188d9-d645-4e0b-be99-e4ec3d9fe356','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh 哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈   \n',NULL,'jan 21 2019,19:54:27'),('c312da95-5727-4ba4-ba35-5a50ff9a1909','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh 哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈   \n',NULL,'jan 21 2019,19:54:25'),('ce55b602-f1ec-4913-9178-89a60463e873','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:40'),('ced39e59-3c22-4ff7-b372-8a6fb1b7f5b9','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:41'),('de6249d9-acfe-4d98-9b00-4678ed5423ff','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:32'),('f2f98fb6-3a63-4578-849f-885211733f70','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh 哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈   \n',NULL,'jan 21 2019,19:54:28'),('f61079f2-3e35-4071-91cd-faf5ab512a25','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:54'),('f6a1549a-36f1-4f96-b59e-3aa05142246d','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'你好阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',NULL,'jan 21 2019,19:28:39'),('fee7bc75-7735-4a61-b347-3008409d996c','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'小白  \n',NULL,'jan 21 2019,19:56:36'),('fef27059-5aaf-4327-987a-fa97c323d916','28456704-5a5c-4f42-b5c2-f35127511cdd',29,'eeeeeeeeeeeeeeeeeeeeee',NULL,'jan 21 2019,19:30:03');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `del_Video_Record`
--

DROP TABLE IF EXISTS `del_Video_Record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `del_Video_Record` (
  `video_id` varchar(64) NOT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `del_Video_Record`
--

LOCK TABLES `del_Video_Record` WRITE;
/*!40000 ALTER TABLE `del_Video_Record` DISABLE KEYS */;
/*!40000 ALTER TABLE `del_Video_Record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sessions` (
  `session_id` varchar(64) NOT NULL,
  `author_id` int(11) DEFAULT NULL,
  `TTL` tinytext,
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
INSERT INTO `sessions` VALUES ('2a7cbb4d-4a40-47ab-94d8-2957887c340e',30,'1563728021466'),('37bbb37f-7a4d-4fc5-bdf5-fc51341ccbf3',30,'1563727928266'),('7d042b90-fc50-40f4-868c-941ee921712f',25,'1563609141607'),('b5434761-1e44-418a-af87-741a64221f6f',29,'1563675949317');
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `pwd` varchar(64) DEFAULT NULL,
  `login_time` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'123','123',NULL),(4,'1234','123',NULL),(5,'nihao','123',NULL),(8,'你好','12345678',NULL),(10,'你好1','12345678',NULL),(11,'','',NULL),(14,'qq','qq',NULL),(22,'78','77',NULL),(23,'l','l',NULL),(24,'we','we',NULL),(25,'we7','we7',NULL),(29,'8','8',NULL),(30,'9','9',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video_info`
--

DROP TABLE IF EXISTS `video_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `video_info` (
  `id` varchar(64) NOT NULL,
  `author_id` int(11) DEFAULT NULL,
  `video_name` text,
  `display_time` text,
  `create_time` datetime DEFAULT NULL,
  `video_path` varchar(64) DEFAULT NULL,
  `image_path` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video_info`
--

LOCK TABLES `video_info` WRITE;
/*!40000 ALTER TABLE `video_info` DISABLE KEYS */;
INSERT INTO `video_info` VALUES ('28456704-5a5c-4f42-b5c2-f35127511cdd',22,'神盾局','jan 21 2019,14:41:51',NULL,'/videos/28456704-5a5c-4f42-b5c2-f35127511cdd','/statics/video_image/f6886375-3f9b-4000-888d-c1da40ee19de.jpg');
/*!40000 ALTER TABLE `video_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-07-22  8:59:09
