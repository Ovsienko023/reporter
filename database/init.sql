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

--Таблица для состояний ивентов
create table if not exists main.event_states
(
    id varchar primary key
);

insert into main.event_states (id)
values ('draft'),
       ('ready')
on conflict do nothing;

-- USERS

create table if not exists main.users
(
    id           uuid primary key     default gen_random_uuid(),
    display_name varchar     not null,
    creator_id   uuid        not null references main.users (id),
    created_at   timestamptz not null default now(),
    deleted_at   timestamptz,
    payload      jsonb
);

-- CREDENTIALS

create table if not exists main.user_passwords
(
    id         uuid primary key     default gen_random_uuid(),
    created_at timestamptz not null default now(),
    creator_id uuid        not null references main.users (id),
    user_id    uuid        not null references main.users (id),
    hash       text        not null,
    deleted_at timestamptz
);

create index if not exists __created_at_idx on main.user_passwords (created_at) where deleted_at is null;
create index if not exists __deleted_at_idx on main.user_passwords (deleted_at) where deleted_at is not null;


create table if not exists main.user_logins
(
    id         uuid primary key     default gen_random_uuid(),
    login      varchar     not null,
    grant_id   uuid        not null references main.user_passwords (id),
    created_at timestamptz not null default now(),
    creator_id uuid        not null references main.users (id),
    deleted_at timestamptz
);

create unique index if not exists __login_idx on main.user_logins (lower(login)) where deleted_at is null;
create index if not exists __created_at_idx on main.user_logins (created_at) where deleted_at is null;
create index if not exists __deleted_at_idx on main.user_logins (deleted_at) where deleted_at is not null;

-- ROLES

create table if not exists main.roles
(
    id          varchar primary key,
    description varchar
);

insert into main.roles (id, description)
values ('administrator', null),
       ('default', null);

create table if not exists main.users_to_roles
(
    user_id uuid    not null references main.users (id),
    role_id varchar not null references main.roles (id),
    primary key (user_id, role_id)
);


-- todo не дупускать дубли в таблицу
create table if not exists main.permissions_users_to_objects
(
    user_id     uuid not null references main.users (id),
    object_type varchar,
    object_id   uuid
);

-- init default user

do
$$
    declare
        _user_id     uuid := gen_random_uuid();
        _password_id uuid;
    begin
        insert into main.users(id, creator_id, display_name)
        values (_user_id, _user_id, 'Administrator')
        returning id into _user_id;
        insert into main.user_passwords(creator_id, user_id, hash)
        values (_user_id, _user_id, '$2a$10$tjATGo3v5KrXQ6.cQz1CbugeBKRyJdmDbwr20rFMtzVJxOHtw3EIi')
        returning id into _password_id;
        insert into main.user_logins(creator_id, login, grant_id) values (_user_id, 'admin', _password_id);
        insert into main.users_to_roles (user_id, role_id) values (_user_id, 'administrator');
        insert into main.permissions_users_to_objects (user_id, object_type, object_id)
        values (_user_id, 'users', _user_id);
        commit;
    end
$$;

-- REPORTS

create table if not exists main.reports
(
    id           uuid primary key                          default gen_random_uuid(),
    display_name varchar,
    state        varchar references main.event_states (id) default 'draft',
    date         timestamptz not null,
    start_time   integer,
    end_time     integer,
    break_time   integer,
    work_time    integer,
    body         varchar,
    creator_id   uuid references main.users (id),
    created_at   timestamptz not null                      default now(),
    updated_at   timestamptz not null                      default now(),
    deleted_at   timestamptz,
    payload      jsonb
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
    deleted_at   timestamptz,
    payload      jsonb
);

create table if not exists main.groups_to_objects
(
    group_id    uuid references main.groups (id),
    object_type varchar,
    object_id   uuid
);


-- Таблица для больничного
create table if not exists main.sick_leave
(
    id          uuid primary key                          default gen_random_uuid(),
    date        timestamptz not null,
    is_paid     bool        not null,
    state       varchar references main.event_states (id) default 'draft',
    status      varchar references main.statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null                      default now(),
    updated_at  timestamptz not null                      default now(),
    deleted_at  timestamptz,
    payload     jsonb
);

create table if not exists main.sick_leave_to_users
(
    user_id       uuid not null references main.users (id),
    sick_leave_id uuid not null references main.sick_leave (id),
    primary key (user_id, sick_leave_id)
);

-- Таблица для отпуска
create table if not exists main.vacation
(
    id          uuid primary key                          default gen_random_uuid(),
    date        timestamptz not null,
    is_paid     bool        not null,
    state       varchar references main.event_states (id) default 'draft',
    status      varchar references main.statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null                      default now(),
    updated_at  timestamptz not null                      default now(),
    deleted_at  timestamptz,
    payload     jsonb
);

create table if not exists main.vacation_to_users
(
    user_id     uuid not null references main.users (id),
    vacation_id uuid not null references main.vacation (id),
    primary key (user_id, vacation_id)
);

-- Вьюха для отображения ивентов
create or replace view main.events as
with tab as (select r.id               as id,
                    rtu.user_id        as user_id,
                    'report':: varchar as event_type,
                    r.date             as date
             from main.reports as r
                      inner join main.reports_to_users rtu on r.id = rtu.report_id
             union all
             select v.id                 as id,
                    vtu.user_id          as user_id,
                    'vacation':: varchar as event_type,
                    v.date               as date
             from main.vacation as v
                      inner join main.vacation_to_users vtu on v.id = vtu.vacation_id
             union all
             select s.id                   as id,
                    stu.user_id            as user_id,
                    'sick_leave':: varchar as event_type,
                    s.date                 as date
             from main.sick_leave as s
                      inner join main.sick_leave_to_users stu on s.id = stu.sick_leave_id)

select a.id         as id,
       a.user_id    as user_id,
       a.event_type as event_type,
       a.date       as date
from tab as a;


drop function if exists main.create_user;

create or replace function main.create_user(
    _invoker_id uuid,
    _login varchar,
    _password varchar,
    _display_name varchar,
    --
    out error jsonb
) as
$$
declare
    _exception   text;
    _user_id     uuid;
    _password_id uuid;
begin

    if not exists(select 1
                  from main.users
                  where id = _invoker_id
                    and deleted_at is null) then
        error = '{"code": 1, "message": "unauthorized", "details": []}'::jsonb;
        return;
    end if;

    if exists(select 1
              from main.user_logins
              where login = _login
                and deleted_at is null) then
        error = '{"code": 3, "message": "", "details": [{"name": "_login", "reason": "exists"}]}'::jsonb;
        return;
    end if;

    insert into main.users(creator_id, display_name)
    values (_invoker_id, _display_name)
    returning id into _user_id;

    insert into main.user_passwords(creator_id, user_id, hash)
    values (_invoker_id, _user_id, _password)
    returning id into _password_id;

    insert into main.user_logins(creator_id, login, grant_id)
    values (_invoker_id, _login, _password_id);

    insert into main.users_to_roles (user_id, role_id)
    values (_user_id, 'default');

    insert into main.permissions_users_to_objects (user_id, object_type, object_id)
    values (_user_id, 'users', _user_id);

    error := null;
    return;

exception
    when others then
        get stacked diagnostics _exception = PG_EXCEPTION_CONTEXT;
        _exception := _exception || ' | ' || SQLERRM || ' | ' || SQLSTATE;
        raise notice 'ERROR: % ', _exception;

        error := _exception;
        return;

end;
$$
    language plpgsql volatile
                     security definer;
