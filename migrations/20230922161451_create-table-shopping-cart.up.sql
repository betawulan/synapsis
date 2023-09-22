create table shopping_cart (
    id int auto_increment not null,
    user_id int not null,
    product_category_id int not  null,
    primary key (id),
    foreign key (user_id) references user(id),
    foreign key (product_category_id) references product_category(id)
);