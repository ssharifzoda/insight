insert into users(full_name, phone, email, role_id, position, password)
select fullname, phone, '', 1, 'test', password
from workers;

insert into roles(name)
values ('Магазин'),
       ('Поставщик');


insert into users(full_name, phone, email, role_id, position, password)
select fullname, phone, email, 4, 'Магазин', password
from shops;

insert into users(full_name, phone, email, role_id, position, password)
select fullname, phone, email, 5, 'Поставщик', password
from suppliers;

insert into user_auth(user_id)
select id
from users;


update user_auth set pass_reset_at = current_timestamp where pass_reset_at is null;