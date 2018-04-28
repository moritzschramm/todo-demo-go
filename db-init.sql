CREATE TABLE IF NOT EXISTS `homestead`.`notes` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `note` VARCHAR(255) NULL,
    `done` BIT(1) NULL,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME NULL,
    PRIMARY KEY (`id`)
)