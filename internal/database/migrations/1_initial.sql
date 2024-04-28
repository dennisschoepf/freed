-- +migrate Up
CREATE table user (
  id text PRIMARY KEY,
  first_name text NOT NULL,
  email text NOT NULL
);

-- +migrate Down
DROP TABLE users;
