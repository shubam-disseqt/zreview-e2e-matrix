-- +migrate Up
ALTER TABLE users ADD COLUMN email VARCHAR(255) NOT NULL DEFAULT '';
CREATE UNIQUE INDEX users_email_idx ON users(email) WHERE email <> '';

-- +migrate Down
DROP INDEX users_email_idx;
ALTER TABLE users DROP COLUMN email;
