create table Users (
    user_id uuid not null,
    username varchar(50) primary key not null,
    password varchar(150),
    firstname varchar(50),
    lastname varchar(50)
);