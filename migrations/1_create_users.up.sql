CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar,
    surname varchar,
    login varchar not null unique,
    password varchar not null,
    birthdate date
);
