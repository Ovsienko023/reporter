CREATE EXTENSION pgcrypto;

create schema hub authorization postgres;
grant all on schema hub to postgres;

create table if not exists hub.users
(
    id           uuid primary key     default gen_random_uuid(),
    display_name varchar     not null,
    created_at   timestamptz not null default now(),
    deleted_at   timestamptz,
    payload      jsonb
);
