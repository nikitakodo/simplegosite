create table cuisine
(
	id serial not null,
	name text not null,
	description text,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index cuisine_id_uindex
	on cuisine (id);

create index cuisine_name_index
	on cuisine (name);

create unique index cuisine_name_uindex
	on cuisine (name);

alter table cuisine
	add constraint cuisine_pk
		primary key (id);

alter table recipe
	add constraint recipe_cuisine_id_fk
		foreign key (cuisine_id) references cuisine;

