create table transaction_detail (
    id int auto_increment not null,
    transaction_id int not null,
    product_category_id int,
    primary key (id),
    foreign key (transaction_id) references transaction(id),
    foreign key (product_category_id) references product_category(id)
);