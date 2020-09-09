-- phpMyAdmin SQL Dump
-- version 4.9.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 09, 2020 at 02:05 PM
-- Server version: 5.7.26
-- PHP Version: 7.4.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+07:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bobobox`
--

--
-- Dumping data for table `hotel`
--

INSERT INTO `hotel` (`id`, `hotel_name`, `address`) VALUES
(1, 'Bobobox Pods Alun-alun Bandung', 'Jl. Kepatihan No.8, Balonggede, Kec. Regol, Kota Bandung, West Java 40251, Indonesia'),
(2, 'Bobobox Pods Paskal', 'Jalan Pasirkaliki No.76a, Pasirkaliki, Bandung City, West Java, Indonesia'),
(3, 'Bobobox Pods Cipaganti', 'Jl. Cipaganti No.14, Pasir Kaliki, Cicendo, Kota Bandung, Jawa Barat 40171, Indonesia');

--
-- Dumping data for table `reservation`
--

INSERT INTO `reservation` (`id`, `order_id`, `room_id`, `customer_name`, `checkin_date`, `checkout_date`, `hotel_id`) VALUES
(1, 1, NULL, 'rojali budi permadi', '2020-01-12', '2020-01-14', 1),
(2, 2, NULL, 'lufiana kurnia', '2020-01-13', '2020-01-14', 1),
(3, 3, NULL, 'irfansyah', '2020-01-16', '2020-01-17', 1),
(4, 4, NULL, 'natasha', '2020-01-12', '2020-09-13', 2);

--
-- Dumping data for table `room`
--

INSERT INTO `room` (`id`, `hotel_id`, `room_number`, `room_status`) VALUES
(1, 1, 101, 'available'),
(2, 1, 102, 'available'),
(3, 1, 103, 'available'),
(4, 1, 104, 'out of service'),
(5, 1, 105, 'available'),
(6, 2, 101, 'available'),
(7, 2, 102, 'available'),
(8, 2, 103, 'available'),
(9, 2, 104, 'available'),
(10, 3, 101, 'available'),
(11, 3, 102, 'available'),
(12, 3, 103, 'available'),
(13, 3, 104, 'available'),
(14, 3, 105, 'available'),
(15, 3, 106, 'available'),
(16, 3, 107, 'available');

--
-- Dumping data for table `stay`
--

INSERT INTO `stay` (`id`, `reservation_id`, `room_id`, `guest_name`) VALUES
(1, 1, 1, 'rojali'),
(2, 1, 2, 'budi'),
(3, 2, 3, 'lufiana'),
(4, 3, 1, 'irfansyah'),
(5, 4, 6, 'natasha');

--
-- Dumping data for table `stay_room`
--

INSERT INTO `stay_room` (`id`, `room_id`, `stay_id`, `date`) VALUES
(1, 1, 1, '2020-01-12'),
(2, 1, 1, '2020-01-13'),
(3, 2, 2, '2020-01-12'),
(4, 2, 2, '2020-01-13'),
(5, 3, 3, '2020-01-13'),
(6, 1, 4, '2020-01-16'),
(7, 6, 5, '2020-01-12');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
