package lesson20acl

type RoleStore interface {
	Create(role *Role) (*Role, error)
	List() ([]Role, error)
	Delete(id string) error
	Get(id string) (*Role, error)
}
