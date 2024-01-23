CREATE TABLE `devbook`.`users` (
    `user_id` INT NOT NULL AUTO_INCREMENT, 
    `name` VARCHAR(30) NOT NULL , 
    `nick` VARCHAR(30) NOT NULL UNIQUE, 
    `email` VARCHAR(30) NOT NULL UNIQUE, 
    `password` VARCHAR(100) NOT NULL, 
    `register` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(), 
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB;

CREATE TABLE `devbook`.`images` (
    `image_id` INT NOT NULL AUTO_INCREMENT, 
    `filename` VARCHAR(200) NOT NULL,
    `user_id` INT NOT NULL,
    PRIMARY KEY (`image_id`)
) ENGINE = InnoDB;

ALTER TABLE `devbook`.`images`
ADD CONSTRAINT `fk_images_user_id`
FOREIGN KEY (`user_id`)
REFERENCES `devbook`.`users` (`user_id`);