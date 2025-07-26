drop table debtors, supplier_shops, users,password_resets;


alter table notifications
    drop foreign key notifications_supplier_id_foreign;

alter table notifications
    drop column supplier_id;


rename table sale_points to sale_point_types;
create table  permissions(
                            id serial primary key,
                            name text not null,
                            active bool default true,
                            created_at timestamp default current_timestamp,
                            updated_at timestamp,
                            deleted_at timestamp
);

create table routes(
                       id serial primary key,
                       path text not null,
                       permission_id bigint unsigned references permissions(id),
                       created_at timestamp default current_timestamp
);

create table role_permission(
                                role_id bigint unsigned references roles(id),
                                permission_id bigint unsigned  references permissions(id)
);

create table users(
                      id serial primary key,
                      full_name text,
                      phone varchar(20),
                      email text,
                      role_id bigint unsigned references roles(id),
                      position text,
                      password text not null,
                      shop_id bigint,
                      supplier_id bigint,
                      active smallint default 1,
                      created_at timestamp default current_timestamp,
                      updated_at timestamp,
                      deleted_at timestamp
);

create table sale_points(
                            id bigint primary key,
                            name text not null,
                            shop_id bigint unsigned references shops(id),
                            user_id bigint unsigned references shops(id),
                            status smallint default 1,
                            created_at timestamp default current_timestamp,
                            updated_at timestamp,
                            deleted_at timestamp
);

create table user_auth(
                          user_id bigint unsigned references users(id),
                          session_id text,
                          pass_reset_at timestamp,
                          temporary_pass smallint default 0,
                          updated_at timestamp default current_timestamp
);

alter table sale_point_supplier
    change sale_point_id sale_point_type bigint unsigned not null;

alter table notifications
    add status smallint default 1;


insert into users(full_name, phone, email, role_id, position, password)
select fullname, phone, '', 1, 'test', password
from workers;

insert into roles(name)
values ('Магазин'),
       ('Поставщик');


insert into users(full_name, phone, email, role_id, position, password, shop_id)
select fullname, phone, email, (select id from roles where name = 'Магазин'), 'Магазин', password, id
from shops;

insert into users(full_name, phone, email, role_id, position, password, supplier_id)
select fullname, phone, email, (select id from roles where name = 'Поставщик'), 'Поставщик', password, id
from suppliers;

insert into user_auth(user_id)
select id
from users;


update user_auth set pass_reset_at = current_timestamp where pass_reset_at is null;

alter table shops
    drop column password;

alter table suppliers drop column password;




