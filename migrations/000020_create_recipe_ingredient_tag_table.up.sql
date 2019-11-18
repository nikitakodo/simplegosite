create table recipe_ingredient_tag
(
	id serial not null,
    recipe_id int not null
		constraint recipe_ingredient_tag_recipe_id_fk
			references recipe,
    ingredient_tag_id int not null
		constraint recipe_ingredient_tag_ingredient_tag_id_id_fk
			references recipe,
	create_time timestamptz default current_timestamp not null,
	update_time timestamptz default current_timestamp not null
);

create unique index recipe_ingredient_tag_id_uindex
	on recipe_ingredient_tag (id);

create index recipe_ingredient_tag_ingredient_tag_id_index
	on recipe_ingredient_tag (ingredient_tag_id);

create index recipe_ingredient_tag_recipe_id_index
	on recipe_ingredient_tag (recipe_id);

create index recipe_ingredient_tag_recipe_id_ingredient_tag_id_index
	on recipe_ingredient_tag (recipe_id, ingredient_tag_id);

alter table recipe_ingredient_tag
	add constraint recipe_ingredient_tag_pk
		primary key (id);

create trigger set_timestamp
before update on recipe_ingredient_tag
for each row
execute procedure trigger_set_timestamp();