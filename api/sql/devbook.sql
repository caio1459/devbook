CREATE TABLE `devbook`.`users` (
    `user_id` INT NOT NULL AUTO_INCREMENT , 
    `name` VARCHAR(30) NOT NULL , 
    `nick` VARCHAR(30) NOT NULL UNIQUE, 
    `email` VARCHAR(30) NOT NULL UNIQUE, 
    `password` VARCHAR(20) NOT NULL UNIQUE, 
    `image` BLOB NULL , 
    `register` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(), 
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB;