-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 28-02-2020 a las 12:29:59
-- Versión del servidor: 10.4.11-MariaDB
-- Versión de PHP: 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `proyecto_go`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id_user` int(10) UNSIGNED NOT NULL,
  `name` varchar(150) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(250) COLLATE utf8_unicode_ci NOT NULL,
  `token` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT 'Token de sesión',
  `register_date` int(8) NOT NULL DEFAULT 0,
  `lostpass_temporal_password` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `lostpass_url_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` (`id_user`, `name`, `email`, `password`, `token`, `register_date`, `lostpass_temporal_password`, `lostpass_url_token`) VALUES
(1, 'Test', 'test@demo.com', '23888723564258e6023dca2dae66cbec', 'c9835695faf518a0583b2a043e844427690e3cfc33a15706a8435d01813839a3ef79d7e59c9ad6aa423ed7f9f8be6d41186b', 1581647963, NULL, NULL),
(2, 'Armando', 'armando@demo.com', 'e10adc3949ba59abbe56e057f20f883e', 'bf276251600f57d136c1e953e475c9ce8c9a0417888d02e827351b01473c53561aa4f282cc0b871f68f0b37c49a7d2f27afe', 1582145591, NULL, NULL),
(3, 'Test 2', 'test2@demo.com', 'e10adc3949ba59abbe56e057f20f883e', '789389027cb328b4daeef98188828a38761d4a8a3a17733f230563f85fb44d49eee8f23bd64382b6450a6a9e452dab556563', 1582732217, NULL, NULL),
(6, 'Test 7', 'test5@demo.com', 'e10adc3949ba59abbe56e057f20f883e', '1593ceb8a5dc299ee1e1c9564c53b19bb97579e0283adfd5786047992de6a3084a5e13c2ea68c77279e6208fbfaf7698f616', 1582738696, NULL, NULL);

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `token` (`token`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
