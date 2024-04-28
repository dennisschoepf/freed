-- +migrate Up
CREATE table users (
  id INTEGER PRIMARY KEY,
  first_name text NOT NULL,
  email text NOT NULL
);

-- +migrate Down
DROP TABLE users;
