-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 28-Out-2025 às 14:49
-- Versão do servidor: 10.4.27-MariaDB
-- versão do PHP: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `produtos`
--

-- --------------------------------------------------------

--
-- Estrutura da tabela `produto`
--

CREATE TABLE `produto` (
  `id_produto` int(11) NOT NULL,
  `descricao` varchar(60) NOT NULL,
  `preco` decimal(10,4) NOT NULL,
  `marca` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Extraindo dados da tabela `produto`
--

INSERT INTO `produto` (`id_produto`, `descricao`, `preco`, `marca`) VALUES
(1, 'Notebook Gamer 15.6\" i7 16GB RAM SSD 512GB', '5899.9000', 'Acer'),
(2, 'Smartphone Galaxy S23 256GB 8GB RAM', '4299.0000', 'Samsung'),
(3, 'Smart TV 55\" 4K UHD HDR', '3499.9900', 'LG'),
(4, 'Fone de Ouvido Bluetooth Noise Cancelling', '899.9000', 'Sony'),
(5, 'Teclado Mecânico RGB Switch Azul', '349.9900', 'Redragon'),
(6, 'Mouse Gamer 16000 DPI RGB', '249.5000', 'Logitech'),
(7, 'Cadeira Gamer Reclinável com Apoio de Pé', '1299.0000', 'ThunderX3'),
(8, 'Monitor 27\" 144Hz Full HD Curvo', '1899.0000', 'AOC'),
(9, 'SSD NVMe 1TB Gen4', '749.9000', 'Kingston'),
(10, 'Placa de Vídeo RTX 4070 12GB GDDR6X', '4999.9900', 'NVIDIA');

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `produto`
--
ALTER TABLE `produto`
  ADD PRIMARY KEY (`id_produto`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `produto`
--
ALTER TABLE `produto`
  MODIFY `id_produto` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
