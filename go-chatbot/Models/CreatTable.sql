CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `phone` varchar(64) NOT NULL DEFAULT 'unknow',
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `creat_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `birthday` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `intr` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '这个人很懒，没有简介',
    `imgBlob` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'user.png',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `friend` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `friend_id` bigint(20) NOT NULL,
    `state` tinyint(4) NOT NULL,
    `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `userIdInFriend` FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
    CONSTRAINT `friendIdInFriend` FOREIGN KEY (`friend_id`) REFERENCES `user`(`user_id`)  ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `singleMessage` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `send_id` bigint(20) NOT NULL,
    `accept_id` bigint(20) NOT NULL,
    `state` bool NOT NULL DEFAULT false,
    `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `message` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `type` tinyint(4) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `sendId` FOREIGN KEY (`send_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
    CONSTRAINT `acceptId` FOREIGN KEY (`accept_id`) REFERENCES `user`(`user_id`)  ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `group` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `group_id` bigint(20) NOT NULL,
     `groupname` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
     `imgBlob` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'group.png',
     `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `notice` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '这个群没有公告',
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_group_id` (`group_id`) USING BTREE,
     UNIQUE KEY `idx_groupname` (`groupname`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `groupMember` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `group_id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `state` tinyint(4) NOT NULL,
    `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `shield` bool NOT NULL DEFAULT false,
    PRIMARY KEY (`id`),
    CONSTRAINT `groupIdInMember` FOREIGN KEY (`group_id`) REFERENCES `group`(`group_id`) ON DELETE CASCADE,
    CONSTRAINT `acceptIdInMember` FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`)  ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `groupMessage` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `group_id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `message` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `type` tinyint(4) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `groupIdInMessage` FOREIGN KEY (`group_id`) REFERENCES `group`(`group_id`) ON DELETE CASCADE,
    CONSTRAINT `acceptIdInMessage` FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`)  ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;