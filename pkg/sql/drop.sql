drop table debtors, supplier_shops, users,password_resets;


alter table notifications
    drop foreign key notifications_supplier_id_foreign;

alter table notifications
    drop column supplier_id;