create table ingredient_tag
(
	id serial not null,
	name text not null,
	create_time timestamptz default current_timestamp  not null,
	update_time timestamptz default current_timestamp
);

create unique index ingredient_tag_id_uindex
	on ingredient_tag (id);

create unique index ingredient_tag_name_uindex
	on ingredient_tag (name);

alter table ingredient_tag
	add constraint ingredient_tag_pk
		primary key (id);

create trigger set_timestamp
before update on ingredient_tag
for each row
EXECUTE procedure trigger_set_timestamp();