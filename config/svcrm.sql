CREATE DATABASE `sv_crm` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;

use sv_crm;

CREATE TABLE `allocation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `ContractID` varchar(45) NOT NULL,
  `CoEntityID` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `ProfileID` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Allocation` float NOT NULL,
  `Status` tinyint(4) NOT NULL DEFAULT '0',
  `ContractType` varchar(45) DEFAULT NULL,
  `Relation` varchar(45) DEFAULT NULL,
  `JsonObject` json DEFAULT NULL,
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`uuid`,`CoEntityID`,`ProfileID`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `coentity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CoEntityID` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CompanyNm` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `AliasNm` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `State` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Country` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Email` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `SecretKey` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Password` varchar(245) COLLATE utf8_unicode_ci NOT NULL,
  `Status` tinyint(4) NOT NULL DEFAULT '0',
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`uuid`,`CoEntityID`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `contracts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ContractID` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CoEntityID` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Version` varchar(245) DEFAULT NULL,
  `EffectiveDate` varchar(256) NOT NULL,
  `ContractType` varchar(45) NOT NULL,
  `JsonObject` json DEFAULT NULL,
  `SecretKey` varchar(245) NOT NULL,
  `Status` tinyint(4) NOT NULL DEFAULT '0',
  `JsonBlock` varchar(745) DEFAULT NULL,
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;





CREATE TABLE `CountryList` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ISOCode` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CountryNm` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ISOCode_UNIQUE` (`ISOCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;




CREATE TABLE `Profile` (
  `id` tinyint(1) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CoEntityID` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `profilename` varchar(18) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `password_hash` char(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `EmailVerified` tinyint(3) unsigned DEFAULT '0',
  `ContactNo` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `PhoneVerified` tinyint(3) unsigned DEFAULT '0',
  `Status` tinyint(4) NOT NULL DEFAULT '0',
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`CoEntityID`,`profilename`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



CREATE TABLE `StateList` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ISOCode` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `StateNm` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `SubDivision` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `CountryISOCode` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ISOCode_UNIQUE` (`ISOCode`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

