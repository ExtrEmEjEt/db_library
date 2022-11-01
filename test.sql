-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Nov 01, 2022 at 01:51 AM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

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
-- Table structure for table `BOOK`
--

CREATE TABLE `BOOK` (
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
-- Dumping data for table `BOOK`
--

INSERT INTO `BOOK` (`name`, `auths`, `pub_house`, `pub_year`, `count_in_hall1`, `count_in_hall2`, `count_in_hall3`, `id`, `count_hall_1_z`, `count_hall_2_z`, `count_hall_3_z`) VALUES
('Ya durachok', 'Daun', 'Ckazheni', '2005-08-25', 53, 11, 0, '1488q6', 0, 0, 0),
('Dodik', 'n', 'biba', '1994-02-14', 54, 33, 22, 'ABOBA1', 0, 0, 0),
('Война и мир', 'Лев Толстой', 'Русские классики', '1995-06-25', 0, 25, 15, 'LEV001', 0, 0, 0),
('ffff', 'ggggg', 'yfyv', '2022-04-15', 10, 12, 3, 'AAAEEE', 0, 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `HALL`
--

CREATE TABLE `HALL` (
  `number` int(10) UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `capacity` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `HALL`
--

INSERT INTO `HALL` (`number`, `name`, `capacity`) VALUES
(1, 'aboba', 50);

-- --------------------------------------------------------

--
-- Table structure for table `USER`
--

CREATE TABLE `USER` (
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
-- Dumping data for table `USER`
--

INSERT INTO `USER` (`id`, `last_name`, `passport_number`, `date_birth`, `adress`, `phone_number`, `academic_degree`, `hall`, `admission_library`, `book_1`, `book_2`, `book_3`, `date_book_1`, `date_book_2`, `date_book_3`) VALUES
(1, 'Makoviy', 12, '2003-07-01', 'Poskot, dom 54', 8228, 'Loh computer science', 1, '2022-10-31', NULL, NULL, NULL, NULL, NULL, NULL),
(5, 'Иванов', 12345670, '2000-06-07', 'Проспект Пушкина, дом 31', 932341235, 'Бакалавр компьютерных наук', 2, '2022-10-31', NULL, NULL, NULL, NULL, NULL, NULL),
(4, 'fgjfffg', 148822800, '1488-09-05', 'kgkug', 1234567890, 'piyt', 1, '1939-09-01', NULL, NULL, NULL, NULL, NULL, NULL),
(6, 'dodik', 12345432, '2003-09-07', 'fffrty', 1234543212, 'phougb', 1, '2034-09-12', NULL, NULL, NULL, NULL, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `BOOK`
--
ALTER TABLE `BOOK`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `HALL`
--
ALTER TABLE `HALL`
  ADD PRIMARY KEY (`number`);

--
-- Indexes for table `USER`
--
ALTER TABLE `USER`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `HALL`
--
ALTER TABLE `HALL`
  MODIFY `number` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `USER`
--
ALTER TABLE `USER`
  MODIFY `id` int(7) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
