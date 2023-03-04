CREATE EXTENSION pgcrypto;

create schema main authorization postgres;
grant all on schema main to postgres;

-- Таблица для статусов
create table if not exists main.report_statuses
(
    id varchar primary key
);

insert into main.report_statuses (id)
values ('approved')
on conflict do nothing;

--Таблица для состояний ивентов
create table if not exists main.reports_states
(
    id varchar primary key
);

insert into main.reports_states (id)
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
    id           uuid primary key                            default gen_random_uuid(),
    display_name varchar,
    state        varchar references main.reports_states (id) default 'draft',
    date         timestamptz not null,
    start_time   integer,
    end_time     integer,
    break_time   integer,
    work_time    integer,
    body         varchar,
    creator_id   uuid references main.users (id),
    created_at   timestamptz not null                        default now(),
    updated_at   timestamptz not null                        default now(),
    deleted_at   timestamptz,
    payload      jsonb
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

-- Таблица для статусов больничных
create table if not exists main.sick_leave_statuses
(
    id varchar primary key
);

insert into main.sick_leave_statuses (id)
values ('approved')
on conflict do nothing;


-- Таблица для больничного
create table if not exists main.sick_leave
(
    id          uuid primary key     default gen_random_uuid(),
    date_from   timestamptz not null,
    date_to     timestamptz not null,
    status      varchar references main.sick_leave_statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted_at  timestamptz,
    payload     jsonb
);

-- Таблица для статусов больничных
create table if not exists main.vacations_statuses
(
    id varchar primary key
);

insert into main.vacations_statuses (id)
values ('approved')
on conflict do nothing;

-- Таблица для отпуска
create table if not exists main.vacations
(
    id          uuid primary key     default gen_random_uuid(),
    date_from   timestamptz not null,
    date_to     timestamptz not null,
    is_paid     bool        not null,
    status      varchar references main.vacations_statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted_at  timestamptz,
    payload     jsonb
);


-- Таблица для отгулов
create table if not exists main.day_off_statuses
(
    id varchar primary key
);

insert into main.day_off_statuses (id)
values ('approved')
on conflict do nothing;


-- Таблица для отгулов
create table if not exists main.day_off
(
    id          uuid primary key     default gen_random_uuid(),
    date_from   timestamptz not null,
    date_to     timestamptz not null,
    status      varchar references main.day_off_statuses (id),
    description varchar,
    creator_id  uuid references main.users (id),
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted_at  timestamptz,
    payload     jsonb
);

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
        error := '{
          "code": 1,
          "message": "unauthorized",
          "details": []
        }'::jsonb;
        return;
    end if;

    if exists(select 1
              from main.user_logins
              where login = _login
                and deleted_at is null) then
        error = '{
          "code": 3,
          "message": "",
          "details": [
            {
              "name": "_login",
              "reason": "exists"
            }
          ]
        }'::jsonb;
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


-- Вьюха для отображения ивентов
create or replace view main.events as
with tab as (select v.id                 as id,
                    v.creator_id         as user_id,
                    'vacation':: varchar as event_type,
                    v.date_from          as date_from,
                    v.date_to            as date_to
             from main.vacations as v
             union all
             select s.id                   as id,
                    s.creator_id           as user_id,
                    'day_off':: varchar as event_type,
                    s.date_from            as date_from,
                    s.date_to              as date_to
             from main.day_off as s
             union all
             select d.id                as id,
                    d.creator_id        as user_id,
                    'sick_leave':: varchar as event_type,
                    d.date_from         as date_from,
                    d.date_to           as date_to
             from main.sick_leave as d)

select a.id         as id,
       a.user_id    as user_id,
       a.event_type as event_type,
       a.date_from  as date_from,
       a.date_to    as date_to
from tab as a;


drop function if exists main.get_events;

create or replace function main.get_events(
    _invoker_id uuid,
    _date_from timestamp,
    _date_to timestamp,
    _page int = 1,
    _page_size int = 60,
    _allowed_to uuid = null,
    --
    out error jsonb,
    out count bigint,
    out event_id uuid,
    out event_type varchar,
    out date_from timestamptz,
    out date_to timestamptz
) returns setof record as
$$
declare
    _exception text;
begin

    if not exists(select 1
                  from main.users
                  where id = _invoker_id
                    and deleted_at is null) then
        error := '{
          "code": 1,
          "message": "unauthorized",
          "details": []
        }'::jsonb;

        return query (values (error::jsonb,
                              null::bigint,
                              null::uuid,
                              null::varchar,
                              null:: timestamptz,
                              null:: timestamptz));
        return;
    end if;

    if _allowed_to is not null and
       not exists(select 1
                  from main.users
                  where id = _allowed_to
                    and deleted_at is null) then

        error = '{
          "code": 3,
          "message": "",
          "details": [
            {
              "name": "_user_id",
              "reason": "not_found"
            }
          ]
        }'::jsonb;

        return query (values (error::jsonb,
                              null::bigint,
                              null::uuid,
                              null::varchar,
                              null:: timestamptz,
                              null:: timestamptz));
        return;
    end if;

    if _allowed_to is not null then
        return query with tab as (select e.id,
                                         e.event_type,
                                         e.date_from,
                                         e.date_to
                                  from main.events as e
                                  where exists(select 1
                                               from main.permissions_users_to_objects
                                               where user_id = _invoker_id
                                                 and object_id = e.user_id
                                                 and object_id = _allowed_to
                                                 and object_id != user_id)
                                    and (
                                              _date_from is null and _date_to is null or
                                              e.date_to >= _date_from and
                                              e.date_from <= _date_to
                                      ))
                     select null::jsonb                as error,
                            (select count(*) from tab) as count,
                            r.id                       as event_id,
                            r.event_type               as event_type,
                            r.date_from                as date_from,
                            r.date_to                  as date_to
                     from tab as r
                     limit _page_size offset _page_size * (_page - 1);
        return;
    end if;
    --
    return query with tab as (select e.id,
                                     e.event_type,
                                     e.date_from,
                                     e.date_to
                              from main.events as e
                              where e.user_id = _invoker_id
                                and (_date_from is null and _date_to is null or
                                     e.date_to >= _date_from and
                                     e.date_from <= _date_to))
                 select null::jsonb                as error,
                        (select count(*) from tab) as count,
                        r.id                       as event_id,
                        r.event_type               as event_type,
                        r.date_from                as date_from,
                        r.date_to                  as date_to
                 from tab as r
                 limit _page_size offset _page_size * (_page - 1);
    return;

exception
    when others then
        get stacked diagnostics _exception = PG_EXCEPTION_CONTEXT;
        _exception := _exception || ' | ' || SQLERRM || ' | ' || SQLSTATE;
        raise notice 'ERROR: % ', _exception;

        return query
            values (format('{"code": -1, "reason": "unknown", "description": "%s"}', _exception)::jsonb,
                    null::bigint,
                    null::uuid,
                    null::varchar,
                    null:: timestamptz,
                    null:: timestamptz);

end;
$$
    language plpgsql volatile
                     security definer;
