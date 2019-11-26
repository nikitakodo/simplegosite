create table bookmarks
(
	id serial not null,
	recipe_id int not null
		constraint bookmarks_recipe_id_fk
			references recipe,
	author_id int not null
		constraint bookmarks_author_id_fk
			references author,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create index bookmarks_author_id_index
	on bookmarks (author_id);

create unique index bookmarks_id_uindex
	on bookmarks (id);

create index bookmarks_recipe_id_author_id_index
	on bookmarks (recipe_id, author_id);

create index bookmarks_recipe_id_index
	on bookmarks (recipe_id);

alter table bookmarks
	add constraint bookmarks_pk
		primary key (id);

create trigger set_timestamp
before update on bookmarks
for each row
execute procedure trigger_set_timestamp();