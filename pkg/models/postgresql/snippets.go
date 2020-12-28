package postgresql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// SnippetModel wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get return a specific snippet.
func (m *SnippetModel) Get(id int) (int, error) {
	return 0, nil
}

// Latest returns the most 10 recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
