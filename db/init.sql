CREATE TABLE IF NOT EXISTS ipv4_networks
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    network      INT UNSIGNED NOT NULL,
    subnet_mask  INT UNSIGNED NOT NULL,
    country_code CHAR(2)      NOT NULL
);

CREATE TABLE ipv6_networks
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    network      VARBINARY(16) NOT NULL,
    subnet_mask  INT UNSIGNED  NOT NULL,
    country_code CHAR(2)       NOT NULL
);

