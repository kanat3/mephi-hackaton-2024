drop schema if exists public cascade;
create schema public;

create table if not exists "user" (
    id uuid primary key,
    login varchar(16) unique
);

create table if not exists meeting  (
    id uuid primary key,
    date timestamp,
    file_path varchar(225)
);

create table if not exists emotion (
    id serial primary key,
    name varchar(20)
);

create table if not exists timepoint (
    id uuid primary key,
    time timestamp not null,
    user_id uuid not null references "user"(id),
    meeting_id uuid not null references meeting(id),
    emotion_id int not null references emotion(id)
);