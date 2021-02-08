package mysql

import (
	"fmt"
	"temsys"

	"github.com/jmoiron/sqlx"
)

// SQLUserRepository is the implementation of a UserRepository using mysql.
type SQLUserRepository struct {
	sqlRepository
}

// NewUserRepo creates a UserRepository implementation using sql.DB.
func NewUserRepo(db *sqlx.DB) SQLUserRepository {
	return SQLUserRepository{newRepository(db)}
}

// Save a user. If something goes wrong (including not finding the requested user)
// it returns an error.
func (repo SQLUserRepository) Save(user temsys.User) error {
	insert := "insert into users (name, password, role) values (?, ?, ?)"
	_, err := repo.db.Exec(insert, user.Name, user.Password, user.Role)
	return err
}

// GetByName a user. Returns a valid user or an error if something goes wrong,
// including the user with the requested name doesnt exists.
func (repo SQLUserRepository) GetByName(name string) (temsys.User, error) {
	selectQuery := "select name, password, role from users where name = ?"
	var user temsys.User
	if err := repo.db.Get(&user, selectQuery, name); err != nil {
		fmt.Println(err)
		return temsys.User{}, err
	}
	return user, nil
}

// ExistsWithName check if a user exists with the requested name. Use this before
// using GetByName or Save.
func (repo SQLUserRepository) ExistsWithName(name string) bool {
	_, err := repo.GetByName(name)
	return err == nil
}
