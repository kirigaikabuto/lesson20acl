package lesson20acl

type CreateRoleCommand struct {
	Name string `json:"name"`
}

func (cmd *CreateRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RoleService).CreateRole(cmd)
}

type GetRoleCommand struct {
	Id string `json:"id"`
}

func (cmd *GetRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RoleService).GetRole(cmd)
}

type DeleteRoleCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteRoleCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(RoleService).DeleteRole(cmd)
}

type ListRoleCommand struct {
}

func (cmd *ListRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RoleService).ListRole(cmd)
}
