-- +migrate Up
CREATE table user (
  id INTEGER PRIMARY KEY,
  first_name text NOT NULL,
  email text NOT NULL
);

-- +migrate Down
DROP TABLE users;
