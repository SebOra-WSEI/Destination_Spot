CREATE DATABASE IF NOT EXISTS destination_spot;

USE destination_spot;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL
);

CREATE TABLE spots (
    id INT PRIMARY KEY AUTO_INCREMENT,
    location INT NOT NULL
);

CREATE TABLE reservations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    spot_id INT NOT NULL,
    reserved_from VARCHAR(255) NOT NULL,
    reserved_to VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (spot_id) REFERENCES spots(id)
);

INSERT INTO users (email, password, role, name, surname)
VALUE
    ('sebastian.oraczek@microsoft.wsei.edu.pl', '$2a$12$BrU4TfK9d64R9auinh3KVu0ZuOXPkrklbWjzx24pqp/0paYqn0UKm', 'admin', 'Sebastian', 'Oraczek');

INSERT INTO spots (location)
VALUES (1),(2),(3)