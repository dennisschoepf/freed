-- +migrate Up
CREATE table feed (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL UNIQUE,
  userId text NOT NULL,
  FOREIGN KEY (userId) REFERENCES user(id)
);

-- +migrate Down
DROP TABLE feed;
