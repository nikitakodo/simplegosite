create table author
(
	id serial not null,
	name text not null,
	login text not null,
	password text not null,
	is_banned bool default false not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index author_id_uindex
	on author (id);

alter table author
	add constraint author_pk
		primary key (id);

alter table recipe
	add constraint recipe_author_id_fk
		foreign key (author_id) references author
			on delete restrict;

alter table mark
	add constraint mark_author_id_fk
		foreign key (author_id) references author
			on delete restrict;

