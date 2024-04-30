-- +migrate Up
CREATE table feed (
  id text PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL,
  user_id text NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE table article (
  id text PRIMARY KEY,
  name text NOT NULL,
  url text NOT NULL,
  read INTEGER NOT NULL DEFAULT 0,
  readAt TIMESTAMP,
  feed_id INTEGER NOT NULL,
  FOREIGN KEY (feed_id) REFERENCES feed (id)
);

-- +migrate Down
DROP TABLE article;
DROP TABLE feed;
