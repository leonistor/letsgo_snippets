package mock

import (
	"snippetbox/pkg/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Ana are mock",
	Content: "Ana are mock und lorem ipsum dolor siat amet ...",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel mock
type SnippetModel struct{}

// Insert mock
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get mock
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest mock
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
