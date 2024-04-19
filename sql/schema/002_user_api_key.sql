-- +goose UP
ALTER TABLE users ADD COLUMN IF NOT EXISTS api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);


-- +goose DOWN
ALTER TABLE users DROP COLUMN IF EXISTS api_key