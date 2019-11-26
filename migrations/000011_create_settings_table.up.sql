create table settings
(
	id serial not null,
	key text not null,
	value text not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
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

