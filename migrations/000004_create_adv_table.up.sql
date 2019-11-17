create table adv
(
	id serial not null,
	title text not null,
	img text not null,
	content text not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index adv_id_uindex
	on adv (id);

alter table adv
	add constraint adv_pk
		primary key (id);

