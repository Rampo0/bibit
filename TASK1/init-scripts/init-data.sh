#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE parentdb;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "parentdb" <<-EOSQL
    CREATE TABLE users (
        id INT,
        username VARCHAR(100),
        parent INT
    );
    INSERT INTO Users (id, username, parent)
    VALUES (1,'Ali',2);
    INSERT INTO Users (id, username, parent)
    VALUES (2,'Budi',0);
    INSERT INTO Users (id, username, parent)
    VALUES (3,'Cecep',1);
EOSQL


  