create table add
(
	id serial not null,
	title text not null,
	first_item text not null,
	second_item text not null,
	third_item text not null,
	fourth_item text not null,
	first_img text not null,
	second_img text not null,
	third_img text not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index add_id_uindex
	on add (id);

alter table add
	add constraint add_pk
		primary key (id);

create trigger add
before update on add
for each row
execute procedure trigger_set_timestamp();