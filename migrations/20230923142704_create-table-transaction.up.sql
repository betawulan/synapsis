create table transaction (
    id int auto_increment not null,
    user_id int not null,
    primary key (id),
    foreign key (user_id) references user(id)
);