package models

import (
	"errors"
	"time"
)

// ErrNoRecord indicates no record found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet represents a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
