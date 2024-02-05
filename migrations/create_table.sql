DROP TABLE IF EXISTS people_data;
CREATE TABLE people_data
(
    id          serial primary key,
    name        varchar(30) not null,
    surname     varchar(30) not null,
    patronymic  varchar(30),
    age         integer     not null,
    gender      varchar(6)  not null,
    country varchar(2)
);

