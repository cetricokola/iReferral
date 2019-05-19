-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 19, 2019 at 02:16 PM
-- Server version: 10.1.40-MariaDB
-- PHP Version: 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ireferral`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin_account`
--

CREATE TABLE `admin_account` (
  `id` varchar(8) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `first_name` varchar(20) NOT NULL,
  `last_name` varchar(20) NOT NULL,
  `position` varchar(6) NOT NULL,
  `emp_id` varchar(6) NOT NULL,
  `code` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `employee_id`
--

CREATE TABLE `employee_id` (
  `emp_id` varchar(6) NOT NULL,
  `email` varchar(60) NOT NULL,
  `phone_no` varchar(10) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `hospital_account`
--

CREATE TABLE `hospital_account` (
  `code` varchar(6) NOT NULL,
  `name` varchar(20) NOT NULL,
  `serial_no` varchar(10) NOT NULL,
  `phone_no` varchar(10) NOT NULL,
  `email` varchar(60) NOT NULL,
  `country` varchar(20) NOT NULL,
  `region` varchar(20) NOT NULL,
  `district` varchar(20) NOT NULL,
  `mgn_id` varchar(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `patient_account`
--

CREATE TABLE `patient_account` (
  `first_name` varchar(20) NOT NULL,
  `last_name` varchar(20) NOT NULL,
  `huduma_no` varchar(11) NOT NULL,
  `do_b` date NOT NULL,
  `phone_no` varchar(10) NOT NULL,
  `password` varchar(128) NOT NULL,
  `sex` varchar(7) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `patient_diagnosis`
--

CREATE TABLE `patient_diagnosis` (
  `id` varchar(128) NOT NULL,
  `huduma_no` varchar(11) NOT NULL,
  `weight` varchar(5) NOT NULL,
  `temperature` varchar(5) NOT NULL,
  `blood_pressure` varchar(5) NOT NULL,
  `diagnosis` text NOT NULL,
  `response` text NOT NULL,
  `prescription` text NOT NULL,
  `reg_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `referrals`
--

CREATE TABLE `referrals` (
  `id` varchar(128) NOT NULL,
  `huduma_no` varchar(11) NOT NULL,
  `service` varchar(20) NOT NULL,
  `hos_name` varchar(20) NOT NULL,
  `r_date` date NOT NULL,
  `r_time` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `service_code` varchar(6) NOT NULL,
  `code` varchar(6) NOT NULL,
  `name` varchar(20) NOT NULL,
  `cost` varchar(10) NOT NULL,
  `slots` varchar(3) NOT NULL,
  `department` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin_account`
--
ALTER TABLE `admin_account`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`emp_id`);

--
-- Indexes for table `employee_id`
--
ALTER TABLE `employee_id`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `hospital_account`
--
ALTER TABLE `hospital_account`
  ADD PRIMARY KEY (`code`),
  ADD UNIQUE KEY `serial_no` (`serial_no`),
  ADD UNIQUE KEY `phone_no` (`phone_no`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `patient_account`
--
ALTER TABLE `patient_account`
  ADD PRIMARY KEY (`huduma_no`),
  ADD UNIQUE KEY `phone_no` (`phone_no`);

--
-- Indexes for table `patient_diagnosis`
--
ALTER TABLE `patient_diagnosis`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `referrals`
--
ALTER TABLE `referrals`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`service_code`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
