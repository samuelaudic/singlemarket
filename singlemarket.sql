-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : ven. 07 juin 2024 à 09:45
-- Version du serveur : 8.3.0
-- Version de PHP : 8.2.18

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `singlemarket`
--

-- --------------------------------------------------------

--
-- Structure de la table `clients`
--

DROP TABLE IF EXISTS `clients`;
CREATE TABLE IF NOT EXISTS `clients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` text COLLATE utf8mb4_unicode_ci,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Déchargement des données de la table `clients`
--

INSERT INTO `clients` (`id`, `first_name`, `last_name`, `phone`, `address`, `email`) VALUES
(1, 'Samuel', 'Audic', '0101001', 'ploerdut', 'samuel@test.fr'),
(2, 'Maximos', 'Petitos', '121212212', 'Paname', 'maximepetit@gmail.com'),
(3, 'Client', 'Test', '00000000', 'adresse test', 'client@test.fr');

-- --------------------------------------------------------

--
-- Structure de la table `orders`
--

DROP TABLE IF EXISTS `orders`;
CREATE TABLE IF NOT EXISTS `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `client_id` int DEFAULT NULL,
  `product_id` int DEFAULT NULL,
  `quantity` int NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `purchase_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `client_id` (`client_id`),
  KEY `product_id` (`product_id`)
) ENGINE=MyISAM AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Déchargement des données de la table `orders`
--

INSERT INTO `orders` (`id`, `client_id`, `product_id`, `quantity`, `price`, `purchase_date`) VALUES
(1, 0, 0, 0, 0.00, '2024-06-06 12:03:56'),
(2, 10, 5, 10, 100.00, '2024-06-06 12:05:54'),
(3, 1, 5, 1, 0.90, '2024-06-06 12:09:06'),
(4, 1, 5, 10, 100.00, '2024-06-06 12:57:47'),
(5, 1, 5, 5, 4.50, '2024-06-06 13:14:58'),
(6, 2, 2, 10, 10.10, '2024-06-06 13:16:57'),
(7, 1, 7, 10, 20.10, '2024-06-07 05:13:08'),
(8, 1, 7, 10, 20.10, '2024-06-07 06:19:35'),
(9, 1, 7, 10, 20.10, '2024-06-07 06:20:45'),
(10, 1, 7, 20, 40.20, '2024-06-07 06:35:38'),
(11, 2, 2, 2, 2.02, '2024-06-07 06:36:15'),
(12, 1, 7, 10, 20.10, '2024-06-07 06:52:04'),
(13, 1, 1, 10, 0.00, '2024-06-07 07:09:19'),
(14, 1, 2, 10, 10.10, '2024-06-07 07:12:35'),
(15, 1, 1, 10, 0.00, '2024-06-07 07:13:30'),
(16, 1, 2, 10, 10.10, '2024-06-07 07:16:06'),
(17, 1, 2, 10, 10.10, '2024-06-07 07:22:26');

-- --------------------------------------------------------

--
-- Structure de la table `products`
--

DROP TABLE IF EXISTS `products`;
CREATE TABLE IF NOT EXISTS `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `price` decimal(10,2) NOT NULL,
  `quantity` int NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Déchargement des données de la table `products`
--

INSERT INTO `products` (`id`, `title`, `description`, `price`, `quantity`, `active`) VALUES
(1, '', '', 0.00, 0, 0),
(2, 'Pomme', 'pomme', 1.01, 10, 1),
(3, 'Fraise', 'Une', 0.00, 0, 1),
(4, '', 'pomme', 10.00, 10, 0),
(5, 'Citron', 'Un citron tron tron', 0.90, 120, 1),
(6, 'Mangue', 'Une mangue', 1.50, 200, 1),
(7, 'Melon', 'un melon', 2.01, 10, 1),
(8, 'Pastèque', 'une papastèque', 1.50, 100, 1);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
