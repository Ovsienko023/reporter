CREATE EXTENSION pgcrypto;

create schema main;

create table if not exists main.users
(
    id             uuid primary key     default gen_random_uuid(),
    created_at     timestamptz not null default now(),
    display_name   varchar     not null,
    login          varchar     not null,
    hash           text        not null,
    deleted_at     timestamptz
);

insert into main.users (display_name, login, hash) values ('Babibtsev', 'babibtsev', 'dp');
insert into main.users (display_name, login, hash) values ('Ovsienko', 'ovsienko', 'dp');

create table if not exists main.reports
(
    id             uuid primary key       default gen_random_uuid(),
    title          varchar,
    date           timestamptz not null,
    start_time     timestamp,
    end_time       timestamp,
    break_time     timestamp,
    work_time      timestamp,
    body           varchar,
    creator_id     uuid        references main.users (id),
    created_at     timestamptz not null   default now(),
    updated_at     timestamptz not null   default now(),
    deleted_at     timestamptz
);

create table if not exists main.reports_to_users
(
    report_id uuid not null references main.reports (id),
    user_id   uuid not null references main.users (id)
);