drop table posts;
drop table threads;
drop table sessions;
drop table users;

CREATE TABLE users (
    id      SERIAL PRIMARY KEY,
    uuid    VARCHAR(64) NOT NULL UNIQUE,
    name    VARCHAR(255),
    email   VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    create_at   TIMESTAMP NOT NULL
);

CREATE TABLE sessions (
    id      SERIAL PRIMARY KEY,
    uuid    VARCHAR(64) NOT NULL UNIQUE,
    email   VARCHAR(255),
    user_id INTEGER REFERENCES users(id),
    created_at  TIMESTAMP NOT NULL
);

