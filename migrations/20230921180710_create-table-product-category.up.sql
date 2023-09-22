create table product_category (
	id int auto_increment not null,
	product_id int,
	category_id int,
	primary key (id),
	foreign key (product_id) references product(id),
	foreign key (category_id) references category(id)
);