create table recipe_info
(
	id serial not null,
	recipe_id int not null
		constraint recipe_info_recipe_id_fk
			references recipe,
	prep int not null,
	cook int not null,
	yields int not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index recipe_info_id_uindex
	on recipe_info (id);

create index recipe_info_recipe_id_index
	on recipe_info (recipe_id);

alter table recipe_info
	add constraint recipe_info_pk
		primary key (id);

create TRIGGER set_timestamp
BEFORE update ON recipe_info
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();