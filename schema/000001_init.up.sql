CREATE TABLE users
(
    id serial not null unique,
    name1 varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null
);