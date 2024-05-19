CREATE TABLE messages
(
    id        SERIAL PRIMARY KEY,
    email     TEXT        NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    message   TEXT        NOT NULL
);
