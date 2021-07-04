-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: airlines
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `airlines`
--

DROP TABLE IF EXISTS `airlines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airlines` (
  `airlinesID` int NOT NULL AUTO_INCREMENT,
  `nameAL` varchar(45) NOT NULL,
  `flightNumber` varchar(15) DEFAULT NULL,
  `flyingTo` varchar(45) DEFAULT NULL,
  `flightDate` date DEFAULT NULL,
  `flightTime` time DEFAULT NULL,
  `numberOfSeat` int DEFAULT '0',
  `seatTaken` int DEFAULT '0',
  `reserved` int DEFAULT '0',
  PRIMARY KEY (`airlinesID`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airlines`
--

LOCK TABLES `airlines` WRITE;
/*!40000 ALTER TABLE `airlines` DISABLE KEYS */;
INSERT INTO `airlines` VALUES (1,'Singapore Airlines','SA101','Taiwan','2021-07-10','23:15:00',200,0,0),(2,'Scoot','SC102','Malaysia','2021-07-10','12:00:00',200,0,0),(3,'Jetstar Airlines','JS103','Hong Kong','2021-07-10','08:00:00',300,0,0),(4,'Singapore Airlines','SA102','Taiwan','2021-07-11','23:15:00',200,0,0),(5,'Scoot','SC103','Malaysia','2021-07-11','12:00:00',150,0,0),(6,'Jetstar Airlines','JS104','Hong Kong','2021-07-11','08:00:00',300,0,0),(7,'Singapore Airlines','SA106','Taiwan','2021-07-12','23:15:00',200,0,0),(8,'Scoot','SC111','Malaysia','2021-07-12','12:00:00',150,0,0),(9,'Jetstar Airlines','JS110','Hong Kong','2021-07-12','08:00:00',300,0,0),(10,'Singapore Airlines','SA112','Taiwan','2021-07-20','23:15:00',190,0,0),(11,'Scoot','SC113','Malaysia','2021-07-20','12:00:00',150,0,0),(12,'Jetstar Airlines','JS105','Hong Kong','2021-07-20','08:00:00',300,0,0),(13,'Singapore Airlines','SA122','Taiwan','2021-07-21','23:15:00',200,0,0),(14,'Scoot','SC121','Malaysia','2021-07-21','12:00:00',150,0,0),(15,'Jetstar Airlines','JS115','Hong Kong','2021-07-21','08:00:00',300,0,0),(16,'Singapore Airlines','SA124','Taiwan','2021-07-22','23:15:00',200,0,0),(17,'Scoot','SC130','Malaysia','2021-07-22','12:00:00',150,0,0),(18,'Jetstar Airlines','JS189','Hong Kong','2021-07-22','08:00:00',300,0,0);
/*!40000 ALTER TABLE `airlines` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-04 23:25:15
