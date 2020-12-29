package postgresql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// UserModel uses DB to store Users.
type UserModel struct {
	DB *sql.DB
}

// Insert inserts User in DB.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate verifies email+pass and returns ID if correct.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get returns User with id.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
