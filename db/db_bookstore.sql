-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema db_bookstore
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema db_bookstore
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `db_bookstore` DEFAULT CHARACTER SET latin1 ;
USE `db_bookstore` ;

-- -----------------------------------------------------
-- Table `db_bookstore`.`category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_bookstore`.`category` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;


-- -----------------------------------------------------
-- Table `db_bookstore`.`publisher`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_bookstore`.`publisher` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL DEFAULT NULL,
  `address` VARCHAR(100) NULL DEFAULT NULL,
  `phone_number` VARCHAR(16) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;


-- -----------------------------------------------------
-- Table `db_bookstore`.`book`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_bookstore`.`book` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `publisher_id` INT(11) NULL DEFAULT NULL,
  `category_id` INT(11) NULL DEFAULT NULL,
  `title` VARCHAR(100) NULL DEFAULT NULL,
  `author` VARCHAR(50) NULL DEFAULT NULL,
  `year_of_publication` VARCHAR(4) NULL DEFAULT NULL,
  `stock` INT(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `FK_CATEGORY_ID`
    FOREIGN KEY (`category_id`)
    REFERENCES `db_bookstore`.`category` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_PUBLISHER_ID`
    FOREIGN KEY (`publisher_id`)
    REFERENCES `db_bookstore`.`publisher` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
