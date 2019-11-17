create table social
(
	id serial not null,
	"order" int not null,
	icon text not null,
	url text not null,
    create_time timestamptz default current_timestamp not null,
	update_time timestamptz
);

create unique index social_id_uindex
	on social (id);

alter table social
	add constraint social_pk
		primary key (id);
