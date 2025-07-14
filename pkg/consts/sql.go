package consts

const (
	GetUserPermissionsByIdSQL = "select permission_id from users as u join role_permission rp on u.role_id = rp.role_id where u.active = 1 and u.id = ?"
)
