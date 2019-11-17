create table settings
(
	id serial not null,
	key text not null,
	value text not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index settings_id_uindex
	on settings (id);

create index settings_key_index
	on settings (key);

create unique index settings_key_uindex
	on settings (key);

alter table settings
	add constraint settings_pk
		primary key (id);

