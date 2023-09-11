create table users (
    id uuid primary key,
    email varchar(50) unique not null,
    username varchar(50) unique not null,
    password varchar(256) not null,
    family_name varchar(40) not null,
    given_name varchar(40) not null,
    profile_picture_id uuid unique,
    created_at timestamp not null
);

create table videos (
    id uuid primary key,
    author_id uuid references users(id) on delete cascade on update cascade not null,
    title varchar(100) not null,
    description varchar(5000) not null,
    preview_id uuid unique,
    views_count bigint not null,
    created_at timestamp not null
);
