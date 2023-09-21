create table user (
    id int auto_increment not null,
    name varchar(100) not null,
    role varchar(10) not null,
    email varchar(100) not null,
    password varchar(10) not null,
    primary key (id)
);