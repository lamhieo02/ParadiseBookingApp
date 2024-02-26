/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `full_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `role` int DEFAULT NULL,
  `status` int DEFAULT '2',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(12) DEFAULT NULL,
  `dob` varchar(50) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `is_email_verified` tinyint(1) NOT NULL DEFAULT '0',
  `bio` varchar(555) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `account_type_id` (`role`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `amenities`;
CREATE TABLE `amenities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `place_id` int DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `config_amenity_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`place_id`,`config_amenity_id`) USING BTREE,
  KEY `place_id` (`place_id`) USING BTREE,
  KEY `config_amenity_id` (`config_amenity_id`) USING BTREE,
  CONSTRAINT `amenities_ibfk_1` FOREIGN KEY (`config_amenity_id`) REFERENCES `config_amenity` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `booking_rating`;
CREATE TABLE `booking_rating` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `rating` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int NOT NULL,
  `booking_id` int NOT NULL,
  `place_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `booking_id` (`booking_id`) USING BTREE,
  KEY `booking_user_id` (`booking_id`,`user_id`) USING BTREE,
  KEY `place_id` (`place_id`) USING BTREE,
  KEY `place_user_id` (`place_id`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `bookings`;
CREATE TABLE `bookings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `place_id` int DEFAULT NULL,
  `status_id` int DEFAULT NULL,
  `checkout_date` datetime DEFAULT NULL,
  `checkin_date` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_email` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `place_id` (`place_id`) USING BTREE,
  KEY `status_id` (`status_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `bookings_detail`;
CREATE TABLE `bookings_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `booking_id` int NOT NULL,
  `full_name` varchar(50) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `type` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `guest_name` varchar(50) DEFAULT NULL,
  `content_to_vendor` varchar(255) DEFAULT NULL,
  `total_price` float DEFAULT NULL,
  `time_to` varchar(25) DEFAULT NULL,
  `time_from` varchar(25) DEFAULT NULL,
  `number_of_guest` int DEFAULT NULL,
  `payment_method` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `booking_id` (`booking_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `config_amenity`;
CREATE TABLE `config_amenity` (
  `id` int NOT NULL AUTO_INCREMENT,
  `icon` varchar(255) DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `group_policy`;
CREATE TABLE `group_policy` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `payment_method`;
CREATE TABLE `payment_method` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `payment_status`;
CREATE TABLE `payment_status` (
  `id` int NOT NULL AUTO_INCREMENT,
  `status` varchar(50) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `booking_id` int DEFAULT NULL,
  `method_id` int DEFAULT NULL,
  `status_id` int DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `request_id` varchar(500) DEFAULT NULL,
  `order_id` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `booking_id` (`booking_id`) USING BTREE,
  KEY `method_id` (`method_id`) USING BTREE,
  KEY `status_id` (`status_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `place_wishlist`;
CREATE TABLE `place_wishlist` (
  `place_id` int NOT NULL,
  `wishlist_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`place_id`,`wishlist_id`),
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
  `id` int NOT NULL AUTO_INCREMENT,
  `vendor_id` int DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price_per_night` double DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `district` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `max_guest` int DEFAULT '1',
  `bed_room` int DEFAULT NULL,
  `num_bed` int DEFAULT NULL,
  `payment_method` int DEFAULT NULL,
  `num_place_original` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `vendor_id` (`vendor_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `policies`;
CREATE TABLE `policies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `place_id` int DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `group_policy_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`place_id`,`group_policy_id`) USING BTREE,
  KEY `place_id` (`place_id`) USING BTREE,
  KEY `group_policy_id` (`group_policy_id`) USING BTREE,
  CONSTRAINT `policies_ibfk_1` FOREIGN KEY (`group_policy_id`) REFERENCES `group_policy` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `reports`;
CREATE TABLE `reports` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `vendor_id` int DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status_id` int DEFAULT NULL,
  `type_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `vendor_id` (`vendor_id`) USING BTREE,
  KEY `status_id` (`status_id`) USING BTREE,
  KEY `type_id` (`type_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `requests`;
CREATE TABLE `requests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `vendor_id` int DEFAULT NULL,
  `status_id` int DEFAULT NULL,
  `title` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `status_id` (`status_id`) USING BTREE,
  KEY `vendor_id` (`vendor_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `verify_emails`;
CREATE TABLE `verify_emails` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `scret_code` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `expired_at` timestamp NULL DEFAULT NULL,
  `type` int DEFAULT NULL COMMENT '1: verify mail; 2: reset code password',
  PRIMARY KEY (`id`),
  KEY `index_email_code_type` (`email`,`scret_code`,`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `wishlists`;
CREATE TABLE `wishlists` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(55) DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE DEFINER=`root`@`%` PROCEDURE `GetAverageRatingByPlaceId`(IN placeId INT)
BEGIN
SELECT AVG(rating) AS average_rating
FROM booking_rating
WHERE place_id = placeId;
END;

CREATE DEFINER=`root`@`%` PROCEDURE `GetBookingsWithinRange`(IN date_from DATETIME, IN date_to DATETIME)
BEGIN
    SELECT * FROM bookings
    WHERE (checkin_date BETWEEN date_from AND date_to)
       OR (checkout_date BETWEEN date_from AND date_to)
       OR (checkin_date <= date_from AND checkout_date >= date_to);
END;

CREATE DEFINER=`root`@`%` PROCEDURE `GetCommentsAndRatingsByVendorId`(IN vendorId INT)
BEGIN
    SELECT *
    FROM booking_rating 
    WHERE place_id IN (SELECT id FROM places WHERE vendor_id = vendorId);
END;

CREATE DEFINER=`root`@`%` PROCEDURE `GetPaymentsForVendor`(IN p_vendor_id INT, IN p_page INT, IN p_limit INT)
BEGIN 

    DECLARE offset_value INT DEFAULT (p_page - 1) * p_limit;

    SELECT payments.* 
    FROM bookings 
    INNER JOIN places ON bookings.place_id = places.id 
    INNER JOIN payments ON payments.booking_id = bookings.id 
    WHERE places.vendor_id = p_vendor_id
    LIMIT p_limit OFFSET offset_value;
END;

CREATE DEFINER=`root`@`%` PROCEDURE `GetPaymentsSizeOfVendor`(IN p_vendor_id INT)
BEGIN 
    SELECT count(1)
    FROM bookings 
    INNER JOIN places ON bookings.place_id = places.id 
    INNER JOIN payments ON payments.booking_id = bookings.id 
    WHERE places.vendor_id = p_vendor_id;
END;

CREATE DEFINER=`root`@`%` PROCEDURE `GetRatingStatisticByPlaceId`(IN placeId INT)
BEGIN
  SELECT
    r.rating,
    COUNT(br.rating) AS count
  FROM
    (SELECT 1 AS rating UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5) r
  LEFT JOIN
    booking_rating br ON r.rating = br.rating AND br.place_id = placeId
  WHERE
    br.place_id IS NULL OR br.place_id = placeId
  GROUP BY
    r.rating;
END;

INSERT INTO `accounts` (`id`, `username`, `email`, `full_name`, `role`, `status`, `password`, `address`, `phone`, `dob`, `avatar`, `created_at`, `updated_at`, `deleted_at`, `is_email_verified`, `bio`) VALUES
(100, '', 'lamlklk2002@gmail.com', '', 2, 2, '$2a$10$XmSVEplY82T6qwTSJ37mgOV2IuchlaFKPOaFFBWphY1WHXXXPu6iG', '', '', '', '', '2024-01-01 14:57:28', '2024-01-01 15:51:56', NULL, 1, '');
INSERT INTO `accounts` (`id`, `username`, `email`, `full_name`, `role`, `status`, `password`, `address`, `phone`, `dob`, `avatar`, `created_at`, `updated_at`, `deleted_at`, `is_email_verified`, `bio`) VALUES
(101, '', 'lamfan011@gmail.com', '', 3, 2, '$2a$10$lI3jysiTusMK9PYVmAWDaelPub8U4RC/cLM4O2EMWNtASvLwsfA62', '', '', '', '', '2024-01-01 15:02:03', '2024-01-01 15:31:02', NULL, 1, '');
INSERT INTO `accounts` (`id`, `username`, `email`, `full_name`, `role`, `status`, `password`, `address`, `phone`, `dob`, `avatar`, `created_at`, `updated_at`, `deleted_at`, `is_email_verified`, `bio`) VALUES
(102, '', 'admin@gmail.com', '', 3, 2, '$2a$10$zMNDhf80wgPYZinTWsUXF.e.HsGUkL1uJEx1JhN3JRYOJjnQDRKFK', '', '', '', '', '2024-01-01 15:23:33', '2024-01-01 15:23:50', NULL, 1, '');
INSERT INTO `accounts` (`id`, `username`, `email`, `full_name`, `role`, `status`, `password`, `address`, `phone`, `dob`, `avatar`, `created_at`, `updated_at`, `deleted_at`, `is_email_verified`, `bio`) VALUES
(103, '', 'leminhtuong09122002@gmail.com', '', 1, 2, '$2a$10$PEI2rcj9jxFCrxTVTDkFReTSfaT4Hm0Tutvf9cn7yK7byLJWybs/O', '', '', '', 'https://d1c6814z4iseiw.cloudfront.net/images/leminhtuong.jpg', '2024-01-01 16:32:16', '2024-01-03 14:27:01', NULL, 1, ''),
(104, 'MTLe', 'mt09122002@gmail.com', 'Vendor so1', 2, 2, '$2a$10$ie//r7vjgJfwBl9CQYIYAOWmyM03rhDo8/jiy5OD0wFNyVKPpvxmO', '', '', '', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh3.jpg', '2024-01-01 16:36:19', '2024-01-01 16:38:24', NULL, 1, 'New Vendor'),
(105, 'lamhieo02', '20110668@student.hcmute.edu.vn', 'Nguyễn Văn Lâm', 2, 2, '$2a$10$ShzkHEC31HwwzyE/IZ0nUu/pCJTyyt6xB4OlOan8SluP3SmIGqRo2', 'Quy Nhon ', '0962981939', '2002-06-02', 'https://d1c6814z4iseiw.cloudfront.net/images/Screenshot 2023-12-25 232919.png', '2024-01-01 16:45:22', '2024-01-02 02:17:31', NULL, 1, 'lamhieooo'),
(106, '', 'hoanglen5@gmail.com', '', 1, 2, '$2a$10$eZjw3VTX3DUWvsUft7JLKu.Do1sqcEi87hk68fEzWSj947KV0cdQ2', '', '', '', '', '2024-02-15 10:21:43', '2024-02-15 10:21:43', NULL, 0, '');

INSERT INTO `amenities` (`id`, `place_id`, `description`, `created_at`, `updated_at`, `config_amenity_id`) VALUES
(79, 87, 'Wifi', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 3);
INSERT INTO `amenities` (`id`, `place_id`, `description`, `created_at`, `updated_at`, `config_amenity_id`) VALUES
(80, 87, 'Garden view', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 1);
INSERT INTO `amenities` (`id`, `place_id`, `description`, `created_at`, `updated_at`, `config_amenity_id`) VALUES
(81, 87, 'Dedicated workspace', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 7);
INSERT INTO `amenities` (`id`, `place_id`, `description`, `created_at`, `updated_at`, `config_amenity_id`) VALUES
(82, 87, 'Hot water', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 2),
(83, 87, 'Safe', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 8),
(84, 87, 'Fire extinguisher', '2024-01-01 17:08:09', '2024-01-01 17:08:09', 10),
(85, 91, 'Security cameras on property', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 5),
(86, 91, 'Free parking on premises', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 9),
(87, 91, 'Wifi', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 3),
(88, 91, 'Garden view', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 1),
(89, 91, 'Dedicated workspace', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 7),
(90, 91, 'Fire extinguisher', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 10),
(91, 91, 'Safe', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 8),
(92, 91, 'Bathtub', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 6),
(93, 91, 'Coffee', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 4),
(94, 91, 'Hot water', '2024-01-01 17:33:02', '2024-01-01 17:33:02', 2),
(95, 77, 'Bathtub', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 6),
(96, 77, 'Coffee', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 4),
(97, 77, 'Safe', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 8),
(98, 77, 'Wifi', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 3),
(99, 77, 'Security cameras on property', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 5),
(100, 77, 'Dedicated workspace', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 7),
(101, 77, 'Free parking on premises', '2024-01-01 17:37:38', '2024-01-01 17:37:38', 9);

INSERT INTO `booking_rating` (`id`, `title`, `content`, `rating`, `created_at`, `updated_at`, `user_id`, `booking_id`, `place_id`) VALUES
(19, 'Đánh giá của tôi', 'Không gian không có gì phải chê, nội thất lý tưởng cho nhu cầu làm việc của tôi 10đ', 5, '2024-01-01 17:28:15', '2024-01-01 17:28:15', 103, 142, 91);
INSERT INTO `booking_rating` (`id`, `title`, `content`, `rating`, `created_at`, `updated_at`, `user_id`, `booking_id`, `place_id`) VALUES
(20, 'Lam Test', 'Place so beautiful, everyone is very good ', 5, '2024-01-01 17:49:52', '2024-01-01 17:59:00', 105, 145, 80);
INSERT INTO `booking_rating` (`id`, `title`, `content`, `rating`, `created_at`, `updated_at`, `user_id`, `booking_id`, `place_id`) VALUES
(21, 'LovelyPlace', 'Everything is good', 4, '2024-01-02 02:19:39', '2024-01-02 02:19:39', 100, 147, 80);

INSERT INTO `bookings` (`id`, `user_id`, `place_id`, `status_id`, `checkout_date`, `checkin_date`, `created_at`, `updated_at`, `user_email`) VALUES
(140, 100, 74, 6, '2024-01-03 00:00:00', '2024-01-01 00:00:00', '2024-01-01 15:42:28', '2024-01-03 00:00:02', NULL);
INSERT INTO `bookings` (`id`, `user_id`, `place_id`, `status_id`, `checkout_date`, `checkin_date`, `created_at`, `updated_at`, `user_email`) VALUES
(141, 100, 73, 2, '2024-01-04 00:00:00', '2024-01-01 00:00:00', '2024-01-01 16:09:10', '2024-01-01 16:09:45', NULL);
INSERT INTO `bookings` (`id`, `user_id`, `place_id`, `status_id`, `checkout_date`, `checkin_date`, `created_at`, `updated_at`, `user_email`) VALUES
(142, 103, 91, 5, '2024-02-06 00:00:00', '2024-01-31 00:00:00', '2024-01-01 17:20:51', '2024-01-01 17:26:35', NULL);
INSERT INTO `bookings` (`id`, `user_id`, `place_id`, `status_id`, `checkout_date`, `checkin_date`, `created_at`, `updated_at`, `user_email`) VALUES
(143, 105, 80, 6, '2024-01-10 00:00:00', '2024-01-02 00:00:00', '2024-01-01 17:34:53', '2024-01-03 00:00:02', NULL),
(144, 103, 77, 2, '2024-01-26 00:00:00', '2024-01-10 00:00:00', '2024-01-01 17:36:12', '2024-01-01 17:36:34', NULL),
(145, 105, 80, 5, '2024-01-02 00:00:00', '2024-01-02 00:00:00', '2024-01-01 17:36:38', '2024-01-01 17:49:11', NULL),
(147, 100, 80, 2, '2024-01-03 00:00:00', '2024-01-02 00:00:00', '2024-01-01 18:03:55', '2024-01-02 05:05:58', NULL),
(148, 105, 86, 6, '2024-01-05 00:00:00', '2024-01-02 00:00:00', '2024-01-02 03:21:04', '2024-01-04 00:00:02', NULL),
(149, 105, 86, 6, '2024-01-05 00:00:00', '2024-01-02 00:00:00', '2024-01-02 03:21:10', '2024-01-04 00:00:02', NULL),
(150, 105, 85, 6, '2024-01-04 00:00:00', '2024-01-02 00:00:00', '2024-01-02 03:21:54', '2024-01-04 00:00:02', NULL);

INSERT INTO `bookings_detail` (`id`, `booking_id`, `full_name`, `phone`, `email`, `type`, `created_at`, `updated_at`, `guest_name`, `content_to_vendor`, `total_price`, `time_to`, `time_from`, `number_of_guest`, `payment_method`) VALUES
(131, 140, 'Nguyễn Văn Lâm', '0962981939', 'phatbfbf@gmail.com', 2, '2024-01-01 15:42:28', '2024-01-01 15:42:28', 'Lam Nguyen ', 'hello vendor', 600000, '', '', 2, 2);
INSERT INTO `bookings_detail` (`id`, `booking_id`, `full_name`, `phone`, `email`, `type`, `created_at`, `updated_at`, `guest_name`, `content_to_vendor`, `total_price`, `time_to`, `time_from`, `number_of_guest`, `payment_method`) VALUES
(132, 141, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-01 16:09:10', '2024-01-01 16:09:10', '', '', 1500000, '', '', 2, 2);
INSERT INTO `bookings_detail` (`id`, `booking_id`, `full_name`, `phone`, `email`, `type`, `created_at`, `updated_at`, `guest_name`, `content_to_vendor`, `total_price`, `time_to`, `time_from`, `number_of_guest`, `payment_method`) VALUES
(133, 142, 'leminhtuong', '0892922911', 'leminhtuong09122002@gmail.com', 1, '2024-01-01 17:20:51', '2024-01-01 17:20:51', '', 'Xin chao Vendor', 3000000, '', '', 1, 1);
INSERT INTO `bookings_detail` (`id`, `booking_id`, `full_name`, `phone`, `email`, `type`, `created_at`, `updated_at`, `guest_name`, `content_to_vendor`, `total_price`, `time_to`, `time_from`, `number_of_guest`, `payment_method`) VALUES
(134, 143, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-01 17:34:53', '2024-01-01 17:34:53', '', 'Nguyen Van Lam 20110668', 8000000, '', '', 2, 2),
(135, 144, 'leminhtuong', '0892922911', 'leminhtuong09122002@gmail.com', 1, '2024-01-01 17:36:12', '2024-01-01 17:36:12', '', 'Xin Chao Nguoi Quen', 4800000, '', '', 1, 1),
(136, 145, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-01 17:36:38', '2024-01-01 17:36:38', '', '20110668 Nguyen Van Lam test', 1000000, '', '', 2, 2),
(137, 146, 'leminhtuong', '0892922911', 'leeminwall091202@gmail.com', 1, '2024-01-01 17:53:10', '2024-01-01 17:53:10', '', '', 4500000, '', '', 2, 2),
(138, 147, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-01 18:03:55', '2024-01-01 18:03:55', '', 'hello lam', 1000000, '', '', 2, 1),
(139, 148, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-02 03:21:04', '2024-01-02 03:21:04', '', '', 3, '', '', 2, 2),
(140, 149, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-02 03:21:10', '2024-01-02 03:21:10', '', '', 3, '', '', 2, 2),
(141, 150, 'Nguyễn Văn Lâm', '0962981939', 'lamlklk2002@gmail.com', 1, '2024-01-02 03:21:54', '2024-01-02 03:21:54', '', '', 1560000, '', '', 4, 2);

INSERT INTO `config_amenity` (`id`, `icon`, `name`, `created_at`, `updated_at`) VALUES
(1, 'GiButterflyFlower', 'Garden view', '2023-12-14 15:18:19', '2023-12-14 17:32:22');
INSERT INTO `config_amenity` (`id`, `icon`, `name`, `created_at`, `updated_at`) VALUES
(2, 'BsFire', 'Hot water', '2023-12-14 15:18:19', '2023-12-14 17:32:22');
INSERT INTO `config_amenity` (`id`, `icon`, `name`, `created_at`, `updated_at`) VALUES
(3, 'AiOutlineWifi', 'Wifi', '2023-12-14 15:18:19', '2023-12-14 17:32:22');
INSERT INTO `config_amenity` (`id`, `icon`, `name`, `created_at`, `updated_at`) VALUES
(4, 'MdOutlineCoffeeMaker', 'Coffee', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(5, 'BiCctv', 'Security cameras on property', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(6, 'MdOutlineBathtub', 'Bathtub', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(7, 'GrWorkshop', 'Dedicated workspace', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(8, 'RiSafeLine', 'Safe', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(9, 'AiOutlineCar', 'Free parking on premises', '2023-12-14 15:18:19', '2023-12-14 15:18:19'),
(10, 'FaFireExtinguisher', 'Fire extinguisher', '2023-12-14 15:18:19', '2023-12-14 15:18:19');

INSERT INTO `group_policy` (`id`, `name`) VALUES
(1, 'House Rules');
INSERT INTO `group_policy` (`id`, `name`) VALUES
(2, 'Safe Rules');
INSERT INTO `group_policy` (`id`, `name`) VALUES
(3, 'Cancel Rules');





INSERT INTO `payments` (`id`, `booking_id`, `method_id`, `status_id`, `amount`, `created_at`, `updated_at`, `request_id`, `order_id`) VALUES
(54, 140, 2, 2, 600000, '2024-01-01 15:42:30', '2024-01-01 15:42:30', '6dbeaffa001b54a', '6dbeaffa000b54a');
INSERT INTO `payments` (`id`, `booking_id`, `method_id`, `status_id`, `amount`, `created_at`, `updated_at`, `request_id`, `order_id`) VALUES
(55, 141, 2, 2, 1500000, '2024-01-01 16:09:11', '2024-01-01 16:09:11', '6dbed713c010002', '6dbed713c000002');
INSERT INTO `payments` (`id`, `booking_id`, `method_id`, `status_id`, `amount`, `created_at`, `updated_at`, `request_id`, `order_id`) VALUES
(56, 142, 1, 1, 3000000, '2024-01-01 17:20:52', '2024-01-01 17:20:52', '', '');
INSERT INTO `payments` (`id`, `booking_id`, `method_id`, `status_id`, `amount`, `created_at`, `updated_at`, `request_id`, `order_id`) VALUES
(57, 143, 2, 2, 8000000, '2024-01-01 17:34:54', '2024-01-01 17:34:54', '6dbf54a44010002', '6dbf54a44000002'),
(58, 144, 1, 1, 4800000, '2024-01-01 17:36:12', '2024-01-01 17:36:12', '', ''),
(59, 145, 2, 2, 1000000, '2024-01-01 17:36:38', '2024-01-01 17:36:38', '6dbf572e4010002', '6dbf572e4000002'),
(60, 146, 2, 2, 4500000, '2024-01-01 17:53:11', '2024-01-01 17:53:11', '6dbf6f6af010002', '6dbf6f6af000002'),
(61, 147, 1, 1, 1000000, '2024-01-01 18:03:56', '2024-01-01 18:03:56', '', ''),
(62, 150, 2, 2, 1560000, '2024-01-02 03:21:54', '2024-01-02 03:21:54', '6dc2b07f8010002', '6dc2b07f8000002');



INSERT INTO `places` (`id`, `vendor_id`, `name`, `description`, `price_per_night`, `address`, `cover`, `lat`, `lng`, `country`, `state`, `district`, `created_at`, `updated_at`, `max_guest`, `bed_room`, `num_bed`, `payment_method`, `num_place_original`) VALUES
(73, 100, 'Home Phu Cat', 'Small and cute home stay ', 500000, 'Tam Binh, Thu Duc, TPHCM', 'https://d1c6814z4iseiw.cloudfront.net/images/Screenshot 2023-11-04 145236.png', 10.8298295, 106.7617899, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 15:20:14', '2024-01-01 15:20:14', 2, 2, 2, NULL, 2);
INSERT INTO `places` (`id`, `vendor_id`, `name`, `description`, `price_per_night`, `address`, `cover`, `lat`, `lng`, `country`, `state`, `district`, `created_at`, `updated_at`, `max_guest`, `bed_room`, `num_bed`, `payment_method`, `num_place_original`) VALUES
(74, 100, 'Quy Nhon Home Stay', 'Small and cute home stay ', 300000, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/9dfaf7d6-40f2-4673-b468-7c5ab3147f86.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 15:40:55', '2024-01-01 15:41:57', 2, 2, 2, NULL, 2);
INSERT INTO `places` (`id`, `vendor_id`, `name`, `description`, `price_per_night`, `address`, `cover`, `lat`, `lng`, `country`, `state`, `district`, `created_at`, `updated_at`, `max_guest`, `bed_room`, `num_bed`, `payment_method`, `num_place_original`) VALUES
(75, 100, 'Home Stay Hihe', 'Small and cute home stay ', 300000, '528 Đường Quang Trung, Quận 3, Thủ Đức, TPHCM', 'https://d1c6814z4iseiw.cloudfront.net/images/23cba0d9-2fcd-4720-a41d-f66092e17a00.webp', 10.778639, 106.6870156, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 3', '2024-01-01 16:16:11', '2024-01-01 16:16:11', 2, 2, 2, NULL, 1);
INSERT INTO `places` (`id`, `vendor_id`, `name`, `description`, `price_per_night`, `address`, `cover`, `lat`, `lng`, `country`, `state`, `district`, `created_at`, `updated_at`, `max_guest`, `bed_room`, `num_bed`, `payment_method`, `num_place_original`) VALUES
(76, 100, 'Home Stay Cat Tien', 'Small and cute home stay ', 600000, 'Xã Cát Tiến, Huyện Phù Cát, Tỉnh Bình Định, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/a177c47e-f7c7-47fa-8506-987230d5ce10.webp', 13.9437796, 109.2386071, 'Vietnam', 'Bình Định', 'Phù Cát', '2024-01-01 16:19:47', '2024-01-01 16:19:47', 2, 2, 2, NULL, 3),
(77, 104, 'Nha Vuon Que', 'Ngoi nha voi kien truc co dien', 300000, 'Ap Phu My, Xa Phu Quy', 'https://d1c6814z4iseiw.cloudfront.net/images/339774131_1063166737974195_1997791306742826169_n.jpg', 10.4065174, 106.1195395, 'Vietnam', 'Tiền Giang', 'Cai Lậy', '2024-01-01 16:41:27', '2024-01-01 16:41:27', 2, 1, 1, NULL, 3),
(78, 104, 'Best Countryside House', 'Best house locates in a countryside', 700000, 'Ap Quy Chanh, Xa Nhi Quy', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh-nen-full-hd-1080-ngoi-nha-don-so-noi-bat-giua-dat-troi_022857575.jpg', 10.4065174, 106.1195395, 'Vietnam', 'Tiền Giang', 'Cai Lậy', '2024-01-01 16:44:27', '2024-01-01 16:44:27', 4, 2, 2, NULL, 2),
(79, 104, 'Vùng băng giá', 'Khí hậu mát mẻ ở miền quê vào mùa đông', 350000, 'Ap Quy Phuoc, Xa Nhi Quy', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh-nen-laptop-full-hd-hinh-anh-rung-cay-ngay-tuyet-roi-dep_034806139.jpg', 10.4065174, 106.1195395, 'Vietnam', 'Tiền Giang', 'Cai Lậy', '2024-01-01 16:47:08', '2024-01-01 16:47:08', 3, 1, 2, NULL, 4),
(80, 105, 'Home Stay Ha Giang', 'Small and cute home stay ', 1000000, 'Hà Giang, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/c3305688-ed60-4af6-8c3e-096434d94bff.webp', 22.7336097, 105.0027271, 'Vietnam', 'Hà Giang', 'Thành phố Hà Giang', '2024-01-01 16:48:22', '2024-01-01 17:36:20', 2, 2, 2, NULL, 2),
(81, 104, 'Góc làm việc', 'Nơi làm việc yên tĩnh cho cá nhân', 300000, 'Hem 530, To Ngoc Van', 'https://d1c6814z4iseiw.cloudfront.net/images/pexels-pixabay-459653.jpg', 10.862434, 106.7441358, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 16:49:11', '2024-01-01 16:49:11', 1, 1, 1, NULL, 6),
(82, 104, 'Hoàng Hôn Tuyệt Nhất', 'Nơi lý tưởng để các cặp đôi ngắm hoàng hôn', 400000, '238/10/6 Hoàng Diệu 2', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh-nen-laptop-full-hd-bai-cat-ven-bien-day-soi-da-dep_034800773.jpg', 10.8586537, 106.7628386, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 16:51:21', '2024-01-01 16:51:21', 4, 2, 2, NULL, 4),
(83, 104, 'Best Mood', 'Nơi để nhóm bạn tụ tập chung vui', 1200000, 'Đường số 7', 'https://d1c6814z4iseiw.cloudfront.net/images/pexels-janez-podnar-1424246.jpg', 10.8596068, 106.7805003, '70000', 'Vietnam', 'Thành phố Hồ Chí Minh', '2024-01-01 16:53:20', '2024-01-01 16:53:20', 5, 3, 3, NULL, 2),
(84, 104, 'Best Night View', 'Nơi tuyệt nhất để ngắm thành phố về đêm', 550000, 'Số 30, Đường Võ Thị Sáu', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh-nen-laptop-full-hd-thanh-pho-da-len-den-dep-lung-linh_034810086.jpg', 10.7773145, 106.6999907, '700000', 'Vietnam', 'Thành phố Hồ Chí Minh', '2024-01-01 16:55:11', '2024-01-01 16:55:11', 2, 1, 2, NULL, 3),
(85, 104, 'Best Camp', 'Nơi lý tưởng cho kỳ nghỉ với nhóm bạn ', 780000, 'Số 50, Đường Nguyễn Tri Phương', 'https://d1c6814z4iseiw.cloudfront.net/images/pexels-aarti-vijay-2693529.jpg', 10.7553616, 106.6685441, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 5', '2024-01-01 16:57:44', '2024-01-01 16:57:44', 4, 2, 2, NULL, 4),
(86, 105, 'Home Stay Thu Duc', 'Small and cute home stay ', 1, 'Tô Ngọc Vân, Thủ Đức, TPHCM, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/23cba0d9-2fcd-4720-a41d-f66092e17a00.webp', 10.8298295, 106.7617899, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 16:58:57', '2024-01-01 16:58:57', 2, 1, 1, NULL, 1),
(87, 105, 'Home Stay Quy Nhon', 'Small and cute home stay ', 8000000, '528 Đường Trần Hưng Đạo, TP.Quy Nhơn, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/7f600b15-cfd7-4b27-9d70-dcd42cbd4b2a.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 17:00:29', '2024-01-01 17:00:29', 2, 2, 2, NULL, 2),
(88, 104, 'White Castle', 'Ngôi nhà lãng mạn cho những cặp đôi', 150000, 'Đường Lê Hồng Phong', 'https://d1c6814z4iseiw.cloudfront.net/images/5306ab09-0e80-494c-8975-08f08fea026b.jpeg', 10.772732, 106.6683666, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 10', '2024-01-01 17:01:22', '2024-01-01 17:01:22', 2, 1, 2, NULL, 2),
(89, 104, 'Best Chill Place', 'Nơi tuyệt vời để trải nghiệm cuộc sống riêng tư', 200000, '669 Đường Tô Ngọc Vân', 'https://d1c6814z4iseiw.cloudfront.net/images/3221922d-7259-43ce-ae73-9602806437c1.png', 10.862434, 106.7441358, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 17:03:08', '2024-01-01 17:03:08', 1, 1, 1, NULL, 2),
(90, 104, 'Sân vườn', 'Không gian sân vườn rộng rãi để trải nghiệm hoạt động ngoài trời', 160000, 'Đường Nguyễn Duy Dương', 'https://d1c6814z4iseiw.cloudfront.net/images/a3857366-a7dd-43de-83f1-5d6fc9aa9423.jpg', 10.7616838, 106.6792631, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 5', '2024-01-01 17:05:10', '2024-01-01 17:05:10', 2, 1, 2, NULL, 1),
(91, 104, 'Không gian thoáng đãng', 'Không gian xung quanh nhà cực kỳ thoáng đãng, phù hợp cho các hoạt động sáng tạo', 500000, 'Đường  Phạm Thái Bường, Phường 9', 'https://d1c6814z4iseiw.cloudfront.net/images/97fabe17-c7b5-4873-8108-136c7683b1c8.jpg', 10.737754800000001, 106.72971280655217, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 7', '2024-01-01 17:08:16', '2024-01-01 17:08:16', 1, 1, 1, NULL, 2),
(92, 104, 'Kiến trúc cổ kính', 'Thiết kế cổ kính, phù hợp cho các hoạt động thiết kế trải nghiệm', 600000, 'Đường Hải Thượng Lãn Ông, Phường 5', 'https://d1c6814z4iseiw.cloudfront.net/images/0d804f06-7c4e-449b-96be-3c312c986da7.png', 10.7553616, 106.6685441, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 5', '2024-01-01 17:09:43', '2024-01-01 17:09:43', 2, 1, 2, NULL, 2),
(93, 104, 'Green Place', 'Màu xanh của cây cối tạo nên sự mát mẻ và rất thiên nhiên', 400000, '20 Đường Phạm Văn Đồng', 'https://d1c6814z4iseiw.cloudfront.net/images/0d603456_original.png', 10.8345635, 106.6739598, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Gò Vấp', '2024-01-01 17:12:16', '2024-01-01 17:12:16', 2, 1, 2, NULL, 2),
(94, 105, 'Home Stay Thu Duc', 'Small and cute home stay ', 1, '528/17/21E Hẻm 530, Tam Phú', 'https://d1c6814z4iseiw.cloudfront.net/images/a177c47e-f7c7-47fa-8506-987230d5ce10.webp', 10.8298295, 106.7617899, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 17:12:41', '2024-01-01 17:12:41', 1, 1, 1, NULL, 1),
(95, 104, 'Best Pool', 'Có hồ bơi phía sau, tuyệt vời cho những ngày nóng nực', 350000, 'Đường Trần Bình Trọng', 'https://d1c6814z4iseiw.cloudfront.net/images/b0ebf8a2_original.png', 10.8345635, 106.6739598, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Gò Vấp', '2024-01-01 17:13:32', '2024-01-01 17:13:32', 1, 1, 1, NULL, 1),
(96, 105, 'Home Stay Nha Lam', 'Small and cute home stay ', 3000000, 'Huyện Phù Cát, Tỉnh Bình Định, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/c3305688-ed60-4af6-8c3e-096434d94bff.webp', 13.95912425, 109.04963163925903, 'Vietnam', 'Bình Định', 'Phù Cát', '2024-01-01 17:14:05', '2024-01-01 17:14:05', 2, 2, 2, NULL, 2),
(97, 104, 'Best Castle', 'Như một lâu đài với thiết kế đặc trưng bắt mắt', 450000, 'Số 30, Đường  Phan Xích Long', 'https://d1c6814z4iseiw.cloudfront.net/images/fc02e72a_original.png', 10.8117887, 106.7039109, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Bình Thạnh', '2024-01-01 17:14:31', '2024-01-01 17:14:31', 1, 1, 1, NULL, 2),
(98, 105, 'Home Stay Nha Lam', 'Small and cute home stay ', 4000000, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/a177c47e-f7c7-47fa-8506-987230d5ce10.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 17:14:57', '2024-01-01 17:14:57', 2, 1, 2, NULL, 2),
(99, 105, 'Home Stay Nha Lam', 'Small and cute home stay ', 1, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/fcb3dc08-0f37-4293-b978-bbd012aff7b5.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 17:17:02', '2024-01-01 17:17:02', 2, 2, 2, NULL, 2),
(100, 105, 'Home Stay Nha Lam', 'Small and cute home stay ', 1, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/17cdb0d7-2d76-4e6f-bf43-b57460d9d682 (1).webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 17:17:47', '2024-01-01 17:17:47', 2, 1, 1, NULL, 5),
(101, 100, 'Home Phu Cat', 'Small and cute home stay ', 500000, 'Tam Binh, Thu Duc, TPHCM', 'https://d1c6814z4iseiw.cloudfront.net/images/Screenshot 2023-11-04 145236.png', 10.8298295, 106.7617899, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 15:20:14', '2024-01-01 15:20:14', 2, 2, 2, NULL, 2),
(102, 100, 'Home Stay Hihe', 'Small and cute home stay ', 300000, '528 Đường Quang Trung, Quận 3, Thủ Đức, TPHCM', 'https://d1c6814z4iseiw.cloudfront.net/images/23cba0d9-2fcd-4720-a41d-f66092e17a00.webp', 10.778639, 106.6870156, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Quận 3', '2024-01-01 16:16:11', '2024-01-01 16:16:11', 2, 2, 2, NULL, 1),
(103, 100, 'Home Stay Cat Tien', 'Small and cute home stay ', 600000, 'Xã Cát Tiến, Huyện Phù Cát, Tỉnh Bình Định, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/a177c47e-f7c7-47fa-8506-987230d5ce10.webp', 13.9437796, 109.2386071, 'Vietnam', 'Bình Định', 'Phù Cát', '2024-01-01 16:19:47', '2024-01-01 16:19:47', 2, 2, 2, NULL, 3),
(105, 100, 'Quy Nhon Home Stay', 'Small and cute home stay ', 300000, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/9dfaf7d6-40f2-4673-b468-7c5ab3147f86.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 15:40:55', '2024-01-01 15:41:57', 2, 2, 2, NULL, 2),
(107, 105, 'Home Stay Thu Duc', 'Small and cute home stay ', 1, 'Tô Ngọc Vân, Thủ Đức, TPHCM, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/23cba0d9-2fcd-4720-a41d-f66092e17a00.webp', 10.8298295, 106.7617899, 'Vietnam', 'Thành phố Hồ Chí Minh', 'Thủ Đức', '2024-01-01 16:58:57', '2024-01-01 16:58:57', 2, 1, 1, NULL, 1),
(108, 105, 'Home Stay Quy Nhon', 'Small and cute home stay ', 8000000, '528 Đường Trần Hưng Đạo, TP.Quy Nhơn, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/7f600b15-cfd7-4b27-9d70-dcd42cbd4b2a.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 17:00:29', '2024-01-01 17:00:29', 2, 2, 2, NULL, 2),
(111, 104, 'Vùng băng giá', 'Khí hậu mát mẻ ở miền quê vào mùa đông', 350000, 'Ap Quy Phuoc, Xa Nhi Quy', 'https://d1c6814z4iseiw.cloudfront.net/images/hinh-nen-laptop-full-hd-hinh-anh-rung-cay-ngay-tuyet-roi-dep_034806139.jpg', 10.4065174, 106.1195395, 'Vietnam', 'Tiền Giang', 'Cai Lậy', '2024-01-01 16:47:08', '2024-01-01 16:47:08', 3, 1, 2, NULL, 4),
(112, 105, 'Home Stay Ha Giang', 'Small and cute home stay ', 1000000, 'Hà Giang, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/c3305688-ed60-4af6-8c3e-096434d94bff.webp', 22.7336097, 105.0027271, 'Vietnam', 'Hà Giang', 'Thành phố Hà Giang', '2024-01-01 16:48:22', '2024-01-01 16:48:22', 2, 2, 2, NULL, 1),
(113, 100, 'Quy Nhon Home Stay', 'Small and cute home stay ', 300000, 'Quy Nhon City, Binh Dinh Province, Viet Nam', 'https://d1c6814z4iseiw.cloudfront.net/images/9dfaf7d6-40f2-4673-b468-7c5ab3147f86.webp', 13.770409, 109.232667, 'Vietnam', 'Bình Định', 'Quy Nhơn', '2024-01-01 15:40:55', '2024-01-01 15:41:57', 2, 2, 2, NULL, 2),
(114, 100, 'Home Stay Nha Trang', 'Small and cute home stay ', 4000000, 'Nha Trang City, Khanh Hoa Province, VietNam', 'https://d1c6814z4iseiw.cloudfront.net/images/7f600b15-cfd7-4b27-9d70-dcd42cbd4b2a.webp', 12.2431693, 109.1898675, '650000', 'Vietnam', 'Khánh Hòa', '2024-01-02 02:14:47', '2024-01-02 02:14:47', 1, 1, 1, NULL, 1);

INSERT INTO `policies` (`id`, `place_id`, `name`, `created_at`, `updated_at`, `group_policy_id`) VALUES
(19, 87, 'Checkin after: 14:00. Checkout before: 12:00', '2024-01-01 17:09:23', '2024-01-01 17:09:55', 1);
INSERT INTO `policies` (`id`, `place_id`, `name`, `created_at`, `updated_at`, `group_policy_id`) VALUES
(20, 87, 'No Smoking', '2024-01-01 17:09:23', '2024-01-01 17:09:23', 2);
INSERT INTO `policies` (`id`, `place_id`, `name`, `created_at`, `updated_at`, `group_policy_id`) VALUES
(21, 87, 'Cancel Before 24 hour', '2024-01-01 17:09:23', '2024-01-01 17:09:23', 3);
INSERT INTO `policies` (`id`, `place_id`, `name`, `created_at`, `updated_at`, `group_policy_id`) VALUES
(22, 91, 'Checkin after: 12:00. Checkout before: 14:00', '2024-01-01 17:35:23', '2024-01-01 17:35:23', 1),
(23, 91, 'Giữ vệ sinh chung\nKhông hút thuốc', '2024-01-01 17:35:23', '2024-01-01 17:35:23', 2),
(24, 91, 'Huỷ phòng khi chưa xác nhận qua email\nLiên hệ với chủ nhà khi có vấn đề', '2024-01-01 17:35:23', '2024-01-01 17:35:23', 3),
(25, 77, 'Checkin after: 14:00. Checkout before: 09:00', '2024-01-01 17:38:43', '2024-01-01 17:38:43', 1),
(26, 77, 'Không gây cháy nổ\nTránh hút thuốc\nKhông mang chất cấm', '2024-01-01 17:38:43', '2024-01-01 17:38:43', 2),
(27, 77, 'Liên hệ chủ nhà\nHuỷ phòng trước khi xác nhận', '2024-01-01 17:38:43', '2024-01-01 17:38:43', 3);





INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(1, 0);


INSERT INTO `verify_emails` (`id`, `email`, `scret_code`, `created_at`, `expired_at`, `type`) VALUES
(57, 'lamlklk2002@gmail.com', 'j7DGHAaC', '2024-01-01 14:57:38', '2024-01-01 15:02:38', 1);
INSERT INTO `verify_emails` (`id`, `email`, `scret_code`, `created_at`, `expired_at`, `type`) VALUES
(58, 'lamfan011@gmail.com', 'nGPE9Xmr', '2024-01-01 15:02:17', '2024-01-01 15:07:17', 1);
INSERT INTO `verify_emails` (`id`, `email`, `scret_code`, `created_at`, `expired_at`, `type`) VALUES
(59, 'admin@gmail.com', '3rKdrb6C', '2024-01-01 15:23:44', '2024-01-01 15:28:44', 1);
INSERT INTO `verify_emails` (`id`, `email`, `scret_code`, `created_at`, `expired_at`, `type`) VALUES
(60, 'leminhtuong09122002@gmail.com', 'ScbHQC0O', '2024-01-01 16:32:27', '2024-01-01 16:37:27', 1),
(61, 'mt09122002@gmail.com', '1RCvr2NI', '2024-01-01 16:36:31', '2024-01-01 16:41:31', 1),
(62, '20110668@student.hcmute.edu.vn', '3TZcrZOO', '2024-01-01 16:45:34', '2024-01-01 16:50:34', 1),
(63, 'hoanglen5@gmail.com', 'qL1HpQE5', '2024-02-15 10:21:57', '2024-02-15 10:26:57', 1);







/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;