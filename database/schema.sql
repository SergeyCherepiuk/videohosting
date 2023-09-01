create table users (
    id uuid primary key default gen_random_uuid(),
    email varchar(50) unique not null,
    username varchar(50) unique not null,
    password varchar(256) not null,
    family_name varchar(40) not null,
    given_name varchar(40) not null,
    profile_picture_id uuid unique,
    created_at timestamp not null default current_timestamp
);

-- insert into users (
--     email, username, password, family_name, given_name, profile_picture_id
-- ) values (
--     'j.doe@gmail.com', 'johndoe', 'hashed_secret', 'Doe', 'John', gen_random_uuid()
-- ) returning *;

create table videos (
    id uuid primary key default gen_random_uuid(),
    author_id uuid references users(id) not null,
    title varchar(100) not null,
    description varchar(5000) not null,
    preview_id uuid unique,
    views_count bigint not null default 0,
    created_at timestamp not null default current_timestamp
);

-- insert into videos (
--     author_id, title, description, preview_id
-- ) values (
--     'a01d4b06-742b-49ef-bbcf-ef75a2d28706', 'Title', 'Description', gen_random_uuid()  
-- ) returning *;