-- MySQL dump 10.13  Distrib 8.0.31, for macos12.6 (x86_64)
--
-- Host: 13.215.228.202    Database: myspotify
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
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_name` varchar(40) NOT NULL,
                            `email` varchar(100) NOT NULL,
                            `password` text NOT NULL,
                            `created_at` datetime DEFAULT NULL,
                            `updated_at` datetime DEFAULT NULL,
                            `role` longtext,
                            `status` longtext,
                            `deleted_at` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (2,'truong','truong@gmail.com','$2a$14$olAQHp.vJvBxIOT1pdV2buB4jNfU/q4eJDLP/5bF2D7l9/dBgU4JO','2022-12-17 00:00:00','2022-12-17 18:15:58','Admin','Active',NULL),(3,'thanh','thanh@gmail.com','$2a$14$olAQHp.vJvBxIOT1pdV2buB4jNfU/q4eJDLP/5bF2D7l9/dBgU4JO','2022-12-18 01:37:35','2022-12-18 01:37:35','User','Active',NULL),(4,'nhon','nhon@gmail.com','$2a$14$olAQHp.vJvBxIOT1pdV2buB4jNfU/q4eJDLP/5bF2D7l9/dBgU4JO','2022-12-18 02:33:04','2022-12-18 02:33:04','User','Active',NULL);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `albums`
--

DROP TABLE IF EXISTS `albums`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `albums` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(40) NOT NULL,
                          `artist_id` int NOT NULL,
                          `created_at` datetime DEFAULT NULL,
                          `updated_at` datetime DEFAULT NULL,
                          `cover_img` varchar(400) DEFAULT NULL,
                          `deleted_at` datetime DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `artist_ID` (`artist_id`),
                          CONSTRAINT `albums_ibfk_1` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `albums`
--

LOCK TABLES `albums` WRITE;
/*!40000 ALTER TABLE `albums` DISABLE KEYS */;
INSERT INTO `albums` VALUES (8,'Theres No One At All ',7,'2022-12-17 22:09:07','2022-12-17 22:09:08','https://cdn.baogiaothong.vn/upload/images/2022-2/article_img/2022-04-20/img-bgt-2021-tung-1650425686-width700height874.jpg',NULL),(9,'Tự Sự',8,'2022-12-17 22:30:18','2022-12-17 22:30:19','https://avatar-ex-swe.nixcdn.com/song/2022/06/23/c/6/c/3/1655946067983_640.jpg',NULL),(10,'Show của Đen',9,'2022-12-17 23:43:57','2022-12-17 23:43:58','https://vov2-media.emitech.vn/sites/default/files/styles/large/public/2021-09/show_cua_den_artwork_copy.jpg',NULL),(11,'3107',10,'2022-12-18 00:58:17','2022-12-18 01:01:44','https://media.viez.vn/prod/2021/10/31/1635701188384_c838a8e069.jpeg',NULL);
/*!40000 ALTER TABLE `albums` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `artists`
--

DROP TABLE IF EXISTS `artists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `artists` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `name` varchar(40) NOT NULL,
                           `description` text,
                           `created_at` datetime DEFAULT NULL,
                           `updated_at` datetime DEFAULT NULL,
                           `cover_img` varchar(400) DEFAULT NULL,
                           `deleted_at` datetime DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artists`
--

LOCK TABLES `artists` WRITE;
/*!40000 ALTER TABLE `artists` DISABLE KEYS */;
INSERT INTO `artists` VALUES (7,'Sơn Tùng MTP','Nguyễn Thanh Tùng, thường được biết đến với nghệ danh Sơn Tùng M-TP, là một nam ca sĩ kiêm sáng tác nhạc, rapper và diễn viên người Việt Nam','2022-12-17 22:08:50','2022-12-17 22:08:52','https://lh3.googleusercontent.com/LBZbzy9NXoY_0vQQOkDQnVSzu27am8yxvcsxOk0CPhfnr7uraTv-9ONUje1b7zcK0bTqTbI1_pY2hVzXu4aGbSQ9',NULL),(8,'Orange','Khương Hoàn Mỹ (sinh ngày 15 tháng 2 năm 1997) thường được biết đến với nghệ danh Orange, là một nữ ca sĩ người Việt Nam. Cô bắt đầu trở nên nổi tiếng sau khi phát hành ca khúc \"Người lạ ơi!\" kết hợp cùng Superbrothers và Karik.','2022-12-17 22:29:24','2022-12-17 22:29:25','https://upload.wikimedia.org/wikipedia/commons/d/d7/Khuong_Hoan_My.png',NULL),(9,'Đen','Nguyễn Đức Cường, thường được biết đến với nghệ danh Đen Vâu hay Đen, là một nam rapper và nhạc sĩ người Việt Nam. Đen Vâu từng giành được giải Cống hiến và là \"một trong số ít nghệ sĩ thành công từ làn sóng underground và âm nhạc indie\" của Việt Nam.','2022-12-17 23:42:32','2022-12-17 23:42:34','https://upload.wikimedia.org/wikipedia/commons/d/d5/Den_Vau_in_Song_Gallery_2021.jpg',NULL),(10,'W/n','W/n là chủ nhân của ca khúc 3107 3 hiện đang thịnh hành với hơn chục triệu view trên Youtube. Anh chàng này còn làm khán giả mong nhớ bởi tài chơi nhạc tuyệt vời của mình và sự bí ẩn của mình. Bởi hoạt động âm nhạc trong thời gian 3 năm nhưng anh chưa một lần lộ mặt.','2022-12-18 00:18:12','2022-12-18 00:30:58','https://media.travelmag.vn/files/content/2021/11/01/after-release_-001-09355240.jpg',NULL);
/*!40000 ALTER TABLE `artists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `interactions`
--

