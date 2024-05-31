-- +migrate Up
ALTER TABLE feed
ADD COLUMN type text;

-- +migrate Down
ALTER TABLE feed
DROP COLUMN type;
