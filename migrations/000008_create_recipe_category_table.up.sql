create table category
(
	id serial not null,
	name text not null,
	description text,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index category_id_uindex
	on category (id);

create index category_name_index
	on category (name);

create unique index category_name_uindex
	on category (name);

alter table category
	add constraint category_pk
		primary key (id);

alter table recipe
	add constraint recipe_category_id_fk
		foreign key (category_id) references category;

