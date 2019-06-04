-- MySQL dump 10.13  Distrib 5.7.20, for Linux (x86_64)
--
-- Host: localhost    Database: project
-- ------------------------------------------------------
-- Server version	5.7.26-0ubuntu0.16.04.1

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
-- Table structure for table `abolition_project`
--

create database project;
use project;

DROP TABLE IF EXISTS `abolition_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `abolition_project` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `organization` varchar(30) NOT NULL,
  `teacher_id` bigint(20) NOT NULL,
  `create_time` date NOT NULL,
  `abolition_organization` varchar(30) NOT NULL,
  `abolition_instruction` varchar(255) NOT NULL,
  `operator` varchar(12) NOT NULL,
  `operator_tel` varchar(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `abolition_project`
--

LOCK TABLES `abolition_project` WRITE;
/*!40000 ALTER TABLE `abolition_project` DISABLE KEYS */;
INSERT INTO `abolition_project` VALUES (1,'ceshi','计算机科学学院',20110001,'2019-05-09','信息化建设推进办公室(专家审核)','btg','欧阳峰','13091234123'),(3,'xx','计算机科学学院',20110001,'2019-05-09','信息化建设推进办公室(专家审核)','btg','欧阳峰','13091234123');
/*!40000 ALTER TABLE `abolition_project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_manager`
--

DROP TABLE IF EXISTS `i_manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_manager` (
  `id` bigint(20) NOT NULL,
  `name` varchar(12) COLLATE utf8_unicode_ci NOT NULL,
  `pwd` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `tel` varchar(11) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_manager`
--

