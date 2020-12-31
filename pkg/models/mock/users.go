package mock

import (
	"snippetbox/pkg/models"
	"time"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Coca",
	Email:   "coca@test.com",
	Created: time.Now(),
	Active:  true,
}

// UserModel mock
type UserModel struct{}

// Insert mock
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@test.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate mock
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "coca@test.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get mock
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
