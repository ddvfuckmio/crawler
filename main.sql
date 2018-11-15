CREATE TABLE `courses` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`courseId` varchar(50) NOT NULL,
	`title` varchar(50) NOT NULL,
	`firstCategory` varchar(50) NOT NULL,
	`secondCategory` varchar(50) NOT NULL,
	`introduction` varchar(100) DEFAULT NULL,
	`playCount` int(11),
	`createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
	 PRIMARY KEY (`id`),
	 UNIQUE KEY `courseId` (`courseId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;