create table about
(
    id serial not null,
    title text not null,
    text text not null,
    img text not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index about_id_uindex
    on about (id);

alter table about
    add constraint about_pk
        primary key (id);

create trigger add
    before update on about
    for each row
execute procedure trigger_set_timestamp();