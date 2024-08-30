-- +migrate Up
CREATE table feed (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL UNIQUE,
  type text NOT NULL
);

-- +migrate Down
DROP TABLE feed;
