go channel 分享

CREATE TABLE `workflow` (
  `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `count` bigint(20) DEFAULT NULL,
  `status` bigint(20) DEFAULT NULL,
  `last_exec_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



SELECT * FROM `apps` WHERE `id` = 1277;

SELECT count(*) FROM `apps`;