CREATE DATABASE IF NOT EXISTS douyin;

Use douyin;

-- 用户表
CREATE TABLE IF NOT EXISTS User
(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_name varchar(64),
    avatar varchar(64),
    follow_count int ,
    follower_count int,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 关注列表
CREATE Table IF NOT EXISTS Relation(
    id int AUTO_INCREMENT PRIMARY KEY,
    from_user int,
    to_user int,
    create_at timestamp,
    UNIQUE KEY `follow`(`from_user`,`to_user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE Table IF NOT EXISTS Friend(
    id int AUTO_INCREMENT PRIMARY KEY,
    from_user int,
    to_user int,
    create_at timestamp,
    UNIQUE KEY `follow`(`from_user`,`to_user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;