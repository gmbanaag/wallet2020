CREATE TABLE IF NOT EXISTS `transactions` (
  `id` varchar(255) NOT NULL,
  `source_user_id` varchar(255) NOT NULL,
  `destination_user_id` varchar(255) DEFAULT NULL,
  `source_wallet_id` varchar(255) NOT NULL,
  `destination_wallet_id` varchar(255) NOT NULL,
  `device_id` varchar(255) NOT NULL,
  `amount` double DEFAULT NULL,
  `message` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `create_ts` bigint NOT NULL,
  `update_ts` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `wallets` (
  `id` varchar(255) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `country_iso` varchar(255) NOT NULL,
  `currency` varchar(255) NOT NULL,
  `balance` double NOT NULL,
  `create_ts` bigint NOT NULL,
  `update_ts` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;