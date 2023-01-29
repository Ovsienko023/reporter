CREATE EXTENSION pgcrypto;

create schema main authorization postgres;
grant all on schema main to postgres;

-- Таблица для статусов
create table if not exists main.statuses
(
    id varchar primary key
);

insert into main.statuses (id)
values ('approved')
on conflict do nothing;

create table if not exists main.users
(
    id           uuid primary key     default gen_random_uuid(),
    created_at   timestamptz not null default now(),
    display_name varchar     not null,
    login        varchar     not null,
    hash         text        not null,
    deleted_at   timestamptz
);

insert into main.users (display_name, login, hash)
values ('Administrator', 'admin', '$2a$10$tjATGo3v5KrXQ6.cQz1CbugeBKRyJdmDbwr20rFMtzVJxOHtw3EIi');


create table if not exists main.reports
(
    id           uuid primary key     default gen_random_uuid(),
    display_name varchar,
    date         timestamptz not null,
    start_time   integer,
    end_time     integer,
    break_time   integer,
    work_time    integer,
    body         varchar,
    creator_id   uuid references main.users (id),
    created_at   timestamptz not null default now(),
    updated_at   timestamptz not null default now(),
    deleted_at   timestamptz
);

create table if not exists main.reports_to_users
(
    report_id uuid not null references main.reports (id),
    user_id   uuid not null references main.users (id)
);


create table if not exists main.groups
(
    id           uuid primary key     default gen_random_uuid(),
    display_name varchar,
    creator_id   uuid references main.users (id),
    created_at   timestamptz not null default now(),
    updated_at   timestamptz not null default now(),
    deleted_at   timestamptz
);

create table if not exists main.groups_to_objects
(
    group_id    uuid references main.groups (id),
    object_type varchar,
    object_id   uuid
);


create table if not exists main.roles
(
    id          varchar primary key,
    description varchar
);
insert into main.roles (id, description)
values ('administrator', null);

create table if not exists main.users_to_roles
(
    user_id uuid    not null references main.users (id),
    role_id varchar not null references main.roles (id),
    primary key (user_id, role_id)
);


insert into main.users_to_roles (user_id, role_id)
values ((select id from main.users where display_name = 'Administrator'),
        (select id from main.roles where id = 'administrator'));


create table if not exists main.permissions_users_to_objects
(
    user_id     uuid not null references main.users (id),
    object_type varchar,
    object_id   uuid
);

-- Таблица для больничного
create table if not exists main.sick_leave
(
    id          uuid primary key     default gen_random_uuid(),
    date        timestamptz not null,
    is_paid     bool        not null,
    status      varchar references main.statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted_at  timestamptz
);

-- Таблица для отпуска
create table if not exists main.vacation
(
    id          uuid primary key     default gen_random_uuid(),
    date        timestamptz not null,
    is_paid     bool        not null,
    status      varchar references main.statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted_at  timestamptz
);

-- Вьюха для отображения ивентов
create or replace view main.events as
with tab as (select r.id               as id,
                    'report':: varchar as event_type,
                    r.date             as date
             from main.reports as r
             union all
             select v.id                 as id,
                    'vacation':: varchar as event_type,
                    v.date               as date
             from main.vacation as v
             union all
             select v.id                   as id,
                    'sick_leave':: varchar as event_type,
                    v.date                 as date
             from main.sick_leave as v)

select a.id         as id,
       a.event_type as event_type,
       a.date       as date
from tab as a;
