create database chats;
create table Chats(
    chat_id serial,
    from_id uuid not null,
    to_id uuid not null,
    unique (chat_id),
    PRIMARY KEY(from_id, to_id)
);
create table Messages(
    Message_id serial primary key ,
    chat_id integer,
    time timestamp,
    body varchar(300),
    sender uuid not null,
    foreign key (chat_id) references Chats (chat_id)
);

