create table ingredients
(
	id serial not null,
	body text not null,
	recipe_id int not null
		constraint ingredients_recipe_id_fk
			references recipe
				on delete restrict,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index ingredients_id_uindex
	on ingredients (id);

create index ingredients_recipe_id_index
	on ingredients (recipe_id);

alter table ingredients
	add constraint ingredients_pk
		primary key (id);

create TRIGGER set_timestamp
BEFORE update ON ingredients
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();