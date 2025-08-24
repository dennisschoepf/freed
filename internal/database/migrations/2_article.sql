-- +migrate Up
CREATE table article (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL UNIQUE,
  readAt DATETIME,
  feedId INTEGER,
  FOREIGN KEY (feedId) REFERENCES feed(id)
);

-- +migrate Down
DROP TABLE article;
