CREATE TABLE users
(
    id              int generated always as identity,
    created_at      timestamptz NOT NULL DEFAULT NOW(),
    updated_at      timestamptz NOT NULL DEFAULT NOW(),
    archived_at     timestamptz,
    forename        text        NOT NULL,
    surname         text        NOT NULL,
    email           text        NOT NULL,
    hashed_password text        NOT NULL
);
