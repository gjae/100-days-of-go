DROP DATABASE IF EXISTS restapi;
CREATE DATABASE restapi;
\c restapi;
/*
Users
*/
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    lastlogin INT,
    admin INT,
    active INT
);
INSERT INTO users (username, password, lastlogin, admin, active) VALUES ('admin', 'admin', 1620922454, 1, 1);
