DROP DATABASE IF EXISTS sample_db;
CREATE DATABASE sample_db;
USE sample_db;

CREATE TABLE books
(
    id     INT         NOT NULL AUTO_INCREMENT,
    title  VARCHAR(20) NOT NULL,
    author VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE users
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    birthday DATE NOT NULL,
    email VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);
