-- phpMyAdmin SQL Dump
-- version 4.9.7
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 12, 2023 at 04:59 PM
-- Server version: 10.3.37-MariaDB-log-cll-lve
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bageurte_store_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `nik` varchar(25) DEFAULT '0',
  `nama` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `no_tlp` varchar(255) DEFAULT NULL,
  `jenis_kelamin` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `nik`, `nama`, `alamat`, `no_tlp`, `jenis_kelamin`) VALUES
(1, '3241203322112233', 'Taupik Pirdian', 'Bandung', '0858877899887', 'LK'),
(2, '3200442233442324', 'Yogi', 'Jakarta', '08133221122388', 'LK'),
(3, '3423441223322122', 'Rico', 'Jakarta', '0873382224932', 'LK');

-- --------------------------------------------------------

--
-- Table structure for table `item`
--

CREATE TABLE `item` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `kategori` varchar(255) DEFAULT NULL,
  `harga` bigint(20) DEFAULT 0,
  `jumlah` int(20) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `item`
--

INSERT INTO `item` (`id`, `nama`, `kategori`, `harga`, `jumlah`) VALUES
(1, 'Service Console Xbox Sereis X', 'Service Console', 4673000, 999),
(2, 'Service Console PS 5', 'Service Console', 4117000, 999),
(3, 'Service Console Xbox Series S', 'Service Console', 3110000, 999),
(4, 'Service Console PS4', 'Service Console', 2227000, 999),
(5, 'Service Console Nintendo Switch', 'Service Console', 973000, 999),
(6, 'Xbox Series X', 'Buy New Console', 9000000, 99),
(7, 'Play Station 5', 'Buy New Console', 8800000, 99),
(8, 'Xbox Series S', 'Buy New Console', 4350000, 99),
(9, 'Play Station 4', 'Buy New Console', 4750000, 99),
(10, 'Nintendo Switch', 'Buy New Console', 3890000, 99),
(11, 'God of War Ragnarok', 'Buy New Game', 1029000, 99),
(12, 'FIFA 23', 'Buy New Game', 759000, 99),
(13, 'Call Of Duty Modern Warfare II', 'Buy New Game', 1500000, 99),
(14, 'Horizon Forbidden West', 'Buy New Game', 799000, 99),
(15, 'NBA 2K23', 'Buy New Game', 1100000, 99),
(16, 'God of War Ragnarok', 'Buy Second Game', 568000, 99),
(17, 'FIFA 23', 'Buy Second Game', 345000, 99),
(18, 'Call Of Duty Modern Warfare II', 'Buy Second Game', 749000, 99),
(19, 'Horizon Forbidden West', 'Buy Second Game', 376000, 99),
(20, 'NBA 2K23', 'Buy Second Game', 547000, 99),
(21, 'Controller PS 5', 'Buy Accessories Console', 1000000, 99),
(22, 'Nintendo Switch Pro Controller', 'Buy Accessories Console', 870000, 99),
(23, 'Oculus Quest 2 Advanced All-In-One VR', 'Buy Accessories Console', 6198000, 99),
(24, 'Ransel PS 5', 'Buy Accessories Console', 149000, 99),
(25, 'Silicon Case Controller PS 5', 'Buy Accessories Console', 249000, 99);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(16) DEFAULT 0,
  `code_transaction` varchar(255) DEFAULT NULL,
  `tanggal_pembelian` date DEFAULT NULL,
  `total` bigint(20) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `customer_id`, `code_transaction`, `tanggal_pembelian`, `total`) VALUES
(1, 1, 'INV-20D12M2022Y14H16M17S', '2022-10-10', 129070000),
(2, 1, 'INV-20D12M2022Y14H16M51S', '2022-10-10', 102700000),
(3, 2, 'INV-21D12M2022Y14H16M51S', '2022-10-10', 4673000),
(84, 1, 'INV-23D12M2022Y17H44M55S', '2022-10-10', 67315000),
(85, 1, 'INV-23D12M2022Y18H03M28S', '2022-10-10', 30000000),
(86, 1, 'INV-23D12M2022Y19H30M21S', '2022-10-10', 30000000),
(87, 1, 'INV-20D12M2022Y14H16M17T', '2022-12-23', 5000000);

