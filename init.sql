CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE message 
(
    id serial not null unique,
    message varchar(255)
);


CREATE TABLE users_messages
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    message_id int references message (id) on delete cascade not null
);

