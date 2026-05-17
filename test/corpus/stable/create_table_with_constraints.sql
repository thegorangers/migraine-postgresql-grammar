CREATE TABLE users (
    id int PRIMARY KEY,
    email text NOT NULL UNIQUE,
    name text NOT NULL
);
