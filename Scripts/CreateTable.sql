CREATE TABLE `User` (
    `id_user` int AUTO_INCREMENT NOT NULL ,
    `Name` varchar(100)  NOT NULL ,
    `Mail` varchar(150)  NOT NULL ,
    `Passaword` VARCHAR(256)  NOT NULL ,
    `CreationDate` datetime  NOT NULL ,
    `IsActive` bit  NOT NULL ,
    `IsValided` bit  NOT NULL ,
    PRIMARY KEY (
        `id_user`
    )
);

CREATE TABLE `Restaurateur` (
    `id_restaurateur` int AUTO_INCREMENT NOT NULL ,
    `id_user` int  NOT NULL ,
    `Description` varchar(255`user`)  NOT NULL ,
    `Telephone` varchar(10)  NOT NULL ,
    `id_address` int  NOT NULL ,
    `SIRET` varchar(10)  NOT NULL ,
    `SIREN` varchar(20)  NOT NULL ,
    PRIMARY KEY (
        `id_restaurateur`
    )
);

CREATE TABLE `Livreur` (
    `id_livreur` int AUTO_INCREMENT NOT NULL ,
    `id_user` int  NOT NULL ,
    `Pseudo` varchar(100)  NOT NULL ,
    `Telephone` varchar(10)  NOT NULL ,
    PRIMARY KEY (
        `id_livreur`
    )
);

CREATE TABLE `Admin` (
    `id_admim` int  NOT NULL ,
    `id_user` int  NOT NULL ,
    PRIMARY KEY (
        `id_admim`
    )
);

CREATE TABLE `Customer` (
    `id_customer` int AUTO_INCREMENT NOT NULL ,
    `id_user` int  NOT NULL ,
    `Pseudo` varchar(100)  NOT NULL ,
    `Telephone` varchar(10)  NOT NULL ,
    PRIMARY KEY (
        `id_customer`
    )
);

CREATE TABLE `Address` (
    `id_address` int AUTO_INCREMENT NOT NULL ,
    `Address` varchar(255)  NOT NULL ,
    `CodePostale` varchar(10)  NOT NULL ,
    `Ville` varchar(100)  NOT NULL ,
    PRIMARY KEY (
        `id_address`
    )
);

CREATE TABLE `UserAddress` (
    `id_address` int  NOT NULL ,
    `id_user` int  NOT NULL 
);

ALTER TABLE `Restaurateur` ADD CONSTRAINT `fk_Restaurateur_id_user` FOREIGN KEY(`id_user`)
REFERENCES `User` (`id_user`);

ALTER TABLE `Livreur` ADD CONSTRAINT `fk_Livreur_id_user` FOREIGN KEY(`id_user`)
REFERENCES `User` (`id_user`);

ALTER TABLE `Admin` ADD CONSTRAINT `fk_Admin_id_user` FOREIGN KEY(`id_user`)
REFERENCES `User` (`id_user`);

ALTER TABLE `Customer` ADD CONSTRAINT `fk_Customer_id_user` FOREIGN KEY(`id_user`)
REFERENCES `User` (`id_user`);

ALTER TABLE `Address` ADD CONSTRAINT `fk_Address_id_adress` FOREIGN KEY(`id_address`)
REFERENCES `UserAddress` (`id_address`);

ALTER TABLE `UserAddress` ADD CONSTRAINT `fk_UserAddress_id_user` FOREIGN KEY(`id_user`)
REFERENCES `User` (`id_user`);
