create table slides
(
	id serial not null,
	first_title text not null,
	second_title text not null,
	third_title text not null,
	img text not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index slides_id_uindex
	on slides (id);

alter table slides
	add constraint slides_pk
		primary key (id);

