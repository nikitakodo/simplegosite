create table social
(
	id serial not null,
	"order" int not null,
	icon text not null,
	url text not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index social_id_uindex
	on social (id);

alter table social
	add constraint social_pk
		primary key (id);
