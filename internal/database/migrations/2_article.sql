-- +migrate Up
CREATE table article (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL UNIQUE,
  read INTEGER DEFAULT 0,
  readAt text,
  feedId INTEGER,
  FOREIGN KEY (feedId) REFERENCES feed(id)
);

-- +migrate Down
DROP TABLE article;
