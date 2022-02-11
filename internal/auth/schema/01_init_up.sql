create table users (
    user_id uuid primary key not null,
    username varchar(50) not null,
    password varchar(150) not null,
    firstname varchar(50) not null,
    lastname varchar(50) not null
);

create unique index on users (username) include (password);

VACUUM ANALYSE;