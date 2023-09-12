CREATE DATABASE go_gin_api_boilerplate_db;

USE go_gin_api_boilerplate_db;

CREATE TABLE
    restaurants (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        description TEXT,
        location VARCHAR(255),
        rating FLOAT
    );