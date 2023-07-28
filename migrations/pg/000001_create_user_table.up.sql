create table users
(
    id              bigserial
        primary key,
    username        varchar(255) not null,
    first_name      varchar(255) not null,
    last_name       varchar(255) not null,
    hashed_password varchar(255) not null
)