DROP TABLE IF EXISTS `interactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `interactions` (
                                `user_id` int NOT NULL,
                                `song_id` int NOT NULL,
                                `liked` tinyint(1) DEFAULT NULL,
                                `created_at` datetime DEFAULT NULL,
                                `updated_at` datetime DEFAULT NULL,
                                `deleted_at` datetime DEFAULT NULL,
                                `id` bigint NOT NULL AUTO_INCREMENT,
                                PRIMARY KEY (`id`),
                                KEY `user_ID` (`user_id`),
                                KEY `song_ID` (`song_id`),
                                CONSTRAINT `interactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `accounts` (`id`),
                                CONSTRAINT `interactions_ibfk_2` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interactions`
--

LOCK TABLES `interactions` WRITE;
/*!40000 ALTER TABLE `interactions` DISABLE KEYS */;
INSERT INTO `interactions` VALUES (2,35,1,'2022-12-17 22:22:41','2022-12-17 22:22:41',NULL,2);
/*!40000 ALTER TABLE `interactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlists`
--

DROP TABLE IF EXISTS `playlists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlists` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(40) NOT NULL,
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             `cover_img` varchar(400) DEFAULT NULL,
                             `user_id` int NOT NULL,
                             `deleted_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             KEY `user_ID` (`user_id`),
                             CONSTRAINT `playlists_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlists`
--

