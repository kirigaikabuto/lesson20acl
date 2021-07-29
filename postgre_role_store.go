package lesson20acl

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

var rolePostgreQueries = []string{
	`create table if not exists roles(
		id text,
		name text,
		primary key(id)
	);`,
}

type roleStore struct {
	db *sql.DB
}

func NewPostgresRoleStore(cfg PostgresConfig) (RoleStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range rolePostgreQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &roleStore{db: db}
	return store, nil
}

func (r *roleStore) Create(role *Role) (*Role, error) {
	query := "insert into roles (id, name) values ($1, $2)"
	result, err := r.db.Exec(query, role.Id, role.Name)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("error during creating of role")
	}
	return role, nil
}

func (r *roleStore) Delete(id string) error {
	query := "delete from roles where id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("no role by this id")
	}
	return nil
}

func (r *roleStore) Get(id string) (*Role, error) {
	role := &Role{}
	query := "select id, name from roles where id = $1 limit 1"
	err := r.db.QueryRow(query, id).Scan(&role.Id, &role.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("no role by this id")
	} else if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleStore) List() ([]Role, error) {
	roles := []Role{}
	query := "select id, name from roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := Role{}
		err = rows.Scan(&item.Id, &item.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, item)
	}
	return roles, nil
}
