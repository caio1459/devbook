CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE `devbook`.`users` (
    `user_id` INT NOT NULL AUTO_INCREMENT, 
    `name` VARCHAR(30) NOT NULL , 
    `nick` VARCHAR(30) NOT NULL UNIQUE, 
    `email` VARCHAR(30) NOT NULL UNIQUE, 
    `password` VARCHAR(100) NOT NULL, 
    `image_url` VARCHAR(300), 
    `register` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(), 
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB;

CREATE TABLE `devbook`.`followers`(
    `user_id` INT NOT NULL,
    FOREIGN KEY (`user_id`)
    REFERENCES `devbook`.`users` (`user_id`)
    ON DELETE CASCADE,

    `follower_id` INT NOT NULL,
    FOREIGN KEY (`follower_id`)
    REFERENCES `devbook`.`users` (`user_id`)
    ON DELETE CASCADE,

    PRIMARY KEY(`user_id`, `follower_id`)
) ENGINE = InnoDB;

CREATE TABLE `devbook`.`publications` (
    `pub_id` INT NOT NULL AUTO_INCREMENT, 
    `user_id` INT NOT NULL,
    FOREIGN KEY (`user_id`)
    REFERENCES `devbook`.`users` (`user_id`)
    ON DELETE CASCADE,
    `title` VARCHAR(30) NOT NULL , 
    `content` VARCHAR(300) NOT NULL, 
    `image_url` VARCHAR(300), 
    `likes` INT DEFAULT 0, 
    `register` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(), 
    PRIMARY KEY (`pub_id`)
) ENGINE = InnoDB;