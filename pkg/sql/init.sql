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
    permission_id int references permissions(id),
    created_at timestamp default current_timestamp
);

create table role_permission(
  role_id int references roles(id),
  permission_id int references permissions(id)
);

drop table users;

create table users(
    id serial primary key,
    full_name text,
    phone text unique not null,
    email text unique,
    role_id int references roles(id),
    position text,
    password text not null,
    active bool default true,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);