LOCK TABLES `playlists` WRITE;
/*!40000 ALTER TABLE `playlists` DISABLE KEYS */;
INSERT INTO `playlists` VALUES (5,'Show của Đen','2022-12-17 23:00:50','2022-12-17 23:17:26','https://vov2-media.emitech.vn/sites/default/files/styles/large/public/2021-09/show_cua_den_artwork_copy.jpg',2,NULL);
/*!40000 ALTER TABLE `playlists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlists_songs`
--

DROP TABLE IF EXISTS `playlists_songs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlists_songs` (
                                   `song_id` int NOT NULL,
                                   `playlist_id` int NOT NULL,
                                   `created_at` datetime DEFAULT NULL,
                                   `deleted_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   PRIMARY KEY (`song_id`,`playlist_id`),
                                   KEY `playlist_ID` (`playlist_id`),
                                   CONSTRAINT `playlists_songs_ibfk_1` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`),
                                   CONSTRAINT `playlists_songs_ibfk_2` FOREIGN KEY (`playlist_id`) REFERENCES `playlists` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlists_songs`
--

LOCK TABLES `playlists_songs` WRITE;
/*!40000 ALTER TABLE `playlists_songs` DISABLE KEYS */;
INSERT INTO `playlists_songs` VALUES (39,5,'2022-12-17 23:48:23','2022-12-17 23:50:54','2022-12-17 23:48:23');
/*!40000 ALTER TABLE `playlists_songs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `songs`
--

DROP TABLE IF EXISTS `songs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `songs` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `name` longtext,
                         `album_id` int NOT NULL,
                         `artist_id` int NOT NULL,
                         `lyrics` text,
                         `length` int DEFAULT NULL,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `youtube_link` text,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `url` text,
                         `song_cloud_id` text NOT NULL,
                         PRIMARY KEY (`id`),
                         KEY `artist_ID` (`artist_id`),
                         KEY `album_ID` (`album_id`),
                         CONSTRAINT `songs_ibfk_1` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`),
                         CONSTRAINT `songs_ibfk_2` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `songs`
--

LOCK TABLES `songs` WRITE;
/*!40000 ALTER TABLE `songs` DISABLE KEYS */;
INSERT INTO `songs` VALUES (35,'Theres No One At All',8,7,'Ah, ah\nDon’t wanna make up\n\n[Verse 1]\nDon’t wanna make up, headstrong\nGonna wake up, so long\nI kinda died hard and grew\n\nSeeing that part of you\nI never knew love\nSomehow pretty sick of that vow\nSwallow it down, that pain\nWearing that crown, unchained\n\n[Pre-Chorus]\nClosing my eyes\nI hear your sigh oh\nCatching the light\nStill hear your whine oh\nWatching you lie\nLosing my mind oh\nWishing you’d die\n\nDon’t even try oh\nThis is goodbye\nDrying that tear oh\nWing it and drive\nAway from here oh\nStill asking why, why, why, why, why, why?\n\n[Chorus]\nThere’s no one at all, ah-ah\nThere’s no one at all, ah-ah\nThere’s no one at all, ah-ah\nThere’s no one at all, oh-oh-oh\nNo one at all, oh-oh\nGot no onе at all, oh-oh\nGot no one at all, ah-ah\nThere’s no onе at all, oh-oh-oh\n\n[Verse 2]\nRan after you, got no doubt\nGod waiting there, knocked me out\nHe let me see whatever\nYou hide from me since the day we loved lesser\nBurned the promise we make faster\nThan hurricane flooding river\nYou took my heart, leaving nothing\nSo did it feel that good, see my soul bleeding?\n\n[Pre-Chorus]\nClosing my eyes\nI hear your sigh oh\nCatching the light\nStill hear your whine oh\nWatching you lie\nLosing my mind oh\nWishing you’d die\n\nDon’t even try oh\nThis is goodbye\nDrying that tear oh\nWing it and drive\nAway from here oh\nStill asking why, why, why, why, why, why?\n\n[Chorus]\nThere’s no one at all, ah-ah\nThere’s no one at all, ah-ah\nThere’s no one at all, ah-ah\nThere’s no one at all, oh-oh-oh\nNo one at all, oh-oh\n\nGot no one at all, oh-oh\nGot no one at all, ah-ah\nThere’s no one at all, oh-oh-oh\nNo one at all (oh-oh-oh)\nGot no one at all (oh-oh-oh)\nGot no one at all (oh-oh-oh)\nThere’s no one at all, oh-oh-oh\n\n[Outro]\nThere’s no one at all\nNo one at all, ah',170,'2022-12-17 22:10:51.836','2022-12-17 22:10:51.836','https://www.youtube.com/watch?v=Bd0Gwku-shI&ab_channel=LyricsStage',NULL,'http://res.cloudinary.com/dbk0cmzcb/video/upload/v1671289851/mca6ls971o8lctnbgf67.mp3','f459617cfe819b8f83f2488e427ace43'),(36,'Cơn mưa ngang qua',8,7,'Cơn mưa ngang qua\nCơn mưa đi ngang qua\nĐừng làm rơi thêm, thêm, thêm, nhiều giọt lệ\nCòn đâu đây bao câu ca anh tặng em\nTình yêu em mang cuốn lấp đi bao nhiêu câu ca\nVà còn lại đây, đôi môi đau thương trong màn đêm\nPhải lẻ loi, gồng mình bước qua niềm đau khi em rời xa\nCơn mưa rồi lại thêm, lại thêm, lại thêm\nXé đi không gian ngập tràn nụ cười\nNhìn lại nơi đây, bao kỉ niệm giờ chìm xuống dưới hố sâu vì em\nChính em đã đổi thay\nVà đôi bàn tay ấm êm, ngày nào còn lại giữ\nVì em vì em\nVì em đã xa rồi, tình anh chìm trong màn đêm\nLà vì em đã quên rồi, tình anh chỉ như giấc mơ\nEm bước đi rồi, ôi bao cơn mưa\nEm bước đi rồi, xin hãy xua tan đi em bóng dáng hình em\nEm đã mãi rời xa\nMy girl, em quên đi bao nhiêu\nMy girl, em quên đi bao lâu\nMy girl, em quên đi cuộc tình mà anh trao em\nThôi thôi em đi đi đã hết rồi\nMy girl em quên đi bao nhiêu\nMy girl, em quên đi bao lâu\nMy girl, em quên đi cuộc tình, em quên quên quên\nCơn mưa ngang qua mang em đi xa\nCơn mưa ngang qua khiến em nhạt nhòa\nChẳng một lời chào người vội rời bỏ đi không chia li cho con tim anh thêm bao yếu mềm\nCơn mưa ngang qua cuốn đi bao yêu thương\nCơn mưa ngang qua khiến con tim mất phương hướng\nCơn mưa Kia nặng hạt, ôi mưa thêm nặng hạt\nEm đã rời xa đôi bàn tay trong con tim của anh\nBuông lơi bàn tay em đi, em đi rời xa bên tôi người ơi\nVà buông lơi giấc mơ em cho, em cho con tim tôi đau biết mấy\nThôi cũng đã đến hồi kết.thật rồi mà người\nThôi cũng đá đến hồi kết\nĐừng nhìn làm gì\nAnh sẽ quên đi một ai, ai, ai, và rồi làm ngơ, ngơ, ngơ\nVì em đã xa rồi, tình anh chìm trong màn đêm\nLà vì em đã quên rồi, tình anh chỉ như giấc mơ\nEm bước đi rồi.Ôi bao cơn mưa\nEm bước đi rồi.Xin hãy xua tan đi em, bóng dáng hình em\nEm đã mãi rời xa...!!!\nMy girl, em quên đi bao nhiêu\nMy girl, em quên đi bao lâu\nMy girl, em quên đi cuộc tình mà anh trao em, thôi thôi em đi đi đã hết rồi\nMy girl em quên đi bao nhiêu\nMy girl em quên đi bao lâu.\nMy girl em quên đi cuộc tình em quên.quên.quên.\nVà rồi em đi qua bước qua\nỞ lại chốn đây với bao u buồn\nĐể rồi tháng trôi qua, rồi năm trôi qua, thoáng qua\nNụ cười em đang ở đâu, người ơi, ở đâu\nVà bờ môi em đang ở đâu, anh tìm\nLục tìm ma không thấy, nụ cười em\nNgười hãy nói trả lời đi\nVì sao em đi đi quên đi bao nhiêu giấc mơ bên anh xưa kia\nCơn mưa cẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nCơn mưa cẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nCơn cơn mưa vẫn rơi rơi rơi\nMy girl, em quên đi bao nhiêu\nMy girl, em quên đi bao lâu\nMy girl, em quên đi cuộc tình mà anh trao em, thôi thôi em đi đi đã hết rồi\nMy girl, em quên đi bao nhiêu\nMy girl, em quên đi bao lâu\nMy girl, em quên đi cuộc tình, em quên.quên.quên\nEm quên mất rồi',231,'2022-12-17 22:25:51.619','2022-12-17 22:25:51.619','https://www.youtube.com/watch?v=HHqbE_Z-tTI&list=RDMM&start_radio=1&rv=Bd0Gwku-shI&ab_channel=gatreduonglinh',NULL,'http://res.cloudinary.com/dbk0cmzcb/video/upload/v1671290750/lundsygajbfouujeqvei.mp3','2de3a9377507ba5499f27c1e44e60515'),(37,'Tự Sự',9,8,'Ừ thì ai cũng có kỉ niệm buồn khó quên\nỪ thì ai cũng có những nỗi thao thức vô hình không tên\nPhải chăng đã yêu quá nhiều, phải chăng thiết tha quá nhiều\nPhải chăng niềm đau giằng xé nhưng sao ta cứ nuôi hi vọng\n\nỪ cuộc đời vẫn thế, ngày và đêm giống nhau\nỪ cuộc đời vẫn thế đâu ai muốn sống mãi trong u sầu\nBiết phải làm sao, biết đi về đâu\nBiết nơi nào đó nắng ấm có ghé ngang qua đời anh\n\nChuyện hợp tan như gió cứ đến rồi đi thật nhanh\nGiống yêu thương kia mỏng manh\nCũng như chuyện về chiếc lá đến lúc phải xa cành\nNếu như một mai thức giấc ngủ mê\nCó quên được ngàn câu ước thề\nVội vàng kí ức ùa về ừ thì cuộc đời vẫn thế\n\nỪ cuộc đời vẫn thế, chỉ còn ta đứng yên\nỪ cuộc đời vẫn thế bon chen xô đẩy ta vẫn hồn nhiên\nLy cà phê vẫn đong đầy, khói thuốc trắng vẫn còn bay\nỪ thì đôi mắt cay\n\nVER 2:\n\nThời gian sẽ chữa lành mọi vết thương, như ai đó đã từng nói\nỪ thì tin vậy đi, cứ tin dẫu cho niềm tin dần xa xôi\nVậy thôi xin hãy giữ lấy những kí ức đẹp một thời đã qua\nDù năm tháng nhạt nhòa phôi pha, tôi vẫn thế\n\nChuyện hợp tan như gió cứ đến rồi đi thật nhanh\nGiống yêu thương kia mỏng manh\nCũng như chuyện về chiếc lá đến lúc phải xa cành\nNếu như một mai thức giấc ngủ mê\nCó quên được ngàn câu ước thề\nVội vàng kí ức ùa về ừ thì cuộc đời vẫn thế\n\nỪ thì tôi vẫn thế, vẫn chỉ yêu mỗi em\nỪ thì tôi vẫn thế, bao nhiêu tiếc nuối vẫn cứ nhiều thêm\nAi chờ ai giữa cuộc đời, ai thương nhớ ai đầy vơi\nỪ thì cũng thế thôi …\n',260,'2022-12-17 22:32:16.362','2022-12-17 22:32:16.362','https://www.youtube.com/watch?v=C4QhZvTqhnI&list=RDMM&index=8&ab_channel=OrangeSingerOfficial',NULL,'http://res.cloudinary.com/dbk0cmzcb/video/upload/v1671291135/ofvb6b864gprgakb2tga.mp3','246b7ed77dc7cc3dd2c7316a5b97b4c9'),(38,'Nắng Ấm Xa Dần',8,7,'Nắng ấm xa dần rồi\nNắng ấm xa dần rồi\nNắng ấm xa dần bỏ rơi để lại những giấc mơ\nGiữ lại đi giữ giữ lại đi\nNắng ấm xa dần rồi\nNắng ấm xa dần rồi\nNắng ấm xa dần xa dần theo những tiếng cười\nHãy mang đi giúp những nỗi buồn\nTheo thời gian những hạt mưa như nặng thêm\nXóa hết thương yêu mặn nồng ngày nào giữa chúng ta\nAnh lục tìm vẫn cứ mãi lục tìm\nGiơ bàn tay cố kìm nén những cảm xúc\nVùi mình vào đêm đen anh chẳng tìm thấy lối ra oh oh\nSau lưng là tiếng nói yêu anh chẳng rời xa anh (uh uh)\nTrước mắt anh điều đấy nó dối trá tại sao người vội quên mau (là vì em)\nBài ca anh viết sẽ không được trọn vẹn đâu em (vội bước đi)\nEm yêu một ai thật rồi mãi chẳng là anh đâu\nVậy thì người cứ bước đi xa nơi này\nÁnh bình minh sẽ không còn nơi đây\nBước đi xa nơi này\nNhững lời yêu sẽ không còn nơi đây\nPhải tự đứng lên mà thôi\nChe nhẹ đi những niềm đau và nỗi buồn\nXung quanh anh giờ đây cô đơn mình anh ôm giấc mơ (vì ai)\nNhìn em bước ra đi xa dần (hoh)\nNhìn em bước ra đi xa dần (hoh)\nNhìn em bước ra đi xa dần (hoh)\nNhìn em bước ra đi xa dần (hoh)\nM T P\nĐến rồi lại đi và cứ vội vàng đi\nTrao cho anh bao yêu thương rồi em lại bỏ đi\nGieo cho anh bao nhiêu niềm đau rồi em mau rời bỏ anh xa anh\nQuay mặt lặng lẽ quên mau\nUh em yêu quên thật rồi\nUh chẳng một lời chia ly\nQuên rồi em yêu quên rồi quên rồi\nVậy thì người cứ bước đi xa nơi này\nÁnh bình minh sẽ không còn nơi đây\nBước đi xa nơi này\nNhững lời yêu sẽ không còn nơi đây\nPhải tự đứng lên mà thôi\nChe nhẹ đi những niềm đau và nỗi buồn\nXung quanh anh giờ đây cô đơn mình anh ôm giấc mơ (là vì ai)\nVậy thì người cứ bước đi xa nơi này\nÁnh bình minh sẽ không còn nơi đây\nBước đi xa nơi này\nNhững lời yêu sẽ không còn nơi đây\nPhải tự đứng lên mà thôi\nChe nhẹ đi những niềm đau và nỗi buồn\nXung quanh anh giờ đây cô đơn mình anh ôm giấc mơ\nNhìn em bước ra đi xa dần (eh)\nNhìn em bước ra đi xa dần (eh)\nNhìn em nhìn em bước đi huh (eh)\nNhìn em bước ra đi xa (eh)\nNắng ấm xa dần rồi\nNắng ấm xa dần rồi\nNắng ấm xa dần bỏ rơi để lại những giấc mơ\nGiữ lại đi giữ giữ lại đi\nNắng ấm xa dần rồi\nNắng ấm xa dần rồi\nNắng ấm xa dần xa dần theo những tiếng cười\nHãy mang đi giúp những nỗi buồn',192,'2022-12-17 22:35:54.147','2022-12-17 22:35:54.147','https://www.youtube.com/watch?v=488ceQWoGGw&ab_channel=S%C6%A1nT%C3%B9ngM-TPOfficial',NULL,'http://res.cloudinary.com/dbk0cmzcb/video/upload/v1671291353/lgxmwzcp4lfq4fcf4pmy.mp3','dd96679b4824c4284f70d2c8d6d4b0c2'),(39,'hai triệu năm',10,9,'Anh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nXung quanh anh toàn là nước, ay\nCơ thể anh đang bị ướt, ay\nMênh mông toàn là nước, ay\nÊm ái như chưa từng trước đây\nTrăm ngàn con sóng xô-ya\nAnh lao vào trong biển cả vì em làm anh nóng khô-ya\nAnh ngâm mình trong làn nước để mặn mòi từ da dẻ (mặn mòi từ da dẻ)\nTa cần tình yêu vì tình yêu làm cho ta trẻ (đúng rồi)\nAnh cũng cần em nhưng không biết em sao\nAnh không care lắm và anh quyết đem trao\nCho em hết nắng cho em hết đêm sao\nNhìn mặt anh đi, em nghĩ anh tiếc em sao?\nTrăm ngàn con sóng từ mọi nơi mà đổ về\nVà đây là cách mà anh đi tìm kiếm sự vỗ về\nEm có quá nhiều bí mật, anh thì không cần gặng hỏi\nEm sâu như là đại dương, anh thì không hề lặn giỏi (anh thì không hề lặn giỏi baby)\nAnh soi mình vào gương cho bõ công lau\nThấy mặt thấy người sao thấy rõ trong nhau\nÁnh mắt nụ cười kia không rõ nông sâu\nTa rồi sẽ là ai, một câu hỏi nhỏ trong đầu\nTa chỉ là hòn đất hay chỉ là cỏ bông lau\nNhư là mấy gã em mới bỏ không lâu\nHay chỉ là đầu thuốc kia cháy đỏ không lâu\nYêu em kiểu nông dân, yêu em kiểu quê mùa\nYêu từ vụ đông xuân, đến hè thu thay mùa\nNhưng em thì trơn trượt như là con cá chuối\nMuốn níu em trong tay, Khá Bảnh cũng khá đuối (Khá Bảnh cũng khá đuối)\nEm giống hệt như biển cả em có nhiều bí mật\nAnh làm rất nhiều thứ, để đồng tiền trong ví chật\nNgười ta không quý con ong, mà người ta chỉ quý mật\nEm hỏi anh nhạc sao hay, anh gọi nó là bí thuật, yo\nEm hỏi anh nhạc sao hay, anh gọi nó là bí thuật, yo\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nNước đã hình thành trong hàng triệu năm (triệu năm)\nCát đã hình thành trong hàng triệu năm (triệu năm)\nBiển cả hình thành trong hàng triệu năm (triệu năm)\nSao em làm anh buồn sau hàng triệu năm? (triệu năm)\nGặp em từ thể đơn bào, rồi tiến hoá\nXa em từ khi thềm lục địa đầy biến hoá\nMuốn được ôm em qua kỷ Ju-ra\nHoá thạch cùng nhau trên những phiến đá (phá đá cùng nhau)\nRồi loài người tìm thấy lửa, anh lại tìm thấy em\nAnh tưởng rằng mọi thứ sẽ được bùng cháy lên\nMuốn được cùng em, trồng rau bên hồ cá\nNhưng tim em lúc đó, đang là thời kì đồ đá\nAnh đã tin vào em như tin vào thuyết nhật tâm\nNhư Ga-li-lê người ta nói anh thật hâm\nCó lẽ Đác-win biết biển cả sẽ khô hơn\nNhưng anh tin ông ta không biết chúng ta đang tiến hoá để cô đơn (tiến hoá để cô đơn)\nVà có lẽ Đác-win biết biển cả sẽ khô hơn\nNhưng anh tin ông ta không biết chúng ta đang tiến hoá để cô đơn (tiến hoá để cô đơn)\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\nAnh cô đơn giữa tinh không này\nMuôn con sóng cuốn xô vào đây\nEm cô đơn giữa mênh mông người\nVà ta cô đơn đã hai triệu năm\n',217,'2022-12-17 23:45:54.158','2022-12-17 23:45:54.158','https://www.youtube.com/watch?v=LSMDNL4n0kM&ab_channel=%C4%90enV%C3%A2uOfficial',NULL,'http://res.cloudinary.com/dbk0cmzcb/video/upload/v1671295553/ozmewl41icn5myjpewp6.mp3','b3e6ab22ceaf5b874123d2e904a26171');
/*!40000 ALTER TABLE `songs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-18  2:50:21
