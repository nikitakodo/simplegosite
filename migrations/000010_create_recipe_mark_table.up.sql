create table mark
(
	id serial not null,
	recipe_id int not null,
	author_id int not null,
	mark int not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index mark_id_uindex
	on mark (id);

alter table mark
	add constraint mark_pk
		primary key (id);

alter table mark
	add constraint mark_recipe_id_fk
		foreign key (recipe_id) references recipe
			on delete restrict;