LOCK TABLES `i_manager` WRITE;
/*!40000 ALTER TABLE `i_manager` DISABLE KEYS */;
INSERT INTO `i_manager` VALUES (20120001,'欧阳峰','password','13091234123');
/*!40000 ALTER TABLE `i_manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master`
--

DROP TABLE IF EXISTS `master`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `master` (
  `id` bigint(20) NOT NULL,
  `name` varchar(12) NOT NULL,
  `pwd` varchar(20) NOT NULL,
  `tel` varchar(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master`
--

LOCK TABLES `master` WRITE;
/*!40000 ALTER TABLE `master` DISABLE KEYS */;
INSERT INTO `master` VALUES (20000001,'张伟','password','13012342121'),(20110011,'李伟','password','13611112222'),(20110012,'张李','password','13312344321'),(20110013,'胡与','password','13412331111');
/*!40000 ALTER TABLE `master` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_audit`
--

DROP TABLE IF EXISTS `master_audit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `master_audit` (
  `master_id` bigint(20) NOT NULL,
  `status` varchar(9) NOT NULL,
  `result` varchar(6) DEFAULT NULL,
  `m_audit_instruction` varchar(255) DEFAULT NULL,
  `fin_funds` int(11) DEFAULT NULL,
  `project_id` int(11) NOT NULL,
  PRIMARY KEY (`project_id`,`master_id`),
  KEY `fk_masteraudit_mid` (`master_id`),
  CONSTRAINT `fk_masteraudit_mid` FOREIGN KEY (`master_id`) REFERENCES `master` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_audit`
--

LOCK TABLES `master_audit` WRITE;
/*!40000 ALTER TABLE `master_audit` DISABLE KEYS */;
INSERT INTO `master_audit` VALUES (20000001,'已完成','pass','tg',1000,1),(20110011,'已完成','fail','instruction',0,1),(20000001,'已完成','pass','tg',1234,2),(20110011,'已完成','pass','tg',100,2),(20110012,'已完成','pass','btg',1000,2),(20000001,'已完成','pass','tg',1000,3),(20110011,'已完成','fail','btg',0,3),(20110012,'已完成','pass','tg',1000,3),(20000001,'已完成','pass','通过',1000,4),(20110011,'已完成','pass','tg',1000,4),(20110012,'已完成','pass','tg',1000,4),(20000001,'已完成','pass','tg',1000,5),(20110011,'已完成','pass','tg',1000,5),(20110012,'已完成','pass','tg',1000,5),(20000001,'已完成','pass','演示系统功能',10000,6),(20110011,'已完成','pass','演示',10000,6),(20110012,'已完成','pass','演示',9000,6),(20000001,'已完成','pass','tongguo',1000,8),(20110011,'已完成','pass','tg',1000,8),(20110012,'已完成','pass','tg',1000,8),(20000001,'已完成','pass','通过',10000,9),(20110011,'已完成','pass','通过',10000,9),(20110012,'已完成','pass','通过',10000,9);
/*!40000 ALTER TABLE `master_audit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `o_manager`
--

DROP TABLE IF EXISTS `o_manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `o_manager` (
  `id` bigint(20) NOT NULL,
  `organization` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `pwd` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(12) COLLATE utf8_unicode_ci NOT NULL,
  `tel` varchar(11) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_manage_org_name` (`organization`),
  CONSTRAINT `fk_manage_org_name` FOREIGN KEY (`organization`) REFERENCES `organization` (`name`),
  CONSTRAINT `fk_manage_teacher_id` FOREIGN KEY (`id`) REFERENCES `teacher` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `o_manager`
--

LOCK TABLES `o_manager` WRITE;
/*!40000 ALTER TABLE `o_manager` DISABLE KEYS */;
INSERT INTO `o_manager` VALUES (20110001,'计算机科学学院','password','李刚','13690001123');
/*!40000 ALTER TABLE `o_manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organization`
--

DROP TABLE IF EXISTS `organization`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `organization` (
  `name` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organization`
--

LOCK TABLES `organization` WRITE;
/*!40000 ALTER TABLE `organization` DISABLE KEYS */;
INSERT INTO `organization` VALUES ('机电工程学院'),('理 学 院'),('电气信息学院'),('石油与天然气工程学院'),('艺术学院'),('计算机科学学院'),('马克思主义学院');
/*!40000 ALTER TABLE `organization` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `organization` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `teacher_id` bigint(20) NOT NULL,
  `create_time` date NOT NULL,
  `budget` int(11) NOT NULL,
  `status` varchar(18) COLLATE utf8_unicode_ci DEFAULT NULL,
  `fin_time` date DEFAULT NULL,
  `invite_way` varchar(18) COLLATE utf8_unicode_ci NOT NULL,
  `instruction` text COLLATE utf8_unicode_ci NOT NULL,
  `run_time` date DEFAULT NULL,
  `o_audit_instruction` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `i_audit_instruction` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `fin_funds` int(11) DEFAULT NULL,
  `purpose` varchar(6) COLLATE utf8_unicode_ci NOT NULL,
  `p_function` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `expect_result` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `i_fin_instruction` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `used_funds` int(11) DEFAULT NULL,
  `completion_status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `self_evaluation` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_teacher_id_teaid` (`teacher_id`),
  CONSTRAINT `fk_teacher_id_teaid` FOREIGN KEY (`teacher_id`) REFERENCES `teacher` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
INSERT INTO `project` VALUES (2,'tse','计算机科学学院',20110001,'2019-05-09',1000,'已完成','2019-05-12','邀请招标','ee','2019-05-11','xx','tg',1000,'教学','ee','ee','tg',900,'cc','cc'),(3,'xx','计算机科学学院',20110001,'2019-05-11',1000,'已完成','2019-05-12','公开招标','xx','2019-05-11','tg','tg',1000,'教学','xx','xx','tg',0,'完成','完美完成'),(4,'测试项目','计算机科学学院',20110001,'2019-05-12',1000,'执行中',NULL,'公开招标','测试','2019-06-03','通过','通过',1000,'教学','测试','测试','tg',0,'',''),(5,'测试','计算机科学学院',20110001,'2019-05-12',10000,'核定参数',NULL,'公开招标','测试',NULL,'通过','通过',10000,'办公','测试','测试','tg',0,'',''),(6,'演示项目','计算机科学学院',20110001,'2019-05-19',10000,'已完成','2019-05-19','公开招标','演示系统功能','2019-05-19','演示系统功能','演示项目功能',10000,'教学','演示系统功能','演示系统功能','演示系统功能',9950,'演示系统功能','演示系统功能'),(7,'演示','计算机科学学院',20110001,'2019-06-02',1000,'网信中心审核',NULL,'公开招标','演示',NULL,'tg','',0,'教学','演示','演示','',0,'',''),(8,'演示','计算机科学学院',20110001,'2019-06-03',1000,'执行中',NULL,'公开招标','演示','2019-06-03','通过','通过',1000,'教学','演示','演示','通过',0,'',''),(9,'视频演示','计算机科学学院',20110001,'2019-06-03',10000,'已完成','2019-06-03','公开招标','视频演示','2019-06-03','通过','通过',10000,'教学','视频演示','视频演示','通过',900,'演示','视频演示');
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_event`
--

DROP TABLE IF EXISTS `project_event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project_event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `time` date NOT NULL,
  `use_funds` int(11) NOT NULL,
  `instruction` varchar(255) NOT NULL,
  `project_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_event`
--

LOCK TABLES `project_event` WRITE;
/*!40000 ALTER TABLE `project_event` DISABLE KEYS */;
INSERT INTO `project_event` VALUES (1,'tse','2019-05-11',100,'测试',2),(2,'tse','2019-05-12',200,'测试',2),(3,'tse','2019-05-12',600,'测试',2),(4,'演示项目','2019-05-19',1000,'演示系统功能',6),(5,'演示项目','2019-05-19',8950,'演示',6),(6,'测试项目','2019-06-03',900,'演示',4),(7,'视频演示','2019-06-03',900,'视频演示',9);
/*!40000 ALTER TABLE `project_event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_invite`
--

DROP TABLE IF EXISTS `project_invite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project_invite` (
  `id` int(11) NOT NULL,
  `begin_time` date NOT NULL,
  `fin_time` date DEFAULT NULL,
  `funds` int(11) NOT NULL,
  `fin_funds` int(11) DEFAULT NULL,
  `company_name` varchar(60) DEFAULT NULL,
  `instruction` varchar(255) DEFAULT NULL,
  `invite_file_name` varchar(60) DEFAULT NULL,
  `invite_way` varchar(18) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `change_reason` varchar(255) DEFAULT NULL,
  `change_apply` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_invite_id` FOREIGN KEY (`id`) REFERENCES `project` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_invite`
--

LOCK TABLES `project_invite` WRITE;
/*!40000 ALTER TABLE `project_invite` DISABLE KEYS */;
INSERT INTO `project_invite` VALUES (2,'2019-05-11','2019-05-11',1000,1000,'测试公司','test','diff (1).txt','邀请招标','tse','','false'),(3,'2019-05-11','2019-05-11',1000,1000,'海尔','ces','研发计划表.xlsx','公开招标','xx','','false'),(4,'2019-05-12','2019-06-03',1000,1000,'微软','测试','diff.txt','公开招标','测试项目','','false'),(6,'2019-05-19','2019-05-19',10000,10000,'演示公司','演示系统功能','招标文件.doc','公开招标','演示项目','演示修改招标参数','true'),(8,'2019-06-03','2019-06-03',1000,1000,'微软','演示','招标文件.doc','公开招标','演示','','false'),(9,'2019-06-03','2019-06-03',10000,10000,'微软','测试','招标文件.doc','公开招标','视频演示','','false');
/*!40000 ALTER TABLE `project_invite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teacher`
--

DROP TABLE IF EXISTS `teacher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teacher` (
  `id` bigint(20) NOT NULL,
  `name` varchar(12) COLLATE utf8_unicode_ci NOT NULL,
  `organization` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `sex` char(1) COLLATE utf8_unicode_ci NOT NULL,
  `birth` date NOT NULL,
  `tel` varchar(11) COLLATE utf8_unicode_ci NOT NULL,
  `pwd` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `professional_title` varchar(18) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_teacher_org_name` (`organization`),
  CONSTRAINT `fk_teacher_org_name` FOREIGN KEY (`organization`) REFERENCES `organization` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher`
--

LOCK TABLES `teacher` WRITE;
/*!40000 ALTER TABLE `teacher` DISABLE KEYS */;
INSERT INTO `teacher` VALUES (20090021,'刘丽','石油与天然气工程学院','f','1988-02-12','13685954844','password',NULL),(20110001,'李小明','计算机科学学院','m','1990-08-12','13011229898','password','教授');
/*!40000 ALTER TABLE `teacher` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-06-04 18:53:46
