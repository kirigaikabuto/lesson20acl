package lesson20acl

import (
	"github.com/google/uuid"
)

type RoleService interface {
	CreateRole(cmd *CreateRoleCommand) (*Role, error)
	GetRole(cmd *GetRoleCommand) (*Role, error)
	ListRole(cmd *ListRoleCommand) ([]Role, error)
	DeleteRole(cmd *DeleteRoleCommand) error
}

type roleService struct {
	store RoleStore
}

func NewRoleService(s RoleStore) RoleService {
	return &roleService{store: s}
}

func (r *roleService) CreateRole(cmd *CreateRoleCommand) (*Role, error) {
	obj := uuid.New()
	role := &Role{Name: cmd.Name}
	role.Id = obj.String()
	return r.store.Create(role)
}

func (r *roleService) GetRole(cmd *GetRoleCommand) (*Role, error) {
	return r.store.Get(cmd.Id)
}

func (r *roleService) ListRole(cmd *ListRoleCommand) ([]Role, error) {
	return r.store.List()
}

func (r *roleService) DeleteRole(cmd *DeleteRoleCommand) error {
	return r.store.Delete(cmd.Id)
}
