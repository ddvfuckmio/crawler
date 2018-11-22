CREATE TABLE `courses` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`courseId` varchar(50) NOT NULL,
	`title` varchar(255),
	`author` varchar(50),
	`firstCategory` varchar(50) NOT NULL,
	`secondCategory` varchar(50) NOT NULL,
	`introduction` varchar(100) DEFAULT NULL,
	`playCount` int(100) unsigned,
	`createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
	 PRIMARY KEY (`id`),
	 UNIQUE KEY `courseId_firstCategory_secondCategory` (`courseId`,`firstCategory`,`secondCategory`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;