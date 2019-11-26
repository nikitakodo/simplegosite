create table comments
(
	id serial not null,
	parent_id int not null,
	recipe_id int not null
		constraint comments_recipe_id_fk
			references recipe,
	author_id int not null
		constraint comments_author_id_fk
			references author,
	text text not null,
	is_published bool default false not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create index comments_author_id_index
	on comments (author_id);

create unique index comments_id_uindex
	on comments (id);

create index comments_recipe_id_author_id_index
	on comments (recipe_id, author_id);

create index comments_recipe_id_index
	on comments (recipe_id);

create index comments_recipe_id_parent_id_index
	on comments (recipe_id, parent_id);

alter table comments
	add constraint comments_pk
		primary key (id);

create trigger set_timestamp
before update on comments
for each row
execute procedure trigger_set_timestamp();