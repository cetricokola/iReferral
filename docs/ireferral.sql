-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 26, 2019 at 07:39 PM
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
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` varchar(128) NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`) VALUES
('', 'cetric', '$2a$10$wnZpeHKilchPtvGx2guopuqV18j4mT3nSD6Un2IqlRXLu6sybmN2u');

-- --------------------------------------------------------

--
-- Table structure for table `admin_account`
--

CREATE TABLE `admin_account` (
  `username` varchar(20) NOT NULL,
  `id` varchar(8) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `admin_account`
--

INSERT INTO `admin_account` (`username`, `id`, `email`, `password`) VALUES
('brian', '33202560', 'x@y1ahoo.com', '$2a$10$InV3dTccVj4qvOthfxignOf3yuiqqP/0uZZWN6VgzGXNujZXfVgTe'),
('cetik', '33202561', 'c@gmail.com', '$2a$10$CnUvCxI2b5qPJv9xtyPcpOBqWAHHooZDOqFfSORoR8vcgXNcxiaJO'),
('okola', '33202564', 'okola@gmail.com', '$2a$10$eQfvRVuOhdznWmLynDVnquaZyZ59M5shFTUrglWfeoBQ7gIQC.wqC'),
('cet', '33202567', 'c3etokola2015@gmail.com', '$2a$10$T9mZqAjPlTPv7QI6wGXXMOkQZHZgZABsdmZ3dU5XZ.R.DG5JywKJa'),
('kcgh', '33202570', 'jon@gmail.com', '$2a$10$tlRbW8aBejB9m23x3/YjZ.eP1b5ZbovTZwXvR5GtKHTYnYwfjcYRq'),
('nomwyo', '33202571', 'jden@gmail.com', '$2a$10$f3pNYZUZLMkO1ExOJs2n5esQt1Tq1wlnFGFNjEbA7hTmtv2ALVvLG'),
('nams', '33202572', 'jden3@gmail.com', '$2a$10$Wcv/0JfyMpODbMqZd6MNYumXF0J3LAbxCBfDvJ551yNFa1q.7v2l6'),
('gracy', '33202573', 'gracy@yahoo.com', '$2a$10$EV2j0gIbXP.8Gjxxumz9K.bUEjOiIddxsPKBB8Jli28MMITCes/ua'),
('kendy', '33202574', 'kendy4@outlook.com', '$2a$10$ujfS5FZ69xVsDEcVlqq5LObUpjQiTp8p4q3zLweKtqJL6TCx0yYRS'),
('pato123', '33202575', 'pat@hotmail.com', '$2a$10$3kyjB3v7raCcnayZGyGzbOTVYmE92IiLxedO7duNUoy2m9WUUOlp6'),
('james', '33202576', 'njames@gmail.com', '$2a$10$c5IVIrNLvrnAOkUdy1RyWefxYWf4KAOxEgd7eTKF3oCS6hEosAcda'),
('elvios', '33202578', 'elvios@yahoo.com', '$2a$10$PyIyqdIGnWQJgifqEgLV7u73uIA0LNOZKS.AI4MtE8hFvwhZ7aZ7e');

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `first_name` varchar(20) NOT NULL,
  `last_name` varchar(20) NOT NULL,
  `emp_id` varchar(6) NOT NULL,
  `code` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `employee`
--

INSERT INTO `employee` (`first_name`, `last_name`, `emp_id`, `code`) VALUES
('justus', 'luganu', '100100', '999333'),
('Daniel', 'Okola', '100101', '999333'),
('Clement', 'Mutai', '100102', '999333'),
('der', 'der', '100103', '999333'),
('der', 'der', '100104', '999333'),
('Bonface', 'Musali', '333666', '');

-- --------------------------------------------------------

--
-- Table structure for table `employee_account`
--

CREATE TABLE `employee_account` (
  `emp_id` varchar(6) NOT NULL,
  `email` varchar(60) NOT NULL,
  `phone_no` varchar(10) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `employee_account`
--

INSERT INTO `employee_account` (`emp_id`, `email`, `phone_no`, `password`) VALUES
('100101', 'cet@gmail.com', '0704145833', '$2a$10$wwXDssYvyMf6GvWJYJin/uxxgzEEh12e20RcZvVBTrDnFOdj2C8HO'),
('100102', 'cetokola@gmail.com', '0704145831', '$2a$10$roaW8IlXodordqmSiF1G2uJ789kopi7t/Da3w77EOsgIUNQ/iyCHS'),
('100100', 'okola@gmail.com', '0704145832', '$2a$10$u3mh46NI5kO8MdA1jPFs.O0ulvBqJVBZ0NrdkrOQmhdRofss68Vja');

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

--
-- Dumping data for table `hospital_account`
--

INSERT INTO `hospital_account` (`code`, `name`, `serial_no`, `phone_no`, `email`, `country`, `region`, `district`, `mgn_id`) VALUES
('', '', '', '', '', '', '', '', '33202578'),
('010102', 'Ben', 'asdfgh', '0704145847', 'x@yahoo.com', 'Kenya', 'Baringo', 'South', '33202561'),
('010103', 'kidn', '5sdfgh', '0704145845', 'x@y1ahoo.com', 'Kenya', 'Baringo', 'East', '33202560'),
('100450', 'Katiku', 'Cet454', '0714581432', 'k@gmail.com', 'Kenya', 'Nakuru', 'Njoro', '00000000'),
('456456', 'Cet', '147cet', '0704145837', 'ceto1kola2015@gmail.com', 'Kenya', 'Baringo', 'North', '33202567'),
('999333', 'Kayle Happy', '123cet', '0704581437', 'cetokola2015@gmail.com', 'Kenya', 'Baringo', 'South', '33202564');

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

--
-- Dumping data for table `patient_account`
--

INSERT INTO `patient_account` (`first_name`, `last_name`, `huduma_no`, `do_b`, `phone_no`, `password`, `sex`) VALUES
('Brian', 'Malala', '12345678910', '2010-06-09', '0704145810', '$2a$10$wsQpr2yQ1.uEKbP/IFIcHeHQ.YnU.wkhKMoBGhjAUjxsxBMPhvPmG', 'Male'),
('cetric', 'okola', '12345678991', '2004-12-14', '0704581433', '$2a$10$E1q2s7ffgZmKshJjVpK45OW9rBhiQrJWOMD2DqbwphqPAhU9MZ4Ha', 'Male'),
('Brian', 'Kemboi', '12345678992', '2019-06-01', '0704145831', '$2a$10$Ttw.sa.W/kXyIYCf58blROWidxF04pGyWXB3MMoxNjVN1/o/iapvm', 'Male'),
('Elvis', 'Mutende', '12345678993', '2019-06-03', '0704145833', '$2a$10$PIpy/FdMGzGRDT3uxPYiYu4F5YiVubZQukjE5JIwjMOiAhz0bgy5e', 'Male'),
('Mercy', 'Mueni', '12345678994', '2019-06-12', '0704145834', '$2a$10$jZD/aJ9kCNcTgiFNatnGUu/ymJtvvCRbOujMYUltaJAPQCR/RSK4y', 'Female'),
('Bridgit', 'Kulubi', '12345678995', '2019-06-17', '0704145835', '$2a$10$E07cwpFOpQDjmOqnS4dm8.sU6XN4XbrfZ/yuPF6M9rhPPVbrKZYYu', 'Female'),
('Raziah', 'Murunga', '12345678996', '2004-06-07', '0704145836', '$2a$10$N/8U2fs48OZKX5h6ETwIAuFLVApotP7v4AOQ4oMpmqoCKdAZXluSy', 'Male'),
('Stellah', 'Mukhwana', '12345678997', '2009-06-09', '0704145837', '$2a$10$HdNnMWku0DmhhbI0WxSczuEH45Tmo.5lrfB0rBptJdBfG./AHPbVq', 'Female'),
('Katiku', 'Jones', '12345678998', '1995-06-07', '0704145838', '$2a$10$WTNwh6aH74vPx4NjCAX4KexlVvNo46b4SNyXVxlVxEbVcdeda6K7K', 'Male'),
('dan', 'james', '12345678999', '1996-06-12', '0704581432', '$2a$10$6lVc5ZfCNnUXOcFM2N16l.tvVToKkgriyT57jLH8EmmVkTqZHZvya', 'Male');

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
  `r_time` time NOT NULL,
  `refer_hos` varchar(20) NOT NULL,
  `refer_by` varchar(6) NOT NULL,
  `email` varchar(60) NOT NULL,
  `phone` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `referrals`
--

INSERT INTO `referrals` (`id`, `huduma_no`, `service`, `hos_name`, `r_date`, `r_time`, `refer_hos`, `refer_by`, `email`, `phone`) VALUES
('237c5fd8-16cd-4231-99ab-ab744df69f37', '12345678999', 'KIDNEY Analysis', 'Ben', '2019-06-05', '00:00:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437'),
('38fad392-02c8-4c9c-89d7-516770608699', '12345678999', 'Kidney analysis', 'Kayle Happy', '2019-05-30', '00:00:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437'),
('784118a1-6dcc-418a-80e4-6e0f52330854', '12345678999', 'KIDNEY Analysis', 'Ben', '2019-06-20', '00:05:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437'),
('81c04ed6-78e2-48f1-8903-7549d0c950be', '12345678999', 'Kidney analysis', 'Kayle Happy', '2019-05-31', '00:00:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437'),
('9c4555a0-b783-417c-94a1-bf1635725d5e', '12345678999', 'kidney analysis', 'Ben', '2019-06-06', '00:00:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437'),
('bd719110-743b-483a-848f-deb2acd90c44', '12345678999', 'KIdney analysis', 'Ben', '2019-06-07', '01:05:00', 'Kayle Happy', '100102', 'cetokola2015@gmail.com', '0704581437');

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `service_code` varchar(6) NOT NULL,
  `code` varchar(6) NOT NULL,
  `name` varchar(20) NOT NULL,
  `department` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`service_code`, `code`, `name`, `department`) VALUES
('010102', '010102', 'kidney analysis', ' Nephrology'),
('010108', '010103', 'den', 'Obstetrics and gynae'),
('456456', '456456', 'den', 'Nutrition and dietet'),
('666998', '999333', 'praisal', ' Nephrology'),
('666999', '999333', 'kidney analysis', 'Urology');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `admin_account`
--
ALTER TABLE `admin_account`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`emp_id`);

--
-- Indexes for table `employee_account`
--
ALTER TABLE `employee_account`
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
