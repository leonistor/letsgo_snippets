package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord indicates no record found.
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials when users login with invalid email or pass
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail when signup with exisiting email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet represents a snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User is a user.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
