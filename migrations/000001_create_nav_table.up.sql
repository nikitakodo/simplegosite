create table nav
(
	id serial not null,
	"order" int not null,
	title text not null,
	uri text not null
);

create unique index nav_id_uindex
	on nav (id);

create unique index nav_title_uindex
	on nav (title);

alter table nav
	add constraint nav_pk
		primary key (id);
