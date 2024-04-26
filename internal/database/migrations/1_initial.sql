-- +migrate Up
CREATE table users (
  id INTEGER PRIMARY KEY,
  name text NOT NULL
);

-- +migrate Down
DROP TABLE users;
