CREATE TABLE `user_classification_words` (
    `id` int NOT NULL AUTO_INCREMENT,
    `word` varchar(10) NOT NULL,
    `weight` int NOT NULL,
    `group_id` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci