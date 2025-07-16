package consts

const (
	GetOrderByIdSQL           = "SELECT     o.id,     sh.fullname AS shop_id,     o.total,     o.deliver_at,     o.comments,     o.sup_comments,     o.status,     o.verified_at,     o.delivered_at,     o.completed_at,     o.canceled,     o.who_canceled,     o.created_at,     o.updated_at,     JSON_ARRAYAGG(             JSON_OBJECT(                     'product_id', p.name,                     'qty', op.qty,                     'price', op.price             )     ) AS products FROM orders AS o          JOIN order_products AS op ON o.id = op.order_id          JOIN shops AS sh ON o.shop_id = sh.id          JOIN products AS p ON op.product_id = p.id where o.id = ? GROUP BY o.id, sh.fullname, o.total, o.deliver_at, o.comments, o.sup_comments, o.status, o.verified_at, o.delivered_at, o.completed_at, o.canceled, o.who_canceled, o.created_at, o.updated_at;"
	GetUserPermissionsByIdSQL = "select permission_id from users as u join role_permission rp on u.role_id = rp.role_id where u.active = 1 and u.id = ?"
)
