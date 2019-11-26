create table admin
(
	id serial not null,
	name text not null,
	login text not null,
	password text not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index admin_id_uindex
	on admin (id);

alter table admin
	add constraint admin_pk
		primary key (id);