-- --------------------------------------------------------

--
-- Table structure for table `transaction_detail`
--

CREATE TABLE `transaction_detail` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `transaction_id` int(20) DEFAULT 0,
  `item_id` int(20) DEFAULT 0,
  `jumlah_pembelian` int(20) DEFAULT 0,
  `harga_pembelian` bigint(20) DEFAULT 0,
  `harga_discount` bigint(20) DEFAULT 0,
  `total` bigint(20) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `transaction_detail`
--

INSERT INTO `transaction_detail` (`id`, `transaction_id`, `item_id`, `jumlah_pembelian`, `harga_pembelian`, `harga_discount`, `total`) VALUES
(1, 1, 1, 10, 4673000, 0, 46730000),
(2, 1, 2, 20, 4117000, 0, 82340000),
(3, 2, 1, 10, 4673000, 7009500, 39720500),
(4, 2, 2, 20, 4117000, 19360500, 62979500),
(5, 3, 1, 1, 4673000, 0, 4673000),
(144, 2, 1, 1, 4360000, 0, 4360000),
(145, 84, 1, 10, 4673000, 0, 46730000),
(146, 84, 2, 5, 4117000, 0, 20585000),
(147, 85, 21, 30, 1000000, 0, 30000000),
(148, 86, 21, 30, 1000000, 0, 30000000);

-- --------------------------------------------------------

--
-- Table structure for table `voucher_customer`
--

CREATE TABLE `voucher_customer` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(16) DEFAULT 0,
  `transaction_id` int(11) DEFAULT 0,
  `code_voucher` varchar(255) DEFAULT NULL,
  `nama_voucher` varchar(255) DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `use_date` date DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `nilai_disc` int(255) DEFAULT 0,
  `type_disc` int(2) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `voucher_customer`
--

INSERT INTO `voucher_customer` (`id`, `customer_id`, `transaction_id`, `code_voucher`, `nama_voucher`, `start_date`, `end_date`, `use_date`, `status`, `nilai_disc`, `type_disc`) VALUES
(102, 1, 0, 'PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6', 'PREMI', '2022-12-20', '2023-12-20', '2022-12-22', 0, 15, 2),
(103, 1, 0, 'PREMI-20D12M2022Y14H16M51SKFGQO4YLR1GWR', 'PREMI', '2022-12-20', '2023-12-20', '0000-00-00', 1, 15, 2),
(106, 1, 0, 'PREMI-23D12M2022Y17H44M55S73NF7JB6CG57X', 'PREMI', '2022-12-23', '2023-12-23', '0000-00-00', 1, 15, 2),
(107, 1, 0, 'ULTI-23D12M2022Y18H03M28S73NF7JB6CG57X', 'ULTI', '2022-12-23', '2023-12-23', '0000-00-00', 1, 30, 1),
(108, 1, 0, 'PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M7', 'PREMI', '2022-12-23', '2022-12-23', '2022-12-23', 1, 15, 2),
(109, 1, 0, 'ULTI-23D12M2022Y19H30M21S73NF7JB6CG57X', 'ULTI', '2022-12-23', '2023-12-23', '0000-00-00', 1, 30, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `item`
--
ALTER TABLE `item`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `voucher_customer`
--
ALTER TABLE `voucher_customer`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=44;

--
-- AUTO_INCREMENT for table `item`
--
ALTER TABLE `item`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=88;

--
-- AUTO_INCREMENT for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=149;

--
-- AUTO_INCREMENT for table `voucher_customer`
--
ALTER TABLE `voucher_customer`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=110;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
