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
    id       INT         NOT NULL AUTO_INCREMENT,
    name     VARCHAR(20) NOT NULL,
    birthday DATE        NOT NULL,
    email    VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE rental_books
(
    id              INT    NOT NULL AUTO_INCREMENT,
    book_id         INT    NOT NULL,
    user_id         INT    NOT NULL,
    loan_date       DATE   NOT NULL,
    return_date     DATE   DEFAULT NULL,
    return_deadline DATE   NOT NULL,
    completion      BIT(1) NOT NULL DEFAULT b'0',
    PRIMARY KEY (id),
    FOREIGN KEY fk_rental_books_book_id (book_id)
        REFERENCES books (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT,
    FOREIGN KEY fk_users_user_id (user_id)
        REFERENCES users (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
);
