rename table sale_points to sale_point_types;
create table permissions(
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
    phone text not null unique,
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

create table users_shop(
                           user_id bigint references users(id),
                           shop_id bigint references shops(id)
);

create table users_supplier(
                               user_id bigint references users(id),
                               supplier_id bigint references suppliers(id)
);