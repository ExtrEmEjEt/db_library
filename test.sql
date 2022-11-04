-- phpMyAdmin SQL Dump
-- version 5.1.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Nov 04, 2022 at 11:00 PM
-- Server version: 5.7.24
-- PHP Version: 8.0.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test`
--

-- --------------------------------------------------------

--
-- Table structure for table `book`
--

CREATE TABLE `book` (
  `name` varchar(256) NOT NULL,
  `auths` text NOT NULL,
  `pub_house` text NOT NULL,
  `pub_year` date NOT NULL,
  `count_in_hall1` int(10) UNSIGNED NOT NULL,
  `count_in_hall2` int(10) UNSIGNED NOT NULL,
  `count_in_hall3` int(10) UNSIGNED NOT NULL,
  `id` varchar(6) NOT NULL,
  `count_hall_1_z` int(10) UNSIGNED DEFAULT '0',
  `count_hall_2_z` int(10) UNSIGNED DEFAULT '0',
  `count_hall_3_z` int(10) UNSIGNED DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `book`
--

INSERT INTO `book` (`name`, `auths`, `pub_house`, `pub_year`, `count_in_hall1`, `count_in_hall2`, `count_in_hall3`, `id`, `count_hall_1_z`, `count_hall_2_z`, `count_hall_3_z`) VALUES
('Кобзар', 'Тарас Григорович Шевченко', 'Україньська класична література', '1989-01-26', 20, 30, 12, 'TARAS0', 0, 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `hall`
--

CREATE TABLE `hall` (
  `number` int(10) UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `capacity` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `hall`
--

INSERT INTO `hall` (`number`, `name`, `capacity`) VALUES
(1, 'aboba', 50),
(2, 'birka', 100),
(3, 'algeciras', 75);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(7) UNSIGNED NOT NULL,
  `last_name` varchar(30) NOT NULL,
  `passport_number` int(9) UNSIGNED NOT NULL,
  `date_birth` date NOT NULL,
  `adress` text NOT NULL,
  `phone_number` int(10) UNSIGNED NOT NULL,
  `academic_degree` varchar(255) NOT NULL,
  `hall` int(5) UNSIGNED NOT NULL,
  `admission_library` date NOT NULL,
  `book_1` varchar(255) DEFAULT NULL,
  `book_2` varchar(255) DEFAULT NULL,
  `book_3` varchar(255) DEFAULT NULL,
  `date_book_1` date DEFAULT NULL,
  `date_book_2` date DEFAULT NULL,
  `date_book_3` date DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `last_name`, `passport_number`, `date_birth`, `adress`, `phone_number`, `academic_degree`, `hall`, `admission_library`, `book_1`, `book_2`, `book_3`, `date_book_1`, `date_book_2`, `date_book_3`) VALUES
(1, 'Makoviy', 12439765, '2003-07-01', 'Poskot, dom 54', 822843238, 'Loh computer science', 1, '2022-10-31', NULL, NULL, NULL, NULL, NULL, NULL),
(5, 'Иванов', 12345670, '2000-06-07', 'Проспект Пушкина, дом 31', 932341235, 'Бакалавр компьютерных наук', 2, '2022-10-31', NULL, NULL, NULL, NULL, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `hall`
--
ALTER TABLE `hall`
  ADD PRIMARY KEY (`number`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hall`
--
ALTER TABLE `hall`
  MODIFY `number` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(7) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
