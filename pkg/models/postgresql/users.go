package postgresql

import (
	"database/sql"
	"errors"
	"snippetbox/pkg/models"

	pq "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// UserModel uses DB to store Users.
type UserModel struct {
	DB *sql.DB
}

// Insert inserts User in DB.
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password)
	VALUES($1, $2, $3)`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			// may use additional check for users_uc_email in pgError.Message
			if pqError.Code == "23505" {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

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
