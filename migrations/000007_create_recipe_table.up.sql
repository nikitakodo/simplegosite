create table recipe
(
	id serial not null,
	title text not null,
	body text not null,
	img text not null,
	category_id int not null,
	cuisine_id int,
	author_id int,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index recipe_id_uindex
	on recipe (id);

alter table recipe
	add constraint recipe_pk
		primary key (id);

