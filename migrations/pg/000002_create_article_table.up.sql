create table articles
(
    id              bigserial
        primary key,
    user_id int,
    title varchar,
    description varchar
)