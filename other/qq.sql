-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: qq
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `qq_user`
--

DROP TABLE IF EXISTS `qq_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `qq_user` (
  `qq` int unsigned NOT NULL,
  `pwd` char(32) NOT NULL DEFAULT '',
  `sign` char(32) NOT NULL DEFAULT '',
  `photoID` int NOT NULL DEFAULT '0',
  `nickName` char(10) NOT NULL DEFAULT '',
  `sex` enum('man','woman') NOT NULL,
  `birthday` int unsigned NOT NULL DEFAULT '0',
  `constellation` enum('Aries','Taurus','Gemini','Cancer','Leo','Virgo','Libra','Scorpio','Sagittarius','Capricorn','Aquarius','Pisces') NOT NULL,
  `bloodType` enum('A','B','O','AB') NOT NULL,
  `diploma` enum('无','小学','初中','高中','大学','研究生','博士') NOT NULL,
  `telephone` char(11) NOT NULL DEFAULT '',
  `email` char(32) NOT NULL DEFAULT '',
  `address` char(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`qq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `qq_user`
--

LOCK TABLES `qq_user` WRITE;
/*!40000 ALTER TABLE `qq_user` DISABLE KEYS */;
INSERT INTO `qq_user` VALUES 
(1234523451,'c4ca4238a0b923820dcc509a6f75849b','pbssert1',2,'SK1','woman',1595128984,'Aries','B','初中','15512214321','1234523491@qq.com','GUDT'),
(1234567890,'e10adc3949ba59abbe56e057f20f883e','okk',1,'SK','man',1594877431,'Aries','A','无','','',''),
(1234567891,'c4ca4238a0b923820dcc509a6f75849b','perfect1',2,'SK1','man',1595128856,'Leo','O','研究生','15512344321','1234567891@qq.com','GUDT'),
(1234567892,'c4ca4238a0b923820dcc509a6f75849b','perfect2',2,'SK2','man',1595128856,'Leo','O','研究生','15512344322','1234567892@qq.com','GUDT'),
(1234567893,'c4ca4238a0b923820dcc509a6f75849b','perfect3',2,'SK3','man',1595128856,'Leo','O','研究生','15512344323','1234567893@qq.com','GUDT'),
(1234567894,'c4ca4238a0b923820dcc509a6f75849b','perfect4',2,'SK4','man',1595128856,'Leo','A','研究生','15512344324','1234567894@qq.com','GUDT'),
(1234567895,'c4ca4238a0b923820dcc509a6f75849b','perfect5',2,'SK5','man',1595128856,'Leo','O','研究生','15512344325','1234567895@qq.com','GUDT'),
(1234567896,'c4ca4238a0b923820dcc509a6f75849b','perfect6',2,'SK6','man',1595128856,'Libra','O','研究生','15512364321','1234567896@qq.com','GUDT'),
(1234567897,'c4ca4238a0b923820dcc509a6f75849b','perfect7',2,'SK7','man',1595128856,'Leo','O','研究生','15512743210','1234567897@qq.com','GUDT'),
(1234567898,'c4ca4238a0b923820dcc509a6f75849b','perfect8',2,'SK8','man',1595128856,'Leo','O','研究生','15582344321','1234567898@qq.com','GUDT'),
(1234867892,'c4ca4238a0b923820dcc509a6f75849b','9-wgvvsd',2,'SK2','woman',1595128984,'Taurus','O','小学','15512744322','1234537892@qq.com','GUDT'),
(1263267893,'c4ca4238a0b923820dcc509a6f75849b','pe99ect3',2,'SK3','woman',1595128984,'Gemini','O','无','15512347323','1234575893@qq.com','GUDT'),
(1742567894,'c4ca4238a0b923820dcc509a6f75849b','per99ct4',2,'SK4','woman',1595128985,'Cancer','A','无','15512414324','1234576894@qq.com','GUDT');
/*!40000 ALTER TABLE `qq_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-19 11:29:55

drop table `qq_groups`;
CREATE TABLE `qq_groups`(
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  user_qq int unsigned NOT NULL,
  group_name varchar(32) NOT NULL DEFAULT "default",
  PRIMARY KEY(id),
  UNIQUE INDEX(user_qq,group_name),
  FOREIGN KEY(user_qq) REFERENCES qq_user(qq)
)ENGINE= InnoDB DEFAULT CHARSET=utf8;

insert into `qq_groups`(user_qq,group_name) value
(1234567890,"好友"),
(1234567890,"家人"),
(1234567890,"同学"),
(1234567891,"好友"),
(1234567891,"大学"),
(1234567891,"小学"),
(1234567892,"学院"),
(1234567892,"社团"),
(1234567892,"老师")
;
insert into `qq_groups`(user_qq,group_name) value
(1234567893,"好友"),
(1234567893,"家人"),
(1234567894,"朋友"),
(1234567893,"同学");

select * from `qq_groups`;

insert into `qq_groups`(user_qq,group_name) value;

drop table `qq_friends`;
create table `qq_friends`(
	id int(11) NOT NULL AUTO_INCREMENT,
	user_qq int unsigned NOT NULL,
	friend_qq int unsigned NOT NULL,
	group_id int(11) unsigned NOT NULL,
	add_time int unsigned NOT NULL ,
	primary key(id),
	UNIQUE INDEX(user_qq,friend_qq),
	FOREIGN KEY(user_qq) REFERENCES qq_user(qq),
	FOREIGN KEY(friend_qq) REFERENCES qq_user(qq),
	FOREIGN KEY(group_id) REFERENCES qq_groups(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `qq_friends` (user_qq,friend_qq,group_id,add_time)values
(1234567890,1234567891,1,unix_timestamp()),
(1234567890,1234567892,2,unix_timestamp()),
(1234567890,1234567893,3,unix_timestamp()),
(1234567891,1234567890,4,unix_timestamp()),
(1234567892,1234567890,7,unix_timestamp()),
(1234567893,1234567890,10,unix_timestamp())
;

insert into `qq_friends` (user_qq,friend_qq,group_id,add_time)values
(1234567890,1234567894,1,unix_timestamp()),
(1234567894,1234567890,,unix_timestamp()),
select * from `qq_friends`;

select qq_user.nickName as "好友名",
qq_groups.group_name as "所属分组"
from qq_user,qq_friends,qq_groups
where qq_friends.user_qq=1234567891 and qq_friends.friend_qq=qq_user.qq and qq_friends.group_id=qq_groups.id
;

