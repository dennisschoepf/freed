-- +migrate Up
CREATE table user (
  id text PRIMARY KEY,
  first_name text NOT NULL,
  email text NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE user;
