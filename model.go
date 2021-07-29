package lesson20acl

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RolePermissions struct {
	Id           string `json:"id"`
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type Permission struct {
	Id       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type UserRoles struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}
